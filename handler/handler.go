package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/im-adarsh/retro-slack-bot/messages/retro"

	"github.com/nlopes/slack"
)

type botHandler struct {
	client *slack.Client
}

func New(client *slack.Client) BotHandler {
	return botHandler{client: client}
}

func (botHandler) SelectOption(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[ERROR] Failed to unespace request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Only accept message from slack with valid token
	//if message.Token != h.verificationToken {
	//	log.Printf("[ERROR] Invalid token: %s", message.Token)
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	if message.ActionCallback.AttachmentActions == nil {
		return
	}
	action := message.ActionCallback.AttachmentActions[0].Value
	switch action {
	case retro.INIT_RETRO_START:
		fmt.Println("START : ", action)
		title := ":ok: Let's start the retro!"
		responseMessage(w, message.OriginalMessage, title, "")
		return
	case retro.INIT_RETRO_END:
		fmt.Println("END : ", action)
		title := ":ok: Ending the retro!"
		responseMessage(w, message.OriginalMessage, title, "")
		return
	case retro.INIT_RETRO_SKIP:
		fmt.Println("SKIP : ", action)
		title := ":ok: skip retro!"
		responseMessage(w, message.OriginalMessage, title, "")
		return
	case retro.INIT_RETRO_DISCARD:
		fmt.Println("DISCARD : ", action)
		title := ":ok: discarding above retro!"
		responseMessage(w, message.OriginalMessage, title, "")
		return
	case retro.HISTORY_RETRO_SHOW_LAST:
		fmt.Println("HISTORY : ", action)
		title := ":ok: showing last retro!"
		showRetroDialogMessage(w, message.OriginalMessage, title, "")
		return
	}

	//action := message.Actions[0]
	//switch action.Name {
	//
	//default:
	//	log.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
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

func showRetroDialogMessage(w http.ResponseWriter, original slack.Message, title, value string) {
	original.Attachments[0].Actions = []slack.AttachmentAction{}
	original.Type = "dialog"
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: "Whats went well ?",
			Value: "",
			Short: false,
		},
		{
			Title: "Things to be improved ?",
			Value: "",
			Short: false,
		},
	}
	original.Type = "textarea"
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&original)
}
