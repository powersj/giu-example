package ui

import (
	"os"

	"github.com/AllenDang/giu"
)

func MenuBar(defaultAppTitle string) *giu.MenuBarWidget {
	return giu.MenuBar().Layout(
		giu.MenuItem(defaultAppTitle),
		giu.MenuItem("Settings").OnClick(settings),
		giu.MenuItem("About").OnClick(about),
		giu.MenuItem("Exit").OnClick(exit),
	)
}

func exit() {
	os.Exit(0)
}

func settings() {
	giu.Msgbox("Settings", "TODO settings box here")
}

func about() {
	giu.Msgbox("About", "devdex v0.0.1")
}
