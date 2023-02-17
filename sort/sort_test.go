package sort

import (
	"log"
	"sort"
	"testing"
)

// sort.Slice 使用例子 测试
func TestDemoSortLice(t *testing.T) {
	demoSortSlice()
}
func demoSortSlice() {
	a := []int{6, 3, 9, 8, 1, 2, 5, 7}
	sort.Slice(a, func(i, j int) bool {
		// 返回值便是降序还是升序
		return a[i] > a[j]
	})
	log.Print(a)
	//[9 8 7 6 5 3 2 1]
	sort.Slice(a, func(i, j int) bool {
		log.Print(i, j)
		return a[i] < a[j]
	})
	log.Print(a)
}
func TestLis(t *testing.T) {
	a := []int{6, 3, 9, 8, 1, 2, 5, 7}
	lis(a)
}
func lis(f []int) int {
	for _, h := range f {
		log.Print(h)
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	log.Print(len(f))
	return len(f)
}

var nums [2][2]int

// 测试二维数组初始化问题
func TestEr(t *testing.T) {
	nums[1][1] = 3
	log.Print(nums)
	a := [len(nums)][len(nums)]int{}
	B := make([][len(nums)]int, len(nums))
	log.Print(B)
	//a[1] = []int{1, 1
	log.Print(len(a))
}
func change() {

}
