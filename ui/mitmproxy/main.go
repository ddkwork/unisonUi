package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/unisonUi/asserts"
	"github.com/ddkwork/unisonUi/ui/mitmproxy/table"
	"github.com/richardwilkes/unison"
)

// todo add like gcs support arg shell run for linux

func main() {
	unison.AttachConsole()
	unison.Start(unison.StartupFinishedCallback(func() {
		w, err := unison.NewWindow(fmt.Sprintf("mitmproxy"))
		if err != nil {
			return
		}
		w.MinMaxContentSizeCallback = func() (min, max unison.Size) {
			return unison.NewSize(1000, 600), unison.NewSize(10000, 1280)
		}
		image, err := unison.NewImageFromBytes(asserts.MitmBuf, 0.5)
		if !mylog.Error(err) {
			return
		}
		w.SetTitleIcons([]*unison.Image{image})
		table.CanvasObject(w)
		w.Pack()
		rect := w.FrameRect()
		rect.Point = unison.PrimaryDisplay().Usable.Point
		w.SetFrameRect(rect)
		w.ToFront()
	}))
}
