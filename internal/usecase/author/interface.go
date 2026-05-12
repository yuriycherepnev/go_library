package author

// interface бизнес логики лежит там где он используется - рядом с handler

import (
	"context"
	"go-library/internal/entity"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *entity.Author
	GetAll(ctx context.Context, limit int, offset int) []*entity.Author
	Create(ctx context.Context, request CreateAuthorDTO) *entity.Author
}
