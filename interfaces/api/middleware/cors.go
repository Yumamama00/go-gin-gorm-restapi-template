package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetCors CORS設定
func SetCors() gin.HandlerFunc {
	f := cors.New(cors.Config{
		// アクセスを許可するOrigin
		// AllowOrigins: []String{},
		AllowAllOrigins: true,

		// 許可するHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"X-Requested-With",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},

		// Cookie等を必要とするかどうか
		AllowCredentials: true,

		// プリフライトリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	})

	return f
}
