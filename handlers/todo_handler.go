package handlers

import (
	"encoding/json"
	"go_todo/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetTodoLists(w http.ResponseWriter, r *http.Request) {
	var todoLists []models.TodoList
	db.Find(&todoLists)
	render.JSON(w, r, todoLists)
}

func CreateTodoList(w http.ResponseWriter, r *http.Request) {
	var todoList models.TodoList
	err := json.NewDecoder(r.Body).Decode(&todoList)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	db.Create(&todoList)
	render.JSON(w, r, todoList)
}

func DeleteTodoList(w http.ResponseWriter, r *http.Request) {
	todoListID := chi.URLParam(r, "todoListID")

	todoListIdInt, err := strconv.Atoi(todoListID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid todoListID"})
		return
	}

	var todoList models.TodoList
	if err := db.Where("id = ?", todoListIdInt).First(&todoList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]interface{}{"error": "Todo not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": "Database error"})
		return
	}

	db.Delete(&todoList)

	render.JSON(w, r, todoList)
}

func GetTodosByTodoListID(w http.ResponseWriter, r *http.Request) {
	todoListID := chi.URLParam(r, "todoListID")
	var todos []models.Todo
	db.Where("todo_list_id = ?", todoListID).Find(&todos)
	render.JSON(w, r, todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todoListID := chi.URLParam(r, "todoListID")
	id, err := strconv.Atoi(todoListID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid todoListID"})
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid request payload"})
		return
	}
	todo.TodoListID = uint(id)
	db.Create(&todo)
	render.JSON(w, r, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoListID := chi.URLParam(r, "todoListID")
	todoID := chi.URLParam(r, "todoID")

	todoListIdInt, err := strconv.Atoi(todoListID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid todoListID"})
		return
	}

	todoIdInt, err := strconv.Atoi(todoID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid todoID"})
		return
	}

	var updatedTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid request payload"})
		return
	}

	var todo models.Todo
	if err := db.Where("id = ? AND todo_list_id = ?", todoIdInt, todoListIdInt).First(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]interface{}{"error": "Todo not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": "Database error"})
		return
	}

	// Update the fields
	todo.Name = updatedTodo.Name
	todo.Completed = updatedTodo.Completed
	db.Save(&todo)

	render.JSON(w, r, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoListID := chi.URLParam(r, "todoListID")
	todoID := chi.URLParam(r, "todoID")

	todoListIdInt, err := strconv.Atoi(todoListID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid todoListID"})
		return
	}

	todoIdInt, err := strconv.Atoi(todoID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": "Invalid todoID"})
		return
	}

	var todo models.Todo
	if err := db.Where("id = ? AND todo_list_id = ?", todoIdInt, todoListIdInt).First(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]interface{}{"error": "Todo not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": "Database error"})
		return
	}

	db.Delete(&todo)

	render.JSON(w, r, todo)
}
