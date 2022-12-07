package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const MAX_INT = 10_000 + 1

type seed struct {
	plant, grow int
}

func CountingSortStruct(container []seed) (sorted []seed) {
	sorted = make([]seed, len(container))
	var counts = make([]int, MAX_INT)
	for _, elem := range container {
		counts[elem.grow] += 1
	}
	for i := len(counts) - 1; i > 0; i-- {
		counts[i-1] += counts[i]
	}
	for _, elem := range container {
		sorted[counts[elem.grow]-1] = elem
		counts[elem.grow]--
	}
	return
}

func earliestFullBloom(plantTime []int, growTime []int) (maxGrow int) {
	n := len(plantTime)
	var seeds = make([]seed, n)
	for i := 0; i < n; i++ {
		seeds[i] = seed{plantTime[i], growTime[i]}
	}
	sorted := CountingSortStruct(seeds)
	var spentTime int
	for _, seed := range sorted {
		spentTime += seed.plant
		maxGrow = max(maxGrow, spentTime+seed.grow)
	}
	return
}

//func main() {
//	a := earliestFullBloom([]int{1, 4, 3}, []int{2, 3, 1})
//}
