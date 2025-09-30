package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	SQL *sql.DB
}

var dbconn = &DB{}

func LoadConnectionForTraining_Db() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		server   = os.Getenv("SERVERTRAINING")
		user     = os.Getenv("USERTRAINING")
		password = os.Getenv("PASSWORDTRAINING")
		database = os.Getenv("DATABASETRAINING")
	)

	return connectSql(server, database, user, password)
}

func connectSql(server, database, user, password string) (*DB, error) {
	connectionString := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", server, database, user, password)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	dbconn.SQL = db
	return dbconn, nil
}
