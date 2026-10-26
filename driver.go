// Written in go, duh. Can probably fit everything in one file, or maybe a driver and a library of modules to load.

package main

import (
	"fmt"
	"time"
	"github.com/rivo/tview"
)

func redraw(app *tview.Application){
	
	// Redraw the app once a second
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		app.Draw()
	}
}


func main() {
	
	// Grab variables
	now := time.Now().Format(time.RFC3339)
	
	// Create a new application
	app := tview.NewApplication()

	// Create a new screen
// 	screen := tview.NewScreen()

	//Initialize modules for grid
	emptyBox := tview.NewTextView().
// 		SetDynamic(true).
		SetText(now)

	menuBar := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Menu bar down here!")
	
	// Create a grid layout
	grid := tview.NewGrid().SetBorders(true).
		SetRows(6).
		SetColumns(6)

	// Add the main content window to the grid
	grid.AddItem(emptyBox, 0, 0, 2, 3, 1, 3, false)
	grid.AddItem(menuBar, 2, 0, 1, 3, 0, 0, false)
	
	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the redraw ticker before running the app
	go redraw(app)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}

	// Clean up
// 	screen.Close()
}
