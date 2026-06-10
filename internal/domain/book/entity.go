// entity - model

package book

import (
	"fmt"
	"go-library/internal/domain/author"
)

type Book struct {
	UUID   string        `json:"uuid,omitempty"`
	Name   string        `json:"name,omitempty"`
	Year   int           `json:"year,omitempty"`
	Busy   bool          `json:"busy,omitempty"`
	Owner  string        `json:"owner,omitempty"`
	Author author.Author `json:"author,omitempty"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
