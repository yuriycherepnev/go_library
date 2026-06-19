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
	BookId   int `json:"book_id"`
	ReaderId int `json:"reader_id"`
}

type ReturnBookDTO struct {
	BookId int `json:"book_id"`
}
