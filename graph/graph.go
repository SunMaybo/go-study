package graph

//岛屿数量
var graph1 = [5][5]int{
	{1, 1, 0, 0, 0},
	{1, 1, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 1, 1},
}
//
//var graph1 = [5][5]int{
//{1, 1, 1, 1, 0},
//{1, 1, 0, 1, 0},
//{1, 1, 0, 0, 0},
//{0, 0, 0, 0, 0},
//}

func SearchCount() int {
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if graph1[i][j] == 1 {
				Search(graph1[:], i, j)
				count++

			}
		}
	}
	return count
}

func Search(graph [][5]int, i, j int) {
	if i < 0 || i >= 5 || j < 0 || j >= 5 || graph[i][j] == 0 {
		return
	}
	graph[i][j] = 0
	Search(graph, i+1, j)
	Search(graph, i-1, j)
	Search(graph, i, j+1)
	Search(graph, i, j-1)
}
