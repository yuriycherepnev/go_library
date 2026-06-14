// adapter - реализация репозитория

package book

import (
	"database/sql"
	"errors"
	"go-library/internal/domain/author"
	"go-library/internal/domain/book"
	"go-library/internal/domain/reader"
)

type bookStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) book.Repository {
	return &bookStorage{
		db: db,
	}
}

func (bs *bookStorage) GetOne(id int) (*book.Book, error) {
	query := `
        SELECT b.id, b.title, b.id_author, b.id_reader,
               a.id, a.name,
               r.id, r.name
        FROM book b
        LEFT JOIN author a ON b.id_author = a.id
        LEFT JOIN reader r ON b.id_reader = r.id
        WHERE b.id = ?
    `
	var b book.Book
	var aID, _ sql.NullInt64
	var aNameStr sql.NullString
	var rID, _ sql.NullInt64
	var rNameStr sql.NullString

	err := bs.db.QueryRow(query, id).Scan(
		&b.Id, &b.Title, &b.IdAuthor, &b.IdReader,
		&aID, &aNameStr,
		&rID, &rNameStr,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if aID.Valid {
		b.Author = &author.Author{Id: int(aID.Int64), Name: aNameStr.String}
	}
	if rID.Valid {
		b.Reader = &reader.Reader{Id: int(rID.Int64), Name: rNameStr.String}
	}
	return &b, nil
}

func (bs *bookStorage) GetAll() ([]book.Book, error) {
	query := `
        SELECT b.id, b.title, b.id_author, b.id_reader,
               a.id, a.name,
               r.id, r.name
        FROM book b
        LEFT JOIN author a ON b.id_author = a.id
        LEFT JOIN reader r ON b.id_reader = r.id
        ORDER BY b.id
    `
	rows, err := bs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []book.Book
	for rows.Next() {
		var b book.Book
		var aID, _ sql.NullInt64
		var aNameStr sql.NullString
		var rID, _ sql.NullInt64
		var rNameStr sql.NullString

		err = rows.Scan(
			&b.Id, &b.Title, &b.IdAuthor, &b.IdReader,
			&aID, &aNameStr,
			&rID, &rNameStr,
		)
		if err != nil {
			return nil, err
		}

		if aID.Valid {
			b.Author = &author.Author{Id: int(aID.Int64), Name: aNameStr.String}
		}

		if rID.Valid {
			b.Reader = &reader.Reader{Id: int(rID.Int64), Name: rNameStr.String}
		}

		books = append(books, b)
	}
	return books, nil
}

func (bs *bookStorage) Create(title string, idAuthor int) (*book.Book, error) {
	query := `
		INSERT INTO book (title, id_author)
		VALUES (?, ?)
	`

	result, err := bs.db.Exec(query, title, idAuthor)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return bs.GetOne(int(id))
}

func (bs *bookStorage) Update(id int, title string, idAuthor *int) error {
	var err error

	if idAuthor == nil {
		query := `
			UPDATE book
			SET title = ?
			WHERE id = ?
		`
		_, err = bs.db.Exec(query, title, id)
	} else {
		query := `
			UPDATE book
			SET title = ?, id_author = ?
			WHERE id = ?
		`
		_, err = bs.db.Exec(query, title, *idAuthor, id)
	}

	if err != nil {
		return err
	}

	return nil
}

func (bs *bookStorage) Delete(id int) error {
	query := `
		DELETE FROM book
		WHERE id = ?
	`

	_, err := bs.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
