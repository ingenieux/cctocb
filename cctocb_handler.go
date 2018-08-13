package cctocb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/gosimple/slug"
	"github.com/shurcooL/go-goon"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type CodeCommitToCodeBuildHandler struct {
	region           string
	sess             *session.Session
	codebuildService *codebuild.CodeBuild
}

func NewCodeCommitToCodeBuildHandler() *CodeCommitToCodeBuildHandler {
	region := "eu-east-1"

	if newRegion := os.Getenv("AWS_REGION"); "" != newRegion {
		region = newRegion
	}

	sess := session.Must(session.NewSession(aws.NewConfig().WithRegion(region)))

	codebuildService := codebuild.New(sess)

	return &CodeCommitToCodeBuildHandler{
		region:           region,
		sess:             sess,
		codebuildService: codebuildService,
	}
}

func (m *CodeCommitToCodeBuildHandler) listProjects(accountId, region string) (map[string]string, error) {
	result := make(map[string]string)
	input := &codebuild.ListProjectsInput{
		SortBy:    aws.String("LAST_MODIFIED_TIME"),
		SortOrder: aws.String("DESCENDING"),
	}

	for {
		listProjectsOutput, err := m.codebuildService.ListProjects(input)

		if nil != err {
			log.Warnf("Oops: %s", err)

			return result, err
		}

		for _, proj := range listProjectsOutput.Projects {
			k := fmt.Sprintf("arn:aws:codecommit:%s:%s:%s", region, accountId, *proj)

			result[k] = *proj
		}

		if nil == listProjectsOutput.NextToken {
			break
		} else {
			input.NextToken = listProjectsOutput.NextToken
		}
	}

	return result, nil
}

func (m *CodeCommitToCodeBuildHandler) Handler(event CodeCommitRepoStateChangeEvent) error {
	fmt.Println(goon.Sdump(event))

	accountId := fmt.Sprintf("%d", event.Account)

	log.Infof("Listing projects for accountId=%s and region=%s", accountId, m.region)

	projectList, err := m.listProjects(accountId, m.region)

	if nil != err {
		log.Warnf("listProjects: %s", err)

		return err
	}

	log.Infof("projectList: %+v", projectList)

	for _, project := range event.Resources {
		if projectRef, ok := projectList[project]; ok {
			log.Infof("Found project %s (%s)", projectRef, project)

			if ("referenceCreated" == event.Detail.Event || "referenceUpdated" == event.Detail.Event) && "branch" == event.Detail.ReferenceType {
				idempotencyToken := fmt.Sprintf("eventid-%s", event.ID)

				rnFeatureBranchP := strings.HasPrefix(event.Detail.ReferenceName, "feature/")

				rnFeatureBranchPValue := fmt.Sprintf("%t", rnFeatureBranchP)

				startBuildInput := &codebuild.StartBuildInput{
					ProjectName:      aws.String(projectRef),
					SourceVersion:    aws.String(event.Detail.CommitID),
					IdempotencyToken: aws.String(idempotencyToken),
					EnvironmentVariablesOverride: []*codebuild.EnvironmentVariable{
						{
							Name:  aws.String("CC2CB_BUILD_BRANCH"),
							Value: aws.String(event.Detail.ReferenceName),
						},
						{
							Name:  aws.String("CC2CB_EVENT_ID"),
							Value: aws.String(event.ID),
						},
						{
							Name:  aws.String("CC2CB_EVENT_TYPE"),
							Value: aws.String(slug.Make(event.Detail.Event)),
						},
					},
				}

				if rnFeatureBranchP {
					rnFeatureName := event.Detail.ReferenceName[len("feature/"):]
					rnFeatureName = slug.Make(rnFeatureName)

					startBuildInput.EnvironmentVariablesOverride = append(startBuildInput.EnvironmentVariablesOverride,
						[]*codebuild.EnvironmentVariable{
							{
								Name:  aws.String("CC2CB_FEATURE_BRANCH"),
								Value: aws.String(rnFeatureBranchPValue),
							},
							{
								Name:  aws.String("CC2CB_FEATURE_NAME"),
								Value: aws.String(rnFeatureName),
							},
						}...)
				}

				fmt.Println(goon.Sdump(*startBuildInput))

				startBuildOutput, err := m.codebuildService.StartBuild(startBuildInput)

				if nil != err {
					log.Warnf("StartBuild: %s", err)

					return err
				} else {
					log.Infof("Created a new Build: %s", startBuildOutput.Build.GoString())
				}
			} else {
				log.Infof(
					"Ignoring event (type: %s, referenceType: %s)",
					event.Detail.Event,
					event.Detail.ReferenceType)
			}
		} else {
			log.Infof("No reference found for project %s. Skipping", project)
		}
	}

	return nil
}
