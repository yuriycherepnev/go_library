package domain

import "errors"

var (
	ErrAuthorHasBooks = errors.New("author has books")
	ErrAuthorNotFound = errors.New("author not found")
	ErrBookNotFound   = errors.New("book not found")
	ErrBookTaken      = errors.New("book already taken")
)
