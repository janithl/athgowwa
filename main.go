package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeScrollTab() fyne.Widget {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(320, 320))
	list := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
	)
	for i := 1; i <= 20; i++ {
		index := i // capture
		list.Append(widget.NewButton(fmt.Sprintf("Button %d", index), func() {
			fmt.Println("Tapped", index)
		}))
	}

	scroll := widget.NewScrollContainer(list)
	scroll.Resize(fyne.NewSize(200, 200))

	return scroll
}

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(makeScrollTab())
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
