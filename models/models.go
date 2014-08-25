package models

type InRequest struct {
	Source Source `json:"source"`
}

type Source struct {
	URI string `json:"uri"`
}

type InResponse struct{}

type CheckRequest struct{}

type CheckResponse []interface{}
