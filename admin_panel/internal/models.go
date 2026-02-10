package internal

import (
	tea "github.com/charmbracelet/bubbletea"
)

func InitialModel() model {
	return model {
		notifications: []Notification{},
		connected: false,
		status: "Connecting",
		width: 80,
		height: 24,
	}
}

func (m model) Init() tea.Cmd {
	return ConnectWebSocket
}

