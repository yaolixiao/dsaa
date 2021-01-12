package main

import "fmt"

// 选择排序（从大到小）
// 1. 假设第一个元素最大
// 2. 找出实际最大元素：拿第一个元素和后面所有元素相比较
// 3. 将实际最大元素和第一个元素互换位置
// 4. 从第二个元素开始到数组末尾，按照1，2，3步重复执行，得到排名第二的元素
// 5. 以此类推...
func selectsort(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxindex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxindex = i
			}
		}
		if maxindex != j {
			arr[j], arr[maxindex] = arr[maxindex], arr[j]
		}
	}
}

func main() {
	ages := [5]int{20, 18, 12, 35, 98}
	fmt.Println(ages)
	selectsort(&ages)
	fmt.Println(ages)
}
