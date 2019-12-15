package singlelinkedlist

import "fmt"

type node struct{
	val int
	next *node
}

type singlelinkedlist struct{
	head *node
}

func newnode(val int) *node {
	return &node{val: val, next: nil}
}

func (ll *singlelinkedlist) addAtBeg(val int) {
	head := newnode(val)
	head.next = ll.head
	ll.head = head
}

func (ll *singlelinkedlist) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}
	fmt.Print("\n")
}

func (ll *singlelinkedlist) addAtEnd(val int) {
	n := newnode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {

	}
	cur.next = n
}

func (ll *singlelinkedlist) delAtBeg() int {
	if ll.head == nil {
		return -1
	}
	cur := ll.head
	ll.head = cur.next
	return cur.val
}

func (ll *singlelinkedlist) delAtEnd() int {
	if ll.head == nil {
		return -1
	}
	if ll.head.next == nil {
		return ll.delAtBeg()
	}
	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {

	}
	ret := cur.next.val
	cur.next = nil
	return ret
}

func (ll *singlelinkedlist) count() int {
	ctr := 0
	
	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}
	return ctr
}

//TODO: reverse()

// func main() {
// 	ll := singlelinkedlist{}

// 	ll.addAtBeg(2)
// 	ll.addAtBeg(3)
// 	ll.addAtEnd(4)
// 	ll.display()
// 	fmt.Print(ll.count(), "\n")
// 	ll.delAtBeg()
// 	fmt.Print(ll.count(), "\n")
// 	ll.display()
// 	ll.delAtEnd()
// 	ll.display()
// 	fmt.Print(ll.count(), "\n")
// }