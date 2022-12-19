package main

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Elem struct {
	val, pos int
}

func dailyTemperatures(temperatures []int) (diffs []int) {
	stack := arraystack.New() // monotonic stack
	diffs = make([]int, len(temperatures))
	stack.Push(Elem{temperatures[len(temperatures)-1], len(temperatures) - 1})
	for i := len(temperatures) - 1; i >= 0; i-- {
		//fmt.Println(stack)
		var top Elem
		tops, _ := stack.Peek()
		top = tops.(Elem)
		cur := Elem{temperatures[i], i}
		if cur.val < top.val {
			stack.Push(cur)
			diffs[i] = 1
		} else {
			for !stack.Empty() && top.val < cur.val {
				stack.Pop()
				tops, _ = stack.Peek()
				top, _ = tops.(Elem)
			}
			if stack.Empty() {
				diffs[i] = 0
			} else {
				diffs[i] = top.pos - i
			}
			stack.Push(cur)
		}
	}
	return
}
