package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ingenieux/cctocb"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	//log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
		DisableColors:    true,
	})

	h := cctocb.NewCodeCommitToCodeBuildHandler()

	if "" == os.Getenv("_LAMBDA_SERVER_PORT") {
		event := cctocb.CodeCommitRepoStateChangeEvent{
			ID:      "a8b4f0ba-484b-2d93-6f1c-a1c44e74e3a9",
			Account: 370868937120,
			Detail: cctocb.CodeCommitRepoStateChangeEventDetail{
				CommitID:          "e5c8c3391a58d058d020ecc9ef970000d7f86cc5",
				Event:             "referenceUpdated",
				ReferenceFullName: "refs/heads/master",
				ReferenceName:     "master",
				ReferenceType:     "branch",
				RepositoryID:      "da5578c6-d2cc-49b8-85fb-3d069ff5e4ec",
				RepositoryName:    "cctocb",
			},
			DetailType: "CodeCommit Repository State Change",
			Region:     "eu-west-2",
			Resources: []string{
				"arn:aws:codecommit:eu-west-2:370868937120:cctocb",
			},
			Source:  "aws.codecommit",
			Time:    time.Now(),
			Version: 0,
		}

		event = (cctocb.CodeCommitRepoStateChangeEvent)(cctocb.CodeCommitRepoStateChangeEvent{
			ID:      (string)("754a7c86-7c3e-8d4c-8362-fa7e4e9b8ffc"),
			Account: (int64)(370868937120),
			Detail: (cctocb.CodeCommitRepoStateChangeEventDetail)(cctocb.CodeCommitRepoStateChangeEventDetail{
				CommitID:          (string)("a2751f3727f48ef56c038455c2145315d5125c88"),
				Event:             (string)("referenceCreated"),
				ReferenceFullName: (string)("refs/heads/feature/dealing-with-gitflow"),
				ReferenceName:     (string)("feature/dealing-with-gitflow"),
				ReferenceType:     (string)("branch"),
				RepositoryID:      (string)("da5578c6-d2cc-49b8-85fb-3d069ff5e4ec"),
				RepositoryName:    (string)("cctocb"),
			}),
			DetailType: (string)("CodeCommit Repository State Change"),
			Region:     (string)("eu-west-2"),
			Resources: ([]string)([]string{
				(string)("arn:aws:codecommit:eu-west-2:370868937120:cctocb"),
			}),
			Source:  (string)("aws.codecommit"),
			Time:    (time.Time)(time.Date(2018, 8, 9, 6, 9, 36, 0, time.UTC)),
			Version: (int64)(0),
		})

		event = (cctocb.CodeCommitRepoStateChangeEvent)(cctocb.CodeCommitRepoStateChangeEvent{
			ID:      (string)("6bba14ba-5e8b-4b4d-ef53-bde8bf68c10a"),
			Account: (int64)(370868937120),
			Detail: (cctocb.CodeCommitRepoStateChangeEventDetail)(cctocb.CodeCommitRepoStateChangeEventDetail{
				CommitID:          (string)("07531777b08a9ac509972930b04fdb1606349d3f"),
				Event:             (string)("referenceUpdated"),
				ReferenceFullName: (string)("refs/heads/feature/slack-build-notifications"),
				ReferenceName:     (string)("feature/slack-build-notifications"),
				ReferenceType:     (string)("branch"),
				RepositoryID:      (string)("da5578c6-d2cc-49b8-85fb-3d069ff5e4ec"),
				RepositoryName:    (string)("cctocb"),
			}),
			DetailType: (string)("CodeCommit Repository State Change"),
			Region:     (string)("eu-west-2"),
			Resources: ([]string)([]string{
				(string)("arn:aws:codecommit:eu-west-2:370868937120:cctocb"),
			}),
			Source:  (string)("aws.codecommit"),
			Time:    (time.Time)(time.Date(2018, 8, 9, 7, 51, 52, 0, time.UTC)),
			Version: (int64)(0),
		})

		h.Handler(event)
	} else {
		lambda.Start(h.Handler)
	}
}
