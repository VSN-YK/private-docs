package main

import (
  "fmt"
  "./packageA"
  "./packageB"
)

func main (){
  fmt.Printf("%s,%s", packageA.DispPackageA(),packageB.DispPackageB())
}
