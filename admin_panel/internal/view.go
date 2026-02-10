package internal

import (
	"fmt"
	"strings"
	"time"
	
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// View renders the UI
func (m model) View() string {
	var s strings.Builder
	
	// Header
	header := headerStyle.Render("🔔 Tourism Platform - Admin Dashboard")
	s.WriteString(header)
	s.WriteString("\n\n")
	
	// Status
	statusText := m.status
	var statusStyled string
	if m.connected {
		statusStyled = connectedStyle.Render(statusText)
	} else {
		statusStyled = disconnectedStyle.Render(statusText)
	}
	s.WriteString("Status: " + statusStyled)
	s.WriteString("\n")
	
	// Statistics
	stats := fmt.Sprintf("Total Notifications: %d", len(m.notifications))
	s.WriteString(stats)
	s.WriteString("\n")
	
	// Divider
	s.WriteString(strings.Repeat("─", m.width))
	s.WriteString("\n")
	
	// Notifications
	if len(m.notifications) == 0 {
		s.WriteString("\n")
		s.WriteString(lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Render("  No notifications yet. Waiting for activity..."))
		s.WriteString("\n")
	} else {
		// Show last 10 notifications
		maxDisplay := 10
		if len(m.notifications) < maxDisplay {
			maxDisplay = len(m.notifications)
		}
		
		for i := 0; i < maxDisplay; i++ {
			notif := m.notifications[i]
			s.WriteString(renderNotification(notif))
		}
	}
	
	// Help
	s.WriteString("\n")
	s.WriteString(strings.Repeat("─", m.width))
	s.WriteString("\n")
	help := helpStyle.Render("Press 'c' to clear | 'q' to quit")
	s.WriteString(help)
	
	return s.String()
}

// renderNotification creates a styled notification
func renderNotification(n Notification) string {
	color := getNotificationColor(n.Type)
	icon := getNotificationIcon(n.Type)
	
	box := notificationBoxStyle.
		BorderForeground(color)
	
	typeStyled := notificationTypeStyle.
		Foreground(color).
		Render(icon + " " + formatType(n.Type))
	
	messageStyled := notificationMessageStyle.Render(n.Message)
	
	timeAgo := formatTimeAgo(n.Timestamp)
	timeStyled := notificationTimeStyle.Render(timeAgo)
	
	content := fmt.Sprintf("%s\n%s\n%s", typeStyled, messageStyled, timeStyled)
	
	return box.Render(content)
}

// formatType converts type to readable format
func formatType(t string) string {
	types := map[string]string{
		"trip_created": "Trip Created",
		"trip_updated": "Trip Updated",
		"trip_deleted": "Trip Deleted",
		"user_signup":  "New User",
		"stats_update": "Statistics",
	}
	if formatted, ok := types[t]; ok {
		return formatted
	}
	return strings.ReplaceAll(t, "_", " ")
}

// formatTimeAgo converts timestamp to "X ago" format
func formatTimeAgo(t time.Time) string {
	diff := time.Since(t)
	
	if diff.Seconds() < 60 {
		return fmt.Sprintf("%d seconds ago", int(diff.Seconds()))
	} else if diff.Minutes() < 60 {
		mins := int(diff.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	} else if diff.Hours() < 24 {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else {
		return t.Format("Jan 02, 15:04")
	}
}

