package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/true7ry/go-gorm-restapi/db"
	"github.com/true7ry/go-gorm-restapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tks []models.Task
	db.DB.Find(&tks)
	json.NewEncoder(w).Encode(&tks)

}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tk models.Task
	//num := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&tk)

	createTk := db.DB.Create(&tk)
	err := createTk.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&tk)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tk models.Task
	params := mux.Vars(r)

	db.DB.First(&tk, params["id"])
	if tk.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //400
		w.Write([]byte("Task not found."))
		return
	}
	json.NewEncoder(w).Encode(&tk)
}
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {

	var tk models.Task
	params := mux.Vars(r)

	db.DB.First(&tk, params["id"])
	if tk.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //400
		w.Write([]byte("Task not found."))
		return
	}
	db.DB.Unscoped().Delete(&tk)
	json.NewEncoder(w).Encode(&tk)
}
