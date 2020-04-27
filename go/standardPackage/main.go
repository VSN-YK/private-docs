package main

import (
	"./pkgOS"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	OS = "os"
)

var (
	pkg = kingpin.Flag("pkg", "Execute OS Package").String()
)

func main() {
	kingpin.Parse()
	switch *pkg {
	case OS:
		pkgOS.FirstOsPackageSummaryCall()
	//TODO: (time, json , ioutil..etc)
	default:
		fmt.Printf("%s", "Another Package")
	}
}
