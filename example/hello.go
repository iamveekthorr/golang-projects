package main

import (
	"fmt"
)

type Adder struct {
	start int
}

type Score struct{}

type HighScore Score

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type LogicProvider struct{}

func (L LogicProvider) Process(data string) string {
	return "Some process"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	c.L.Process("Some long running data!!!")
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}
func main() {
	c := Client{
		L: LogicProvider{},
	}

	c.Program()
	// variables and types
	var name string = "victor"
	var age int = 24
	isMarried := false

	fmt.Println(age, name, isMarried)

	// Functions
	var sum int = addTwoNumbers(2, 3)
	fmt.Println("Sum is: ", sum)

	product, message := multiplyTwoNumbersAndReturnString(2, 3)
	fmt.Println(message, product)

	x := []int{1, 2, 3, 4, 5}
	fmt.Println("x:", x[:2])

	var s string = "Hello 🌞"
	fmt.Println(len(s), "length of the string")

	true := true
	fmt.Println(true)

	// higher order functions / closures
	closure(2)(3, 4, 5)
	fu := closure(2)
	fu(3, 2, 1)

	adderInst := Adder{start: 10}
	f2 := Adder.AddTo

	fmt.Println(f2(adderInst, 20))

	var i int = 200
	var s Score = 100
	var hs HighScore = 300

	s = Score(i)
	hs = HighScore(s)

}

// if a function should return a datatype,
// there is a need to name the type being returned
func addTwoNumbers(a, b int) int {
	return a + b
}

// if a function should return multiple datatypes,
// there is a need to name the type being returned
func multiplyTwoNumbersAndReturnString(a, b int) (int, string) {
	return a + b, "Product is: "
}

func closure(a int) func(args ...int) {
	return func(args ...int) {
		fmt.Printf("%v %v", a, args)
	}
}
