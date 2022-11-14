package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sudharshan3/sql-edtech/models"
)

func LoginStudents(w http.ResponseWriter, r *http.Request) {
	var reqloginstudent models.LoginStudent
	var resloginstudent models.LoginStudent
	var response models.Response
	var loginstudents []models.LoginStudent
	// var students []models.Student
	// var student models.Student
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter some Data...")
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&reqloginstudent)
	username := reqloginstudent.Username
	password := reqloginstudent.Password
	result, err := db.Query("SELECT student_email,student_pass FROM students WHERE student_email=?", username)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		for result.Next() {
			err = result.Scan(&resloginstudent.Username, &resloginstudent.Password)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
				return
			} else {
				loginstudents = append(loginstudents, resloginstudent)
			}
		}
		if len(loginstudents) == 0 {
			json.NewEncoder(w).Encode("User Doesnt Exist. Please Register")
			return
		} else {
			if len(loginstudents) == 1 && CheckPasswordHash(password, resloginstudent.Password) {
				response.Status = 200
				response.Message = "Login Successful!..."
				json.NewEncoder(w).Encode(response)
				return
			} else {
				json.NewEncoder(w).Encode("Incorrect Password")
				return
			}
		}

	}
}
