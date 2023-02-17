package DFS

import (
	"github.com/stretchr/testify/assert"
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
