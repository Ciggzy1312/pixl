package ui

import (
	"fyne.io/fyne/v2"
	"github.com/Ciggzy1312/pixl/apptype"
	"github.com/Ciggzy1312/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
