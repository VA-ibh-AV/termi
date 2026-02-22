package ui

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

var frames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func StartSpinner(ctx context.Context, text string) func() {
	ticker := time.NewTicker(80 * time.Millisecond)
	done := make(chan struct{})

	lineLen := len(text) + 4 // spinner + space + buffer

	go func() {
		i := 0
		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(os.Stdout, "\r%s %s", frames[i%len(frames)], text)
				i++
			case <-ctx.Done():
				return
			case <-done:
				return
			}
		}
	}()

	return func() {
		ticker.Stop()
		close(done)

		// HARD clear line (Windows-safe)
		fmt.Fprintf(os.Stdout, "\r%s\r", strings.Repeat(" ", lineLen))
	}
}
