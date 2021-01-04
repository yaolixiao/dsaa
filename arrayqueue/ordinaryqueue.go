package main

import (
	"fmt"
	"os"
	"errors"
)

// 队列是一个有序列表，可以用数组或链表来实现
// 遵循先入先出

// 数组模拟队列
// 1. maxSize 是队列的最大容量
// 2. front 和 rear 表示队首和队尾，front随着数据的输出而改变，rear随着数据的输入而改变

// 使用数组实现队列（非环形）
// 1. 创建一个数组array，是作为队列的一个字段
// 2. front初始化为-1
// 3. rear初始化为-1
// 4. 完成队列的基本操作：addQueue加入、getQueue取出、showQueue显示

type Queue struct {
	maxSize int
	array 	[4]int
	front 	int
	rear 	int
}

// 添加数据到队列
func (this *Queue) addQueue(val int) error {

	// 先判断队列是否已满
	// rear队尾（含最后一个元素）
	if this.rear == this.maxSize - 1 {
		return errors.New("队列已满，不能再添加。")
	}

	// rear向后移动
	this.rear++
	this.array[this.rear] = val
	return nil
}

// 从队列取出数据
func (this *Queue) getQueue() (int, error) {

	// 先判断队列是否为空
	if this.front == this.rear {
		return -1, errors.New("队列空了... ")
	}

	// front向后移动
	this.front++
	return this.array[this.front], nil
}

// 显示队列
func (this *Queue) showQueue() {
	
	// 思路：从队首遍历到队尾
	// front不包含队首元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, this.array[i])
	}
}

func main() {
	
	// 初始化一个队列
	queue := &Queue{
		maxSize: 4,
		front: -1,
		rear: -1,
	}

	var key string
	var val int

	for {
		fmt.Println("1. 输入add 添加")
		fmt.Println("2. 输入get 获取")
		fmt.Println("3. 输入show 显示队列")
		fmt.Println("4. 输入exit 退出")

		fmt.Scanln(&key)
		switch key {
			case "add":
				fmt.Println("输入你要入队的数据")
				fmt.Scanln(&val)
				err := queue.addQueue(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("add succsss.")
				}
			case "get":
				val, err := queue.getQueue()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Printf("取出的数据=%d\n", val)
				}
			case "show":
				queue.showQueue()
			case "exit":
				os.Exit(0)
		}
	}
}