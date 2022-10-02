package main

import (
	"github.com/ddkwork/unisonUi/packet"
	"github.com/richardwilkes/unison"
)

func main() {
	unison.AttachConsole()
	unison.Start(unison.StartupFinishedCallback(func() { packet.CanvasObject(unison.PrimaryDisplay().Usable.Point) }))
}
