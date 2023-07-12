package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/true7ry/go-gorm-restapi/db"
	"github.com/true7ry/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	db.DB.Find(&users)                // guarda los usuarios del DB en la variable "users"
	json.NewEncoder(w).Encode(&users) // muestra por pantalla la variable "users" que contienen los users(lol).
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	params := mux.Vars(r)            // recibe la info(el ID) y la guarda en la variable "params".
	db.DB.First(&user, params["id"]) // sirve para buscar el "user" con el ID indicado.

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found."))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks) // sirve para mostrar la tarea(PERO SOLO CUANDO ESPECIFICAS el ID)

	json.NewEncoder(w).Encode(&user)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	json.NewDecoder(r.Body).Decode(&user) // esto guarda la info que recibe de "r" a la variable "user"

	createUser := db.DB.Create(&user) // esto creara el user en la DB, utilizando los datos de la variable.
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	params := mux.Vars(r)

	db.DB.First(&user, params["id"]) // sirve para buscar el "user" con el ID indicado.

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found."))
		return
	}
	//db.DB.Delete(&user) << esto haria que desaparesca al intentar mostrar TODOS los users, pero no lo elimina de la DB.
	db.DB.Unscoped().Delete(&user) // esto ELIMINARIA el usuario de la DB.
	w.WriteHeader(http.StatusOK)

}
