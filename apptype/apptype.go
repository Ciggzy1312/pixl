package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int // No of pixels in each row and column
	PxSize         int // Size of each pixel
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int    // Index of the selected swatch
	FilePath       string // Path of the file to be opened
}

type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
