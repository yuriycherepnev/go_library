package book

// interface бизнес логики лежит там где он используется - рядом с handler
import (
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *Book
	GetAll(ctx context.Context, limit int, offset int) []*Book
	Create(ctx context.Context, request CreateBookDTO) *Book
}
