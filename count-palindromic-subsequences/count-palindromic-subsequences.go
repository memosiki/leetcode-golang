package main

type Counter map[byte]int

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func NewCounter(s string) Counter {
    counts := make(map[byte]int)
    var letter byte
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
    counts := make(map[byte]int)
    for k, v := range counter {
        counts[k] = v
    }
    return counts
}
func (counter Counter) Difference(other Counter) (counts Counter) {
    counts = make(map[byte]int)
    for key, value := range counter {
        counts[key] = min(value, other.Get(key))
    }
    return
}
func (counter Counter) Get(key byte) (value int) {
    value, ok := counter[key]
    if !ok {
        return 0
    }
    return
}
func (counter Counter) Add(key byte, increment int) {
    if _, ok := counter[key]; ok {
        counter[key] += increment
    } else {
        counter[key] = increment
    }
}

func countPalindromes(s string) (answer int) {
    n := len(s)
    left := make([]Counter, n)
    right := make([]Counter, n)
    left[0] = NewCounter("")
    right[0] = NewCounter(s)
    right[0][s[0]] -= 1

    for i := 1; i < n; i++ {
        left[i] = left[i-1].Copy()
        right[i] = right[i-1].Copy()
        left[i].Add(s[i], 1)
        right[i].Add(s[i-1], -1)
    }
    for i := 0; i < n; i++ {
        for j := 1; j < n; j++ {
        }
    }
    for i in range(n):
        for j in range(n):
            if i!=j and s[i] == s[j]:
                for k in range(i + 2, j-1):
                    left_inner = left[k] - left[i+1]
                    right_inner = right[k] - right[j-1]
                    for t in range(10):
                        answer += left_inner[str(t)] * right_inner[str(t)]
                        if answer > MOD:
                            answer = answer - MOD
                            return answer % MOD
}
