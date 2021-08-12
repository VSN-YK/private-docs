package main

import (
	"fmt"
	"time"
	"github.com/comail/colog"
	"log"
   // "runtime"
   // "channel"
)
func setLogCondition () {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag: log.Ldate | log.Ltime | log.Lshortfile,		
	})
	colog.Register()
}
func receive(name string , ch <-chan int){
	st := time.Now()
	for {
		//check channel
		i , ok := <-ch
		if ok == false {
			break
		}
		ed := time.Now()
		log.Printf("info: %s\t%d\t%fms\n", name , i , (ed.Sub(st)).Seconds() *1000)
		
	}
	fmt.Println(name + "is done.")
}

func main (){
	setLogCondition()
	ch := make(chan int , 10)
	go receive("1st Goroutine", ch)
	go receive("2nd Goroutine", ch)
	go receive("3rd Goroutine", ch)
	
	for i:= 0;  i < 100; i++ {
		ch <-i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}
	/*channel.ChannelPractice()
    fmt.Println(runtime.NumGoroutine())
    ch := make(chan int, 5)
	send channel
    ch <- 5 
    // Recieved Channel
    rcv := <-ch
	fmt.Println(rcv)
	
	ch := make(chan int)
	go Reciever(ch)
	
	i := 0
	for i < 10 {
		ch <- i
		i++
	}
	ch :=make(chan string, 3)
	ch <- "Apple"
	ch <- "Orange"
	close(ch)  
    ch <- "Banana" //restrict of Send Channel
	for i := range ch {
		 _ , ok := <-ch
		fmt.Println(i,ok)
		 if ok == false {
			fmt.Println(ok)
		 }
	}
	fmt.Printf("%d" , len(ch))
 	
	
}*/

