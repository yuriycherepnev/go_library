package author

// interface бизнес логики лежит там где он используется - рядом с handler

import (
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *Author
	GetAll(ctx context.Context, limit int, offset int) []*Author
	Create(ctx context.Context, request CreateAuthorDTO) *Author
}
