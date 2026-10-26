// Written in go, duh. Can probably fit everything in one file, or maybe a driver and a library of modules to load.

package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"unicode"
	"strings"
)


func main() {	
	// Grab variables
// 	time := 0
	
	// Create a new application
	app := tview.NewApplication()

	// Create a new screen
// 	screen := tview.NewScreen()

	//Initialize modules for grid
	emptyBox := tview.NewTextView().
// 		SetTextAlign(tview.AlignCenter).
		SetText("skdhgskhdgkhsdgkdsg")		
		
// 	emptyBox.SetBorder(true)
		
	
	menuBar := tview.NewTextView().
	// 	menuBar.SetBorder(true)
		SetTitle("Keybindings")
		
	
	// Create a grid layout
	//AI! Move the menubar to the bottom of the screen, with only about 10% of screen space, and give the rest to emptyBox
	grid := tview.NewGrid().
		SetBorders(true).
		AddItem(menuBar, 1, 0, 1, 1, 0, 0, false).
		AddItem(emptyBox, 3, 0, 3, 1, 0, 0, false)
		
		
		
	// Testing text writing as was hoped.
	textToBreak := "Wow look heres a long string of text look at all the words I can load in\n\n\n wow neat look new lines wheeeeeee"
	slicedText := strings.Split(textToBreak, " ")
	fmt.Print(slicedText)
	loadedText := ""
	
	
	// Unbind default keys
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey{
		switch event.Rune(){
			case 'j':
				return tcell.NewEventKey(tcell.KeyRune, 'j', tcell.ModNone)
			case 'k':
				return tcell.NewEventKey(tcell.KeyRune, 'k', tcell.ModNone)
		}			
		return event		
	})
	
	
	// Handles inputs
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune{
			key := event.Rune()
			if unicode.IsLetter(key) || unicode.IsDigit(key) {
						
				
			// Add one word to the string at a time, so it simulates typing. Maybe there should also be a per letter setting?
			loadedText += slicedText[0]
			loadedText += " "
			slicedText = slicedText[1:]							
			// Load the array into the textbox
			emptyBox.SetText(loadedText)
				
			}
		}
			
		return event
	})
	
	
		
	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}


	// Clean up
// 	screen.Close()
}
