package exception

// システム例外
type systemError interface {
	// 内部エラーか否か
	IsInternal() bool
	// エラーコード
	Code() string
}
