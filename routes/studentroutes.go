package routes

import (
	"github.com/sudharshan3/sql-edtech/controllers"
	"github.com/sudharshan3/sql-edtech/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetEdtechHome).Methods("GET")
	r.HandleFunc("/students", middleware.ValidateMiddleware(controllers.GetAllStudents)).Methods("GET")
	r.HandleFunc("/students/{id}", middleware.ValidateMiddleware(controllers.GetStudent)).Methods("GET")
	r.HandleFunc("/students", controllers.InsertStudent).Methods("POST")
	r.HandleFunc("/students/{id}", middleware.ValidateMiddleware(controllers.UpdateStudent)).Methods("POST")
	r.HandleFunc("/students/{id}", middleware.ValidateMiddleware(controllers.DeleteStudent)).Methods("DELETE")
	r.HandleFunc("/students", middleware.ValidateMiddleware(controllers.DeleteAllStudents)).Methods("DELETE")
	r.HandleFunc("/login", controllers.LoginStudents).Methods("POST")
	return r

}
