package main

import (
	"fmt"
	"errors"
	"os"
)

// 数组模拟环形队列（模拟环形队列，队列的容量是数组实际大小-1）
// 什么时候表示队列满 (rear + 1) % maxSize == front
// 什么时候表示队列空 rear == front
// 初始化时 front = 0  rear = 0
// 怎么统计该队列有多少个元素 (rear + maxSize - front) % maxSize

// 使用结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array [4]int
	front int
	rear int
}

func (this *CircleQueue) Add(val int) error {
	if (this.rear + 1) % this.maxSize == this.front {
		return errors.New("队列已满了")
	}

	// this.rear 不包含队列尾部元素，如果包含就不好做了
	this.array[this.rear] = val
	this.rear = (this.rear + 1) % this.maxSize
	return nil
}

func (this *CircleQueue) Get() (int, error) {
	if this.rear == this.front {
		return 0, errors.New("队列已空")
	}

	// this.front 包含队首元素
    val := this.array[this.front]
	this.front = (this.front + 1) % this.maxSize
	return val, nil
}

func (this *CircleQueue) showQueue() {
	if this.rear == this.front {
		fmt.Println("队列是空的，没有可显示的元素")
		return
	}

	// 求出队列一共多少元素，遍历队列
	// 设计一个临时队首，根据临时队首开始遍历
	size := (this.rear + this.maxSize - this.front) % this.maxSize
	var tempFront = this.front
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempFront, this.array[tempFront])
		tempFront = (tempFront + 1) % this.maxSize
	}
	fmt.Println()
}

func main() {

	// 初始化一个队列
	queue := &CircleQueue{
		maxSize: 4,
		front: 0,
		rear: 0,
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
				err := queue.Add(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("add succsss.")
				}
			case "get":
				val, err := queue.Get()
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