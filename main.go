package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"golang-gin-sample/api/v1/todo"

	"github.com/sirupsen/logrus"
)

var DB *gorm.DB

func init() {
	fmt.Println("init db")
	DB, err := gorm.Open("mysql", "root:123456@/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Panicf("failed to connect database %v", err)
	}
	defer DB.Close()
	DB.AutoMigrate(&todo.TodoModel{})
	todo.DB = DB
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	v1.POST("/todos", todo.CreateTodo)
	v1.GET("/todos", todo.GetTodos)
	//	v1.GET("/todos/:id", todo.GetTodo)
	//	v1.PUT("/todos/:id", todo.UpdateTode)
	//	v1.DELETE("/todos/:id", todo.DeleteTodo)
}
