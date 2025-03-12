package main

import (
	"math"
)

// dpResult holds the total items (sum) and the total number of packs (count).
type dpResult struct {
	sum   int
	count int
}

// calculatePacks returns how many packs of each size are needed to satisfy
// an order with minimal overage (sum) and then minimal pack count.
func calculatePacks(order int) map[int]int {
	// dp[t] = best way to cover t items (or more) with minimal sum, then minimal count.
	// chosen[t] = which pack was chosen to improve dp[t].
	dp := make([]dpResult, order+1)
	chosen := make([]int, order+1)

	// Initialize dp[0].
	dp[0] = dpResult{sum: 0, count: 0}
	for t := 1; t <= order; t++ {
		dp[t] = dpResult{sum: math.MaxInt, count: math.MaxInt}
	}

	// Fill dp table.
	for t := 1; t <= order; t++ {
		for _, packSize := range packSizes {
			prevIndex := t - packSize
			if prevIndex < 0 {
				// If packSize > t, we just take one pack of that size.
				candidateSum := packSize
				candidateCount := 1

				if (candidateSum < dp[t].sum) ||
					(candidateSum == dp[t].sum && candidateCount < dp[t].count) {
					dp[t] = dpResult{sum: candidateSum, count: candidateCount}
					chosen[t] = packSize
				}
			} else {
				// Use dp[prevIndex] + this packSize.
				prev := dp[prevIndex]
				candidateSum := prev.sum + packSize
				candidateCount := prev.count + 1

				if (candidateSum < dp[t].sum) ||
					(candidateSum == dp[t].sum && candidateCount < dp[t].count) {
					dp[t] = dpResult{sum: candidateSum, count: candidateCount}
					chosen[t] = packSize
				}
			}
		}
	}

	// Reconstruct which packs were used.
	result := make(map[int]int)
	t := order
	for t > 0 {
		packUsed := chosen[t]
		result[packUsed]++
		if packUsed > t {
			break
		}
		t -= packUsed
	}

	return result
}
