package main

import(
	"fmt"
)
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
	fmt.Printf("ByRef: %v" , p)
}
