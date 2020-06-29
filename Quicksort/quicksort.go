package Quicksort

func Sort(start, end int, c []int) {
	flag := c[start]
	right := end
	left := start
loop:
	for {
		if left == right {
			c[left] = flag
			if left-1 > start {
				Sort(start, left-1, c)
			}
			if left+1 < end {
				Sort(left+1, end, c)
			}
			break
		}
		if c[right] > flag {
			c[left] = c[right]
			for {
				if left == right {
					c[left] = flag
					if left-1 > start {
						Sort(start, left-1, c)
					}
					if left+1 < end {
						Sort(left+1, end, c)
					}

					break loop
				}
				if c[left] < flag {
					c[right] = c[left]
					break
				}
				left++
			}
		}
		right--
	}
}
