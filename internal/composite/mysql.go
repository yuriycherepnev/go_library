package composite

import (
	"context"
	"database/sql"
	"go-library/pkg/client/mysql"
)

type MysqlComposite struct {
	db *sql.DB
}

func NewMysqlComposite(ctx context.Context, Host, Port, User, Password, DBName string) (*MysqlComposite, error) {
	client, err := mysql.NewClient(ctx, Host, Port, User, Password, DBName)

	if err != nil {
		return nil, err
	}

	return &MysqlComposite{
		db: client,
	}, nil
}
