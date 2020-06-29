package main

import (
	"fmt"
	"algorithm-study/tree"
	"algorithm-study/substring"
)

type M struct {
	A *int
}

func main() {
	//c := []int{6, 3, 8, 1, 0, 8, 4, 8, 9, 5, 8, 7}
	//heap.HeapSort(c)
	//Quicksort.Sort(0, len(c)-1, c)
	//fmt.Println(substring.Substring("pwwkew"))
	//fmt.Println(substring.LengthOfLongestBack("abcba"))
	//p,q:=binary.BuildTree()
	//fmt.Println(binary.IsSameTree(p,q))
	//r := binary.BuildBinaryTree()
	//binary.FrontScan(r)
	//binary.MiddleScan(r)
	//max := new(int)
	//binary.Depth(r, 0, max)
	//fmt.Println(*max)
	//fmt.Println(dp.Dp(6))
	//root := list.New()
	//////root:=list.NewTest()
	//root = list.Flip(root, 3)
	//
	//for root.Next != nil {
	//	fmt.Println(root.Val)
	//	root = root.Next
	//}

	//list.NewTest()
	//a := []int{1, 2, 3, 1}
	//fmt.Println(dp.Rob(a))
	//fmt.Println(graph.SearchCount())
	//graph.Matrix()
	//b:=890
	//a:=43
	//fmt.Println(strconv.FormatInt(int64(a),2))
	//fmt.Println(strconv.FormatInt(int64(b),2))
	//fmt.Println("\n")
	//for b!=0{
	//	sum := a^b;
	//	carry := (a&b)<<1;
	//	a = sum;
	//	b = carry;
	//	fmt.Println(strconv.FormatInt(int64(sum),2))
	//	fmt.Println(strconv.FormatInt(int64(carry),2))
	//	fmt.Println("\n")
	//}
	//fmt.Println(a)
	//n:=2048*2*2*2*2
	//fmt.Println(strconv.FormatInt(int64(n),2))
	//fmt.Println(n&(n-1))
	//	graph.TranM()

	//var a int
	//var m M
	//m.A=a
	//test(&a, &m)
	//test2(a)

	//fmt.Println(dp.NumSquares(64))

	//trie := tree.Constructor()
	//trie.Insert("abs")
	//trie.Insert("abc")
	//trie.Insert("dbc")
	//fmt.Println(trie.Search("dbca"))
	//fmt.Println(trie.StartsWith("e"))
	//print(&trie)
	fmt.Println(substring.Trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))

	fmt.Println(substring.MinDistance("我喜欢你", "我不喜欢"))

}
func print(trie *tree.Trie) {
	if trie.Next == nil {
		return
	}
	for _, t := range trie.Next {
		fmt.Println(string(*t.Val))
		print(t)
	}
}
func test(a *int, m *M) {
	m.A = a
}

type S struct {
	M *int
}

func ref(y *int, z *S) {
	z.M = y
}

func test2(i interface{}) {

}
