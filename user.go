package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func (s *SlackListener) getUserInfo(user string) *slack.User {
	u, err := s.client.GetUserInfo(user)
	if err != nil {
		fmt.Println("error fetching user info")
	}
	return u
}
