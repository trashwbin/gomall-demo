package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/dal"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/model"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
	dal.Init()
	// Create a user
	mysql.DB.Create(&model.User{
		Email:    "demo@example.com",
		Password: "123456",
	})

	// Update a user
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "666")

	// Query a user
	var user model.User
	mysql.DB.Where("email = ?", "demo@example.com").First(&user)
	fmt.Printf("user: %+v\n", user)

	// Delete a user
	// simple delete
	mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})
	// complete delete
	mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})

}
