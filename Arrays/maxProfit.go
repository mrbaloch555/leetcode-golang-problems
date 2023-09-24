package main

import "fmt"

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	min_price := prices[0]
	max_profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min_price {
			min_price = prices[i]
		} else {
			if max_profit < prices[i]-min_price {
				max_profit = prices[i] - min_price
			}
		}
	}
	return max_profit
}
