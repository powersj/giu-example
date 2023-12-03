package main

import (
	"devdex/ui"

	"github.com/AllenDang/giu"
)

var (
	defaultAppTitle             = "devdex"
	defaultFontSize     float32 = 20
	defaultNavWidth     float32 = 200
	defaultWindowWidth          = 1200
	defaultWindowHeight         = 1000

	contentPanel giu.Layout
)

func loop() {
	giu.SingleWindowWithMenuBar().Layout(
		giu.Style().SetFontSize(defaultFontSize).To(
			ui.MenuBar(defaultAppTitle),
			giu.SplitLayout(giu.DirectionVertical, &defaultNavWidth,
				giu.Layout{
					giu.Button("Hash").OnClick(func() {
						contentPanel = ui.Hash()
					}),
					giu.Button("JSON").OnClick(func() {
						contentPanel = ui.JSON()
					}),
				},
				contentPanel,
			),
		),
		giu.PrepareMsgbox(),
	)
}

func main() {
	wnd := giu.NewMasterWindow(defaultAppTitle, defaultWindowWidth, defaultWindowHeight, 0)
	wnd.Run(loop)
}
