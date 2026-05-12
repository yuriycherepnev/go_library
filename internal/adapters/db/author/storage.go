package author

import "ca-template/internal/domain/author"

type authorStorage struct {
}

func NewStorage() author.Storage {
	return &authorStorage{}
}

func (as *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}
func (as *authorStorage) GetAll(limit int, offset int) []*author.Author {
	return nil
}
func (as *authorStorage) Create(author *author.Author) *author.Author {
	return author
}
func (as *authorStorage) Delete(author *author.Author) error {
	return nil
}
