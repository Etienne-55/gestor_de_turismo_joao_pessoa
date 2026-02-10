package internal

import "time"


type Notification struct {
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	Data      any       `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

type NotificationMsg struct {
	Notification Notification
}

type ErrorMsg struct {
	Err error
}

type ConnectedMsg struct{}

type DisconnectedMsg struct{}

type reconnectMsg struct{}

type model struct {
	notifications []Notification
	connected     bool
	status        string
	width         int
	height        int
}

