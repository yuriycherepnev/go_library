package reader

type CreateReaderDTO struct {
	Name string `json:"name"`
}

type UpdateReaderDTO struct {
	Id   int    `json:"uuid"`
	Name string `json:"name"`
}

type DeleteReaderDTO struct {
	Id int `json:"uuid"`
}
