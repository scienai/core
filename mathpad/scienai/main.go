package main

import (
	"fmt"
	"image"
	"cogentcore.org/core/core"
)

var bb=[]int{1:3,3:4,8:9}

func a() int {
	return 33
}

var dd,ee=[]int{23432,32,4,321,432,4},[]int{3241,4,32,4,3214,4}
var lla = []int{334}
var kdskf=[]map[int]int{{3 : 34}}
type sdk struct {
	a int
}

var KKAA []  []     []int

var ujf=map[int]int{5*6:833,23-3243/324:834}
var kdf=[][]*image.Point{}

func main() {
	mainwin := core.NewBody("Scienai")
	mathpad := core.NewMathpad(mainwin)
	mathpad.SetTooltip("mathpad")
	fmt.P("bb",bb,a(),dd,ee,lla,kdskf,&sdk{333}, [] [] int{{3,343}}, KKAA,kdf)
	mainwin.RunMainWindow()
}
