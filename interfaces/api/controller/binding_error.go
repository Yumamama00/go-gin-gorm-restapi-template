package controller

import "net/http"

// GinのBind機能でBindingに失敗した際のエラー
type bindingError struct {
	errMsg        string
	originalError error
}

// Code エラーコード
func (be *bindingError) Code() string {
	return "binding_error"
}

// Message エラーメッセージ
func (be *bindingError) Message() []string {
	return []string{"不正な入力値が含まれています。入力内容を再度確認してください。"}
}

// HTTPStatus ステータスコード
func (be *bindingError) HTTPStatus() int {
	return http.StatusBadRequest
}

func (be *bindingError) Error() string {
	return be.errMsg + " , " + be.originalError.Error()
}
