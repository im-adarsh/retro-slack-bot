package listener

import (
	"fmt"
	"log"

	user2 "github.com/im-adarsh/retro-slack-bot/user"

	"github.com/im-adarsh/retro-slack-bot/messages/retro"

	"github.com/nlopes/slack"
)

type slackListener struct {
	client *slack.Client
	botID  string
	user   user2.SlackUser
}

func New(client *slack.Client, botId string) BotListener {
	user := user2.New(client)
	return &slackListener{
		client: client,
		botID:  botId,
		user:   user,
	}
}

func (s *slackListener) ListenAndResponse() {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		}
	}
}

func (s *slackListener) handleMessageEvent(ev *slack.MessageEvent) error {

	if ev.Msg.User == s.botID {
		return nil
	}

	u := s.user.GetUserInfo(ev.Msg.User)
	_, _, err := s.client.PostMessage(ev.Channel, retro.GetInitRetroMessage(u.Name)...)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return nil
}
