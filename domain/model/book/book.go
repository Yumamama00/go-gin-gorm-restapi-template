package book

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model"
)

// Book Entity
type Book struct {
	BookID    string `gorm:"primary_key"`
	Title     string `gorm:"size:100;not null;"`
	Content   string `gorm:"size:255;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewBook コンストラクタ
func NewBook(bookDto *PublicBook) (*Book, error) {
	var validateErrMsg []string
	if utf8.RuneCountInString(bookDto.Title) == 0 || utf8.RuneCountInString(bookDto.Title) > 20 {
		validateErrMsg = append(validateErrMsg, "本のタイトルは1〜20文字の間で入力してください")
	}

	if utf8.RuneCountInString(bookDto.Content) == 0 || utf8.RuneCountInString(bookDto.Content) > 255 {
		validateErrMsg = append(validateErrMsg, "本の内容は1〜255文字の間で入力してください")
	}

	if len(validateErrMsg) != 0 {
		return nil, &model.ValidationError{
			MsgList: validateErrMsg,
			ErrMsg:  fmt.Sprintf("[domain.model.book.NewBook] failed to create new Book. validate error. ID : %s", bookDto.BookID),
		}
	}

	return &Book{
		BookID:    bookDto.BookID,
		Title:     bookDto.Title,
		Content:   bookDto.Content,
		UpdatedAt: bookDto.UpdatedAt,
	}, nil
}

// PublicBook 公開用のDTO
type PublicBook struct {
	BookID    string    `json:"bookID"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updatedAt"`
}
