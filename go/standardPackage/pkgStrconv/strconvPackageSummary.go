package pkgStrconv

import (
	"fmt"
	"log"
	"strconv"
)

const (
	BINARY_BASE = 2
	HEX_BASE    = 16
)

func StrconvPackageSummary() {
	//signner string convert type in boolen
	b := "true"
	if s, err := strconv.ParseBool(b); err == nil {
		fmt.Printf("%T\n", s)
	}
	//convert intger to string
	var buff string
	if resultBuffer, err := fmt.Scan(&buff); err == nil {
		log.Printf("[before]%T\n[After]%T\n", resultBuffer, strconv.Itoa(resultBuffer))
	}

	var inputHundred int64 = 100
	fmt.Printf("%s\n%s\n", strconv.FormatInt(inputHundred, BINARY_BASE), strconv.FormatInt(inputHundred, HEX_BASE))
}
