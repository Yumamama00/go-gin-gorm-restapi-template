package book

import "context"

// IBookRepository Bookを取り扱うRepositoryのインターフェース
type IBookRepository interface {
	AddBook(context.Context, *Book) error
	FindBook(context.Context, string) (*Book, error)
	GetAllBook(context.Context) ([]*Book, error)
	UpdateBook(context.Context, *Book) error
	DeleteBook(context.Context, string) error
	CheckDuplicateTitle(context.Context, string, string) (bool, error)
}
