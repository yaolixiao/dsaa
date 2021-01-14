package main

import (
	"fmt"
)

// 递归算法
// 案例：走迷宫

func playing(maze *[8][7]int, i int, j int) bool {

	if maze[6][5] == 2 {
		return true
	} else {

		// 每次递归进来这个点，如果是0则可以向前走
		if maze[i][j] == 0 {

			// 假设能走通，设置为2
			maze[i][j] = 2

			// 向前走，下右上左
			if playing(maze, i+1, j) {
				return true
			} else if playing(maze, i, j+1) {
				return true
			} else if playing(maze, i-1, j) {
				return true
			} else if playing(maze, i, j-1) {
				return true
			} else {
				// 这里表示走不通，设置为3
				maze[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {

	// 定义迷宫，是个二维数组
	// 0 代表没走过的路
	// 1 代表一堵墙
	// 2 代表走过，且能走通的路
	// 3 代表走过，但走不通的路
	var maze [8][7]int

	// 初始化迷宫
	for i := 0; i < 8; i++ {
		maze[i][0] = 1
		maze[i][6] = 1
	}
	for i := 0; i < 7; i++ {
		maze[0][i] = 1
		maze[7][i] = 1
	}

	maze[3][1] = 1
	maze[3][2] = 1

	// 显示迷宫
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d\t", maze[i][j])
		}
		fmt.Println()
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	playing(&maze, 1, 1)
	// 显示迷宫
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d\t", maze[i][j])
		}
		fmt.Println()
		fmt.Println()
	}
}
