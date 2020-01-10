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
	// 从文件中读取第一行的内容(根据第二个参数format读取文件内容) 赋值给row col (注意要取地址)
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			// fmt.Fscanf()函数是对文件的扫描读取
			// 从文件中扫描读取一个数字 赋值给 maze[i][j]
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

// 移动方向： 上 左 下 右
var dirs = [4]point {
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 待探索的点 存放在队列中
	Q := []point{ start }

	// 队列为空之后则停止探索 结束循环
	for len(Q) > 0 {
		// 待探索的当前点
		cur := Q[0]
		// 将当前要探索的点推出队列
		Q = Q[1:]

		// 到达迷宫终点
		if cur == end {
			break
		}
		
		for _, dir := range dirs {
			next := cur.add(dir)

			// next 必须满足条件： 1.不能越界maze数组 2.不能是墙壁 3.不能是已经走过的点
			val, ok := next.at(maze)
			// 越界或墙壁 则直接进入下一个循环
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			// next这个点是已经走过的点
			if val != 0 {
				continue
			}
			// 回到了起点
			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			// 将满足条件的next点添加到待探索的队列中
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println("----------")

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
