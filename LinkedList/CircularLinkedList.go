package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}


func(cll *LinkedList)Display(){
	current := cll.head
	for  {
		fmt.Printf("%d->",current.data)
		current = current.next
		if current == cll.head {
			break
		}
	}
	fmt.Println()
}

func(cll *LinkedList)InsertAtHead(data int){
	newNode := &ListNode{data:data}
	if cll.head == nil {
		cll.head = newNode
		cll.head.next = cll.head
		cll.tail = newNode
	} else {
		cll.tail.next = newNode
		newNode.next = cll.head
		cll.head = newNode
	}
	cll.size++
}

func(cll *LinkedList)InsertAtTail(data int){
	newNode := &ListNode{data:data}
	newNode.next = cll.head
	cll.tail.next = newNode
	cll.tail = newNode
	cll.size++
}

func(cll *LinkedList)InsertAtPosition(data int , position int){
	if position == 1 {
		cll.InsertAtHead(data)
		return
	} else if position > cll.size {
		cll.InsertAtTail(data)
		return
	} else {
		newNode := &ListNode{data:data}
		current := cll.head
		var prev *ListNode
		for pos := 1; pos < position; pos++ {
			prev = current
			current = current.next
		}
		prev.next = newNode
		newNode.next = current
		current.next = current.next
		cll.size++
	}
}

func(cll *LinkedList)DeleteHead(){
	cll.head = cll.head.next
	cll.tail.next = cll.head
	cll.size--
}

func(cll *LinkedList)DeleteTail(){
	current := cll.head
	for current.next != cll.tail {
		current = current.next
	}

	current.next = cll.tail.next
	cll.tail = current
	cll.size--
}

func(cll *LinkedList)DeleteAtPosition(position int){
	if position == 1 {
		cll.DeleteHead()
		return
	} else if position >= cll.size {
		cll.DeleteTail()
		return
	} else {
		current := cll.head
		var prev *ListNode
		for pos := 1; pos < position; pos++{
			prev = current
			current = current.next
		}
		prev.next = current.next
		current = current.next
		cll.size--
	}

}

func main(){
	cll := &LinkedList{}
	cll.InsertAtHead(40)
	cll.InsertAtHead(30)
	cll.InsertAtHead(20)
	cll.InsertAtHead(10)
	cll.Display()
	cll.InsertAtTail(70)
	cll.InsertAtTail(80)
	cll.InsertAtTail(90)
	cll.InsertAtTail(100)
	cll.Display()
	cll.InsertAtPosition(50,5)
	cll.InsertAtPosition(60,6)
	cll.InsertAtPosition(70,7)
	cll.InsertAtPosition(80,8)
	cll.Display()
	cll.DeleteHead()
	cll.DeleteHead()
	cll.Display()
	cll.DeleteTail()
	cll.DeleteTail()
	cll.Display()
	fmt.Println()
	cll.DeleteAtPosition(2)

	cll.Display()
	cll.DeleteAtPosition(3)
	cll.Display()
}