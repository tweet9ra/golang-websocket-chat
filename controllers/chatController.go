package controllers

import (
	"net/http"
	"websocketsProject/models"
	u "websocketsProject/utils"
)

func ChatsOfCurrentUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	resp := u.Message(true, "success")
	chats := getUserChatsWithMessages(userId)
	for _, chat := range chats {
		for _, message := range chat.Messages {
			message.ChatID = 0
		}
	}
	resp["chats"] = chats

	u.Respond(w, resp)
}

func getUserChatsWithMessages(userId uint) []*models.Chat {
	db := models.GetDB()

	var chats []*models.Chat

	db.Where("user_id = ?", userId).
		Joins("left join user_chats uc on uc.chat_id = chats.id").
		Preload("Users").
		Preload("Messages").
		Find(&chats)

	return chats
}