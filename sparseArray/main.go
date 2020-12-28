package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"strconv"
)

// 稀疏数组
// 1. 记录数组一共有几行几列，有多少个不同的值
// 2. 将具有不同的值的行、列及值记录在一个小规模的数组中，从而缩小规模

type Node struct {
	row int
	col int
	val int
}

// 定义全局文件路径
const filePath = "/Users/yaolixiao/code/goproject/src/app/dsaa/sparseArray/sparse.data"

// 原始数组 转 稀疏数组 => 相当于压缩
// 1. 因为稀疏数组规模不确定，只能使用切片存储
// 2. 将每一组数据存为 node结构体，插入到切片中
// 3. 将转化好的稀疏数组存储为文件 sparse.data
func originalToSparse(origin [11][11]int) {

	// 1. 存储行数、列数、默认值
	node := Node{
		row: 11,
		col: 11,
		val: 0,
	}
	sparse := []Node{node}

	// 2. 转化
	for i, v := range origin {
		for j, v2 := range v {
			if v2 != 0 {
				node = Node{
					row: i,
					col: j,
					val: v2,
				}
				sparse = append(sparse, node)
			}
		}
	}

	// 3. 存盘
	// 创建或打开一个文件，写入内容
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("os.OpenFile err=", err)
		return
	}

	// 及时关闭文件句柄
	defer file.Close()

	// 使用带缓存的 *writer
	writer := bufio.NewWriter(file)

	fmt.Println("稀疏数组：")
	for i, node := range sparse {
		writer.WriteString(fmt.Sprintf("%d\t%d\t%d\n", node.row, node.col, node.val))
		fmt.Printf("%d:  %d\t%d\t%d\n", i, node.row, node.col, node.val)
	}
	// writer.WriteString将内容写入缓存，需要Flush才能真正写入文件
	writer.Flush()
}

// 稀疏数组 转 原始数组 => 相当于解压
// 1. 从文件中读取稀疏数组
// 2. 恢复原始数据
func recoverOriginFromFile() {
	//
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err=", err)
		return
	}

	// 及时关闭文件句柄
	defer file.Close()

	// 定义原始数组，在读取文件内容时恢复数据
	var origin [11][11]int
	// 标记文件的第几行内容，从0开始
	var n int = 0

	// 带缓冲的reader
	reader := bufio.NewReader(file)
	// 循环读取文件内容，按行读取，io.EOF表示文件末尾
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// 踩坑
		str = strings.Replace(str, "\n", "", -1)

		// 第0行表示原始数组规模和默认值，所以跳过
		if n == 0 {
			// 踩坑
			n = 1
			continue
		}

		// str用 \t 分开有三个值，第一个是行，第二个是列，第三个是值
		strs := strings.Split(str, "\t")
		if len(strs) != 3 {
			panic("文件内容损坏")
		}

		fmt.Println(strs)
		fmt.Println(strs[2])

		// 赋值，恢复原始数组
		// 踩坑
		row, err := strconv.Atoi(strs[0])
		col, err := strconv.Atoi(strs[1])
		val, err := strconv.Atoi(strs[2])
		fmt.Printf("数据： %d %d %d \n", row, col, val)
		// todo 判断错误
		origin[row][col] = val

		// 增加行数
		n++
	}

	fmt.Println("恢复后的原始数组：")
	for _, v := range origin {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println("\n")
	}
}

func main() {

	// 案例：存储棋盘（0表示无子，1表示黑子，2表示白子）

	// 定义棋盘 11行11列
	var chess [11][11]int
	chess[1][2] = 1
	chess[2][3] = 2

	// 输出棋盘
	fmt.Println("原始数组：")
	for _, v := range chess {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println("\n")
	}

	// 原始数组转稀疏数组
	// originalToSparse(chess)

	// 恢复原始数组
	recoverOriginFromFile()
}