package slice

// Range 根据条件遍历删除切片内元素，返回一个新切片
func Range(data []int, f func(n int) bool) []int {
	k := 0
	res := make([]int, len(data))
	for _, item := range data {
		if !f(item) {
			res[k] = item
			k++
		}
	}
	res = res[:k]
	return res
}
