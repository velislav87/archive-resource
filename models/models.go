package models

type InRequest struct {
	Source Source `json:"source"`
}

type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

type OutParams struct {
	DirectoriesToInclude []string `json:"dirs"`
}

type Source struct {
	URI           string `json:"uri"`
	Authorization string `json:"authorization"`
}

type InResponse struct{}

type CheckRequest struct{}

type CheckResponse []interface{}
