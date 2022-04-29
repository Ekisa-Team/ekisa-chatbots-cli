package data

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
)

func getConnection() (*sql.DB, error) {
	conn := viper.GetString("connection_string")
	return sql.Open("mssql", conn)
}
