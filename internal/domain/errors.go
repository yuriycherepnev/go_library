package domain

import "errors"

var (
	ErrAuthorHasBooks  = errors.New("author has books")
	ErrAuthorNotFound  = errors.New("author not found")
	ErrBookNotFound    = errors.New("book not found")
	ErrReaderHasBooks  = errors.New("reader has books")
	ErrReaderNotFound  = errors.New("reader not found")
	ErrBookTaken       = errors.New("book already taken")
	ErrBookNotBorrowed = errors.New("book not borrowed")
)
