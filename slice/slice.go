package slice

import "fmt"

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

// DelDuplicate 列表去重
func DelDuplicate(data interface{}) interface{} {
	m := make(map[string]struct{})
	switch slice := data.(type) {
	case []string:
		res := make([]string, 0, len(data.([]string)))
		for _, item := range slice {
			if _, ok := m[item]; !ok {
				m[item] = struct{}{}
				res = append(res, item)
			}
		}
		return res
	case []int64:
		res := make([]int64, 0, len(data.([]int64)))
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := m[key]; !ok {
				m[key] = struct{}{}
				res = append(res, item)
			}
		}
		return res
	default:
		return nil
	}

}
