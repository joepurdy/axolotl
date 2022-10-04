package main

import (
	"os"

	"github.com/joepurdy/awswitch/cli"

	"github.com/alecthomas/kingpin"
)

// Version is provided at compile time
var Version = "dev"

func main() {
	app := kingpin.New("awswitch", "A helper utility for switching AWS profiles in subshells.")
	app.Version(Version)

	a := cli.ConfigureGlobals(app)
	cli.ConfigureExecCommand(app, a)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
