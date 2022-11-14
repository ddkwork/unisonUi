package main

import (
	"github.com/richardwilkes/unison"
)

func main() {
	unison.AttachConsole()
	unison.Start(unison.StartupFinishedCallback(func() { CanvasObject(unison.PrimaryDisplay().Usable.Point) }))
}
