package main

import (
	"admin_panel/internal"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		internal.InitialModel(),
		tea.WithAltScreen(),      
		tea.WithMouseCellMotion(), 
	)

	internal.Program = p

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
