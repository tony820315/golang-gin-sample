package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"golang-gin-sample/pkg/resp"
)

var DB *gorm.DB

type TodoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"Completed"`
}

type TransformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"Completed"`
}

func CreateTodo(c *gin.Context) {
	respBody := resp.NewResponseBody(resp.NewBaseError(http.StatusCreated, "Todo item created successfully", nil))

	for {
		completed, _ := strconv.Atoi(c.PostForm("completed"))
		todo := TodoModel{Title: c.PostForm("title"), Completed: completed}
		DB.Save(&todo)
		respBody.Result = todo.ID
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
