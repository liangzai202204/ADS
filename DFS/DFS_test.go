package DFS

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"strings"
	"testing"
)

func TestNumIsLands(t *testing.T) {
	tastCase := []struct {
		name  string
		input [][]byte
		out   int
	}{
		{
			name: "1",
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			out: 1,
		},
		{
			name: "1",
			input: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '1', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '0', '1'},
			},
			out: 4,
		},
		{
			name: "1",
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'0', '0', '0', '0', '1'},
				{'0', '0', '1', '0', '0'},
			},
			out: 3,
		},
	}
	for _, ts := range tastCase {
		t.Run(ts.name, func(t *testing.T) {
			res := numIslands(ts.input)
			assert.Equal(t, ts.out, res)
		})
	}

}
func numIslandsNor(grid [][]byte) (res int) {
	res = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				DFS(grid, i, j)
				res++
			}

		}
	}
	return
}
func DFS(grid [][]byte, x, y int) {
	if !inArea(grid, x, y) {
		return
	}

	if grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	DFS(grid, x+1, y)
	DFS(grid, x-1, y)
	DFS(grid, x, y+1)
	DFS(grid, x, y-1)

}
func inArea(grid [][]byte, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}

// numIslands leetcode-200 岛屿数量问题
func numIslands(grid [][]byte) (res int) {
	res = 0
	var dfs func(int, int)
	dfs = func(x int, y int) {
		if !(x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])) {
			return
		}

		if grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				res++
			}

		}
	}
	return
}

func TestFinfsameFS(t *testing.T) {
	cs := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	res := findDuplicate(cs)
	log.Print(res)

}

// 重复文件
func findDuplicate(paths []string) (res [][]string) {
	record := make(map[string][]string, len(paths))
	for _, v := range paths {
		splis := strings.Split(v, " ")
		dir := splis[0]
		for _, v2 := range splis[1:] {
			indexPosition := strings.IndexByte(v2, '(')
			content := v2[indexPosition : len(v2)-1]
			record[content] = append(record[content], dir+"/"+v2[:indexPosition])
		}
	}
	for _, dirF := range record {
		if len(dirF) >= 2 {
			res = append(res, dirF)
		}
	}
	return
}

func TestMaxSlidingWindow(t *testing.T) {

	tastCase := []struct {
		name  string
		input []int
		out   []int
		big   int
	}{
		{
			name:  "1",
			input: []int{1, 3, -1, -3, 5, 3, 6, 7},
			out:   []int{3, 3, 5, 5, 6, 7},
			big:   3,
		},
		{
			name:  "1",
			input: []int{7, 2, 4},
			out:   []int{7, 4},
			big:   2,
		},
	}
	for _, ts := range tastCase {
		t.Run(ts.name, func(t *testing.T) {
			res := maxSlidingWindowV1(ts.input, ts.big)
			assert.Equal(t, ts.out, res)
		})
	}
}

// 滑动窗口
func maxSlidingWindow(nums []int, k int) []int {
	left, right := k-1, k-1
	var mk int = maxList(nums[:k-1])
	res := make([]int, 0, len(nums)-k+1)
	res = append(res, mk)
	for right < len(nums)-1 {

		right++
		mk = max(nums[right], mk)
		// 增大窗口
		if right-left == 1 {
			res = append(res, mk)
			left++
		}
	}
	return res
}
func max(r int, y int) int {
	if r > y {
		return r
	}
	return y
}
func maxList(n []int) (res int) {
	res = n[0]
	for i := 0; i < len(n); i++ {
		res = max(res, n[i])
	}
	return res
}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func maxSlidingWindowV1(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
func pivotIndex(nums []int) int {
	left, right := 0, len(nums)-1
	leftNum, rightNum := 0, 0
	for right < left-1 {
		if leftNum <= rightNum {
			leftNum = nums[left] + leftNum
			left++
		}
		if leftNum <= rightNum {
			rightNum = nums[right] + rightNum
			right++
		}
	}
	return left
}
