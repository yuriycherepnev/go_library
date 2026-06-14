package book

type CreateBookDTO struct {
	Title    string `json:"title"`
	IdAuthor int    `json:"id_author"`
}

type UpdateBookDTO struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	IdAuthor int    `json:"id_author"`
	IdReader int    `json:"id_reader"`
}

type DeleteBookDTO struct {
	Id int `json:"id"`
}
