package pkgLog

import (
	"log"
	"os"
)

func LogPackageSummary() {
	log.Println("stdout")
	//change file discriptor
	//log.SetOutput(os.Stdout)
	log.Println("stdin")

	//outputlog custom file
	f, err := os.Create("OutputLog.log")
	if err != nil {
		log.Fatal("File Create Error")
	}
	log.SetOutput(f)
	log.Println("This is Log Data")

	//set custom format
	log.SetFlags(log.Ldate | log.Llongfile)
	log.Println("After Cutom My Log")
}
