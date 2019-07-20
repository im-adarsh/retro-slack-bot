package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/nlopes/slack/slackevents"

	"github.com/im-adarsh/retro-slack-bot/messages/retro"

	"github.com/nlopes/slack"
)

// interactionHandler handles interactive message response.
type interactionHandler struct {
	slackClient *slack.Client
}

func (h interactionHandler) EventListener(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>>>>> hellow ")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: "TOKEN"}))
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			_, _, _ = h.slackClient.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		}
	}
}

func (h interactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>>>>>>>>>>> click >>>>>>>>>>>>>")
	if r.Method != http.MethodPost {
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(">>>>>>>>>>>> click1 >>>>>>>>>>>>>")
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(">>>>>>>>>>>> click2 >>>>>>>>>>>>>")
	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[ERROR] Failed to unespace request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(">>>>>>>>>>>> click3 >>>>>>>>>>>>>")
	var message slack.InteractionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(">>>>>>>>>>>> click4 >>>>>>>>>>>>>")
	switch message.CallbackID {
	case retro.INIT_CALLBACK_ID:
		fallthrough
	case retro.HISTORY_CALLBACK_ID:
		fallthrough
	default:
		title := ":ok: your selection was submitted! yay!"
		responseMessage(w, message.OriginalMessage, title, "")
		return
	}
}

// responseMessage response to the original slackbutton enabled message.
// It removes button and replace it with message which indicate how bot will work
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
