package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList)IsPalindrome(){
	fast, slow := ll.head,ll.head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	firstTail := slow
	slow = slow.next
	firstTail.next = nil

	reversedList := Reverse(slow)
	current := ll.head
	IsPalindrome := true
	for reversedList.next != nil  && current.next != nil {
		if reversedList.data != current.data {
			IsPalindrome = false
			break
		}
		reversedList = reversedList.next
		current = current.next
	}
	fmt.Println(IsPalindrome)
}

func Reverse(head *ListNode)*ListNode{
	var prev *ListNode
	current := head

	for current != nil {
		temp := current.next
		current.next = prev
		prev = current		
		current = temp
	}

	return prev
}

func main(){
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = &ListNode{data:30}
	ll.head.next.next.next.next.next = &ListNode{data:20}
	ll.head.next.next.next.next.next.next = &ListNode{data:10}
	ll.IsPalindrome()

	ll1 := &LinkedList{}
	ll1.head = &ListNode{data:10}
	ll1.head.next = &ListNode{data:20}
	ll1.head.next.next = &ListNode{data:30}
	ll1.head.next.next.next = &ListNode{data:30}
	ll1.head.next.next.next.next = &ListNode{data:20}
	ll1.head.next.next.next.next.next = &ListNode{data:10}
	ll1.IsPalindrome()

	ll2 := &LinkedList{}
	ll2.head = &ListNode{data:10}
	ll2.head.next = &ListNode{data:20}
	ll2.head.next.next = &ListNode{data:30}
	ll2.head.next.next.next = &ListNode{data:30}
	ll2.head.next.next.next.next = &ListNode{data:20}
	ll2.head.next.next.next.next.next = &ListNode{data:100}
	ll2.IsPalindrome()
}