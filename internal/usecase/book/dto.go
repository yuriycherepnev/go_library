package book

type CreateBookDTO struct {
	Title    string `json:"title"`
	IdAuthor int    `json:"id_author"`
}

type UpdateBookDTO struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	IdAuthor *int   `json:"id_author"`
}

type DeleteBookDTO struct {
	Id int `json:"id"`
}

type BorrowBookDTO struct {
	BookID   int `json:"book_id"`
	ReaderID int `json:"reader_id"`
}

type ReturnBookDTO struct {
	BookID int `json:"book_id"`
}
