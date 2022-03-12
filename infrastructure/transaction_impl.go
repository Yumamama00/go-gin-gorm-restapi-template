package infrastructure

import (
	"context"

	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/logger"
	"github.com/Yumamama00/go-gin-gorm-restapi-sample/utils/transaction"
	"gorm.io/gorm"
)

// TxCtxKey Keyの型を独自定義
type TxCtxKey string

// TransactionをContextから取得する際のKey
const txKey TxCtxKey = "TRANSACTION_CONTEXT_ATTRIBUTE_KEY"

// txImpl Transactionを制御する実装構造体
type txImpl struct{}

// txImplのシングルトンインスタンス
var sharedTx = &txImpl{}

// Transaction txImplのシングルトンインスタンス取得メソッド
func Transaction() transaction.Transaction {
	return sharedTx
}

// Required Transaction制御
func (t *txImpl) Required(ctx context.Context, txFunc func(ctx context.Context) (interface{}, error)) (data interface{}, err error) {

	tx := GetConn().Begin()
	if tx.Error != nil {
		return nil, &RdbRuntimeError{
			ErrMsg:        "[infrastructure.Required] failed to begin Transaction",
			OriginalError: tx.Error,
		}
	}
	logger.Logger.Debug("Transaction Begin")

	ctx = context.WithValue(ctx, txKey, tx)

	defer func() {
		if p := recover(); p != nil {
			logger.Logger.Debug("Transaction Rollback")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			logger.Logger.Debug("Transaction Rollback")
			tx.Rollback()
		} else {
			logger.Logger.Debug("Transaction Commit")
			if commitErr := tx.Commit().Error; commitErr != nil {
				err = &RdbRuntimeError{
					ErrMsg:        "[infrastructure.Required] failed to commit Transaction",
					OriginalError: commitErr,
				}
			}
		}
	}()

	data, err = txFunc(ctx)
	return
}

// GetTx Contextに格納したトランザクションを取得する
func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	return tx, ok
}
