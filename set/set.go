package set

// Intersect 求交集
func Intersect(slice1, slice2 []int64) []int64 {
	var m = make(map[int64]int)
	for _, v := range slice1 {
		m[v] = 1
	}

	var result []int64
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v)
		}
	}
	return result
}

// Difference 求差集
func Difference(slice1 []int64, slice2 []int64) []int64 {
	var m = make(map[int64]int)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v] = 1
	}

	var result []int64
	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = 1
		}
	}
	return result
}
