module example

go 1.25.0

replace github.com/lontten/ldb/v2 => ../

require github.com/lontten/ldb/v2 v2.23.15

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/lontten/lcore/v2 v2.22.0
	github.com/lontten/lutil v0.1.8
	gorm.io/gorm v1.31.1
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/exp v0.0.0-20260218203240-3dfff04db8fa // indirect
	golang.org/x/text v0.34.0 // indirect
	gorm.io/plugin/soft_delete v1.2.1 // indirect
)
