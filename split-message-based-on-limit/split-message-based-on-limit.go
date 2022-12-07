package main

import (
	//	"fmt"
	"math"
	"strconv"
	"strings"
)

func lenInt(i int) (count int) {
	/* Количество цифр в числе */
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
func additionalSymbols(i int) (ans int) {
	/* количество дополнительных символов для count сообщений */
	// <a/b>
	ans += i * 3         //  Длина < / >
	ans += i * lenInt(i) // Длина b
	remaining := i
	// 1 2 3 4 5 6 7 8 9 10 11 12 13
	//	switch {
	//	case i < 10: // 1-9 (9 вариантов)
	//		ans += i * 1
	//	case i < 100:
	//		ans += 9*1 + (i-9)*2
	//	case i < 1_000:
	//		ans += 9*1 + 90*2 + (i-90-9)*3
	//	case i < 10_000:
	//		ans += 9*1 + 90*2 + 900*3 + (i-900-90-9)*4
	//	case i < 100_000:
	//		ans += 9*1 + 90*2 + 900*3 + 900*4 + (i-9_000-900-90-9)*5
	//	default:
	//		panic(i)
	//	}
	for powerLen, power := range []int{10, 100, 1_000, 10_000, 100_000, 1_000_000} {
		if i < power {
			ans += remaining * (powerLen + 1)
			break
		} else {
			remaining -= power * 9 / 10
			ans += power * 9 / 10 * (powerLen + 1)
		}
	}

	return
}
func ceilDiv(a int, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}
func splitMessage(message string, limit int) (answer []string) {
	m := 1 // количество сообщений
	found := false
	for ; m <= len(message); m++ {
		add := additionalSymbols(m)
		payloadPerMessage := limit*m - add

		println(m, add, payloadPerMessage, len(message), payloadPerMessage*m)
		/* add = все доп символы,
		   ceil(message + add) / m ==  количество символов в сообщении
		   ищем минимальное количество сообщений */
		if limit*m-add >= len(message) {
			found = true
			break
		}
	}
	if !found {
		return []string{}
	}
	answer = make([]string, m)
	var closing = "/" + strconv.Itoa(m) + ">"
	var curPos int
	for i := range answer[:m-1] {
		/* ._.  */
		suffixBuilder := strings.Builder{}
		suffixBuilder.WriteByte('<')
		suffixBuilder.WriteString(strconv.Itoa(i + 1))
		suffixBuilder.WriteString(closing)
		suffix := suffixBuilder.String()
		answer[i] = message[curPos:curPos+limit-len(suffix)] + suffix
		/*
		               длина значимой части :
		           len(message[x:y] + suffix) == limit
		   		   y-x + len(suffix) == limit
		   		   x <- curpos
		   		   y = limit-len(suffix) + curpos
		*/

		curPos += limit - len(suffix)

		//        msgPart.WriteString()
		//        suffixBuilder = append(strconv.AppendInt([]byte{'<'}, i, 10), closing...)
		//        answer[i] = string(append([]byte(message[curPos:curPos+len(suffixBuilder)]), suffixBuilder...))
	}
	//    fmt.Printf("%v, %v\n", s1, s2)
	suffixBuilder := strings.Builder{}
	suffixBuilder.WriteByte('<')
	suffixBuilder.WriteString(strconv.Itoa(m))
	suffixBuilder.WriteString(closing)
	suffix := suffixBuilder.String()
	answer[m-1] = message[curPos:] + suffix
	return answer
}

//func main() {
//	//    fmt.Println(splitMessage(strings.Repeat("a", 10_000), 5))
//	println(splitMessage("this is really a very awesome message", 9))
//}
