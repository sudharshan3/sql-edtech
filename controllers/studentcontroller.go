package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sudharshan3/sql-edtech/config"

	"github.com/sudharshan3/sql-edtech/models"

	"github.com/gorilla/mux"
)

var db *sql.DB = config.GetDB()

func GetEdtechHome(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Edtech APP"))

}
func GetAllStudents(w http.ResponseWriter, r *http.Request) {

	var student models.Student
	var response models.Response
	var studentcollection []models.Student

	result, err := db.Query("SELECT student_id,student_name,student_email,student_phone FROM edtech.students")
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		for result.Next() {
			err = result.Scan(&student.Id, &student.Name, &student.Email, &student.Phone)
			if err != nil {
				log.Println(err.Error())
				return
			} else {
				studentcollection = append(studentcollection, student)
			}

		}
		response.Status = 200
		response.Message = "Success"
		response.Data = studentcollection
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}

}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	var response models.Response
	var studentcollection []models.Student

	param := mux.Vars(r)
	result, err := db.Query("SELECT student_id,student_name,student_email,student_phone FROM edtech.students WHERE student_id=" + param["id"])
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next() {
		err = result.Scan(&student.Id, &student.Name, &student.Email, &student.Phone)
		if err != nil {
			log.Println(err.Error())
			return
		} else {
			studentcollection = append(studentcollection, student)
		}
	}
	if len(studentcollection) == 0 {
		json.NewEncoder(w).Encode("No Record Found with requested ID")
		return
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = studentcollection

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(response)
}

func InsertStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	var response models.Response
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter some Data!")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&student)

	if student.Password == "" || student.Email == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Email and Password are required")
		return
	} else {
		name := student.Name
		pass, _ := HashPassword(student.Password)
		email := student.Email
		phone := student.Phone
		_, err := db.Exec("INSERT INTO students (student_name,student_email,student_phone,student_pass) VALUES (?,?,?,?)", name, email, phone, pass)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error Inserting Data")
			log.Println(err.Error())
			json.NewEncoder(w).Encode("Error Inserting Data...")
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		response.Status = 200
		response.Message = "Student Created Successfully"
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)

	}

}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	var studentcollection []models.Student
	var response models.Response
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter some Data!")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&student)
	name := student.Name
	pass := student.Password
	email := student.Email
	phone := student.Phone
	params := mux.Vars(r)
	count, err := db.Query("SELECT student_id FROM students WHERE student_id=" + params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error Updating Data...")
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		for count.Next() {
			err = count.Scan(&student.Id)
			if err != nil {
				log.Println(err.Error())

			} else {
				studentcollection = append(studentcollection, student)
			}

		}
		if len(studentcollection) == 0 {
			json.NewEncoder(w).Encode("No records found with ID")
			return
		} else {
			if pass != "" {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode("Password Cannot be Updated Here...")
				return
			}
			if name != "" {
				_, err := db.Exec("UPDATE students SET student_name=? WHERE student_id=?", name, params["id"])
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode("Error Updating Data...")
					json.NewEncoder(w).Encode(err.Error())
					return
				} else {
					response.Status = 200
					response.Message = "Student Name Updated Successfully"
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(response)
				}
			}

			if email != "" {
				_, err := db.Exec("UPDATE students SET student_email=? WHERE student_id=?", email, params["id"])
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode("Error Updating Data...")
					json.NewEncoder(w).Encode(err.Error())
					return
				} else {
					response.Status = 200
					response.Message = "Student Email Updated Successfully"
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(response)
				}
			}
			if phone != "" {
				_, err := db.Exec("UPDATE students SET student_phone=? WHERE student_id=?", phone, params["id"])
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode("Error Updating Data...")
					json.NewEncoder(w).Encode(err.Error())
					return
				} else {
					response.Status = 200
					response.Message = "Student Phone Updated Successfully"
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(response)
				}
			}

		}

	}

}
func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	_, err := db.Query("DELETE FROM students ")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response.Status = 200
	response.Message = "All Records Deleted Successfully"
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(response)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	var response models.Response
	var studentcollection []models.Student

	param := mux.Vars(r)
	result, err := db.Query("SELECT student_id,student_name,student_email,student_phone FROM edtech.students WHERE student_id=" + param["id"])
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next() {
		err = result.Scan(&student.Id, &student.Name, &student.Email, &student.Phone)
		if err != nil {
			log.Println(err.Error())
			return
		} else {
			studentcollection = append(studentcollection, student)
		}
	}
	if len(studentcollection) == 0 {
		json.NewEncoder(w).Encode("No Record Found with requested ID")
		return
	} else {
		_, err = db.Exec("DELETE FROM students WHERE student_id=?", param["id"])
		response.Status = 200
		response.Message = "ID-" + param["id"] + " deleted Successfully"

	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(response)
}
