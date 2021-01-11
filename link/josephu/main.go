package main

import "fmt"

// 约瑟夫问题

type Boy struct {
	No   int
	Next *Boy
}

// 根据传入数量创建环形链表，返回第一个元素
func createBoyLink(n int) *Boy {

	first := &Boy{}  // 第一个元素，一旦确定不能改变
	curBoy := &Boy{} // 当前元素，辅助作用，随着遍历方法向前移动

	if n < 1 {
		fmt.Println("n必须大于等于1")
		return first
	}

	for i := 1; i <= n; i++ {
		boy := &Boy{
			No: i,
		}

		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first // 构成环形链表
		} else {
			// 这里的curBoy是上一轮的boy
			curBoy.Next = boy
			// 变换curBoy为本轮的boy
			curBoy = boy
			// 构造环形链表
			curBoy.Next = first
		}
	}
	return first
}

// 显示链表
func showBoyLink(first *Boy) {

	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}

	temp := first
	for {
		fmt.Printf("编号%d -> ", temp.No)
		if temp.Next == first {
			break
		}
		temp = temp.Next
	}
	fmt.Println()
}

// 约瑟夫问题实现
func play(first *Boy, start int, count int) {

	// 验证链表是否为空
	if first.Next == nil {
		fmt.Println("链表为空，不能继续了")
		return
	}

	// todo 验证start是否大于链表长度，如果大于则不能继续

	// 定义辅助元素tail，因单向链表删除元素需要借助辅助元素
	tail := first

	// 将tail指向first前一个元素，跟随first一起移动，保证tail始终是firs前一个元素
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	// 根据start将first移动到开始位置
	// 为什么start-1？因为first本身就占一个位置，移动start-1次的位置是start位置
	for i := 1; i <= start-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 循环执行，直到链表只剩下一个元素为止
	for {
		// 根据count将first移动到要删除元素的位置
		// 为什么count-1？因为从first本身开始数，first本身也算一次
		for i := 1; i <= count-1; i++ {
			first = first.Next
			tail = tail.Next
		}

		// 删除first指向的元素
		fmt.Printf("删除元素：%d\n", first.No)
		first = first.Next
		tail.Next = first

		// 当tail==first时候，表示只剩下了一个元素
		if tail == first {
			break
		}
	}

	fmt.Println("剩下的元素编号 =", first.No)
}

func main() {

	first := createBoyLink(50)
	showBoyLink(first)
	play(first, 20, 33)
}
