package author

// Storage - repository

type Storage interface {
	GetOne(uuid string) *Author
	GetAll(limit int, offset int) []*Author
	Create(author *Author) *Author
	Delete(author *Author) error
}
