//So from here.
// Need to move text input/word writing logic to a separate function and add error handling. We want this to import text and text styling, so creating a class and instaciating at runtime is the best option.
//Figure out why vim navigation is still working and stop it
//Start writing modules to load?
//Polish and style

package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	// 	"time"
	"kasio.xyz/m/v2/textImport"
)


type module struct{
	name string
	textToType string
	primitive tview.Primitive
	gridRef *tview.Grid
}


func moduleKeyInput(application *tview.Application, moduleIn module, sugarModule module, counterIn int, chatIn string){
	
	module := moduleIn
	app := application
	splitByLetter := true
	counter := counterIn
	textToRead := chatIn
	
	
	textToBreak := module.textToType
	    var slicedText []string
	    if splitByLetter {
	        for _, char := range textToBreak {
	            slicedText = append(slicedText, string(char))
	        }
	    } else {
	        slicedText = strings.Split(textToBreak, " ")
	    }
	    loadedText := ""
	
	// Handles key capture
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey{
		switch event.Rune(){
				//Alphas get captured to simulate typing
			case 'a','b','c','d','e','f','g','i','m','n','o','p','q','r','s','t','u','v','w','x','y','z','h','j','k','l':
				// Add one word to the string at a time, so it simulates typing. Maybe there should also be a per letter setting?
			
				// Activates when there is no more text to type.
				if len(slicedText) == 0 {
						
					// Clear the text and reset so more text can appear
					if module, err := module.primitive.(*tview.TextView); 	err {
						module.SetText("End of sliced text")
						loadedText = ""
						module.SetText("loadedText")						
						textToBreak := textImport.ReturnString()
						module.SetText("textToBreak")
	    					
	    					if splitByLetter {
						        for _, char := range textToBreak {
						            slicedText = append(slicedText, string(char))
						        }
						}else {
						        slicedText = strings.Split(textToBreak, " ")}	
					}
				}
				loadedText += slicedText[0]
		
				if splitByLetter{
					loadedText += ""
				}else{
					loadedText += " "
				}
		
				slicedText = slicedText[1:]
				
				//Cursor logic
				loadedText = strings.ReplaceAll(loadedText, "█", "")
				loadedText += "█"
				
				// Reset the text being written
				if module, err := module.primitive.(*tview.TextView); err {
					module.SetText(loadedText)
				}
			
				
				
				// Count keystrokes, and on a good number of keystrokes apply an update to the sugar module				
				
				randCount := rand.Intn(100)


				if randCount >= 99{
					counter += 1		
					moduleAutoUpdate(app,sugarModule, counter, textToRead)
				}
				
				
				
				
				//Return nil is important to prevent default key behavior	
				return nil
		}			
		return event		
	})
	
}


func moduleAutoUpdate(application *tview.Application, moduleIn module, counterIn int, textIn string){
	
	module := moduleIn
// 	app := application
	counter := counterIn
	textToRead := textIn	


	
	var slicedText []string	
	slicedText = strings.Split(textToRead,"\n")
	loadedText := ""
	
	//AI! Add error handling to this specific for loop. Currently it crashes when slicedText runs out. Have it continue the program please
	for i := 0; i <= counter && len(slicedText) > 0; i++{
		loadedText += slicedText[0]
		loadedText += "\n"
		slicedText = slicedText[1:]	
	}

	if module, err := module.primitive.(*tview.TextView); err {
		module.SetText(loadedText)				
	}
}



func main() {
	
	counter := 0
	textToRead := textImport.ReturnChat()
	
	moduleArray := []module{}
	fmt.Print(moduleArray)
	
	// Create a new application
	app := tview.NewApplication()

	// Create a new screen

	//Initialize modules for grid
	textBox := tview.NewTextView().
		SetText("Start Typing!").
		SetScrollable(false).
		SetDynamicColors(true)
		
		
	box2 := tview.NewTextView().
		SetText("Chat Module").
		SetScrollable(false).
		SetDynamicColors(true)
		
		
	
	menuBar := tview.NewTextView().
		SetText("Ctrl+C to exit.")
	// 	menuBar.SetBorder(true)
		
	
	// Create a grid layout
	grid := tview.NewGrid().
		SetBorders(true).
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
			//RowInt, ColumnInt, RowSpan Int, ColSpan Int, MinGrid Height, MinGridWidth, Focus
		AddItem(menuBar, 2, 0, 1, 10, 0, 0, false).
		AddItem(tview.NewGrid(), 0,0,2,10,0,0,false).
			AddItem(textBox, 0, 0, 2, 4, 0, 0, false).
			AddItem(box2, 0,4,2,6,0,0,false)
		
		
	
	// Load primitives into struct
	
	var boxStruct module
	boxStruct.textToType = textImport.ReturnString()
	boxStruct.gridRef = grid
	boxStruct.primitive = textBox
	
	var extraStruct module
	extraStruct.gridRef = grid
	extraStruct.primitive = box2
	
	moduleKeyInput(app,boxStruct,extraStruct,counter, textToRead)	
		
	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}


	// Clean up
// 	screen.Close()
}
