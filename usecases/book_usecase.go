package usecases

import (
	"context"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model/book"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/services"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/idutil"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/transaction"
)

// IBookUsecase Bookを取り扱うUsecaseのインターフェース
type IBookUsecase interface {
	AddBook(context.Context, *book.PublicBook) (string, error)
	FindBook(context.Context, string) (*book.PublicBook, error)
	GetAllBook(context.Context) ([]*book.PublicBook, error)
	UpdateBook(context.Context, *book.PublicBook, string) error
	DeleteBook(context.Context, string) error
}

// BookUsecaseImpl Bookを取り扱うUsecaseの実装
type BookUsecaseImpl struct {
	br          book.IBookRepository
	transaction transaction.Transaction
	bs          services.IBookService
}

// NewBookUsecase コンストラクタ
func NewBookUsecase(br book.IBookRepository, transaction transaction.Transaction) IBookUsecase {
	return &BookUsecaseImpl{br: br, transaction: transaction, bs: services.NewBookService(br)}
}

// AddBook Book追加処理
func (bu *BookUsecaseImpl) AddBook(ctx context.Context, bookDto *book.PublicBook) (string, error) {

	id, err := bu.transaction.Required(ctx, func(ctx context.Context) (interface{}, error) {
		newID := idutil.XID()
		newBook, err := book.NewBook(bookDto)
		if err != nil {
			return newID, err
		}

		if err = bu.bs.IsDuplicatedTitleWhenInsert(ctx, newBook.Title); err != nil {
			return newID, err
		}

		newBook.BookID = newID
		return newID, bu.br.AddBook(ctx, newBook)
	})

	return id.(string), err
}

// FindBook Book1件取得処理
func (bu *BookUsecaseImpl) FindBook(ctx context.Context, ID string) (*book.PublicBook, error) {
	ebook, err := bu.br.FindBook(ctx, ID)

	if err != nil {
		return nil, err
	}

	return &book.PublicBook{
		BookID:    ebook.BookID,
		Title:     ebook.Title,
		Content:   ebook.Content,
		UpdatedAt: ebook.UpdatedAt,
	}, nil
}

// GetAllBook Book全件取得処理
func (bu *BookUsecaseImpl) GetAllBook(ctx context.Context) ([]*book.PublicBook, error) {

	books, err := bu.br.GetAllBook(ctx)
	if err != nil {
		return nil, err
	}

	var bookDtoList []*book.PublicBook
	for _, b := range books {
		bookDtoList = append(bookDtoList, &book.PublicBook{
			BookID:    b.BookID,
			Title:     b.Title,
			Content:   b.Content,
			UpdatedAt: b.UpdatedAt,
		})
	}

	return bookDtoList, err
}

// UpdateBook Book更新処理
func (bu *BookUsecaseImpl) UpdateBook(ctx context.Context, bookDto *book.PublicBook, ID string) error {

	_, err := bu.transaction.Required(ctx, func(ctx context.Context) (interface{}, error) {
		bookDto.BookID = ID
		book, err := book.NewBook(bookDto)
		if err != nil {
			return nil, err
		}

		if err = bu.bs.IsDuplicatedTitleWhenUpdate(ctx, book.Title, ID); err != nil {
			return nil, err
		}

		return nil, bu.br.UpdateBook(ctx, book)
	})

	return err
}

// DeleteBook Book削除処理
func (bu *BookUsecaseImpl) DeleteBook(ctx context.Context, ID string) error {
	_, err := bu.transaction.Required(ctx, func(ctx context.Context) (interface{}, error) {
		return nil, bu.br.DeleteBook(ctx, ID)
	})

	return err
}
