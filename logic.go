package main

import "sort"

// Request represents a JSON request containing the number of items to order.
type Request struct {
	Items int `json:"items"`
}

// Response represents the JSON response containing the calculated pack distribution.
type Response struct {
	Packs map[int]int `json:"packs"`
}

// calculatePacks determines the optimal number of packs to fulfill an order.
// It prioritizes minimizing the number of items sent, then minimizing the number of packs used.
func calculatePacks(items int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	result := make(map[int]int)

	remaining := items

	for _, pack := range packSizes {
		if remaining == 0 {
			break
		}
		count := remaining / pack
		if count > 0 {
			result[pack] = count
			remaining -= count * pack
		}
	}

	if remaining > 0 {
		result[packSizes[len(packSizes)-1]]++
	}

	return result
}
