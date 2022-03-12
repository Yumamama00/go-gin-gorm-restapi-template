package infrastructure

import "net/http"

// OptimisticLockError 楽観的排他制御のエラー（更新対象が見つからなかった場合にも用いる）
type OptimisticLockError struct {
	Msg    string
	ErrMsg string
}

// Code エラーコード
func (ole *OptimisticLockError) Code() string {
	return "optimistic_lock_error"
}

// Message エラーメッセージ
func (ole *OptimisticLockError) Message() []string {
	return []string{ole.Msg}
}

// HTTPStatus ステータスコード
func (ole *OptimisticLockError) HTTPStatus() int {
	return http.StatusConflict
}

func (ole *OptimisticLockError) Error() string {
	return ole.ErrMsg
}
