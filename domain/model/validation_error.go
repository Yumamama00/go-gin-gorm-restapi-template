package model

import (
	"net/http"
)

// ValidationError ドメイン層でのバリデーションや重複チェックに引っかかった場合の例外
type ValidationError struct {
	MsgList []string
	ErrMsg  string
}

// Code エラーコード
func (ve *ValidationError) Code() string {
	return "validation_error"
}

// Message エラーメッセージ
func (ve *ValidationError) Message() []string {
	return ve.MsgList
}

// HTTPStatus ステータスコード
func (ve *ValidationError) HTTPStatus() int {
	return http.StatusBadRequest
}

func (ve *ValidationError) Error() string {
	return ve.ErrMsg
}
