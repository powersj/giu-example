package ui

import "github.com/AllenDang/giu"

var splitWidth float32 = 500

func JSON() giu.Layout {
	jsonInput := giu.CodeEditor().
		ShowWhitespaces(false).
		TabSize(2).
		Text("select * from greeting\nwhere date > current_timestamp\norder by date").
		Border(true)

	jsonOutput := giu.CodeEditor().
		ShowWhitespaces(false).
		TabSize(2).
		Text("").
		Border(true)

	return giu.Layout{
		giu.Label("JSON Layout"),
		giu.Button("Format").OnClick(func() {
			jsonOutput.Text(jsonInput.GetText())
		}),
		giu.SplitLayout(giu.DirectionVertical, &splitWidth,
			giu.Layout{
				giu.Label("Input"),
				jsonInput.Text("input"),
			},
			giu.Layout{
				giu.Label("Output"),
				jsonOutput.Text("output"),
			},
		),
	}
}
