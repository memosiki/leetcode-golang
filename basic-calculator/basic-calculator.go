package main

import (
	"fmt"
	"log"
	"strconv"
)

type (
	Token struct {
		kind  TokenKind
		value TokenValue
	}
	TokenKind  byte
	TokenValue int
)

const (
	bracket TokenKind = iota
	numeric
	operation
)
const (
	openBracket TokenValue = iota
	closeBracket

	plus
	minus
	mul
	div
)

var (
	precedence = map[TokenValue]int{
		mul:   1,
		div:   1,
		plus:  0,
		minus: 0,
	}
	do = map[TokenValue]func(TokenValue, TokenValue) TokenValue{
		mul:   func(b, a TokenValue) TokenValue { return a * b },
		div:   func(b, a TokenValue) TokenValue { return a / b },
		plus:  func(b, a TokenValue) TokenValue { return a + b },
		minus: func(b, a TokenValue) TokenValue { return a - b },
	}
	translate = map[TokenKind]map[byte]TokenValue{
		numeric: {
			'0': 0,
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
		},
		operation: {
			'+': plus,
			'-': minus,
			'*': mul,
			'/': div,
		},
		bracket: {
			'(': openBracket,
			')': closeBracket,
		},
	}
	untranslate = map[TokenValue]string{
		plus:         "+",
		minus:        "-",
		mul:          "*",
		div:          "/",
		openBracket:  "(",
		closeBracket: ")",
	}
)

func (token Token) String() string {
	if token.kind == numeric {
		return strconv.Itoa(int(token.value))
	}
	return untranslate[token.value]
}

func ofAKind(char byte, kind TokenKind) (ok bool) {
	_, ok = translate[kind][char]
	return
}

func tokenize(s string) (tokens []Token) {
	for i := 0; i < len(s); i++ {
		switch char := s[i]; {
		case ofAKind(char, numeric):
			var num TokenValue = 0
			for ; i < len(s) && ofAKind(s[i], numeric); i++ {
				num = num*10 + translate[numeric][s[i]]
			}
			i--
			tokens = append(tokens, Token{numeric, num})
		case ofAKind(char, operation):
			value := translate[operation][char]
			// resolve unary minus
			if value == minus && (len(tokens) == 0 ||
				tokens[len(tokens)-1].kind == operation ||
				(tokens[len(tokens)-1].kind == bracket && tokens[len(tokens)-1].value == openBracket)) {
				tokens = append(tokens, Token{numeric, 0})
			}
			tokens = append(tokens, Token{operation, value})
		case ofAKind(char, bracket):
			tokens = append(tokens, Token{bracket, translate[bracket][char]})
		case char == ' ':
			// noop
		default:
			log.Panicln("Invalid symbol", char)
		}
	}
	return
}

func convertPostfix(tokens []Token) (postfix []Token) {
	var stack Stack
	for _, token := range tokens {
		switch token.kind {
		case numeric:
			postfix = append(postfix, token)
		case operation:
			for !stack.Empty() {
				stackToken := stack.Peek()
				if stackToken.kind == operation && precedence[stackToken.value] >= precedence[token.value] {
					postfix = append(postfix, stack.Pop())
				} else {
					break
				}
			}
			stack.Push(token)
		case bracket:
			if token.value == openBracket {
				stack.Push(token)
			} else {
				for !stack.Empty() && stack.Peek().kind != bracket {
					postfix = append(postfix, stack.Pop())
				}
				if !stack.Empty() {
					stack.Pop()
				}
			}
		}
	}
	for !stack.Empty() {
		postfix = append(postfix, stack.Pop())
	}
	return
}
func calcPostfix(tokens []Token) int {
	var stack Stack
	for _, token := range tokens {
		if token.kind == operation {
			stack.Push(Token{numeric, do[token.value](stack.Pop().value, stack.Pop().value)})
		} else {
			stack.Push(token)
		}
	}
	if stack.Size() != 1 {
		log.Panicln("Incorrect tokens", stack)
	}
	return int(stack.Pop().value)
}
func calculate(s string) int {
	fmt.Println(s)
	tokens := tokenize(s)
	fmt.Println(tokens)
	tokens = convertPostfix(tokens)
	fmt.Println(tokens)
	return calcPostfix(tokens)
}

type Stack struct {
	data []Token
}

func (stack *Stack) Push(v Token) {
	stack.data = append(stack.data, v)
}

func (stack *Stack) Pop() (popped Token) {
	popped = stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return
}

func (stack *Stack) Peek() Token {
	return stack.data[len(stack.data)-1]
}
func (stack *Stack) Empty() bool {
	return len(stack.data) == 0
}
func (stack *Stack) Size() int {
	return len(stack.data)
}

//func main() {
//	a := calculate(" 3+5 / 2 ")
//	fmt.Println(a)
//}
