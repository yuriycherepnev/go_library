// entity - model

package book

import (
	"go-library/internal/domain/author"
	"go-library/internal/domain/reader"
)

type Book struct {
	Id       int            `json:"id"`
	Title    *string        `json:"title"`
	IdAuthor *int           `json:"id_author"`
	IdReader *int           `json:"id_reader"`
	Author   *author.Author `json:"author,omitempty"`
	Reader   *reader.Reader `json:"reader,omitempty"`
}
