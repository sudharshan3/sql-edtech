package routes

import (
	"github.com/sudharshan3/sql-edtech/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetEdtechHome).Methods("GET")
	r.HandleFunc("/students", controllers.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", controllers.GetStudent).Methods("GET")
	r.HandleFunc("/students", controllers.InsertStudent).Methods("POST")
	r.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students", controllers.DeleteAllStudents).Methods("DELETE")
	return r

}
