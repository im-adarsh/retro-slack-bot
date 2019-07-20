package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

type botHandler struct {
	client *slack.Client
}

func New(client *slack.Client) BotHandler {
	return botHandler{client: client}
}

func (botHandler) SelectOption(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>>>>> hellow ")
}

func responseMessage(w http.ResponseWriter, original slack.Message, title, value string) {
	original.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: title,
			Value: value,
			Short: false,
		},
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&original)
}
