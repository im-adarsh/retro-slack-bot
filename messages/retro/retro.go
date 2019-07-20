package retro

import (
	"fmt"

	"github.com/nlopes/slack"
)

const INIT_CALLBACK_ID = "scrum-master-init-retro-selected-action"

const HISTORY_CALLBACK_ID = "scrum-master-history-retro-selected-action"

func GetInitRetroMessage(username string) slack.Attachment {
	return slack.Attachment{
		Text:       fmt.Sprintf("Hey @%s! Please choose what you want to do below : ", username),
		Color:      "#3AA3E3",
		Fallback:   "You are unable to choose an action",
		CallbackID: INIT_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{Name: "sprint", Text: "Start Retro", Type: "button", Value: "scrum-master-retro-action-start-retro", Style: "primary"},
			{Name: "sprint", Text: "End Retro", Type: "button", Value: "scrum-master-retro-action-end-retro"},
			{Name: "sprint", Text: "Discard Retro", Type: "button", Value: "scrum-master-retro-action-discard-retro", Style: "danger"},
			{Name: "sprint", Text: "Skip Retro", Type: "button", Value: "scrum-master-retro-action-skip-retro"},
		},
	}
}

func ShowRetroHistoryMessage(username string) slack.Attachment {
	return slack.Attachment{
		Color:      "#3AA3E3",
		Fallback:   "You are unable to choose an action",
		CallbackID: HISTORY_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{Name: "sprint", Text: "Show Last Retro", Type: "button", Value: "scrum-master-retro-action-show-last-retro"},
		},
	}
}

func GetRetroCallbackMessage(callbackID string) slack.Attachment {
	return slack.Attachment{
		Color:    "#3AA3E3",
		Fallback: fmt.Sprintf("Triggering %v", callbackID),
	}
}
