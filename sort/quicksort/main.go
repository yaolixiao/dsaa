package main

import (
	"fmt"
)

// 快速排序（从大到小）
// 选中一个元素作为中间值，将比他大的放到左边，比他小的放到右边
// 递归这个过程，就实现了快速排序
func quicksort(left int, right int, arr *[11]int) {
	l := left
	r := right

	// 取中间值
	middleVal := arr[(l+r)/2]
	// fmt.Println("middleVal:", middleVal)
	// fmt.Println("arr:", arr)

	for l < r {
		// 从右边开始找比他大的
		for arr[r] < middleVal {
			r--
		}
		// 从左边开始找比他小的
		for arr[l] > middleVal {
			l++
		}

		if l != r {
			arr[l], arr[r] = arr[r], arr[l]
		}

		// 避免重复数据造成死循环
		if l != r && arr[l] == arr[r] {
			l++
		}
	}

	// 找到此时中间值的位置
	var mid int
	if arr[l] == middleVal {
		mid = l
	} else if arr[r] == middleVal {
		mid = r
	} else {
		fmt.Println("出错了！")
	}

	// fmt.Println(left, mid, right)

	if left < mid-1 {
		quicksort(left, mid-1, arr)
	}
	if right > mid+1 {
		quicksort(mid+1, right, arr)
	}
}

func main() {
	//
	ages := [11]int{6, 7, 6, 10, 2, 11, 1, 2, 3, 5, 2}
	fmt.Println(ages)
	quicksort(0, len(ages)-1, &ages)
	fmt.Println(ages)
}
