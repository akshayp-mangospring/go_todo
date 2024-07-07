package main

import (
	"go_todo/handlers"
	"go_todo/migrations"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	db := migrations.MigrateDB()

	r := chi.NewRouter()

	// Configure CORS middleware
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Change this to your frontend origin to be more secure
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsConfig.Handler)

	handlers.SetDB(db)

	// Define routes
	r.Route("/api/todolists", func(r chi.Router) {
		r.Get("/", handlers.GetTodoLists)
		r.Post("/", handlers.CreateTodoList)
		r.Delete("/{todoListID}", handlers.DeleteTodoList)

		r.Route("/{todoListID}/todos", func(r chi.Router) {
			r.Get("/", handlers.GetTodosByTodoListID)
			r.Post("/", handlers.CreateTodo)
			r.Put("/{todoID}", handlers.UpdateTodo)
			r.Delete("/{todoID}", handlers.DeleteTodo)
		})
	})

	// Start server
	http.ListenAndServe(":3000", r)
}
