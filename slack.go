package main

import (
	"fmt"
	"log"

	"github.com/im-adarsh/retro-slack-bot/messages/retro"

	"github.com/nlopes/slack"
)

type SlackListener struct {
	client *slack.Client
	botID  string
}

// LstenAndResponse listens slack events and response
// particular messages. It replies by slack message button.
func (s *SlackListener) ListenAndResponse() {
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
		}
	}
}

// handleMesageEvent handles message events.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {

	if ev.Msg.User == s.botID {
		return nil
	}

	u := s.getUserInfo(ev.Msg.User)
	_, _, err := s.client.PostMessage(ev.Channel,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(retro.GetInitRetroMessage(u.Name), retro.ShowRetroHistoryMessage(u.Name)))
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return nil
}
