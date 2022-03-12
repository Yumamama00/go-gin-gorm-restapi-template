package routers

import (
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/controller"
	"github.com/gin-gonic/gin"
)

// BookRouter Bookに関するRoutingを取り扱う
type BookRouter struct {
	bc *controller.BookController
}

// NewBookRouter コンストラクタ
func NewBookRouter(bc *controller.BookController) Router {
	return &BookRouter{bc: bc}
}

// ルーティング
func (b *BookRouter) routing(e *gin.RouterGroup) {
	be := e.Group("/books")
	{
		be.POST("", b.bc.AddBook())
		be.GET("/:id", b.bc.FindBook())
		be.GET("", b.bc.GetAllBook())
		be.PUT("/:id", b.bc.UpdateBook())
		be.DELETE("/:id", b.bc.DeleteBook())
	}
}
