package main

import (
	"os"
	"time"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/infrastructure"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/middleware"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/interfaces/api/routers"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/di"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	mode := os.Getenv("MODE")
	apiPort := os.Getenv("API_PORT")

	// DotEnv初期化プロセス
	dotEnvInit(mode)

	// Logger初期化プロセス
	logger.Init()
	logger.Logger.Debug("Logger init succeeded")

	// Gin初期化プロセス
	r := ginInit(mode)
	logger.Logger.Debug("Gin init succeeded")

	// 依存性注入関数の登録
	c := di.RegisterDIFunction()
	logger.Logger.Debug("Register all DI function succeeded")

	// DB接続
	infrastructure.OpenDB()
	logger.Logger.Debug("Connect DB succeeded")

	// 全てのRouter初期化、起動
	if err := c.Invoke(func(s *routers.Service) {
		baseRoute := r.Group("/api")

		// 存在しないURLは404を返す
		r.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})

		s.RouterInit(baseRoute)
		logger.Logger.Debug("All router init succeeded")
	}); err != nil {
		logger.Logger.Panic("Failed to resolve dependency: " + err.Error())
	}

	// StartUp Server
	logger.Logger.Info("Listening and serving " + "HTTP on :" + apiPort + " , " + "MODE:" + mode)
	r.Run(":" + os.Getenv("API_PORT"))
}

func dotEnvInit(mode string) {
	if mode == "PROD" {
		if err := godotenv.Load(".prod_env"); err != nil {
			panic("Can not loading .prod_env. error:" + err.Error())
		}
	} else {
		if err := godotenv.Load(".dev_env"); err != nil {
			panic("Can not loading .dev_env. error:" + err.Error())
		}
	}
}

func ginInit(mode string) *gin.Engine {
	if mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	// gin-zap middleware
	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	// logging all panic to error log
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))
	// CORS middleware
	r.Use(middleware.SetCors())

	return r
}
