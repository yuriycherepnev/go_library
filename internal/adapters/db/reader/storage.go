// adapter - реализация репозитория

package reader

import (
	"database/sql"
	"errors"
	"go-library/internal/domain"
	"go-library/internal/domain/reader"
)

type readerStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) reader.Repository {
	return &readerStorage{
		db: db,
	}
}

func (s *readerStorage) GetOne(id int) (*reader.Reader, error) {
	var r reader.Reader
	err := s.db.QueryRow("SELECT id, name FROM reader WHERE id = ?", id).Scan(&r.Id, &r.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *readerStorage) GetAll() ([]reader.Reader, error) {
	rows, err := s.db.Query("SELECT id, name FROM reader ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readers []reader.Reader
	for rows.Next() {
		var r reader.Reader
		if err := rows.Scan(&r.Id, &r.Name); err != nil {
			return nil, err
		}
		readers = append(readers, r)
	}
	return readers, nil
}

func (s *readerStorage) Create(name string) (*reader.Reader, error) {
	result, err := s.db.Exec("INSERT INTO reader(name) VALUES(?)", name)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	r := &reader.Reader{Id: int(id), Name: name}
	return r, nil
}

func (s *readerStorage) Update(id int, name string) error {
	result, err := s.db.Exec("UPDATE reader SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (s *readerStorage) Delete(id int) error {
	var count int
	s.db.QueryRow("SELECT COUNT(*) FROM book WHERE id_reader = ?", id).Scan(&count)
	if count > 0 {
		return domain.ErrReaderHasBooks
	}
	result, err := s.db.Exec("DELETE FROM reader WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
