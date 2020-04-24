package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/siller174/meetingHelper/pkg/logger"
	"github.com/siller174/meetingHelper/pkg/notification/service"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = 2 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type NotificationHandler struct {
	notificationManager *service.NotificationManager
}

func (notificationHandler *NotificationHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			logger.Error("Could not create websocket. Error %v", err)
		}
		return
	}

	go writer(connection)
	reader(connection)
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		mesType, mes, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Mestype:%v mes %v \n", mesType, string(mes))
	}
}

func writer(ws *websocket.Conn) {
	pingTicker := time.NewTicker(pingPeriod)

	defer func() {
		pingTicker.Stop()
		ws.Close()
	}()
	for {
		timer := time.After(5 * time.Second)
		select {
		case <-timer:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			ws.WriteMessage(websocket.TextMessage, []byte(`{"status": "UP"}`))
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
