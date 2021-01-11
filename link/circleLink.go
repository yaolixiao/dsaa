package main

import "fmt"

// 环形链表

type CircleNode struct {
	no       int
	name     string
	nickname string
	next     *CircleNode
}

func Add(head *CircleNode, newNode *CircleNode) {

	temp := head

	if temp.next == nil {
		// 这里说明环形链表还是空的，添加的是第一个元素
		head.no = newNode.no
		head.name = newNode.name
		head.nickname = newNode.nickname
		head.next = head
		return
	}

	// 先找到环形链表末尾，将末尾的next指向head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	newNode.next = head
	temp.next = newNode

}

func List(head *CircleNode) {

	temp := head
	if temp.next == nil {
		fmt.Println("链表是空的")
		return
	}

	for {
		fmt.Printf("[%d %s %s] ==> ", temp.no, temp.name, temp.nickname)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	// 环形链表，添加元素时，第一个元素应该给head
	head := &CircleNode{}

	node1 := &CircleNode{
		no:       1,
		name:     "tom",
		nickname: "nick-tom",
	}

	node2 := &CircleNode{
		no:       2,
		name:     "Lucy",
		nickname: "nick-Lucy",
	}

	Add(head, node1)
	Add(head, node2)
	List(head)

}
