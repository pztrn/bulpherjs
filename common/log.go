package common

import (
	"github.com/gopherjs/gopherjs/js"
)

// Log logs everything passed to Javascript console.
func Log(args ...interface{}) {
	js.Global.Get(JSConsole).Call("log", args...)
}
