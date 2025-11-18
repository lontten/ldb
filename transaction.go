package ldb

import (
	"fmt"
	"log"
)

// 回滚结构体
type RollbackWithResult[T any] struct {
	Result T
}

// 辅助函数，让业务代码更清晰
func Rollback[T any](result T) {
	panic(RollbackWithResult[T]{Result: result})
}

func Transaction[T any](db Engine, fn func(tx Engine) (T, error)) (res T, err error) {
	tx, err := db.Begin()
	if err != nil {
		return res, err
	}

	defer func() {
		r := recover()
		if r != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				log.Printf("rollback failed: %v", rbErr)
			}

			// 检查是否是自定义的回滚类型
			switch v := r.(type) {
			case RollbackWithResult[T]:
				// 如果是自定义回滚，返回指定的结果，没有错误
				res = v.Result
				err = nil
			case error:
				err = v
			default:
				err = fmt.Errorf("%v", err)
			}
			return
		}
	}()

	res, err = fn(tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			log.Printf("rollback failed: %v", rbErr)
		}
		return res, err
	}

	if err = tx.Commit(); err != nil {
		return res, fmt.Errorf("commit failed: %v", err)
	}

	return res, nil
}
