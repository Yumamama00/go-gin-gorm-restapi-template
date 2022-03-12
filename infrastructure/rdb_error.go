package infrastructure

// RdbRuntimeError GORMからエラーを返却された場合のラッパー
type RdbRuntimeError struct {
	ErrMsg        string
	OriginalError error
}

// Code エラーコードを返却
func (re *RdbRuntimeError) Code() string {
	return "rdb_runtime_error"
}

// IsInternal 内部エラーか否かを返却
func (re *RdbRuntimeError) IsInternal() bool {
	return true
}

func (re *RdbRuntimeError) Error() string {
	return re.ErrMsg + " , " + re.OriginalError.Error()
}
