package main

import (
	"fmt"
)

func main() {

	/*送受信の両方を行うことができるチャネルを定義*/
	passingMessageNeutralCh := make(chan string)
	fmt.Printf("%v", passingMessageNeutralCh)
	go func() { passingMessageNeutralCh <- "sendData1" }()

	rcvMsg1 := <-passingMessageNeutralCh
	fmt.Println(rcvMsg1) // エラーにはならず値を取得し、データの受信も可能

	/*送信専用チャネルを定義*/
	passingMessageSendCh := make(chan<- string)
	fmt.Printf("%v", passingMessageSendCh)
	go func() { passingMessageSendCh <- "sendData2" }()

	// 送信専用のChannelのためデータの受信でエラーが発生する。
	rcvMsg2 := <-passingMessageSendCh
	fmt.Println(rcvMsg2)

	/*受信専用チャネルを定義*/
	passingMessageReceiveCh := make(<-chan string)
	//受信専用のチャネルのためデータの送信でエラーが発生する。
	go func() { passingMessageReceiveCh <- "sendData3" }()
	rcvMsg3 := <-passingMessageReceiveCh
	fmt.Println(rcvMsg3)
}
