package substring

import "fmt"

func Substring(text string) int {
	max := 0
	dataMap := make(map[string]int)
	length := 0
	for _, item := range text {
		if _, ok := dataMap[string(item)]; ok {
			if max < length {
				max = length
			}
			length = 1
			continue
		}
		dataMap[string(item)] = 1
		length++
		if max < length {
			max = length
		}

	}
	return max
}

func LengthOfLongestBack(text string) int {
	max := 0
	for i := 0; i < len(text); i++ {
		odd := 0
		even := 0
		for j := 1; j <= i && i > 0 && i+j < len(text); j++ {
			if text[i-j] == text[i+j] {
				odd++
			} else {
				break
			}
		}

		if odd*2+1 > max {
			max = odd*2 + 1
		}
		for j := 0; j <= i && i+j+1 < len(text); j++ {
			if text[i-j] == text[i+j+1] {
				even++
			} else {
				break
			}
		}
		if even*2 > max {
			max = even * 2
		}
	}
	return max
}

func MinDistance(a, b string) int {
	_a := []rune(a)
	_b := []rune(b)
	dp := make([][]int, len(_a))
	for i := 0; i < len(_a); i++ {
		dp[i] = make([]int, len(_b))
	}
	if _a[0] == _b[0] {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 0; i < len(_a); i++ {
		for j := 0; j < len(_b); j++ {
			if i == j && i == 0 {
				continue
			}
			min := 2 << 32
			if i-1 >= 0 && dp[i-1][j]+1 <= min {
				min = dp[i-1][j] + 1
			}
			if j-1 >= 0 && dp[i][j-1]+1 <= min {
				min = dp[i][j-1] + 1
			}
			if i-1 >= 0 && j-1 >= 0 {
				if _a[i] == _b[j] && dp[i-1][j-1] <= min {
					min = dp[i-1][j-1]
				}
				if _a[i] != _b[j] && dp[i-1][j-1]+1 <= min {
					min = dp[i-1][j-1] + 1
				}
			}
			dp[i][j] = min
		}
	}
	for i := 0; i < len(_a); i++ {
		fmt.Println(dp[i])
	}

	return dp[len(_a)-1][len(_b)-1]
}

//func Trap(height []int) int {
//	max := 0
//	for i := 0; i < len(height)-1; i++ {
//		for j := i + 1; j < len(height); j++ {
//			min := 0
//			if height[i] < height[j] {
//				min = height[i]
//			} else {
//				min = height[j]
//			}
//			if (j-i)*min >= max {
//				max = (j - i) * min
//			}
//		}
//	}
//	return max
//}

func Trap(height []int) int {
	max := 0
	for i := 1; i < len(height)-1; i++ {
		left := i - 1
		right := i + 1
		leftMax := height[i]
		rightMax := height[i]
		for j := left; j >= 0; j-- {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}
		for h := right; h < len(height); h++ {
			if height[h] > rightMax {
				rightMax = height[h]
			}
		}
		min := 0
		if leftMax <= height[i] || rightMax <= height[i] {
			min = 0
		} else if leftMax > rightMax {
			min = rightMax
		} else {
			min = leftMax
		}
		if min > 0 {
			max += min - height[i]
		}
	}
	return max
}


//
