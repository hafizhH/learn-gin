package models

type BookDetails struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	Author     string `json:"author"`
	Summary    string `json:"summary"`
	Publisher  string `json:"publisher"`
	PageCount  int    `json:"pageCount"`
	ReadPage   int    `json:"readPage"`
	Reading    bool   `json:"reading"`
	Finished   bool   `json:"finished"`
	InsertedAt string `json:"insertedAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type BookBody struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}
