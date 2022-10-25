package main

import (
	_ "embed"
	"os"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/example/demo"
)

func main() {
	cmdline.AppName = "Example"
	cmdline.AppCmdName = "example"
	cmdline.CopyrightStartYear = "2021"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.AppIdentifier = "com.trollworks.unison.example"

	unison.AttachConsole()

	cl := cmdline.New(true)
	cl.Parse(os.Args[1:])

	unison.Start(unison.StartupFinishedCallback(func() {
		_, err := demo.NewDemoWindow(unison.PrimaryDisplay().Usable.Point)
		jot.FatalIfErr(err)
	})) // Never returns
}
