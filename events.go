package cctocb

import "time"

type CodeCommitRepoStateChangeEventDetail struct {
	CommitID          string `json:"commitId"`
	Event             string `json:"event"`
	ReferenceFullName string `json:"referenceFullName"`
	ReferenceName     string `json:"referenceName"`
	ReferenceType     string `json:"referenceType"`
	RepositoryID      string `json:"repositoryId"`
	RepositoryName    string `json:"repositoryName"`
}

type CodeCommitRepoStateChangeEvent struct {
	ID         string                               `json:"id"`
	Account    int64                                `json:"account,string"`
	Detail     CodeCommitRepoStateChangeEventDetail `json:"detail"`
	DetailType string                               `json:"detail-type"`
	Region     string                               `json:"region"`
	Resources  []string                             `json:"resources"`
	Source     string                               `json:"source"`
	Time       time.Time                            `json:"time"`
	Version    int64                                `json:"version,string"`
}
