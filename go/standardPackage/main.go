package main

import (
	"./pkgOS"
	"./pkgTime"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	OS   = "os"
	TIME = "time"
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
	//TODO: (time, json , ioutil..etc)
	default:
		fmt.Printf("%s", "Another Package")
	}
}
