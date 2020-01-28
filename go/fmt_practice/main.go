package main

import(
  "fmt"
)
// TypeStruct is a maltipe struct
type TypeStruct struct {
  A string
  B int
  C float64
  D map[string]int
  F []string
}

func main (){

  var(
    a = "A"
    l = map[string]map[string]int{
      "Language":{"Golang":2009},
    }
  )

  fmt.Println(a,l)

  var st TypeStruct
  st.A = "A"
  st.B =10
  st.C = 3.14
  st.D = map[string]int{"Golang" : 2009}
  st.F = []string{"A","B", "C"}
  fmt.Printf("Binary\t%b\nOctal\t%o\nDecimal\t%d\nHex\t%X\nString\t%s\nStruct()\t%v" ,  10,10,10,10,"Golang",st,
  )
  resolveType(10,"Golang",st,2.96032)
}

func resolveType(types...interface{}){
  for _ , v := range types{
    fmt.Printf("%T\n", v)
  }
}
