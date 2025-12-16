package ws

import (
	"log"
	"net/http"

	"fourinarow/internal/game"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var currentGame = game.NewGame()

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		var col int
		if err := conn.ReadJSON(&col); err != nil {
			log.Println(err)
			return
		}
		currentGame.Drop(col)
		conn.WriteJSON(currentGame)
	}
}
