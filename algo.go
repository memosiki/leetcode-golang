package main

// LenInt возращает количество цифр в числе
func LenInt(i int) (count int) {
	if i >= 1e18 {
		return 19
	}
	x := 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

/*
BinarySearch двоичный поиск в отсортированном слайсе без повторений.
Полный аналог Sort.Search и bisect_right из Python.

- если ключ не найден возвращает индекс элемента большего чем key
- если ключ больше максимального значения в массиве -- вернёт len(arr)
*/
func BinarySearch(arr []int, key int) (index int, ok bool) {
	var left, right, mid int
	left = 0
	right = len(arr)

	for left < right {
		mid = (right + left) / 2
		if arr[mid] > key {
			right = mid
		} else if arr[mid] < key {
			left = mid + 1
		} else {
			return mid, true
		}
	}
	if left != right {
		panic(key)
	}
	return left, false
}

func CountingSortInplace(container []Element) {
	var counts = make([]int, Max(container...)+1)
	for _, elem := range container {
		counts[elem] += 1
	}
	var i int
	for number, count := range counts {
		for k := 0; k < count; k++ {
			container[i+k] = Element(number)
		}
		i += count
	}
}

func CountingSort(container []Element) (sorted []Element) {
	const DESC = false
	counts := make([]int, Max(container...)+1)
	sorted = make([]Element, len(container))
	for _, elem := range container {
		counts[elem] += 1
	}
	if DESC {
		for i := len(counts) - 1; 0 < i; i-- {
			counts[i-1] += counts[i]
		}
	} else {
		for i, count := range counts[:len(counts)-1] {
			counts[i+1] += count
		}
	}
	for _, elem := range container {
		sorted[counts[elem]-1] = elem
		counts[elem]--
	}
	return
}

const radix int = 0xff
const MaxCount = 100_000

var counts, zeroes [radix + 1]int
var buffer [MaxCount]int

func CountingSortForRadix(container []int, order int) {
	copy(counts[:], zeroes[:])
	for _, elem := range container {
		counts[(elem>>order)&radix]++
	}
	for i := 1; i <= radix; i++ {
		counts[i] += counts[i-1]
	}
	for i := len(container) - 1; i >= 0; i-- {
		elem := container[i]
		counts[(elem>>order)&radix]--
		buffer[counts[(elem>>order)&radix]] = elem
	}
	copy(container, buffer[:])
}
func RadixSort(container []int) {
	for order := 0x0; order <= 0x3; order++ {
		CountingSortForRadix(container, order*0b1000)
	}
}

// Permutations returns a channel that will receive all permutations of a container
func Permutations(container []Element) <-chan none {
	next := make(chan none)
	p := make([]int, len(container))
	go func() {
		defer close(next)
		for p[0] < len(p) {
			// get permutation
			for i, v := range p {
				container[i], container[i+v] = container[i+v], container[i]
			}
			next <- none{}
			// next permutation
			for i := len(p) - 1; i >= 0; i-- {
				if i == 0 || p[i] < len(p)-i-1 {
					p[i]++
					break
				}
				p[i] = 0
			}
		}
	}()
	return next
}

// Permutation implementation of https://stackoverflow.com/a/30230552
// Slice holds intermediate state as offsets in a Fisher-Yates shuffle algorithm
type Permutation []int

func (diffs Permutation) Next() (exists bool) {
	for i := len(diffs) - 1; i >= 0; i-- {
		if i == 0 || diffs[i] < len(diffs)-i-1 {
			diffs[i]++
			break
		}
		diffs[i] = 0
	}
	if diffs[0] >= len(diffs) {
		return
	}
	return true
}

func (diffs Permutation) Get(container []StructureElement) {
	for i, v := range diffs {
		container[i], container[i+v] = container[i+v], container[i]
	}
}
