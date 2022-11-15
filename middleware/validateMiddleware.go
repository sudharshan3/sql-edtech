package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	u "github.com/sudharshan3/sql-edtech/utils"
)

var jwt_secret = os.Getenv("API_SECRET")

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(jwt_secret), nil
				})
				if error != nil {
					u.Respond(w, u.Message(false, error.Error()))
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					u.Respond(w, u.Message(false, "Invalid authorization token"))
					return
				}
			}
		} else {
			u.Respond(w, u.Message(false, "An authorization header is required"))
			return
		}
	})
}
