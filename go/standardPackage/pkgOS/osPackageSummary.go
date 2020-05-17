package pkgOS

import (
	"../consoleLogger"
	"fmt"
	"log"
	"os"
	"strings"
)

func FirstOsPackageSummaryCall() {
	consoleLogger.CustomConsoleLog()
	log.Printf("info: os.Hostname()")
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("[HostName]:[%s]\nOS package Called by  main Entry\n", host)

	log.Printf("info: File Operation")
	var fileName string
	fmt.Scan(&fileName)
	log.Printf("debug: %s", fileName)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("info: os.Stat()")
		fInfo, err := f.Stat()

		if fInfo.IsDir() {
			fmt.Println("is Dir")
		} else {
			fmt.Println("is File")
		}
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Printf("%s\n", fInfo.Name())
		}
	}
	log.Printf("info: os.Getwd()")
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(pwd)
	}
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	files, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err.Error())
	}

	var buffer string
	log.Printf("info: os.Environ()")
	envList := os.Environ()
	for _, env := range envList {
		buffer += env + "\n"
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			f, err := os.OpenFile(file.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer f.Close()
			readBuffer := []byte(buffer)
			_, err = f.Write(readBuffer)
			if err != nil {
				panic(err)
			}
		}
	}
}
