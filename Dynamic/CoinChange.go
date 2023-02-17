// Package Dynamic 动态规划 零钱问题

package Dynamic

//import "math"
//
//func coinChange(coins []int, amount int) int {
//	return dp(coins, amount)
//
//}
//func dp(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//	if amount < 0 {
//		return -1
//	}
//	res := math.MaxInt
//	for coin := range coins {
//		r := dp(coins, amount-coin)
//		if res == -1 {
//			continue
//		}
//		if res >= r+1 {
//			res = r + 1
//		}
//	}
//	if res == math.MaxInt {
//		return -1
//	}
//	return res
//
//}
