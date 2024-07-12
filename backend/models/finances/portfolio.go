package models

type Assets struct {
	ID         string     `json:"id"`
	Properties []Property `json:"properties"`
}

type Portfolio struct {
	ID     string   `json:"id"`
	Assets []Assets `json:"assets"`
}
