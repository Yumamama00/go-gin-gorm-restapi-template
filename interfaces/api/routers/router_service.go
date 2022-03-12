package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Router 全Routerの基底インターフェース
type Router interface {
	routing(e *gin.RouterGroup)
}

// Service 全Routerを保持・管理するService
type Service struct {
	routers []Router
}

// AllRouter Dig.Inで全てのRouterの依存性を保持
type AllRouter struct {
	dig.In
	BookRouter Router `name:"book"`
}

// NewService コンストラクタ
func NewService(allRouter AllRouter) *Service {
	routers := []Router{allRouter.BookRouter}

	return &Service{
		routers: routers,
	}
}

// RouterInit 全Routerの初期化プロセス
func (rs *Service) RouterInit(e *gin.RouterGroup) {

	for _, r := range rs.routers {
		r.routing(e)
	}

}
