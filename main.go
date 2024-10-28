package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Modes for the editor
type Mode int

const (
	Normal Mode = iota
	Insert
	Visual
)

var currentMode Mode = Normal

// Main function to run the editor
func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().SetDynamicColors(true)
	textView.SetBorder(true).SetTitle("Go Vim Editor").SetTitleAlign(tview.AlignLeft)
	textView.SetScrollable(true)

	// Initial content
	textContent := "Hello, this is a basic Vim-like editor with Goroutines!\n\nPress 'i' to enter Insert mode, 'v' for Visual mode, and 'Esc' to return to Normal mode."
	textView.SetText(textContent)

	// Track cursor position
	cursorX, cursorY := 0, 0

	// Update editor mode and display mode status
	updateMode := func() {
		app.QueueUpdateDraw(func() {
			switch currentMode {
			case Normal:
				textView.SetTitle("Go Vim Editor [Normal Mode]")
			case Insert:
				textView.SetTitle("Go Vim Editor [Insert Mode]")
			case Visual:
				textView.SetTitle("Go Vim Editor [Visual Mode]")
			}
		})
	}

	// Set up keybindings
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch currentMode {
		case Normal:
			switch event.Rune() {
			case 'h':
				cursorX = max(0, cursorX-1)
			case 'j':
				cursorY++
			case 'k':
				cursorY = max(0, cursorY-1)
			case 'l':
				cursorX++
			case 'i':
				currentMode = Insert
				updateMode()
			case 'v':
				currentMode = Visual
				updateMode()
			}
			if event.Key() == tcell.KeyEsc {
				currentMode = Normal
				updateMode()
			}
		case Insert:
			if event.Key() == tcell.KeyEsc {
				currentMode = Normal
				updateMode()
			} else if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
				// Backspace logic (remove last character from text)
				if len(textContent) > 0 {
					textContent = textContent[:len(textContent)-1]
					textView.SetText(textContent)
				}
			} else {
				textContent += string(event.Rune())
				textView.SetText(textContent)
			}
		case Visual:
			if event.Key() == tcell.KeyEsc {
				currentMode = Normal
				updateMode()
			}
		}
		return event
	})

	// Start a background goroutine to update the status periodically
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			// Queue an update only if the application is running
			app.QueueUpdateDraw(func() {
				textView.SetText(fmt.Sprintf("%s\n\nStatus updated at: %s", textContent, time.Now().Format("15:04:05")))
			})
		}
	}()

	updateMode() // Initialize mode display
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}

// Helper function to prevent negative cursor positions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
