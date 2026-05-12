package book

// interface бизнес логики лежит там где он используется - рядом с handler
import (
	"context"
	"go-library/internal/entity"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *entity.Book
	GetAll(ctx context.Context, limit int, offset int) []*entity.Book
	Create(ctx context.Context, request CreateBookDTO) *entity.Book
}
