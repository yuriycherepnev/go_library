package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

type Client struct {
	db *sql.DB
}

func NewClient(ctx context.Context, Host, Port, User, Password, DBName string) (*sql.DB, error) {
	config := Config{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DBName:   DBName,
	}

	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to ping mysql: %w", err)
	}

	log.Println("MySQL client initialized successfully")

	return db, nil
}

func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}
