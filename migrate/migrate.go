package migrate

import (
	"backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Migrate() {
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
	connectionString := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", server, database, user, password)

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
    		TablePrefix: "anandhu2.", 
    	},
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&models.Employee{})
}