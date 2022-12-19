package main

func longestConsecutive(nums []int) (ans int) {
	numsMap := make(Set, len(nums))
	for _, num := range nums {
		numsMap.Add(num)
	}
	for _, num := range nums {
		if !numsMap.In(num - 1) {
			nextNum := num + 1
			for ; numsMap.In(nextNum); nextNum++ {
			}
			ans = Max(ans, nextNum-num)
		}
	}
	return
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type StructureElement = int
type none = struct{}
type Set map[StructureElement]none

func (s Set) Add(val StructureElement) {
	s[val] = none{}
}
func (s Set) In(val StructureElement) bool {
	_, ok := s[val]
	return ok
}
func (s Set) Remove(val StructureElement) {
	delete(s, val)
}
