## Goのチャネル入門

### 1.そもそもスレッドとは?

```
一連のプログラムの流れである。
```
ちなみにスレッドには大きく分けて2種類のタイプが存在します。

- [x] `シングルスレッド` : 単一のプロセスだけで動作するもの...MainThreadだけで動くプログラ
- [x] `マルチスレッド`   : MainThreadとサブスレッドで動作するもの...`今回のメイン`

この説明だけではイメージがし辛いかもしれないので実際に`Main`スレッドを書いてみます。

```go
package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
)

func main() {
	fmt.Println("This is Main Thread !")
	pid := os.Getpid()
	fmt.Printf("%d", pid)

	pidInfo, _ := ps.FindProcess(pid)
	fmt.Println("%s", pidInfo.Executable())
}
```
上記のプログラムはメインスレッドを使い、コンソールに`This is Main Thread !`を表示させているものになります。
```sh
This is Main Thread !
8670%s mainThread
```
### Gorlangにおけるサブスレッドとは

下記のプログラムは`メインスレッド`と`サブスレっド`が存在します。goでサブスレっドを実現する時は`GoRoutine`と呼ばれるGo特有のオブジェクトを使用します。

 **メインスレッドもGoRoutineで動作している!。**

```go
package mthread

import (
	"fmt"
	"runtime"
)

func MainAlsoGoRoutine() {
	fmt.Printf("Active GoRutine:%d\n", runtime.NumGoroutine())
}

```

コンソールログ

```sh
$ go run  mainThread/mainThread.go
Active GoRutine:1
```

プログラムの流れとして...

メインスレッド内で定義したprint文:`This is Call by Main Thread`が1回実行された後に、goルーチンでサブスレッド化した`firstGoRoutine関数`をコールし、関数内部で定義したfor文が3回実行されます。その後、メインスレッドに戻ってきてこの繰り返しをあと2回行います。

```go
package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
)

const (
	RUNNING_COUNTER = 3
)

func main() {

	for i := 1; i <= RUNNING_COUNTER; i++ {
		fmt.Printf("%d:This is Call by Main Thread\n", i)
		go firstGoRoutine("This is Call by SubThread")
		showCuurentPidInfo()
	}
}

func firstGoRoutine(title string) {
	for j := 1; j <= RUNNING_COUNTER; j++ {
		fmt.Printf("%d:%s\n", j, title)
	}
}

func showCuurentPidInfo() {
	pid := os.Getpid()
	pidInfo, _ := ps.FindProcess(pid)
	fmt.Printf("%v\n", pidInfo)
}


```
**※この例であれば期待した処理が順序を守り実行しているので問題なし**

コンソールログ
```sh
1:This is Call by Main Thread
1:This is Call by SubThread
2:This is Call by SubThread
3:This is Call by SubThread
&{12897 12882 mainThread}
2:This is Call by Main Thread
1:This is Call by SubThread
2:This is Call by SubThread
3:This is Call by SubThread
&{12897 12882 mainThread}
3:This is Call by Main Thread
1:This is Call by SubThread
2:This is Call by SubThread
3:This is Call by SubThread
&{12897 12882 mainThread}
```

ではChannelが必要な例をみていきます。
こんな感じでSleepを1s仕込んでみました。では実際に動かしてみましょう。
期待する結果のフローを下記に示しておきます。

期待するフロー

- [ ] 1.メインスレッドでloopProc関数がコールされ1s毎に`This is Called by MainThread`が3回表示される。

- [ ] 2.1の処理が終了次第、サブスレッドでloopProc関数がコールされ1s毎に`This is Called by SubThread`が3回表示される。

```go
package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"time"
)

const (
	RUNNING_COUNTER = 3
)

func main() {
	loopProc("This is Called by MainThread")
	go loopProc("This is Called by SubThread")
	showCuurentPidInfo()
	fmt.Println("Done!")
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

```
実際のコンソールログ

```sh
1:This is Called by MainThread
2:This is Called by MainThread
3:This is Called by MainThread
1:This is Called by SubThread
&{13630 13616 mainThread}
```
上記のログをみると`SubThread`が一回しか表示されていないことがわかると思います。
つまりは`SubThread`の処理を全て実施する前にプログラムが終了してしまったっということです。
ではどうすれば`SubThread`の処理を期待通り実施できるのでしょうか。

ここでようやく`Channel`の出番です。

#### 2.Channelの定義の実際の使い方

```
Channel は goroutine 間でのメッセージパッシングをするためのもの
```

メッセージパッシングとは?
```
```

Channelの初期化について

```go
変数 := make(chan 型)

/******Chennlの向きについて*****/

//送信専用のChannel
sndCh := make(chan<- int)
//受信専用のChanel
rcvCh := make(<-chan int)
```

### 1.メインスレッドとGoルーチン間でデータの授受を行う入門コード

```go
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
```

### 2. チャネルバッファリングによるChannel制御

チャネル[バッファリング](http://e-words.jp/w/%E3%83%90%E3%83%83%E3%83%95%E3%82%A1%E3%83%AA%E3%83%B3%E3%82%B0.html)とは...

パフォーマンスの向上を図る目的で使用されるケースが多く、`複数のメッセージをテャネルに対して送受信できるようにする機能`

```go
package main

import (
	"fmt"
)

func main() {
	// ここでバッファー付きのチャネルを初期化している。
	messages := make(chan string, 2)
	fmt.Printf("%T", messages)

	messages <- "sendData1"
	messages <- "sendData2"
	messages <- "sendData3" // bufferの領域が不足するためエラーが発生

	receiveMessage1 := <-messages
	receiveMessage2 := <-messages

	fmt.Println(receiveMessage1)
	fmt.Println(receiveMessage2)
}
```

### 3. Goルーチンの同期

最初に`Channel Synchronization`がプリント文として実行される。その後、はgoルーチンとして
worker関数が実行される。worker関数は引数にbool型のchannelを持つため、この値に`true`のステータスを送信することで処理の終了を促すことができる。
```go
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Now Loading.....")
	time.Sleep(time.Second * 10)
	fmt.Println("Done")

	done <- true
}

func main() {

	fmt.Println("Channel Synchronization")
	done := make(chan bool)
	go worker(done)
	<-done
	close(done)
}
```

### 4. select構文にによる複数のチャネルを制御

```go

/* 複数のChannelをselect文で制御し100までの奇数を出力するプログラム
Goルーチンは全体で3台の構成
奇数算出公式: 2n-1*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//ch1から受信した整数を2倍してch2に送信するGoルーチン
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	//ch2から受信した整数から1減算してch3へ送信する。
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()
	n := 1
LOOP:
	for {
		select {
		case ch1 <- n: // nの値がch1に送信されたときに行う処理
			n++
		case i := <-ch3: //ch3からデータの受信を行うとき
			fmt.Println("Received!", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}

	fmt.Println(runtime.NumGoroutine())
}
```

### 5. Goルーチンのタイムアウト制御

`time.After(duration)`関数を使用することにより、特定のGoルーチン内の処理をタイムアウトさせることができます。下記の例は10秒スリープするGoルーチンを定義し、select構文で内部で2秒waiteし、それでもレスポンスを得られない場合はタイムアウトするものです。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	m1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 10)
		m1 <- "Done!"
	}()

	select {
	case passingMessage := <-m1:
		fmt.Println(passingMessage)
	case <-time.After(2 * time.Second):
		fmt.Println("m1 goroutine was.......Time Out......")
	}
}

```
