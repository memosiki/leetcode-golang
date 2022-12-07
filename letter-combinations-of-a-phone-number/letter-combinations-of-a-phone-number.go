package main

func letterCombinations(digits string) []string {
	var (
		combs  = make([]string, 256) // 4**4 max combinations
		memory = make([]string, 256)
		dial   = map[rune]string{
			'2': "abc",
			'3': "def",
			'4': "ghi",
			'5': "jkl",
			'6': "mno",
			'7': "pqrs",
			'8': "tuv",
			'9': "wxyz",
		}
		digit   rune
		letter  rune
		curInd  int
		curComb int = 1
	)
	for _, digit = range digits {
		curInd = 0
		for _, letter = range dial[digit] {
			for k := 0; k < curComb; k++ {
				combs[curInd] = memory[k] + string(letter)
				curInd++
			}
		}
		curComb = curInd
		copy(memory, combs)
	}
	return combs[:curInd]
}

//func main() {
//	fmt.Println(letterCombinations("9999"))
//}
