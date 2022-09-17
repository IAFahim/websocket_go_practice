package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MassageType string `json:"massageType"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Client Successfully Connected")
	}
	var response WsJsonResponse = WsJsonResponse{
		Action:  "connected",
		Message: `<em><small>Connected to server</small></em>`,
	}
	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	views, err := views.GetTemplate(tmpl)
	if err != nil {
		return err
	}
	err = views.Execute(w, data, nil)
	if err != nil {
		return err
	}
	return nil
}
