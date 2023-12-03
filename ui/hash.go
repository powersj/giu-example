package ui

import (
	"devdex"

	"github.com/AllenDang/giu"
)

var (
	hashInput      string
	hashMD5        string
	hashSHA256     string
	hashBlake2b512 string
)

func Hash() giu.Layout {
	return giu.Layout{
		giu.Label("Hash Layout"),
		giu.Label("Input"),
		giu.InputText(&hashInput).OnChange(func() {
			hashMD5 = devdex.MD5Sum(hashInput)
			hashSHA256 = devdex.SHA256Sum(hashInput)
			hashBlake2b512 = devdex.Blake2b512Sum(hashInput)
		}),
		giu.Label("MD5"),
		giu.InputText(&hashMD5),
		giu.Label("SHA256"),
		giu.InputText(&hashSHA256),
		giu.Label("Blake2b512"),
		giu.InputText(&hashBlake2b512),
	}
}
