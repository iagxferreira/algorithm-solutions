package leetcode

import "math"

func maxProfit(prices []int) int {
	profit := 0
	price := math.MaxInt32
	for _, current := range prices {
		price = min(price, current)
		profit = max(profit, current-price)

	}

	return profit
}
