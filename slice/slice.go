package slice

// Range 根据条件遍历删除切片内元素，返回一个新切片
func Range[T any](data []T, f func(n T) bool) []T {
	k := 0
	res := make([]T, len(data))
	for _, item := range data {
		if !f(item) {
			res[k] = item
			k++
		}
	}
	res = res[:k]
	return res
}
