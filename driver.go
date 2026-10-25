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
	
	fmt.Println("Executing")
	boxframe := tview.NewBox().
		SetBorder(true).
		SetTitle("Keyboard Tester")
	if err := tview.NewApplication().SetRoot(boxframe, true).Run(); err != nil {
		panic(err)
	}
	
	newbox := boxframe.tview.NewBox().
		SetBorder(true)
}
