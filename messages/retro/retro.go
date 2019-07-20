package retro

import (
	"fmt"

	"github.com/nlopes/slack"
)

const (
	INIT_RETRO_START        = "scrum-master-retro-action-start-retro"
	INIT_RETRO_END          = "scrum-master-retro-action-end-retro"
	INIT_RETRO_DISCARD      = "scrum-master-retro-action-discard-retro"
	INIT_RETRO_SKIP         = "scrum-master-retro-action-skip-retro"
	HISTORY_RETRO_SHOW_LAST = "scrum-master-retro-action-show-last-retro"

	INIT_CALLBACK_ID    = "scrum-master-init-retro-selected-action"
	HISTORY_CALLBACK_ID = "scrum-master-history-retro-selected-action"
)

func GetInitRetroMessage(username string) slack.Attachment {
	return slack.Attachment{
		Text:       fmt.Sprintf("Hey @%s! Please choose what you want to do below : ", username),
		Color:      "#3AA3E3",
		Fallback:   "You are unable to choose an action",
		CallbackID: INIT_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{Name: "sprint", Text: "Start Retro", Type: "button", Value: INIT_RETRO_START, Style: "primary"},
			{Name: "sprint", Text: "End Retro", Type: "button", Value: INIT_RETRO_END},
			{Name: "sprint", Text: "Discard Retro", Type: "button", Value: INIT_RETRO_DISCARD, Style: "danger"},
			{Name: "sprint", Text: "Skip Retro", Type: "button", Value: INIT_RETRO_SKIP},
			{Name: "sprint", Text: "Show Last Retro", Type: "button", Value: HISTORY_RETRO_SHOW_LAST},
		},
	}
}

func ShowRetroDialogMessage(username string) slack.Attachment {
	return slack.Attachment{
		Text:       fmt.Sprintf("Hey @%s! Please choose what you want to do below : ", username),
		Color:      "#3AA3E3",
		Fallback:   "You are unable to choose an action",
		CallbackID: INIT_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{Name: "sprint", Text: "Start Retro", Type: "button", Value: INIT_RETRO_START, Style: "primary"},
			{Name: "sprint", Text: "End Retro", Type: "button", Value: INIT_RETRO_END},
			{Name: "sprint", Text: "Discard Retro", Type: "button", Value: INIT_RETRO_DISCARD, Style: "danger"},
			{Name: "sprint", Text: "Skip Retro", Type: "button", Value: INIT_RETRO_SKIP},
			{Name: "sprint", Text: "Show Last Retro", Type: "button", Value: HISTORY_RETRO_SHOW_LAST},
		},
	}
}

func ShowRetroHistoryMessage(username string) slack.Attachment {
	return slack.Attachment{
		Color:      "#3AA3E3",
		Fallback:   "You are unable to choose an action",
		CallbackID: HISTORY_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{Name: "sprint", Text: "Show Last Retro", Type: "button", Value: HISTORY_RETRO_SHOW_LAST},
		},
	}
}

func GetRetroCallbackMessage(callbackID string) slack.Attachment {
	return slack.Attachment{
		Color:    "#3AA3E3",
		Fallback: fmt.Sprintf("Triggering %v", callbackID),
	}
}
