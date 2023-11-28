package permutation

// Subsets Given an integer array nums of unique elements, return all possible subsets.
func Subsets(nums []int) [][]int {
	var result [][]int
	for _, num := range nums {
		var newSubsets [][]int
		for _, item := range result {
			tempItem := make([]int, len(item))
			copy(tempItem, item)
			newSubsets = append(newSubsets, append(tempItem, num))
		}
		result = append(result, newSubsets...)
		result = append(result, []int{num})
	}
	result = append(result, []int{})
	return result
}
