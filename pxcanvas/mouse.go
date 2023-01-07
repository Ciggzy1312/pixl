package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/Ciggzy1312/pixl/pxcanvas/brush"
)

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	pxCanvas.TryPan(pxCanvas.mouseState.prevcoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.prevcoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                      {}

func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
