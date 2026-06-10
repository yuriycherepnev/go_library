package author

// interface бизнес логики лежит там где он используется - рядом с handler

import (
	"context"
	"go-library/internal/domain/author"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *author.Author
	GetAll(ctx context.Context, limit int, offset int) []*author.Author
	Create(ctx context.Context, request CreateAuthorDTO) *author.Author
}
