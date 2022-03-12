package di

import (
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/infrastructure"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/infrastructure/persistence"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/controller"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/routers"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/usecases"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/logger"

	"go.uber.org/dig"
)

var errors []error
var errCount = 0
var errMsg = ""

// RegisterDIFunction 依存性注入関数の登録
func RegisterDIFunction() *dig.Container {
	c := dig.New()

	registerTransaction(c)
	registerRepository(c)
	registerUsecase(c)
	registerController(c)
	registerRouter(c)

	for _, e := range errors {
		if e != nil {
			errMsg = errMsg + e.Error() + ", "
			errCount = errCount + 1
		}
	}
	if errCount > 0 {
		logger.Logger.Panic("Can not register DI function. error:" + errMsg)
	}

	return c
}

// Transactionの依存性注入関数の登録
func registerTransaction(c *dig.Container) {
	errors = append(errors, c.Provide(infrastructure.Transaction))
}

// Repositoryの依存性注入関数の登録
func registerRepository(c *dig.Container) {
	errors = append(errors, c.Provide(persistence.NewBookRepository))
}

// Usecaseの依存性注入関数の登録
func registerUsecase(c *dig.Container) {
	errors = append(errors, c.Provide(usecases.NewBookUsecase))
}

// Controllerの依存性注入関数の登録
func registerController(c *dig.Container) {
	errors = append(errors, c.Provide(controller.NewBookController))
}

// Routerの依存性注入関数の登録
func registerRouter(c *dig.Container) {
	// Routerを列挙、dig.Name必須
	errors = append(errors, c.Provide(routers.NewBookRouter, dig.Name("book")))

	// RouterService
	errors = append(errors, c.Provide(routers.NewService))
}
