// IDEA!!! Keyboard tester with a tui, throws a bunch of screens to interact with but mostly its just for typing.
// Maybe an actual keyboard test mode?? Scancodes/ASCII
// or a screensaver mode?

//Written in go, duh. Can probably fit everything in one file, or maybe a driver and a library of modules to load.

//AI! Import tview, and write a basic tui program. Dont need it to do anhything just display a screen.

package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	// Create a new application
	app := tview.NewApplication()

	// Create a new screen
	screen := tview.NewScreen()

	// Create a grid layout
	grid := tview.NewGrid().SetBorders(true)

	// Add the main content window to the grid
	mainWindow := tview.NewTextView().
		SetBorder(true).
		SetTitle("Main Content")
	grid.SetCell(0, 0, mainWindow, tview.Grid{}.DefaultSize())

	// Add the keybindings row to the grid
	keybindingsRow := tview.NewTextView().
		SetBorder(true).
		SetTitle("Keybindings")
	grid.SetCell(1, 0, keybindingsRow, tview.Grid{}.DefaultSize())

	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}

	// Clean up
	screen.Close()
}
