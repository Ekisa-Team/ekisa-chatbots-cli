package data

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
)

func getConnection() (*sql.DB, error) {
	conn := viper.GetString("connection_string")
	fmt.Printf("Connection: %v\n", conn)
	return sql.Open("mssql", conn)
}
