package ui

import "github.com/AllenDang/giu"

var (
	hashInput string
)

func Hash() giu.Layout {
	return giu.Layout{
		giu.Label("Hash Layout"),
		giu.Label("Input"),
		giu.InputText(&hashInput),
		giu.Label("MD5"),
		giu.InputText(&jsonOutput),
		giu.Label("SHA256"),
		giu.InputText(&jsonOutput),
		giu.Label("Blake2b512"),
		giu.InputText(&jsonOutput),
	}
}
