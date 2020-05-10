package main

import (
	"./pkgJson"
	"./pkgLog"
	"./pkgMath"
	"./pkgOS"
	"./pkgRegexp"
	"./pkgStrconv"
	"./pkgTime"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	OS      = "os"
	TIME    = "time"
	MATH    = "math"
	LOG     = "log"
	STRCONV = "strconv"
	REGEXP  = "regexp"
	JSON    = "json"
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
	case LOG:
		pkgLog.LogPackageSummary()
	case STRCONV:
		pkgStrconv.StrconvPackageSummary()
	case REGEXP:
		pkgRegexp.RegexpPackageSummary()
	case JSON:
		pkgJson.JsonPackageSummary()
	//TODO: (time, json , ioutil..etc)
	default:
		fmt.Printf("%s", "Another Package")
	}
}
