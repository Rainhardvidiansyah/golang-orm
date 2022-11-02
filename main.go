package main

import (
	"fmt"
	"go-orm/repository"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open(dsn("")), &gorm.Config{})
	if err != nil {
		log.Fatal("There's an error while you are connecting to database", err.Error())
	}

	repository := repository.NewUserRepository(db)
	id := repository.FindUserByID(3)
	fmt.Printf("User Id with %v \n", id)

	// err = repository.DeleteUserById(3)
	// if err != nil {
	// 	fmt.Println("Fail to delete User", err.Error())
	// } else {
	// 	fmt.Println("Deleting user success")
	// }

}

func dsn(con string) string {
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
