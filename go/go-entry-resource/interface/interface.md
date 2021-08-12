# Golang Interfaceの思想設計を紐解く

### Accept interface Return Structを紐解く
- [ ] step0 interfaceの性質について理解する
- どんな型でも吸収することができる。

  ```go
  package main

  import (
  	"fmt"
  	"reflect"
  )

  func main() {
  	var itf interface{}

  	PutObject("ABC", 0xccff, 3.1415, itf, map[string]int{"Apple": 200},nil)

  }

  func PutObject(objects ...interface{}) {
  	for _, obj := range objects {
  		DebugPrint(obj)
  	}
  }

  func DebugPrint(obj interface{}) {
  	fmt.Printf("Type=%v\tValue=%v\n", reflect.TypeOf(obj), obj)
  }
  ```
  Console
  ```
  Type=string	Value=ABC
  Type=int	Value=52479
  Type=float64	Value=3.1415
  Type=<nil>	Value=<nil>
  Type=map[string]int	Value=map[Apple:200]
  Type=<nil>	Value=<nil>
  ```
- 構造体が振る舞い（メソッド）を持つことを保障するためのもの
```
sttuct T は interface Iに依存する
```
つまり構造体Tが以下のメンバを持つ場合は必ず
interfaceでメンバーの振る舞いを担保する必要がある。
  ```go
  type T struct {
    a string
    b int
    c float64
  }

  type TI interface {
    setFeild(x string , y int , z float64)
  }

  func(t * T) setFeild(x string , y int , z float64){
    t.a = x
    t.b = y
    t.c = z
  }
  ```


- [ ] step1. golangのinterfaceの制約について理解する
```go
Go の interface の定義の内容は、単なるメソッドリストである。
リスト中のメソッドを全て実装していれば、その Interface を満たす(satisfy) とみなす
```
下記の例は Accessor interfaceでは `GetContext()` , `SetContext()`メソッドのプロトタイプを宣言
しているので実装が強要される。
この規則に反するとコンパイラに指摘される。
```go
  package main

  import (
    "fmt"
  )

  type Accessor interface {
      GetContext() string
      SetContext(string , string , string)
  }

  type Context struct {
      header string
      body string
      footer string
  }

  /*************must implementation*******************/
  func (c *Context) GetContext() string {
    return c.header + "\n" + c.body + "\n" + c.footer
  }

  func (c *Context) SetContext( h , b ,f string) {
    c.header = h
    c.body = b
    c.footer = f
  }
  /**************************************************/
```

- [ ] step2. duck typingについての理解(型を保障できない可能性があるので推奨されない)

  アヒルのように鳴くのであればそれはアヒルであると限定できるようにように
  あるinterfaceに存在する振る舞いでオブジェクトを判断するタイピング手法のこと

  ```go
  type Human interface {
    LikeHumanSpeak() string
  }
  type Monkey interface {
    LikeMonkeySpeak() string
  }

  type SpeakFormat struct {
      soundOfHumanSpeak string
      soundOfMonkeySpeak string
  }

  func (sf *SpeakFormat) LikeHumanSpeak () string{
    return sf.soundOfHumanSpeak
  }
  func (sf *SpeakFormat) LikeMonkeySpeak() string{
    return sf.soundOfMonkeySpeak
  }

  func LetHumanSpeak(h Human) string{
    return h.LikeHumanSpeak()
  }
  func LetMonkeySpeak(m Monkey) string {
    return m.LikeMonkeySpeak()
  }

  func main () {
    speakFormat := &SpeakFormat{
      "Hello",
      "Uki-",
    }
    fmt.Printf("%+v\n" , speakFormat)
    //fource to only human
    FourcetHumanSpeak(speakFormat)
    FourceMonkeySpeak(speakFormat)

  }
  ```
- [ ] Accept interface Return Struct

  https://qiita.com/weloan/items/de3b1bcabd329ec61709

## Appendix

- 構造体の初期化についての定石 コンストラクタ設ける

  ```go
  func newStructure (x string , y int) *Structure {
    s := new(Structure)
    s.x = x
    s.y = y
    return s
  }
  //こちらのほうが一般的
  func newStructure (x string , y int) *Structure {
      return &Structure{
        soundOfHumanSpeak : x,
        soundOfMonkeySpeak : y,
      }
  }
  ```
