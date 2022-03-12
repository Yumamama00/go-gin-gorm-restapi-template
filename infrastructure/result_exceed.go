package infrastructure

import "net/http"

// DefaultResultExceedMsg 検索結果が多すぎる場合のデフォルトメッセージ
const DefaultResultExceedMsg string = "検索結果が多すぎます。検索条件を変更して再度検索してください。"

// ResultExceedError 検索結果が多すぎる場合の場合のエラー
type ResultExceedError struct {
	Msg    string
	ErrMsg string
}

// Code エラーコード
func (ree *ResultExceedError) Code() string {
	return "result_exceed_error"
}

// Message エラーメッセージ
func (ree *ResultExceedError) Message() []string {
	return []string{ree.Msg}
}

// HTTPStatus ステータスコード
func (ree *ResultExceedError) HTTPStatus() int {
	return http.StatusBadRequest
}

func (ree *ResultExceedError) Error() string {
	return ree.ErrMsg
}
