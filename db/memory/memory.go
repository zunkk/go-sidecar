package memory

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/marcboeker/go-duckdb/v2"
)

func OpenSQLDB() (*sqlx.DB, error) {
	return sqlx.Connect("duckdb", "")
}
