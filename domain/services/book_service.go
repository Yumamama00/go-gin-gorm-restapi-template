package services

import (
	"context"
	"fmt"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model/book"
)

// IBookService BookドメインサービスのIF
type IBookService interface {
	IsDuplicatedTitleWhenInsert(ctx context.Context, title string) error
	IsDuplicatedTitleWhenUpdate(ctx context.Context, title string, ID string) error
}

// Service Bookドメインサービス
type Service struct {
	br book.IBookRepository
}

// NewBookService コンストラクタ
func NewBookService(br book.IBookRepository) IBookService {
	return &Service{br: br}
}

const bookDuplicateMsg = "既に登録されている本のタイトルです。タイトルを変更してください。"

// IsDuplicatedTitleWhenInsert 保存時のタイトル重複チェック
func (bs *Service) IsDuplicatedTitleWhenInsert(ctx context.Context, title string) error {
	ok, err := bs.br.CheckDuplicateTitle(ctx, title, "")
	if err != nil {
		return err
	} else if !ok {
		return &model.ValidationError{
			MsgList: []string{bookDuplicateMsg},
			ErrMsg:  fmt.Sprintf("[domain.services.IsDuplicatedTitleWhenInsert] failed to create new Book. Title duplicated. Title : %s", title),
		}
	}

	return nil
}

// IsDuplicatedTitleWhenUpdate 更新時のタイトル重複チェック
func (bs *Service) IsDuplicatedTitleWhenUpdate(ctx context.Context, title string, ID string) error {
	ok, err := bs.br.CheckDuplicateTitle(ctx, title, ID)
	if err != nil {
		return err
	} else if !ok {
		return &model.ValidationError{
			MsgList: []string{bookDuplicateMsg},
			ErrMsg:  fmt.Sprintf("[domain.services.IsDuplicatedTitleWhenUpdate] failed to update new Book. Title duplicated. Title : %s", title),
		}
	}

	return nil
}
