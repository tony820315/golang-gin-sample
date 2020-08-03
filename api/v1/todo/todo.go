package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"golang-gin-sample/pkg/resp"
)

var DB *gorm.DB

type TodoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

type TransformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateTodo(c *gin.Context) {
	respBody := resp.NewResponseBody(resp.NewBaseError(http.StatusCreated, "Todo item created successfully", nil))

	for {
		var todo TodoModel
		if err := c.BindJSON(&todo); err != nil {
			respBody.SetExtendError(resp.NewBaseError(http.StatusBadRequest, "", err))
			break
		}
		DB.Save(&todo)
		respBody.Result = todo
		break
	}
	c.JSON(respBody.StatusCode(), respBody)
}

func GetTodos(c *gin.Context) {
	var todos []TodoModel
	var _todos []TransformedTodo
	respBody := resp.NewResponseBody(resp.NewBaseError(http.StatusOK, "", nil))
	for {
		DB.Find(&todos)
		if len(todos) <= 0 {
			respBody.SetExtendError(resp.NewBaseError(http.StatusNotFound, "", nil))
			break
		}
		for _, item := range todos {
			completed := false
			if item.Completed == 1 {
				completed = true
			}
			_todos = append(_todos, TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
		}
		respBody.Result = _todos
		break
	}
	c.JSON(respBody.StatusCode(), respBody)
}

func GetTodo(c *gin.Context) {
	var todo TodoModel
	respBody := resp.NewResponseBody(resp.NewBaseError(http.StatusOK, "", nil))
	for {
		id := c.Param("id")
		DB.First(&todo, id)
		if todo.ID == 0 {
			respBody.SetExtendError(resp.NewBaseError(http.StatusNotFound, "No todo found!", nil))
			break
		}
		completed := false
		if todo.Completed == 1 {
			completed = true
		}
		respBody.Result = TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
		break
	}
	c.JSON(respBody.StatusCode(), respBody)
}

func UpdateTodo(c *gin.Context) {
	var todo TodoModel
	respBody := resp.NewResponseBody(resp.NewBaseError(http.StatusOK, "", nil))
	for {
		id := c.Param("id")
		DB.First(&todo, id)

		if todo.ID == 0 {
			respBody.SetExtendError(resp.NewBaseError(http.StatusNotFound, "No todo found!", nil))
			break
		}
		var item TodoModel
		if err := c.BindJSON(&item); err != nil {
			respBody.SetExtendError(resp.NewBaseError(http.StatusBadRequest, "body not found", err))
			break
		}

		DB.Model(&todo).Updates(TodoModel{Title: item.Title, Completed: item.Completed})
		respBody.Message = "Todo updated successfully!"
		break
	}
	c.JSON(respBody.StatusCode(), respBody)
}

func DeleteTodo(c *gin.Context) {
	var todo TodoModel
	respBody := resp.NewResponseBody(resp.NewBaseError(http.StatusOK, "", nil))
	for {
		id := c.Param("id")
		DB.First(&todo, id)

		if todo.ID == 0 {
			respBody.SetExtendError(resp.NewBaseError(http.StatusNotFound, "No todo found!", nil))
			break
		}
		fmt.Println(todo)
		DB.Delete(&todo)
		respBody.Message = "Todo deleted successfully!"
		break
	}
	c.JSON(respBody.StatusCode(), respBody)
}
