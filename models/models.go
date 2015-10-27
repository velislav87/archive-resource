package models

type InRequest struct {
	Source Source `json:"source"`
}

type Source struct {
	URI           string `json:"uri"`
	Authorization string `json:"authorization"`
}

type InResponse struct{}

type CheckRequest struct{}

type CheckResponse []interface{}
