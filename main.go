package main

import (
	"fmt"
	"fyne.io/fyne"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("test", func() {
			fmt.Println("test")
		}),
		widget.NewButton("panic", func() {
			panic("panic")
		}),
		//widget.NewButton("Quit", func() {
		//	a.Quit()
		//}),
	))

	w.Resize(fyne.Size{
		Width:800,
		Height:400,
	})

	w.ShowAndRun()
}
