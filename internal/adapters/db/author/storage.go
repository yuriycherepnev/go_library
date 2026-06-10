// реализация репозитория

package author

import (
	"database/sql"
	"errors"
	"go-library/internal/domain"
	"go-library/internal/domain/author"
)

type authorStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) author.Repository {
	return &authorStorage{
		db: db,
	}
}

func (as *authorStorage) GetOne(id int) (*author.Author, error) {
	var a author.Author
	err := as.db.QueryRow("SELECT id, name FROM author WHERE id = ?", id).Scan(&a)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (as *authorStorage) GetAll() (authors []author.Author, err error) {
	rows, err := as.db.Query("SELECT id, name FROM author ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()
	for rows.Next() {
		var a author.Author
		if err = rows.Scan(&a); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}

func (as *authorStorage) Update(id int, name string) error {
	result, err := as.db.Exec("UPDATE author SET name = ? WHERE id = ?", name, id)

	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (as *authorStorage) Create(name string) (*author.Author, error) {
	result, err := as.db.Exec("INSERT INTO author(name) VALUES(?)", name)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &author.Author{Id: int(id), Name: name}, nil
}

func (as *authorStorage) Delete(id int) error {
	var count int
	as.db.QueryRow("SELECT COUNT(*) FROM book WHERE id_author = ?", id).Scan(&count)
	if count > 0 {
		return domain.ErrAuthorHasBooks
	}
	result, err := as.db.Exec("DELETE FROM author WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
