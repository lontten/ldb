package ldb

import "github.com/pkg/errors"

var (
	ErrNil          = errors.New("nil")
	ErrContainEmpty = errors.New("slice empty")
	ErrNoPkOrUnique = errors.New(" ERROR: there is no unique or exclusion constraint matching the ON CONFLICT specification (SQLSTATE 42P10) ")
	ErrNoPk         = errors.New("no primary key")
	ErrTypePkArgs   = errors.New("type of args is err")
	ErrNeedMultiPk  = errors.New("need multi primary key")
)
