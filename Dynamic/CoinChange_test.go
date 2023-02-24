package Dynamic

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"sort"
	"testing"
)

func TestCoin(t *testing.T) {
	var coins []int = []int{1, 2, 5}
	var a int = 11
	res := CoinChange(coins, a)
	log.Print(res)
}

func CoinChange(coins []int, amount int) int {
	initDp(amount)
	return dp1(coins, amount)

}

var Dp []int

func initDp(amount int) {
	Dp = make([]int, amount+1)
	for i := 0; i < len(Dp); i++ {
		Dp[i] = math.MaxInt
	}

}
func dp1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if Dp[amount] != math.MaxInt {
		return Dp[amount]
	}
	res := math.MaxInt
	for _, i := range coins {
		// log.Print(i)
		r := dp1(coins, amount-i)
		if r == -1 {
			continue
		}
		if res > r+1 {
			res = r + 1
		}
	}
	if res == math.MaxInt {
		Dp[amount] = -1
	} else {
		Dp[amount] = res
	}
	return Dp[amount]

}

func fib(n int) int {
	var nums = make([]int, n+1)
	return fibs(nums, n)
}

func fibs(nums []int, n int) int {
	if n == 1 || n == 0 {
		nums[n] = n
		return nums[n]
	}
	if nums[n] != 0 {
		return nums[n]
	}
	nums[n] = fibs(nums, n-1) + fibs(nums, n-2)
	return nums[n]
}
func TestFib(t *testing.T) {
	var n int = 8
	res := fib(n)
	log.Print(res)
}
func TestLengthOfLIS(t *testing.T) {
	var n = []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	res := lengthOfLIS(n)
	log.Print(res)
}

//func lengthOfLIS1(nums []int) int {
//	var nl = 1
//	var n []int
//	var alln []int
//	n = append(n, nums[0])
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			continue
//		}
//		if n[len(n)-1] > nums[i+1] {
//			if isOk(n, nums[i+1]) {
//				alln = append(alln, nl)
//				nl++
//			} else {
//				alln = append(alln, nl)
//				nl = 1
//			}
//		}
//		if n[len(n)-1] < nums[i+1] {
//			n = append(n, nums[i+1])
//			nl++
//		}
//	}
//	return maxq(alln)
//}
//
//func isOk(n []int, j int) bool {
//	for i := 0; i < (len(n) - 1); i++ {
//		if j == n[i] {
//			return false
//		}
//		if j > n[i] {
//			return true
//		}
//	}
//	return false
//}
//func maxq(nums []int) int {
//	var in = nums[0]
//	for i := 1; i < (len(nums) - 1); i++ {
//		if nums[i] > nums[i-1] {
//			in = nums[i]
//		}
//	}
//	return in
//}

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				k := dp[j] + 1
				if dp[i] < k {
					dp[i] = k
				}
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res + 1
}

// 俄罗斯套娃，基于二分查找的动态规划求解
func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	var f []int
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
func TestMin(t *testing.T) {
	tastCase := []struct {
		name  string
		input [][]int
		out   int
	}{
		{
			name:  "[][]int{{5, 1, 3}, {6, 5, 4}, {7, 8, 9}}",
			input: [][]int{{5, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			out:   13,
		},
		{
			name:  "{{82,81},{69,33}}",
			input: [][]int{{82, 81}, {69, 33}},
			out:   114,
		},
		{
			name:  "{{82,81},{69,33}}",
			input: [][]int{{-48}},
			out:   -48,
		},
		{
			name:  "[][]int{{5, 1, 3}, {6, 5, 4}, {7, 8, 9}}",
			input: [][]int{{-80, -13, 22}, {83, 94, -5}, {73, -48, 61}},
			out:   -66,
		},
		{
			name: "[][]int{{5, 1, 3}, {6, 5, 4}, {7, 8, 9}}",
			input: [][]int{
				{-51, -35, 74},
				{-62, 14, -53},
				{94, 61, -10}},
			out: -98,
		},
	}
	for _, ts := range tastCase {
		t.Run(ts.name, func(t *testing.T) {
			res := minFallingPathSum(ts.input)
			assert.Equal(t, ts.out, res)
		})
	}
}
func minFallingPathSum(matrix [][]int) int {
	l := len(matrix)
	if l == 1 {
		return matrix[0][0]
	}
	er := len(matrix[0])
	re = make([][]int, l)
	for id, _ := range re {
		re[id] = make([]int, er)
	}
	res := math.MaxInt
	for i := 0; i < l; i++ {
		res = min(res, dpm(matrix, 0, i))
	}
	return res
}

var re [][]int

func dpm(matrix [][]int, k, h int) int {
	if h < 0 || h >= len(matrix) || k >= len(matrix) {
		return math.MaxInt
	}
	if (k + 1) == len(matrix) {
		re[k][h] = matrix[k][h]
		return re[k][h]
	}
	if re[k][h] != 0 {
		return re[k][h]
	}
	r := dpm(matrix, k+1, h)
	left := dpm(matrix, k+1, h-1)
	d := dpm(matrix, k+1, h+1)
	re[k][h] = min(r, min(left, d)) + matrix[k][h]
	return re[k][h]
}
func min(r, p int) int {
	if r > p {
		return p
	}
	return r
}
func maxProfit(k int, prices []int) int {
	var zuan []int
	res := 0
	now := 0
	end := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			now = prices[i]
		}
		if prices[i] > prices[i+1] {
			end = prices[i]
			zuan = append(zuan, end-now)
		}
	}
	if len(zuan) == 0 {
		return 0

	}
	for j := 0; j < len(zuan); j++ {
		if j > k-1 {
			break
		}
		res = res + zuan[j]
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
