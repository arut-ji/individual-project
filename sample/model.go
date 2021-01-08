package sample

type CodeContent struct {
	FileName      string `json:"fileName,omitempty" bson:"fileName,omitempty"`
	Path          string `json:"path,omitempty" bson:"path,omitempty"`
	Repository    string `json:"repository,omitempty" bson:"repository,omitempty"`
	RepositoryId  int64  `json:"repositoryId,omitempty" bson:"repositoryId,omitempty"`
	Fork          bool   `json:"fork,omitempty" bson:"fork,omitempty"`
	LintingResult bool   `json:"lintingResult,omitempty" bson:"lintingResult,omitempty"`
	Content       string `json:"content,omitempty" bson:"content,omitempty"`
	CommitCount   int64  `json:"commitCount,omitempty" bson:"commitCount,omitempty"`
	Owner         string `json:"owner,omitempty" bson:"owner,omitempty"`
}
