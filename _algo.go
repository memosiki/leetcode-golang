package main

/*
	Используй Sort.Search

Бинарный поиск ключа в отсортированном массиве без повторений,
если ключ не найден возвращает индекс элемента большего чем key
если ключ больше максимального значения в массиве -- вернёт len(arr)
bisect_right из питона
*/
func binarySearch(arr []int, key int) (index int, ok bool) {
	var left, right, mid int
	left = 0
	right = len(arr)

	for left < right {
		fmt.Println(left, right, mid)
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

func CountingSort(nums []int) {
	var counts = make([]int, MaxList(nums)+1)
	for _, elem := range nums {
		counts[elem] += 1
	}
	var i int
	for number, count := range counts {
		for k := 0; k < count; k++ {
			nums[i+k] = number
		}
		i += count
	}
}

// CountingSort для tuple
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
	for i, count := range counts[:len(counts)-1] {
		counts[i+1] += count
	}
	for _, elem := range container {
		sorted[counts[elem.grow]-1] = elem
		counts[elem.grow]--
	}
	return
}

/* Количество цифр в числе */
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
