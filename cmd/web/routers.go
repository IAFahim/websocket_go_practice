package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"websocket_go_practice/internal/handlers"
)

func router() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}
