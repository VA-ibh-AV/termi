package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"golang.design/x/clipboard"
)

func PrintAIResponse(command string, elapsed time.Duration) {
	// Copy to clipboard (best effort)
	_ = clipboard.Write(clipboard.FmtText, []byte(command))

	cmdStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("229"))

	metaStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	fmt.Println()
	fmt.Println(cmdStyle.Render(command))
	fmt.Println()
	fmt.Println(
		metaStyle.Render(
			fmt.Sprintf("⏱  %s   📋 copied", elapsed.Round(time.Millisecond)),
		),
	)
}
