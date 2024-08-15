package go_001

// 查询数组中某个元素的下标
func ListFindIndex(list []int, target int) int {
	index := -1
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			index = i
		}
	}
	return index
}
