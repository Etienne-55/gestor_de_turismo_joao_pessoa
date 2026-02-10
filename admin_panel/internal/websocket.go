package internal

import (
	"encoding/json"
	"log"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

var conn *websocket.Conn

func ConnectWebSocket() tea.Msg {
	var err error
	conn, _, err = websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Printf("Connection error: %v", err)
		return ErrorMsg{err}
	}
	
	go ListenForMessages()
	
	return ConnectedMsg{}
}

func ListenForMessages() {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			sendMsg(DisconnectedMsg{})
			return
		}
		
		var notification Notification
		if err := json.Unmarshal(message, &notification); err != nil {
			log.Printf("Parse error: %v", err)
			continue
		}
		
		sendMsg(NotificationMsg{Notification: notification})
	}
}

var Program *tea.Program

func sendMsg(msg tea.Msg) {
	if Program != nil {
		Program.Send(msg)
	}
}

