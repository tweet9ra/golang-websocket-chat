package app

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"websocketsProject/models"
	u "websocketsProject/utils"
)

 func JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/user/store", "/api/user/login"} //Список эндпоинтов, для которых не требуется авторизация
		requestPath := r.URL.Path //текущий путь запроса

		// Use middleware only on api requests and skip check when method is OPTIONS
		if strings.Split(requestPath, "/")[1] != "api" || r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Получение токена

		// token missing
		if tokenHeader == "" {
			respondError("Missing auth token", w)
			return
		}

		// token format check
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			respondError("Invalid/Malformed auth token", w)
			return
		}

		// Parse token
		tokenPart := splitted[1]
		userId, err := GetUserId(tokenPart)
		if err != nil {
			respondError(err.Error(), w)
			return
		}

		ctx := context.WithValue(r.Context(), "user", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// Send http response with 403 status and error msg
func respondError(error string, w http.ResponseWriter)  {
	response := u.Message(false, error)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	u.Respond(w, response)
}

type invalidToken struct {}
func (err *invalidToken) Error() string {
	return "Token is not valid."
}

// Validate JWT token and decode user id
func GetUserId(token string) (uint, error) {
	tk := &models.Token{}

	parsedToken, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, &invalidToken{}
	}

	return tk.UserId, nil
}