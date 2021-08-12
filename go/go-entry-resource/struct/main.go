package main

import(
	"fmt"
	"reflect"
)

type UseTagStruct struct {
	Id    string   `Language ID`
	Name  string   `Language Name`
	Age   uint	   `Language Age` 
	Address string `Language Address`
} 

type MapKeyStruct struct {
	ID string
	NAME string
}

type PointZ struct {x , y , z int}

// Normal Difinition
type GoStruct struct {	
	x int
	y int
}
// Anonymouis Difinition
type AnonymouisT struct {
	int
	float64
}

// Nest Struct Super
type Animal struct {
	Name string
	Feed Feed
}

type Feed struct {
	Name string
	Amount uint
}

type Point struct {
	X int
	Y int
}

func swap(p Point){
	fmt.Println(p) //ByValue
	x , y := p.Y , p.X
	p.X = x
	p.X = y
}
func swapByRef(p *Point){
	fmt.Println(p)
	x , y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main () {
	st := &GoStruct{
		x : 0xcb,
		y : 0xaacc,
	}
	fmt.Println(st)

	ast := AnonymouisT{0xcced,3.14,}
	fmt.Println(ast)

	animal := &Animal{
		Name : "Dog",
		Feed : Feed{
			Name : "DogFood",
			Amount : 200,
		} ,
	}

	fmt.Println(animal,animal.Name , animal.Feed, animal.Feed.Name, animal.Feed.Amount)
	p := Point{X : 1, Y : 2}
	fmt.Println(p)
	swap(p)
	fmt.Printf("ByValue: %v\n", p)
	swapByRef(&p)
	fmt.Printf("ByRef: %v\n" , p)

	fmt.Println("********Struct Array*****") 	
	ps := make([]PointZ , 5)
	fmt.Println(ps)
	for idx , val := range ps {
		fmt.Println(idx ,  val)
	}
	fmt.Println("************Map Variation*************")
	m := map[MapKeyStruct]int {
		{ ID : "A0001", NAME: "APPLE", } : 500,
	}

	for k , v:=range m {
		fmt.Println(k ,v)
		subk := reflect.TypeOf(k)
		fmt.Println(subk.NumField())
		for i := 0; i < subk.NumField(); i++ {
			skf:=subk.Field(i)
			fmt.Println(skf.Name)
			// Key Param Judgement 
			if skf := subk.Field(i) ; skf.Name == "ID" {
			 	fmt.Println("This is ID Member")
			}
		}
	}
	
	//K=uint V=[]string 		
	mSlice := map[uint][]string {
		1 : {"A" , "B" , "C"},
		2 : {"D" , "E" , "F"},
	}
	fmt.Println(mSlice)

	for k , v:=range mSlice {
		fmt.Println(k , v)
	}
	mm := map[string]map[string]int {
		"A0001" : {"Apple" : 300 },
	}
	fmt.Println(mm)

	t := UseTagStruct {
		Id : "LANG01",
		Name : "Golang",
		Age  : 2009,
		Address : "Golang@gmail.com",
	}
	fmt.Println(t)
	typ := reflect.TypeOf(t)
	fmt.Println(typ , typ.NumField()) // 4	
	// Loop by Struc
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		fmt.Println(f.Name , f.Tag)
	}
}
