package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// func SetupModels() *gorm.DB {
// 	env := godotenv.Load(".env")
// 	user_name := goDotEnvVariable("POSTGRES_USER")
// 	password := goDotEnvVariable("POSTGRES_PASSWORD")
// 	db := goDotEnvVariable("POSTGRES_DB")
// 	host := goDotEnvVariable("POSTGRES_HOST")
// 	port := goDotEnvVariable("POSTGRES_PORT")

// 	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user_name, db, password)
// 	fmt.Println("conname is\t\t", prosgret_conname)
// 	db, err := gorm.Open("postgres", prosgret_conname)
// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}
// 	db.AutoMigrate() // Initialise value

// 	return db
// }

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
