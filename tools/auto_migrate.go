package tools

import (
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/domain/model/book"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/infrastructure"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/logger"
	"github.com/joho/godotenv"
)

// RDBのauto migrateを行う
func main() {
	if err := godotenv.Load("../../.dev_env"); err != nil {
		panic("Can not loading .dev_env" + err.Error())
	}

	logger.Init()

	infrastructure.OpenDB()

	db := infrastructure.GetConn()

	if db != nil {
		db.Debug().AutoMigrate(
			&book.Book{},
		)
	}
}
