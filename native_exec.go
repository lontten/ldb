package ldb

func Exec(db Engine, query string, args ...any) (int64, error) {
	db = db.init()
	exec, err := db.exec(query, args...)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}
