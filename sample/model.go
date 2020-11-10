package sample

type CodeContent struct {
	FileName   string `json:"fileName,omitempty"`
	Path       string `json:"path,omitempty"`
	Repository string `json:"repository,omitempty"`
	Fork       bool   `json:"fork,omitempty"`
	Content    string `json:"content:omitempty"`
}
