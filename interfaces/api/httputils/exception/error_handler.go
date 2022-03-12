package exception

import (
	"fmt"
	"net/http"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Handle エラーのハンドリング
func Handle(err error, c *gin.Context) {
	// アプリケーション例外のチェック
	ae, ok := errors.Cause(err).(applicationError)
	if ok {
		handleAppError(ae, c, err)
		return
	}

	// システム例外のチェック
	se, ok := errors.Cause(err).(systemError)
	if ok {
		handleSysError(se, c, err)
		return
	}

}

// アプリケーション例外のハンドリング
func handleAppError(ae applicationError, c *gin.Context, err error) {
	logger.Logger.Info(fmt.Sprintf("application error occured: %s", err.Error()), zap.String("errorCode", ae.Code()))

	c.JSON(ae.HTTPStatus(),
		gin.H{"code": ae.Code(), "message": ae.Message(), "status": ae.HTTPStatus()})
	return
}

// システム例外のハンドリング
const systemErrMsg string = "予期しないシステムエラーが発生しました。システム管理者に連絡してください"

func handleSysError(se systemError, c *gin.Context, err error) {

	if se.IsInternal() {
		logger.Logger.Warn(fmt.Sprintf("system error occurred: %s", err.Error()), zap.String("errorCode", se.Code()))
	} else {
		logger.Logger.Warn(fmt.Sprintf("unexpected error occurred: %s", err.Error()), zap.String("errorCode", se.Code()))
	}

	c.JSON(http.StatusInternalServerError,
		gin.H{"code": se.Code(), "message": systemErrMsg, "status": http.StatusInternalServerError})
	return
}
