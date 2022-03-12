package transaction

import "context"

// Transaction Transactionインタフェース
type Transaction interface {
	Required(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
