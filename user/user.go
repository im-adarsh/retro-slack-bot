package user

import (
	"fmt"

	"github.com/nlopes/slack"
)

type slackUser struct {
	client *slack.Client
}

func New(client *slack.Client) SlackUser {
	return &slackUser{
		client: client,
	}
}

func (s *slackUser) GetUserInfo(user string) *slack.User {
	u, err := s.client.GetUserInfo(user)
	if err != nil {
		fmt.Println("error fetching user info")
	}
	return u
}
