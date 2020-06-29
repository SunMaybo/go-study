package graph

import "fmt"

//太平洋和大西洋交汇坐标
var graph = [][]int{
	{1, 2, 2, 3, 5},
	{3, 2, 3, 4, 4},
	{2, 4, 5, 3, 1},
	{6, 7, 1, 4, 5},
	{5, 1, 1, 2, 4},
}

func Matrix() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			isOk1 := IsLeftTopInput(graph, i, j)
			isOk2 := IsRightBottomInput(graph, i, j)
			if isOk1 && isOk2 {
				fmt.Println(i, j)
			}
		}
	}
}

func IsLeftTopInput(graph [][]int, i, j int) bool {
	if (i < 5 && j == 0) || (i == 0 && j < 5) {
		return true
	}
	if i-1 >= 0 && graph[i-1][j] <= graph[i][j] {
		isOk1 := IsLeftTopInput(graph, i-1, j)
		if isOk1 {
			return true
		}
	}
	if j-1 >= 0 && graph[i][j-1] <= graph[i][j] {
		isOk2 := IsLeftTopInput(graph, i, j-1)
		if isOk2 {
			return true
		}
	}
	return false
}
func IsRightBottomInput(graph [][]int, i, j int) bool {
	if (i < 5 && j == 4) || (i == 4 && j < 5) {
		return true
	}
	if i+1 <= 4 && graph[i+1][j] <= graph[i][j] {
		isOk3 := IsRightBottomInput(graph, i+1, j)
		if isOk3 {
			return true
		}
	}
	if j+1 <= 4 && graph[i][j+1] <= graph[i][j] {
		isOk4 := IsRightBottomInput(graph, i, j+1)
		if isOk4 {
			return true
		}
	}
	return false
}

//旅行商问题
var MAX int = 2 << 30

var tran = [][]int{
	{MAX, 3, MAX, 8, 9},
	{3, MAX, 3, 10, 5},
	{MAX, 3, MAX, 4, 3},
	{8, 10, 4, MAX, 20},
	{9, 5, 3, 20, MAX},
}


func tranMaxtrix(start, i int, a ... int) int {

	if len(a) == 1 {
		return tran[i][a[0]] + tran[start][a[0]]
	}

	min := MAX
	positon := 0
	for j := 0; j < len(a); j++ {
		index := a[j]
		var b []int
		for i := 0; i < len(a); i++ {
			if j == i {
				continue
			}
			b = append(b, a[i])
		}
		_min := tran[i][a[j]] + tranMaxtrix(start, index, b...)
		if _min < min {
			min = _min
			positon = j
		}

	}
	fmt.Println(i, a[positon])
	return min

}
func TranM() {
	fmt.Println(tranMaxtrix(0, 0, 1, 2, 3, 4))

}
