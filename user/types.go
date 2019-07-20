package user

import "github.com/nlopes/slack"

type SlackUser interface {
	GetUserInfo(user string) *slack.User
}
