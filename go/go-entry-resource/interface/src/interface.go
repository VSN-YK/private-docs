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

func FourcetHumanSpeak(h Human) string{
  return h.LikeHumanSpeak()
}
func FourceMonkeySpeak(m Monkey) string {
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
  //fource to only Monkey
  FourcetHumanSpeak(speakFormat)

  fmt.Printf("%v", newSpeakFormat())

}


func newSpeakFormat() *SpeakFormat {
  return &SpeakFormat{
    soundOfHumanSpeak : "A",
    soundOfMonkeySpeak : "B",
  }
}



new
