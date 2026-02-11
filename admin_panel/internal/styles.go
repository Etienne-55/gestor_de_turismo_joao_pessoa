package internal

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor   = lipgloss.Color("#2563EB")
	successColor   = lipgloss.Color("#10b981")
	warningColor   = lipgloss.Color("#f59e0b")
	dangerColor    = lipgloss.Color("#ef4444")
	infoColor      = lipgloss.Color("#8b5cf6")
	mutedColor     = lipgloss.Color("#6b7280")
	
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(primaryColor).
		Padding(0, 1)
	
	connectedStyle = lipgloss.NewStyle().
		Foreground(successColor).
		Bold(true)
	
	disconnectedStyle = lipgloss.NewStyle().
		Foreground(dangerColor).
		Bold(true)
	
	notificationBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		MarginTop(1)
	
	notificationTypeStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryColor)
	
	notificationMessageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#333333"))
	
	notificationTimeStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		Italic(true)
	
	helpStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		Italic(true).
		MarginTop(1)
)

func getNotificationColor(notifType string) lipgloss.Color {
	switch notifType {
	case "trip_created":
		return successColor
	case "trip_updated":
		return warningColor
	case "trip_deleted":
		return dangerColor
	case "user_signup":
		return infoColor
	default:
		return primaryColor
	}
}

func getNotificationIcon(notifType string) string {
	icons := map[string]string{
		"trip_created": "",
		"trip_updated": "",
		"trip_deleted": "",
		"user_signup":  "",
		"stats_update": "",
	}
	if icon, ok := icons[notifType]; ok {
		return icon
	}
	return ""
}
