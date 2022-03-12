package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessHandle Query、Update処理のResponseハンドル (200を返却)
func SuccessHandle(data interface{}, c *gin.Context) {
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
	return
}

// CreatedHandle Create処理のResponseハンドル (201を返却)
func CreatedHandle(ID string, c *gin.Context) {
	c.Header("Location", "https://"+c.Request.Host+c.FullPath()+"/"+ID)
	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
	return
}

// NoContentHandle Delete処理のResponseハンドル (204を返却)
func NoContentHandle(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
	return
}
