package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sudharshan3/sql-edtech/models"
)

var jwt_secret = os.Getenv("API_SECRET")

func LoginStudents(w http.ResponseWriter, r *http.Request) {
	var reqloginstudent models.LoginStudent
	var resloginstudent models.LoginStudent
	var logresponse models.LoginResponse
	var loginfo models.LoginInfo
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
				ttl := 2 * time.Minute
				accessTokenExpire := os.Getenv("ACESS_TOKEN_EXPIRE")
				min, err := strconv.Atoi(accessTokenExpire)
				if err != nil {
					log.Println(err)
				}
				if accessTokenExpire != "" {
					ttl = time.Duration(min) * time.Minute
				}
				loginfo.Token = CreateToken(w, reqloginstudent.Username, resloginstudent.Password, ttl)
				loginfo.Email = resloginstudent.Username

				logresponse.Status = 200
				logresponse.Message = "Login Successful!..."
				logresponse.Data = loginfo
				json.NewEncoder(w).Encode(logresponse)
				return
			} else {
				json.NewEncoder(w).Encode("Incorrect Password")
				return
			}
		}

	}
}

func CreateToken(w http.ResponseWriter, username string, password string, ttl time.Duration) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      ttl,
	})

	tokenString, error := token.SignedString([]byte(jwt_secret))
	if error != nil {
		json.NewEncoder(w).Encode(error.Error())
	}

	return tokenString
}
