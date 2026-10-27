//So from here.
// Need to move text input/word writing logic to a separate function and add error handling. We want this to import text and text styling, so creating a class and instaciating at runtime is the best option.
//Figure out why vim navigation is still working and stop it
//Start writing modules to load?
//Polish and style

package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
	"kasio.xyz/m/v2/textImport"
)


type module struct{
	name string
	textToType string
	primitive tview.Primitive
	gridRef *tview.Grid
}


func moduleKeyInput(application *tview.Application, moduleIn module){
	
	module := moduleIn
	app := application
	splitByLetter := false
	
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
						
					// Clear the text to make room for more.
					if module, err := module.primitive.(*tview.TextView); 	err {
						module.SetText("")
					}
					
					return nil
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
				
				// This mutates the primitive type used by the module struct, allowing the SetText function		
				if module, err := module.primitive.(*tview.TextView); err {
// 					fmt.Println("boxStuct sucessfully set to TextView

					module.SetText(loadedText)
				}
			
				//Return nil is important to prevent default key behavior	
				return nil
		}			
		return event		
	})
	
}



func main() {
		
// 	textIn := textImport.ReturnString()
// 	fmt.Println(textIn)
	
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
		SetText("Wow look box2!").
		SetScrollable(false)
		
		
	
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
	
	moduleKeyInput(app,boxStruct)
	
		
	// Testing text writing as was hoped.
// 	textToBreak := "Wow look heres a long string of text look at all the words I can load in\n\n\n wow neat look new lines wheeeeeee"
// 	slicedText := strings.Split(textToBreak, " ")
// 	fmt.Print(slicedText)
// 	loadedText := ""
	
	// Handles key capture
// 	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey{
// 		switch event.Rune(){
// 				//Alphas get captured to simulate typing
// 			case 'a','b','c','d','e','f','g','i','m','n','o','p','q','r','s','t','u','v','w','x','y','z','h','j','k','l':
// 				// Add one word to the string at a time, so it simulates typing. Maybe there should also be a per letter setting?
// 				
// 				loadedText += slicedText[0]
// 				loadedText += " "
// 				slicedText = slicedText[1:]							
// 				// Load the array into the textbox
// 				emptyBox.SetText(loadedText)
// 			
// 				//Return nil is important to prevent default key behavior	
// 				return nil
// 		}			
// 		return event		
// 	})
	
	
	
	
		
	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}


	// Clean up
// 	screen.Close()
}
