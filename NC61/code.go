package main

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	// write code here
	cache := map[int]int{}

	for i, n := range numbers {
		if x, ok := cache[target-n]; ok {
			return []int{x + 1, i + 1}
		} else {
			cache[n] = i
		}
	}
	return []int{}
}
