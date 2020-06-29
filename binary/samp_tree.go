package binary

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Input:     1         1
/// \       / \
//2   3     2   3
//
//[1,2,3],   [1,2,3]
//
//Output: true
//Example 2:
//
//Input:     1         1
///           \
//2             2
//
//[1,2],     [1,null,2]
//
//Output: false
//Example 3:
//
//Input:     1         1
/// \       / \
//2   1     1   2
//
//[1,2,1],   [1,1,2]
//
//Output: false

func BuildTree() (*TreeNode, *TreeNode) {

	rootp := new(TreeNode)
	rootq := new(TreeNode)
	rootp.Val = 1
	rootq.Val = 1

	rootp.Left = new(TreeNode)
	rootp.Right = new(TreeNode)
	rootq.Left = new(TreeNode)
	rootq.Right = new(TreeNode)
	rootp.Left.Val = 3
	rootq.Left.Val = 2
	rootq.Right.Val = 3

	return rootp, rootq
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		isLeft := IsSameTree(p.Left, q.Left)
		isRight := IsSameTree(p.Right, q.Right)
		if isLeft && isRight {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

func BuildBinaryTree() (*TreeNode) {

	rootp := new(TreeNode)

	rootp.Val = 1

	rootp.Left = new(TreeNode)
	rootp.Right = new(TreeNode)
	rootp.Left.Val = 2
	rootp.Right.Val = 3
	rootp.Right.Left = new(TreeNode)
	rootp.Right.Left.Val = 5
	rootp.Left.Left = new(TreeNode)
	rootp.Left.Left.Val = 4
	rootp.Left.Left.Right = new(TreeNode)
	rootp.Left.Left.Right.Val = 6
	rootp.Left.Left.Right.Right = new(TreeNode)
	rootp.Left.Left.Right.Right.Val = 7

	return rootp
}

func FrontScan(p *TreeNode) {

	if p == nil {
		return
	}
	fmt.Println(p.Val)
	FrontScan(p.Left)
	FrontScan(p.Right)

}
func Depth(p *TreeNode, depth int, max *int) {

	if p == nil {
		return
	}
	depth++
	if depth > *max {
		*max = depth
	}
	Depth(p.Left, depth, max)
	Depth(p.Right, depth, max)

}
func MiddleScan(p *TreeNode) {
	if p == nil {
		return
	}

	MiddleScan(p.Left)
	fmt.Println(p.Val)
	MiddleScan(p.Right)
}
func BackScan(p *TreeNode) {
	if p == nil {
		return
	}
	MiddleScan(p.Left)
	MiddleScan(p.Right)
	fmt.Println(p.Val)
}
