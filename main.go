package main

import (
	"github.com/Thospol/go-todo-clean-arch/api/route"
	"github.com/Thospol/go-todo-clean-arch/database"
	"github.com/Thospol/go-todo-clean-arch/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	defer database.DB.Close()

	// run the migrations: todo struct
	database.DB.AutoMigrate(&domain.Todo{})

	//setup routes
	r := route.SetupRouter()

	// running
	err = r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
