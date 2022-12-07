package main

import "github.com/emirpasic/gods/containers"

/* Counter */
type Counter map[rune]int

func NewCounter(s string) Counter {
	counts := make(map[rune]int)
	var letter rune
	for _, letter = range s {
		if _, ok := counts[letter]; ok {
			counts[letter] += 1
		} else {
			counts[letter] = 1
		}
	}
	return counts
}
func (counter Counter) Copy() Counter {
	counts := make(map[rune]int)
	for k, v := range counter {
		counts[k] = v
	}
	return counts
}
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func (counter Counter) Intersection(other Counter) (counts Counter) {
	counts = make(map[rune]int)
	for key, value := range counter {
		if min := min(value, other.Get(key)); min > 0 {
			counts[key] = min
		}
	}
	return
}
func (counter Counter) Get(key rune) (value int) {
	value, ok := counter[key]
	if !ok {
		return 0
	}
	return
}
func (counter Counter) Add(key rune, increment int) {
	if _, ok := counter[key]; ok {
		counter[key] += increment
	} else {
		counter[key] = increment
	}
}

// Получить произвольный элемент
func (counter Counter) GetAny() (key rune, value int, ok bool) {
	for k, val := range counter {
		return k, val, true
	}
	return 0, 0, false
}

// Все символы в counter уникальны
func (counter Counter) IsUnique() bool {
	for _, count := range counter {
		if count > 1 {
			return false
		}
	}
	return true
}

/* Stack */
type Stack []int

func (s Stack) Push(v int) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, int) {
	return s[:len(s)-1], s[len(s)-1]
}

/* Set */
type Set map[int]struct{}

func (s Set) Add(val int) {
	s[val] = struct{}{}
}

func (s Set) In(val int) bool {
	_, ok := s[val]
	return ok
}

func (s Set) Remove(val int) {
	delete(s, val)
}
