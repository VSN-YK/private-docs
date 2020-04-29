package main

import (
	"./pkgMath"
	"./pkgOS"
	"./pkgTime"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	OS   = "os"
	TIME = "time"
	MATH = "math"
)

var (
	pkg = kingpin.Flag("pkg", "Execute [X] Package").String()
)

func main() {
	kingpin.Parse()
	switch *pkg {
	case OS:
		pkgOS.FirstOsPackageSummaryCall()
	case TIME:
		pkgTime.TimePackageSummary()
	case MATH:
		pkgMath.MathPackageSummary()
	//TODO: (time, json , ioutil..etc)
	default:
		fmt.Printf("%s", "Another Package")
	}
}
