package Initialize

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	model "github.com/marlon/golangsimplecrud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitializeDBConnection() {
	var err error
	// dbCreds := os.Getenv("DB_URL")
	// db, err := gorm.Open(postgres.Open(dbCreds), &gorm.Config{})3
	db, err := gorm.Open(postgres.Open("postgres://postgres:secret@10.43.196.47:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(model.Payment{})
	DB = db

	fmt.Printf("\nSuccessfully connected to database!\n")
}
