package main

import "fmt"

// 双向链表

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode
}

func InsertHero(head *HeroNode, hero *HeroNode) {
	temp := head

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = hero
	hero.pre = temp
}

func InsertHero2(head *HeroNode, hero *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > hero.no {
			break
		} else if temp.next.no == hero.no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Printf("hero no = %d 已存在", hero.no)
	} else {
		hero.pre = temp
		hero.next = temp.next
		if temp.next != nil {
			temp.next.pre = hero
		}
		temp.next = hero
	}
}

func DelHero(head *HeroNode, id int) {
	temp := head
	flag := false

	// 需要找到编号是id的这个hero
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		// 代表找到了, temp.next是要删除的元素
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Printf("删除错误：id = %d 的英雄不存在\n", id)
	}

}

func ListHero(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}

	for {
		fmt.Printf("[%d %s %s] ==> ", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

func ListHero2(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}

	// 先让temp指向链表尾部
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 向前遍历
	for {
		fmt.Printf("[%d %s %s] ==> ", temp.no, temp.name, temp.nickname)
		if temp.pre == nil || temp.pre.no == 0 {
			break
		}
		temp = temp.pre
	}
	fmt.Println()
}

func main() {

	head := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "武松",
		nickname: "打虎英雄",
	}

	InsertHero2(head, hero2)
	InsertHero2(head, hero4)
	InsertHero2(head, hero1)
	InsertHero2(head, hero3)

	DelHero(head, 3)
	DelHero(head, 2)
	DelHero(head, 1)
	DelHero(head, 4)
	DelHero(head, 4)

	ListHero(head)
	fmt.Println("逆向英雄列表：")
	ListHero2(head)
}
