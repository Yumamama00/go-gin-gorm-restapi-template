package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model/book"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/infrastructure"
	"gorm.io/gorm"
)

// BookRepositoryImpl Bookを取り扱うRepositoryの実装
type BookRepositoryImpl struct {
}

// NewBookRepository コンストラクタ
func NewBookRepository() book.IBookRepository {
	return &BookRepositoryImpl{}
}

// AddBook Bookの保存処理
func (br *BookRepositoryImpl) AddBook(ctx context.Context, newBook *book.Book) error {
	db, _ := infrastructure.GetTx(ctx)

	if result := db.Create(newBook); result.Error != nil {
		return &infrastructure.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.AddBook] failed to insert Book to RDB. ID : %s", newBook.BookID),
			OriginalError: result.Error,
		}
	}

	return nil
}

// FindBook Book1件取得処理
func (br *BookRepositoryImpl) FindBook(ctx context.Context, ID string) (*book.Book, error) {
	db := infrastructure.GetConn()

	book := &book.Book{}
	if err := db.First(book, "book_id = ?", ID).Error; err != nil {
		return nil, &infrastructure.NotFoundError{
			Msg:           infrastructure.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.FindBook] failed to find Book from rdb. ID : %s", ID),
			OriginalError: err,
		}
	}

	return book, nil
}

// GetAllBook Book全件取得処理
func (br *BookRepositoryImpl) GetAllBook(ctx context.Context) ([]*book.Book, error) {
	db := infrastructure.GetConn()

	var books []*book.Book
	result := db.Order("title").Limit(51).Find(&books)

	if result.RowsAffected == 51 {
		return nil, &infrastructure.ResultExceedError{
			Msg:    infrastructure.DefaultResultExceedMsg,
			ErrMsg: "[infrastructure.persistence.GetAllBook] failed to get Books from rdb. Too many results.",
		}
	}

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, &infrastructure.NotFoundError{
			Msg:           infrastructure.DefaultNotFoundMsg,
			ErrMsg:        "[infrastructure.persistence.GetAllBook] failed to get Books from rdb",
			OriginalError: result.Error,
		}
	}

	return books, nil
}

// UpdateBook Book更新処理
func (br *BookRepositoryImpl) UpdateBook(ctx context.Context, book *book.Book) error {
	db, _ := infrastructure.GetTx(ctx)
	updatedAt := book.UpdatedAt

	result := db.Where("updated_at <= ?", updatedAt).Model(&book).Updates(book)

	if result.Error != nil {
		return &infrastructure.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.UpdateBook] failed to update Book in RDB. ID : %s", book.BookID),
			OriginalError: result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &infrastructure.OptimisticLockError{
			Msg:    "更新に失敗しました。その本は既に削除されているか、他のユーザによって更新されている可能性があります。",
			ErrMsg: fmt.Sprintf("[infrastructure.persistence.UpdateBook] failed to update Book in RDB. May be optimistic lock. ID : %s", book.BookID),
		}
	}

	return nil
}

// DeleteBook Book削除処理
func (br *BookRepositoryImpl) DeleteBook(ctx context.Context, ID string) error {
	db, _ := infrastructure.GetTx(ctx)

	result := db.Where("book_id = ?", ID).Delete(&book.Book{})

	if result.Error != nil {
		return &infrastructure.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.DeleteBook] failed to delete Book in RDB. ID : %s", ID),
			OriginalError: result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &infrastructure.NotFoundError{
			Msg:           "削除対象の本が見つかりませんでした。その本は既に削除されている可能性があります。",
			ErrMsg:        fmt.Sprintf("[infrastructure.persistence.DeleteBook] failed to delete Book in RDB. Record not found. ID : %s", ID),
			OriginalError: nil,
		}
	}

	return nil
}

// CheckDuplicateTitle タイトルの重複チェック
func (br *BookRepositoryImpl) CheckDuplicateTitle(ctx context.Context, title string, ID string) (bool, error) {
	db, _ := infrastructure.GetTx(ctx)

	book := &book.Book{}
	var err error
	if ID == "" {
		err = db.First(book, "title = ?", title).Error
	} else {
		err = db.First(book, "title = ? AND id <> ?", title, ID).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
