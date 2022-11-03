package main

import (
	"fmt"
	"go-orm/migration"
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

	migration.AutoMigration(db) //used to migrate a data

	repository := repository.NewUserRepository(db)
	// id := repository.FindUserByID(3)
	// fmt.Printf("User Id with %v \n", id)

	// var newuser = models.User{
	// 	UserName: "Bintang",
	// 	Email:    "bintang@gmail.com",
	// }
	// x, err := repository.UpdateUser(10, &newuser)
	// log.Printf("Email: %s \n", x.Email)

	users, _ := repository.FindAllUser()
	for _, data := range *users {
		fmt.Println(data)
	}

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
