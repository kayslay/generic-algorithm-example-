package generic

import (
	"fmt"
	"math"
	"sort"
)

//Population type
type Population struct {
	entities _DNASet
	mutation float64
	finalVal []byte
	popNum   int
}

type _DNASet []DNA

func (d _DNASet) Len() int {
	return len(d)
}
func (d _DNASet) Less(i, j int) bool {
	return d[i].Score() > d[j].Score()
}

func (d _DNASet) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

//NewPopulation create a new population. finalVal is the []byte(string to find), popNum the number of DNAs in the population
//mutation is the rate of mutation of the population
func NewPopulation(finalVal []byte, popNum int, mutation float64) Population {
	dna := make([]DNA, popNum)
	for i := 0; i < popNum; i++ {
		dna[i] = NewDNA(finalVal)
	}
	return Population{entities: dna, mutation: mutation, popNum: popNum, finalVal: finalVal}
}

//GenerateString generates the sting using genetic algorithm
func GenerateString(b []byte, p int, m float64) {
	d := NewPopulation(b, p, m)
	str := ""
	step := 0
	for str != string(b) {
		child := d.Reproduce()
		str = child.String()
		fmt.Println(str, fmt.Sprintf("| %.2f%% close to the real deal |", (child.Score()*100)/float64(len(b))), fmt.Sprintf(" %d steps |", step))
		// d.V()
		step++
	}
	fmt.Println("finally")
}

//Reproduce creates a new child DNA
func (p Population) Reproduce() DNA {
	var score float64
	var child DNA
	parents := p.pickN(2)

	for score == 0 {
		child = crossOver(p.entities[parents[0]], p.entities[parents[1]])
		p.mutate(child)
		score = child.Score()
	}
	p.entities = append(p.entities[:len(p.entities)-1], child)
	p.sortPopulation()

	return child
}

//pickN picks n random index from the p.entities. panic if n is greater than len(p.entities)
func (p Population) pickN(n int) []int {
	if n > len(p.entities) || n < 0 {
		panic("n in Population.PickN is greater than the number of DNAs or less than 0")
	}
	dnaIndex := []int{}
	for len(dnaIndex) < n {
		i := p.pick()
		if !inIntSlice(dnaIndex, i) {
			dnaIndex = append(dnaIndex, i)
		}
	}
	return dnaIndex
}

//inIntSlice check if y exist in x
func inIntSlice(x []int, y int) bool {
	for v := range x {
		if x[v] == y {
			return true
		}
	}
	return false
}

//pick picks the index of an entity at random
func (p Population) pick() int {
	r := r1.Float64() * 100
	probSum := p.genProbability()
	index := sort.Search(len(probSum), func(i int) bool { return probSum[i] >= r })
	if index == len(probSum) {
		return index
	}
	return index
}

//genProbability retruns the probabilities selecting each DNA in the _DNASet
func (p Population) genProbability() []float64 {
	total := p.getTotalScore() + float64(len(p.entities))
	probSum := []float64{}
	sum := 0.0
	for v := range p.entities {
		score := p.entities[v].Score() + 1

		sum += (score / total) * 100
		probSum = append(probSum, sum)
	}
	return probSum
}

//sortPopulation sorts the entities in the Population by its score
func (p *Population) sortPopulation() {
	sort.Sort(p.entities)
}

//getTotalScore gets the summation of all the scores of the entities
func (p Population) getTotalScore() float64 {
	total := 0.0
	for k := range p.entities {
		total += p.entities[k].Score()
	}
	return total
}

func (p Population) GetTotalScore() float64 {
	total := 0.0
	for k := range p.entities {
		total += p.entities[k].Score()
	}
	return total
}

func crossOver(i, j DNA) DNA {
	dna := NewDNA(i.finalVal)
	crossOverNum := int(math.Abs(float64(r1.Intn(len(dna.finalVal) - 2))))
	var p1, p2 []byte = make([]byte, crossOverNum), make([]byte, len(dna.finalVal)-crossOverNum)
	copy(p1, i.chromosome[:crossOverNum])
	copy(p2, j.chromosome[crossOverNum:])
	dna.chromosome = append(p1, p2...)
	return dna
}

func (p Population) mutate(child DNA) {
	if r1.Float64() <= p.mutation {
		i, b := r1.Intn(len(p.finalVal)), random(32, 122)
		child.chromosome[i] = b
	}
}

func (p Population) GetGreatest() string {
	return fmt.Sprint(p.entities[0], p.entities[1])
}
