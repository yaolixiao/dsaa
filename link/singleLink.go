package main

import "fmt"

// 单链表 增删改查

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func InsertHeroNode(head *HeroNode, heroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = heroNode
}

func InsertHeroNode2(head *HeroNode, heroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > heroNode.no {
			break
		} else if temp.next.no == heroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Printf("hero no = %d 重复！\n", heroNode.no)
	} else {
		heroNode.next = temp.next
		temp.next = heroNode
	}
}

func DelHero(head *HeroNode, id int) {
	temp := head
	flag := false
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
		temp.next = temp.next.next
	} else {
		fmt.Printf("id = %d 不存在\n", id)
	}
}

func ListHero(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("英雄空空如也...")
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

	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero4)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	ListHero(head)

	DelHero(head, 3)
	ListHero(head)

	DelHero(head, 2)
	ListHero(head)

	DelHero(head, 3)
	ListHero(head)

	DelHero(head, 4)
	ListHero(head)

	DelHero(head, 1)
	ListHero(head)
}
