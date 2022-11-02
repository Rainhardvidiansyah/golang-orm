package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	_, err := gorm.Open(mysql.Open(dsn("")), &gorm.Config{})
	if err != nil {
		log.Fatal("There's an error while you are connecting to database", err.Error())
	}

	// repository := repository.NewUserRepository(db)
	// id := repository.FindUserByID(2)
	// fmt.Printf("User Id with %v \n", id)

}

func dsn(des string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There's an error in your env file")
	}
	username := os.Getenv("username")
	password := os.Getenv("password")
	hostname := os.Getenv("hostname")
	dbname := os.Getenv("dbname")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
