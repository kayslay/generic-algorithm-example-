package main

import (
	"flag"
	"fmt"
	"learn/generic/generic"
)

var popNum int
var str string
var mutation float64

func init() {
	flag.IntVar(&popNum, "pop-num", 20, "the poplation of the system")
	flag.IntVar(&popNum, "p", 20, "the poplation of the system (shorthand)")
	flag.StringVar(&str, "str", "genetic", "the string to generate")
	flag.StringVar(&str, "s", "genetic", "the string to generate (shorthand)")
	flag.Float64Var(&mutation, "mutation", 0.01, "the rate of mutation")
	flag.Float64Var(&mutation, "m", 0.01, "the rate of mutation (shorthand)")
	fmt.Println("hello")
}

func main() {
	flag.Parse()
	generic.GenerateString([]byte(str), popNum, mutation)
	// fmt.Println(str, popNum, mutation)
}
