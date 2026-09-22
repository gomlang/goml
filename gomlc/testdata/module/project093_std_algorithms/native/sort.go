package native

import "sort"

func StableOrder(keys []int) []int {
	order := make([]int, len(keys))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(i, j int) bool { return keys[order[i]] < keys[order[j]] })
	return order
}

func FindIndex(length, target int) (int, bool) {
	return sort.Find(length, func(index int) int {
		if target < index {
			return -1
		}
		if target > index {
			return 1
		}
		return 0
	})
}
