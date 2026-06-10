package book

// interface бизнес логики лежит там где он используется - рядом с handler
import (
	"context"
	"go-library/internal/domain/book"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *book.Book
	GetAll(ctx context.Context, limit int, offset int) []*book.Book
	Create(ctx context.Context, request CreateBookDTO) *book.Book
}
