package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	evenHead := head.Next

	cursor := head
	secondaryCursor := head.Next

	node := head.Next.Next

	var iterations int
	for node != nil {
		cursor.Next = node
		cursor = node
		node = node.Next
		cursor, secondaryCursor = secondaryCursor, cursor
		iterations++
	}
	if iterations%2 == 0 {
		cursor.Next = evenHead
	} else {
		secondaryCursor.Next = evenHead
	}

	return head
}

func main() {

}
