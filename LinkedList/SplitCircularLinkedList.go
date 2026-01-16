package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList)Display(){
	current := ll.head

	for {
		fmt.Printf("%d->", current.data)
		current = current.next
		if current == ll.head {
			break
		}
	}
	fmt.Println()
}

func(ll *LinkedList)Split()(head1 *ListNode, head2 *ListNode){

	current := ll.head
	fast,slow := current, current

	
	for fast.next != nil && fast.next.next != nil{
		fast = fast.next.next
		slow = slow.next
		if fast.next == current {
			break
		}
	}


	firstTail := slow
	slow = slow.next
	firstTail.next = ll.head

	
	currentSlow := slow
	for currentSlow.next != ll.head {
		currentSlow = currentSlow.next
	}
	currentSlow.next = slow

	
	return firstTail,slow
}


func main() {
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = &ListNode{data:50}
	ll.head.next.next.next.next.next = &ListNode{data:60}
	ll.head.next.next.next.next.next.next = &ListNode{data:70}
	ll.head.next.next.next.next.next.next.next = &ListNode{data:80}
	ll.head.next.next.next.next.next.next.next.next = &ListNode{data:90}
	ll.head.next.next.next.next.next.next.next.next.next = ll.head
	ll.Display()
	res1,res2 := ll.Split()

	current1 := res1
	for {
		fmt.Printf("%d->",res1.data)
		res1 = res1.next
		if res1 == current1{
			break
		}
	}
	fmt.Println()
	current2 := res2
	for {
		fmt.Printf("%d->",res2.data)
		res2 = res2.next
		if res2 == current2{
			break
		}
	}


}