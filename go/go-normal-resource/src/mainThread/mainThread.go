package main

import (
	//	"./channelBuffer"
	"./mthread"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"time"
)

const (
	RUNNING_COUNTER  = 3
	GO_PROCEE_SUFFIX = "go"
)

func main() {
	fmt.Printf("Go Process:%s\n", showGoProcess())
	mthread.MainIsGoRoutine()
	loopProc("This is Called by MainThread")
	go loopProc("This is Called by SubThread")
	showCuurentPidInfo()
	fmt.Println("Done!")

	fmt.Println("Channel Buffer Sample")
	//channelBuffer.ChannelBufferSummary()
}

func loopProc(title string) {
	for j := 1; j <= RUNNING_COUNTER; j++ {
		fmt.Printf("%d:%s\n", j, title)
		time.Sleep(1 * time.Second)
	}
}

func showCuurentPidInfo() {
	pid := os.Getpid()
	pidInfo, _ := ps.FindProcess(pid)
	fmt.Printf("%v\n", pidInfo)
}

func showGoProcess() string {
	processList, err := ps.Processes()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, process := range processList {
		if process.Executable() == GO_PROCEE_SUFFIX {
			return process.Executable()
		}
	}
	return ""

}
