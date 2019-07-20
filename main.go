package main

import (
	"fmt"
	"os"

	"github.com/im-adarsh/retro-slack-bot/messages/retro"

	"github.com/nlopes/slack"
)

var (
	slackClient *slack.Client
)

func main() {
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {

		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(ev)
		case *slack.AttachmentAction:
			go handleCallback(ev)

		default:
			fmt.Println()
		}
	}
}

func handleCallback(event *slack.AttachmentAction) {
	fmt.Println(fmt.Sprintf("%v\n", event))
	//logUserInfo(event.Value)

	go replyAck(event)
}

func replyAck(ev *slack.AttachmentAction) {
	_, _, err := slackClient.PostMessage(ev.Name,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(retro.GetRetroCallbackMessage(ev.Value)))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	return
}

func handleMessage(event *slack.MessageEvent) {
	fmt.Println(fmt.Sprintf("%v\n", event))
	if event.Msg.User == "ULNQ5BJ7Q" {
		return
	}
	logUserInfo(event.Msg.User)
	go replyEmpty(event)
}

func replyEmpty(ev *slack.MessageEvent) {

	u := getUserInfo(ev.Msg.User)
	_, _, err := slackClient.PostMessage(ev.Channel,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(retro.GetInitRetroMessage(u.Name), retro.ShowRetroHistoryMessage(u.Name)))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	return
}

func getUserInfo(user string) *slack.User {
	u, err := slackClient.GetUserInfo(user)
	if err != nil {
		fmt.Println("error fetching user info")
	}
	return u
}

func logUserInfo(user string) {
	u := getUserInfo(user)
	fmt.Println("############################################################")
	fmt.Println(u.Name)
	fmt.Println("############################################################")
}
