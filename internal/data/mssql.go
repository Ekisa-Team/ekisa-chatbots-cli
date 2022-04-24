package data

import (
	"database/sql"
	"os"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/config"
	_ "github.com/denisenkom/go-mssqldb"
)

func getConnection() (*sql.DB, error) {
	conn := os.Getenv(config.ENV_CONN_STRING)
	return sql.Open("mssql", conn)
}
