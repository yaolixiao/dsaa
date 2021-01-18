package main

import (
	"fmt"
)

type Emp struct {
	id   int
	name string
	next *Emp
}

type empLink struct {
	head *Emp
}

func (this *empLink) insert(emp *Emp) {
	cur := this.head
	var pre *Emp

	if cur == nil {
		this.head = emp
		return
	}

	for {
		if cur != nil {
			if cur.id > emp.id {
				break
			}
			pre = cur
			cur = cur.next
		} else {
			// 这里说明到了链表的末尾，cur==nil
			break
		}
	}

	if pre != nil {
		pre.next = emp
		emp.next = cur
	}
}

func (this *empLink) showLink(no int) {
	if this.head == nil {
		fmt.Println("link empty id=", no)
		return
	}

	cur := this.head
	for {
		if cur != nil {
			fmt.Printf("link=%d id=%d name=%s -> ", no, cur.id, cur.name)
			cur = cur.next
		} else {
			break
		}
	}
	fmt.Println()
}

type hashTable struct {
	empArr [7]*empLink
}

func (this *hashTable) insert(emp *Emp) {
	no := getHashNo(emp.id)
	this.empArr[no].insert(emp)
}

func (this *hashTable) showAll() {
	for i := 0; i < len(this.empArr); i++ {
		this.empArr[i].showLink(i)
	}
}

func getHashNo(id int) int {
	return id % 7
}

func main() {
	//
}
