package handler

import "net/http"

type BotHandler interface {
	SelectOption(w http.ResponseWriter, r *http.Request)
}
