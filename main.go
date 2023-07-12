package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/true7ry/go-gorm-restapi/db"
	"github.com/true7ry/go-gorm-restapi/models"
	"github.com/true7ry/go-gorm-restapi/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Users{})
	db.DB.AutoMigrate(models.Task{})

	r := mux.NewRouter()

	// Users routes
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	// Tasks routes
	r.HandleFunc("/task", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	// Iniciar el server (HAY QUE TENERLO SIEMPRE DEBAJO DE TODO)
	http.ListenAndServe(":3000", r)
}
