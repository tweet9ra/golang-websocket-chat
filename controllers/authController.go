package controllers

import (
	"encoding/json"
	"net/http"
	"websocketsProject/models"
	u "websocketsProject/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decodedResponse := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&decodedResponse)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	user := &models.User{
		Email: decodedResponse["email"],
		Password: decodedResponse["password"],
	}

	resp := user.Create() //Создать аккаунт
	u.Respond(w, resp)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	decodedResponse := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&decodedResponse)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	user := &models.User{
		Email: decodedResponse["email"],
		Password: decodedResponse["password"],
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

func Current(w http.ResponseWriter, r *http.Request)  {
	userId := r.Context().Value("user").(uint)

	resp := u.Message(true, "success")
	resp["data"] = models.GetUser(userId)

	u.Respond(w, resp)
}