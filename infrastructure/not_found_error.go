package infrastructure

import "net/http"

// DefaultNotFoundMsg デフォルトメッセージ
const DefaultNotFoundMsg string = "検索結果が見つかりませんでした。検索条件を変更して再度お試しください。"

// NotFoundError 検索結果が0件の場合のエラー
type NotFoundError struct {
	Msg           string
	ErrMsg        string
	OriginalError error
}

// Code エラーコード
func (nfe *NotFoundError) Code() string {
	return "not_found_error"
}

// Message エラーメッセージ
func (nfe *NotFoundError) Message() []string {
	return []string{nfe.Msg}
}

// HTTPStatus ステータスコード
func (nfe *NotFoundError) HTTPStatus() int {
	return http.StatusNotFound
}

func (nfe *NotFoundError) Error() string {
	if nfe.OriginalError != nil {
		return nfe.ErrMsg + " , " + nfe.OriginalError.Error()
	}
	return nfe.ErrMsg
}
