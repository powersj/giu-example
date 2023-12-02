package main

import (
	"github.com/AllenDang/giu"
)

var (
	navigationWidth float32 = 200
)

func loop() {
	giu.SingleWindowWithMenuBar().Layout(
		giu.Style().SetFontSize(20).To(
			giu.MenuBar().Layout(
				giu.Menu("File").Layout(
					giu.MenuItem("Open").Shortcut("Ctrl+O"),
					giu.MenuItem("Save"),
					// You could add any kind of widget here, not just menu item.
					giu.Menu("Save as ...").Layout(
						giu.MenuItem("Excel file"),
						giu.MenuItem("CSV file"),
						giu.Button("Button inside menu"),
					),
					giu.MenuItem("Exit"),
				),
				giu.MenuItem("Settings"),
				giu.MenuItem("About"),
			),
			giu.SplitLayout(giu.DirectionVertical, &navigationWidth,
				giu.Layout{
					giu.Label("devdex"),
				},
				giu.Layout{
					giu.Label("Right panel"),
				},
			),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("devdex", 1200, 1000, 0)
	wnd.Run(loop)
}
