package heap

func maxHeap(root int, end int, c []int) {
	n := 2*root + 1
	if n > end {
		return
	}

	if n+1 <= end && c[n+1] >= c[n] {
		n++
	}
	if c[n] >= c[root] {
		c[root], c[n] = c[n], c[root]
		maxHeap(n, end, c)
	}
}

//降序排序
func HeapSort(c []int) {
	var n = len(c) - 1
	for root := n / 2; root >= 0; root-- {
		maxHeap(root, n, c)
	}
	for i := n; i >= 0; i-- {
		if c[0] >= c[i] {
			c[0], c[i] = c[i], c[0]
			maxHeap(0, i-1, c)
		}
	}

}
