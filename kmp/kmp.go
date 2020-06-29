package kmp

func KmpSearch(key, text string) (bool, int) {
	length := len(key)
	total := 0
	isExist := false
	if length > len(text) {
		return false, 0
	}
	var arrays []int
	for i := 0; i < length; i++ {
		arrays = append(arrays, useCacheArrays(i, key))
	}
	i := 0
	for i <= len(text)-length {
		itemCount := 0
		for j := 0; j < len(key); j++ {
			if key[j] == text[i+j] {
				itemCount++
			} else {
				//简历数组优化调位置
				i = i + arrays[j]
				break
			}
		}
		if itemCount == length {
			total++
			isExist = true
			i++
		}
	}

	return isExist, total
}
func useCacheArrays(postion int, key string) int {
	if postion == 0 || postion == 1 {
		return 1
	}
	max := 1
	for i := 1; i < postion; i++ {
		if string(key[i:postion]) == string(key[:(postion - i)]) {
			if max < i {
				max = i
			}
		}
	}
	return max
}
