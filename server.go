package gohotty

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func New() {
	// TODO
}

func Server(w http.ResponseWriter, req *http.Request) {
	websocket.Handler(func(ws *websocket.Conn) {
		fmt.Println("🔥 GoHotty enable")
		defer ws.Close()

		err := websocket.Message.Send(ws, req.URL.Path)
		if err != nil {
			// TODO
			fmt.Println(err)
		}

		msg := ""
		err = websocket.Message.Receive(ws, &msg)
		if err != nil {
			// TODO
			fmt.Println(err)
		}
	}).ServeHTTP(w, req)
}
