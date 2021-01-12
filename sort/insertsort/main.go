package main

import (
	"fmt"
)

// 插入排序（从大到小）
// 思想：
// 将第一个元素视为有序数组，将第二个元素到末尾视为无序数组
// 从无序数组向有序数组插入元素，插入前，先找到位置
// 给第n个元素找位置，从有序数组末尾开始，循环有序数组，找到大于第n个元素的下标，该下标的下一个位置就是第n个元素的位置
func insertsort(arr *[7]int) {

	// 外层循环的是无序数组-正向循环
	// 随着循环的次数增加，无序数组的元素数量减少
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		// 内层循环的是有序数组-反向循环
		// 随着循环次数的增加，有序数组的元素数量增加
		for insertIndex >= 0 && insertVal > arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		// 如果insertIndex+1 == i则表示找到的位置 和 本来的位置相同，无需插入，进行下一轮即可
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}

func main() {
	//
	ages := [7]int{20, 0, 33, 12, -1, 35, 98}
	fmt.Println(ages)
	insertsort(&ages)
	fmt.Println(ages)
}
