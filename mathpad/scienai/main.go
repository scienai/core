package main

import (
	"cogentcore.org/core/core"
)

func main() {
	mainwin := core.NewBody("Scienai")
	mathpad := core.NewMathpad(mainwin)
	mathpad.SetTooltip("mathpad")
	mainwin.RunMainWindow()
}
