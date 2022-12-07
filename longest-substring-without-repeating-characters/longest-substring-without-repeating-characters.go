package main

//func Max(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}

func lengthOfLongestSubstring(s string) (substLen int) {
	var met = make(map[rune]int)
	var left int
	for right, letter := range s {
		pos, ok := met[letter]
		if ok && pos >= left {
			left = pos + 1
		}
		met[letter] = right
		substLen = Max(substLen, right-left+1)
	}
	return
}

//
//func main() {
//    println(lengthOfLongestSubstring("abcabcbb"))
//}
