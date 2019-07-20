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

func GetInitRetroMessage(username string) []slack.MsgOption {
	u1 := slack.MsgOptionAsUser(true)
	u2 := slack.MsgOptionAttachments(
		slack.Attachment{
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
		})

	return []slack.MsgOption{u1, u2}
}

func ShowAddRetroDialogMessage(username string) slack.Dialog {
	return slack.Dialog{
		Title:       "Request a coffee",
		SubmitLabel: "Submit",
		CallbackID:  username + "coffee_order_form",
		Elements: []slack.DialogElement{
			slack.DialogInputSelect{
				DialogInput: slack.DialogInput{
					Label:       "Coffee Type",
					Type:        slack.InputTypeSelect,
					Name:        "mealPreferences",
					Placeholder: "Select a drink",
				},
				Options: []slack.DialogSelectOption{
					{
						Label: "Cappuccino",
						Value: "cappuccino",
					},
					{
						Label: "Latte",
						Value: "latte",
					},
					{
						Label: "Pour Over",
						Value: "pourOver",
					},
					{
						Label: "Cold Brew",
						Value: "coldBrew",
					},
				},
			},
			slack.DialogInput{
				Label:    "Customization orders",
				Type:     slack.InputTypeTextArea,
				Name:     "customizePreference",
				Optional: true,
			},
			slack.DialogInput{
				Label:       "Time to deliver",
				Type:        slack.InputTypeText,
				Name:        "timeToDeliver",
				Placeholder: "hh:mm",
			},
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
