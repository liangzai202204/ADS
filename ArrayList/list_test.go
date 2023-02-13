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

//func TestCon(t *testing.T) {
//	var nums []int = []int{1, 2, 4, 5, 6}
//	c := Constructor2(nums)
//	res := c.PickIndex()
//	log.Print(res)
//	l := len(nums) - 1
//	r := rand.Intn(l) + 1
//	log.Print(r)
//}
