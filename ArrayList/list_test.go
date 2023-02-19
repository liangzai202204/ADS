package ArrayList

import (
	"log"
	"math/rand"
	"testing"
)

func FuzzName(f *testing.F) {
	f.Fuzz(func(t *testing.T) {

	})
}
func TestList_List(t *testing.T) {

}

func TestLeft_bound(t *testing.T) {
	var nums []int = []int{1, 2, 4, 5, 6}
	tg := leftBound(nums, 3)
	log.Print(tg)
	l := len(nums) - 1
	rand.Seed(int64(l))
	r := rand.Intn(l) + 1
	log.Print(r)

}
func leftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r - l) / 2
		if target == nums[mid] {
			return mid
		} else if target >= nums[mid] {
			r--
		} else if target <= nums[mid] {
			l++
		}
	}
	return l + 1
}

//	func TestCon(t *testing.T) {
//		var nums []int = []int{1, 2, 4, 5, 6}
//		c := Constructor2(nums)
//		res := c.PickIndex()
//		log.Print(res)
//		l := len(nums) - 1
//		r := rand.Intn(l) + 1
//		log.Print(r)
//	}
func maxSlidingWindow(nums []int, k int) (res []int) {
	left, right := 0, 0
	mk := nums[0]
	res = make([]int, 0, len(nums)-k+1)
	var mmax func(int, int) int
	mmax = func(r int, y int) int {
		if r > y {
			return r
		}
		return y
	}
	for right < len(nums) {
		if k == 1 {
			mk = nums[left]
		}
		mk = mmax(nums[right], mk)
		right++
		// 增大窗口
		if right >= k {
			res = append(res, mk)
		}

		for (right - left) >= k {
			left++
		}
	}
	return
}
