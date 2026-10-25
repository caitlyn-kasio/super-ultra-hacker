// IDEA!!! Keyboard tester with a tui, throws a bunch of screens to interact with but mostly its just for typing.
// Maybe an actual keyboard test mode?? Scancodes/ASCII
// or a screensaver mode?

//Written in go, duh. Can probably fit everything in one file, or maybe a driver and a library of modules to load.

//AI! Import tview, and write a basic tui program. Dont need it to do anhything just display a screen.

package main

import (
	"fmt"
// 	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	// Create a new application
	app := tview.NewApplication()

	// Create a new screen
// 	screen := tview.NewScreen()

	//Initialize modules for grid
	emptyBox := tview.NewTextView().
		SetText("skudgfskdjghskjdg")
	menuBar := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Menu bar down here!")
	
	// Create a grid layout
	grid := tview.NewGrid().SetBorders(true)

	// Add the main content window to the grid
	grid.AddItem(emptyBox, 0, 0, 2, 3, 1, 3, false)
	grid.AddItem(menuBar, 2, 0, 1, 3, 0, 0, false)
	
	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}

	// Clean up
// 	screen.Close()
}
