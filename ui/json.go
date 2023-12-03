package ui

import "github.com/AllenDang/giu"

var (
	jsonInput  string
	jsonOutput string
)

func JSON() giu.Layout {
	return giu.Layout{
		giu.Label("JSON Layout"),
		giu.Label("Input"),
		giu.InputText(&jsonInput),
		giu.Label("Output"),
		giu.InputText(&jsonOutput),
	}
}
