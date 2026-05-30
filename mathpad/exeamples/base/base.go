package main

import (
	"cogentcore.org/core/core"
)

func main() {
	mainwin := core.NewBody("Mathpad")
	mathpad := core.NewMathpad(mainwin)
	mathpad.SetTooltip("mathpad")
	mainwin.RunMainWindow()
}
