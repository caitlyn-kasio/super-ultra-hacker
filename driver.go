// Written in go, duh. Can probably fit everything in one file, or maybe a driver and a library of modules to load.

package main

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
)

	//AI! Fix this please, the appIn is undefined. This is an object used by main() passed to this function at runtime
func updateTime(appIn){
	app := appIn
	for{
		time.Sleep(1)
		app.QueueUpdateDraw(func(){
			app.emptyBox.SetText("wow look!!!")
		})
	}
}

func main() {	
	
	// Grab variables
// 	time := 0
	
	// Create a new application
	app := tview.NewApplication()

	// Create a new screen
// 	screen := tview.NewScreen()

	//Initialize modules for grid
	emptyBox := tview.NewTextView().
// 		SetTextAlign(tview.AlignCenter.
		SetText("skdhgskhdgkhsdgkdsg")

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

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}

	go updateTime(app)

	// Clean up
// 	screen.Close()
}
