package db

import (
	"fmt"
	"gin_docker/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	// env := godotenv.Load(".env")
	user_name := goDotEnvVariable("POSTGRES_USER")
	password := goDotEnvVariable("POSTGRES_PASSWORD")
	db_name := goDotEnvVariable("POSTGRES_DB")
	host := goDotEnvVariable("POSTGRES_HOST")
	port := goDotEnvVariable("POSTGRES_PORT")

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user_name, db_name, password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open(postgres.Open(prosgret_conname), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.Books{}) // Initialise value

	return db
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
