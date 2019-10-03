package main

import (
	"github.com/andlabs/ui"
)

func setupUI() {
	mainwin := ui.NewWindow("Athgowwa", 360, 640, true)

	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
