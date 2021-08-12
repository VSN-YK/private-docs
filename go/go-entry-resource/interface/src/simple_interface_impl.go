
/*************************************************************************
interfaceの最もpopullerな特性を表わしたものである。
interfaceは特定の型に依存しないという性格を持っているため、異なるオブジェクトに
共通の性質を与えたいときによく使われる。

下記のコードは全く別のオブジェクトであるが、どのOSもBoot(起動)という振る舞いを
保有しているため、interfaceを用い汎用的に実装したものである。
/************************************************************************/

package main

import (
  "fmt"
)

type ComputerAction interface{
  boot()
}

type MacOS struct {}
type WindowsOS struct{}
type LinuxOS struct{}

// implementaion of Mac
func (m *MacOS) boot(){
  fmt.Println("Mac")
}

// implementaion of Windows
func (w *WindowsOS) boot(){
  fmt.Println("Windows")
}

// implementaion of Linux
func (l *LinuxOS) boot(){
  fmt.Println("Linux")
}

// ComputerAction : interface
func DoingComputerAction(ca ComputerAction){
  ca.boot()
}

func main (){
  //definate on computerAction Struct of Array
  cpAction :=[]ComputerAction{
    &MacOS{},
    &WindowsOS{},
    &LinuxOS{},
  }

  fmt.Println(&cpAction)


  for _ , action:= range cpAction {
    fmt.Println(action , &action)
    DoingComputerAction(action)
  }

}

/*************************Console Log******************************
$ go run sample.go
&[0x593c18 0x593c18 0x593c18] // interface
&{} 0xc00004c1f0
Mac
&{} 0xc00004c1f0
Windows
&{} 0xc00004c1f0
Linux
***************************************************************/
