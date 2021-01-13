package main

import (
	"errors"
	"fmt"
)

// 入栈、出栈、遍历
// 先进后出

type Stack struct {
	MaxTop int    // 栈的最大容量
	Top    int    // 栈顶，默认-1表示空栈
	arr    [5]int // 用数组表示栈空间
}

func (this *Stack) push(val int) error {
	if this.Top == this.MaxTop {
		return errors.New("stack full not push")
	}

	this.Top++
	this.arr[this.Top] = val
	return nil
}

func (this *Stack) pop() (int, error) {
	if this.Top == -1 {
		return 0, errors.New("stack empty not pop")
	}

	val := this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈
func (this *Stack) list() {
	if this.Top == -1 {
		fmt.Println("stack is empty")
		return
	}

	for i := this.Top; i >= 0; i-- {
		fmt.Println(this.arr[i])
	}
}

func main() {
	//
	// s := &Stack{
	// 	MaxTop: 5,
	// 	Top:    -1,
	// }

	str := "3+2+4"
	s1 := str[0:1]
	s2 := str[1:2]
	fmt.Println(s1)
	fmt.Println(s2)
	b1 := []byte(s1)
	b2 := []byte(s2)
	fmt.Println(b1)
	fmt.Println(b2)
	i1 := int(b1[0])
	i2 := int(b2[0])
	fmt.Println(i1)
	fmt.Println(i2)
}
