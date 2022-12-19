package main

func main() {

}
func maxScoreWords(words []string, letters []byte, price []int) int {
	score := make([]int, len(words))
	counters := make([]Counter, len(words))
	for i, word := range words {
		counters[i] = NewCounter([]byte(word))
		score[i] = counters[i].Score(price)
	}
}

// combinations of k element from n set
// returns indexies on included elements
func combinationsGen(n, k int) <-chan []bool {
	combs := make(chan []bool)
	go func() {
		defer close(combs)
	}()
	return combs
}
func combinations(n, k int, combs chan<- []bool) {

	if k == 1 {
		for i := 0; i < n; i++ {
			trivial := make([]bool, n)
			trivial[i] = true
			combs <- trivial
		}
		return
	}
	for i := 0; i < n; i++ {
		for subset := range combinations(n, k-1, combs) {
			if !subset[i] {
				subset[i] = true
				combs <- subset
			}
		}
	}
}
func powerset(n int) <-chan []bool {

}
func (counter Counter) Score(price []int) (score int) {
	for elem, count := range counter {
		score += price[elem] * count
	}
	return
}

type StructureElement = byte

// Counter is a limited port of collections.Counter from Python
type Counter map[StructureElement]int

// Intersection returns counter that is an intersection of both element-wise
func (counter Counter) Intersection(other Counter) (counts Counter) {
	counts = make(map[StructureElement]int)
	for key, value := range counter {
		if other[key] > 0 {
			counts[key] = Min2(value, other[key])
		}
	}
	return
}
func (counter Counter) IntersectionInplace(other Counter) {
	for key, value := range counter {
		if other[key] > 0 {
			counter[key] = Min2(value, other[key])
		} else {
			delete(counter, key)
		}
	}
}
func (counter Counter) Union(other Counter) (counts Counter) {
	counts = make(map[StructureElement]int)
	for key, value := range counter {
		counts[key] = value
	}
	for key, value := range other {
		counts[key] += value
	}
	return
}
func (counter Counter) UnionInplace(other Counter) {
	for key, value := range other {
		counter[key] += value
	}
}

// GetAny returns any element
func (counter Counter) GetAny() (key StructureElement, value int, ok bool) {
	for k, val := range counter {
		return k, val, true
	}
	return 0, 0, false
}

// IsUnique returns if all elements are unique
func (counter Counter) IsUnique() bool {
	for _, count := range counter {
		if count > 1 {
			return false
		}
	}
	return true
}

func NewCounter(container []StructureElement) Counter {
	counts := make(map[StructureElement]int, len(container))
	for _, element := range container {
		counts[element] += 1
	}
	return counts
}

type Element = int

func Max2(a, b Element) Element {
	if a > b {
		return a
	}
	return b
}
func Min2(a, b Element) Element {
	if a < b {
		return a
	}
	return b
}
