package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main() {
	var mtx sync.Mutex
	fmt.Printf("%+v", mtx)
	c := make(chan bool)

	for i := 0; i < 5; i++ {
		mtx.Lock()
		fmt.Printf("%+v", mtx)
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
			mtx.Unlock()
			c <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
	}
}

func showStructObjField(obj interface{}) {
	typ := reflect.TypeOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		fmt.Println(f.Name)
	}
}
