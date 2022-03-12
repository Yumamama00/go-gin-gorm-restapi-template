package controller

import (
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model/book"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/httputils"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/httputils/exception"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/usecases"
	"github.com/gin-gonic/gin"
)

// BookController Bookを取り扱うController
type BookController struct {
	bu usecases.IBookUsecase
}

// NewBookController コンストラクタ
func NewBookController(bu usecases.IBookUsecase) *BookController {
	return &BookController{bu: bu}
}

// AddBook Book追加リクエストのハンドリング
func (bc *BookController) AddBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		var bookDto book.PublicBook
		if err := c.ShouldBind(&bookDto); err != nil {
			bindingError := &bindingError{
				errMsg:        "[interfaces.api.controller.AddBook] failed to bind Book from request.",
				originalError: err,
			}
			exception.Handle(bindingError, c)
			return
		}

		id, err := bc.bu.AddBook(c, &bookDto)
		if err != nil {
			exception.Handle(err, c)
			return
		}

		httputils.CreatedHandle(id, c)
	}

}

// FindBook Book1件取得リクエストのハンドリング
func (bc *BookController) FindBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		ID := c.Param("id")
		book, err := bc.bu.FindBook(c, ID)

		if err != nil {
			exception.Handle(err, c)
			return
		}

		httputils.SuccessHandle(*book, c)
	}

}

// GetAllBook Book全件取得リクエストのハンドリング
func (bc *BookController) GetAllBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		books, err := bc.bu.GetAllBook(c)

		if err != nil {
			exception.Handle(err, c)
			return
		}

		httputils.SuccessHandle(books, c)
	}

}

// UpdateBook Book更新リクエストのハンドリング
func (bc *BookController) UpdateBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		ID := c.Param("id")
		var bookDto book.PublicBook
		if err := c.ShouldBind(&bookDto); err != nil {
			bindingError := &bindingError{
				errMsg:        "[controller.UpdateBook] failed to bind Book from request",
				originalError: err,
			}
			exception.Handle(bindingError, c)
			return
		}

		if err := bc.bu.UpdateBook(c, &bookDto, ID); err != nil {
			exception.Handle(err, c)
			return
		}

		httputils.SuccessHandle(nil, c)
	}

}

// DeleteBook Book削除リクエストのハンドリング
func (bc *BookController) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		ID := c.Param("id")
		if err := bc.bu.DeleteBook(c, ID); err != nil {
			exception.Handle(err, c)
			return
		}

		httputils.NoContentHandle(c)
	}
}
