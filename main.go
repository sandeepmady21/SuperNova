package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Mode int

const (
	Normal Mode = iota
	Insert
	Command
)

func main() {
	app := tview.NewApplication()
	editor := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	// Initial content and editor state
	content := "Press 'i' to enter Insert mode, ':' for Command mode, and 'ESC' to return to Normal mode.\n"
	editor.SetText(content)
	mode := Normal

	// Capture key events for Vim-like behavior
	editor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch mode {
		case Normal:
			switch event.Rune() {
			case 'i': // Enter Insert Mode
				mode = Insert
			case 'h': // Move cursor left
				// Handle left movement (tview doesn't support precise cursor manipulation directly)
			case 'j': // Move cursor down
				// Handle down movement
			case 'k': // Move cursor up
				// Handle up movement
			case 'l': // Move cursor right
				// Handle right movement
			case ':': // Enter Command Mode
				mode = Command
				content += "\n:" // Display ":" as command prompt
				editor.SetText(content)
			}
		case Insert:
			switch event.Key() {
			case tcell.KeyEsc: // Exit Insert Mode
				mode = Normal
			case tcell.KeyRune: // Add typed character
				content += string(event.Rune())
				editor.SetText(content)
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(content) > 0 {
					content = content[:len(content)-1]
					editor.SetText(content)
				}
			}
		case Command:
			switch event.Key() {
			case tcell.KeyEsc: // Exit Command Mode
				mode = Normal
				content = content[:len(content)-1] // Remove ":"
				editor.SetText(content)
			case tcell.KeyRune:
				content += string(event.Rune())
				editor.SetText(content)
			case tcell.KeyEnter: // Execute Command (e.g., :q to quit)
				if content[len(content)-2:] == ":q" {
					app.Stop()
				}
				mode = Normal
				content += "\n" // Clear command line
				editor.SetText(content)
			}
		}
		return event
	})

	if err := app.SetRoot(editor, true).Run(); err != nil {
		panic(err)
	}
}
