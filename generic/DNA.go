package generic

import (
	"math"
	"math/rand"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

//DNA type
type DNA struct {
	len        int
	chromosome []byte
	finalVal   []byte
}

//NewDNA creates a new DNA
func NewDNA(final []byte) DNA {
	l := len(final)
	b := make([]byte, l)
	return DNA{len: l, chromosome: generateChromosome(b, l), finalVal: final}
}

func generateChromosome(b []byte, len int) []byte {
	for i := 0; i < len; i++ {
		b[i] = Random(32, 122)
	}
	return b
}

//Score calculate the score
func (d DNA) Score() float64 {
	score := 0.0
	for k := range d.chromosome {
		if d.finalVal[k] == d.chromosome[k] {
			score++
		}
	}
	return score
}

func Random(from, to int) byte {
	diff := math.Abs(float64(to - from))
	return byte(r1.Intn(int(diff)) + from)
}

func (d DNA) String() string {
	return string(d.chromosome)
}
