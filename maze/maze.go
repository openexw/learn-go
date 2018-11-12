package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			// 读取文件
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 点，坐标
type point struct {
	i, j int
}

// 4个方向
var dirs = [4]point{
	{-1, 0},	// 上
	{0, -1},	// 左
	{1, 0},		// 下
	{0, 1},		// 右
}

// 扩展 point 方法
// 两个点相加
// @return 新的 point
func (p point) add(r point) point {
	return point{p.i+r.i, p.j+r.j}
}

// 获取点的值
// @return int grid的值
// @return bool 是否有值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i > len(grid) {
		return  0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return  grid[p.i][p.j], true
}
// 寻路
// @return [][]
func walk(maze [][]int,
	start,
	end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 队列 slice来实现对象
	Q := []point{start}
	for len(Q) > 0  {
		cur := Q[0]
		Q = Q[1:]

		// 退出条件
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 1 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	steps := walk(maze,
		point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})
	for _,row := range steps{
		for _,val := range row  {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}