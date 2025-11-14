package ldb

import (
	"fmt"
	"log"
)

func Transaction[T any](db Engine, fn func(tx Engine) T) (res T, err error) {
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

			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
			return
		}
	}()

	res = fn(tx)
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

func TransactionErr[T any](db Engine, fn func(tx Engine) (T, error)) (res T, err error) {
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

			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("panic: %v", r)
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
