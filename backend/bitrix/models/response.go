package models

type APIResponse struct {
	Result []OpenLine `json:"result"`
}

type OpenLine struct {
	ID   int    `json:"ID"`
	Name string `json:"NAME"`
}
