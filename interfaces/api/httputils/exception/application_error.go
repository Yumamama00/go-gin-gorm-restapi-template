package exception

// アプリケーション例外
type applicationError interface {
	// エラーコード
	Code() string
	// エラーメッセージ
	Message() []string
	// HTTPステータス
	HTTPStatus() int
}
