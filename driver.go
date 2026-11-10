// TODO
// Args, allow for text by letter or by word


package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"kasio.xyz/m/v2/textImport"
)


type module struct{
	name string
	textToType string
	primitive tview.Primitive
	gridRef *tview.Grid
}


func moduleKeyInput(application *tview.Application, moduleIn module, sugarModule module, counterIn int, chatIn string, candyModule module, candyInt int){
	
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
				
				if len(slicedText) != 0 {
					loadedText += slicedText[0]
				}
		
				if splitByLetter{
					loadedText += ""
				}else{
					loadedText += " "
				}
		
				if len(slicedText) != 0{
					slicedText = slicedText[1:]
				}
				
				//Cursor logic
				loadedText = strings.ReplaceAll(loadedText, "█", "")
				loadedText += "█"
				
				// Reset the text being written
				if module, err := module.primitive.(*tview.TextView); err {
					module.SetText(loadedText)
				}
			

								
				
				randCount := rand.Intn(100)

				
				if randCount >= 99{
					counter += 1		
					if moduleAutoUpdate(app,sugarModule,counter, textToRead) == 1{
// 						println() // Reaches here 
						textToRead=textImport.ReturnChat()

						textToRead=textImport.ReturnChat()
						counter = 0
						
						
						
					}
// 					candyInt = 0
				}
				
							
				if randCount >= (95){
					// If pulseCandy returns one, select a new int
					if pulseCandy(app,candyModule, candyInt) == 1{
						//TODO: Make this select a random int	
						randInt := rand.Intn(39)
						if candyInt == randInt{
							rand.Intn(39)
						}
						candyInt = randInt
					}
				}
				
				
				
				
				//Return nil is important to prevent default key behavior	
				return nil
		}			
		return event		
	})
	
}


func moduleAutoUpdate(application *tview.Application, moduleIn module, counterIn int, textIn string) int{
	
	module := moduleIn
// 	app := application
	counter := counterIn
	textToRead := textIn	


	
	var slicedText []string	
	slicedText = strings.Split(textToRead,"\n")
	loadedText := ""
	
	

	for i := 0; i <= counter && len(slicedText) > 0; i++{
		loadedText += slicedText[0]
		loadedText += "\n"
		slicedText = slicedText[1:]	
	}
	
		// Detect if the chat has reached the end
	loadedTextCompare := strings.ReplaceAll(loadedText, "\n", "")
	textToReadCompare := strings.ReplaceAll(textToRead, "\n", "")
	
	if loadedTextCompare == textToReadCompare{
		return 1
	}

	
	if module, err := module.primitive.(*tview.TextView); err {
		module.SetText(loadedText)				
	}
	return 0
}

func pulseCandy(application *tview.Application, moduleIn module, currentCandyInt int) int {
	
// 	app := application
	module := moduleIn
	candySliceInt := currentCandyInt
	
	
	// Detect if current text differs from int
// 	if module, err := module.primitive.(*tview.TextView); 	err {
// 		if textImport.ReturnCandy(candySliceInt) != module.GetText(false){
// 			module.SetText("here")
// 	}}

	// Get current candy string
	if module, err := module.primitive.(*tview.TextView); 	err {
		compareText := module.GetText(false)
		currentText := module.GetText(false)
		compareText = strings.ReplaceAll(currentText, ".", "")
		
		if compareText == textImport.ReturnCandy(candySliceInt){
			// String matches current int
		}else{
			module.SetText(textImport.ReturnCandy(candySliceInt))
		}
		
			// Detect if theres more than 6 periods, returns 1 if true
		trailingPeriods := len(currentText) -
	len(strings.TrimRight(currentText, "."))
		
		if trailingPeriods > 6{
			return 1
		}else{
		
			// Add to current string, separated by semicolon
			currentText += "."
			module.SetText(currentText)
			return 0
		}	
		
		
	}

		
	
	

	return 0	
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
		SetText("").
		SetScrollable(false).
		SetDynamicColors(true)
		
	candy := tview.NewTextView().
		SetText(textImport.ReturnCandy(1)).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

		
		
	
	menuBar := tview.NewTextView().
		SetText("Ctrl+C to exit. Any alphabetic key will interact with the program.").
		SetTextAlign(tview.AlignCenter)
	// 	menuBar.SetBorder(true)
		
	
	// Create a grid layout
	grid := tview.NewGrid().
		SetBorders(true).
		SetRows(4, 0, 4).
		SetColumns(30, 0, 30).
			//RowInt, ColumnInt, RowSpan Int, ColSpan Int, MinGrid Height, MinGridWidth, Focus
		AddItem(menuBar, 2, 0, 1, 10, 0, 0, false).
		AddItem(tview.NewGrid(), 0,0,2,10,0,0,false).
			AddItem(textBox, 0, 0, 2, 6, 0, 0, false).
			AddItem(tview.NewGrid(), 1,6,1,4,0,0,false).
				AddItem(candy, 0,6,1,4,0,0,false).
				AddItem(box2, 1,6,1,4,0,0,false)
		
		
	
	// Load primitives into struct
	
	var boxStruct module
	boxStruct.textToType = textImport.ReturnString()
	boxStruct.gridRef = grid
	boxStruct.primitive = textBox
	
	var extraStruct module
	extraStruct.gridRef = grid
	extraStruct.primitive = box2
	
	var candyStruct module
	candyStruct.gridRef = grid
	candyStruct.primitive = candy
	
	candyInt := 1
	
	moduleKeyInput(app,boxStruct,extraStruct,counter, textToRead, candyStruct, candyInt )	
		
	// Set the grid for the application
	app.SetRoot(grid, true)

	// Start the application
	if err := app.Run(); err != nil {
		fmt.Println("Application exited with error:", err)
	}


	// Clean up
// 	screen.Close()
}
