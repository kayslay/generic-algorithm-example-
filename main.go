package main

import (
	"learn/generic/generic"
	"math/rand"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func main() {
	generic.GenerateString([]byte("gentic algorithm is cool"), 1500, 0.1)
	// fmt.Println(' ')
}
