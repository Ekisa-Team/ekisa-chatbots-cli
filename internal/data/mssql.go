package data

import (
	"database/sql"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func getConnection() (*sql.DB, error) {
	conn := os.Getenv("CONN_STRING")
	return sql.Open("mssql", conn)
}
