package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int // i = row, j = col
}
type grids [][]int

type result struct {
	grids
	steps  int
	points []point
}

// 根据迷宫规则，每一个点周围都有4个点包围，设置初始点,按照逆时针旋转(不一定要逆时针，只是逆时针更加有规律，随便什么写都可以)
// 定义dirs数组，放入四个以(0,0)为中心的初始点~
var dirs = [4]point{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

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

func readMaze() grids {
	file, err := os.Open("Algorithms/maze/maze.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &col, &row) // 6行5列
	maze := make(grids, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//func walk(maze [][]int, start, end point) [][]int {
//	steps := make([][]int, len(maze))
//	for i := range steps {
//		steps[i] = make([]int, len(maze[i]))
//	}
//	Q := []point{start}
//	for len(Q) > 0 {
//		cur := Q[0]
//		Q = Q[1:]
//		if cur == end {
//			break
//		}
//		for _, dir := range dirs {
//			next := cur.add(dir)
//			// maze at next is 0
//			// and steps at next is 0
//			// next != start
//			val, ok := next.at(maze)
//			if !ok || val == 1 {
//				continue
//			}
//			val, ok = next.at(steps)
//			if !ok || val != 0 {
//				continue
//			}
//			if next == start {
//				continue
//			}
//			curSteps, _ := cur.at(steps)
//			steps[next.i][next.j] = curSteps + 1
//			Q = append(Q, next)
//		}
//	}
//	return steps
//}
// walk 广度搜索优先
// 三个状态：已知已探索，已知为探索，未知
func walk(maze grids, start, end point) result {
	var result result
	steps := make(grids, len(maze)) // 构建基础二维数组(复刻一个maze但是初始值为0的二维数组)
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start} // 初始化任务队列，队列中装有已知未探索的Point
	for len(Q) > 0 {    // 循环队列
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, v := range dirs {
			// 判断条件
			// 边界：(i, j)必须大于0且小于i < len(maze), j < len(maze(0)) // 下一步必须是0
			// 墙: 遇见墙(1)无法行走
			// 起点：走回起点，无法行走
			// 在已走过的路，不会往回走// steps必须是0，非0说明走过
			//next := point{cur.i + v.i, cur.j + v.j}
			next := cur.add(v)
			if next == start {
				continue
			}
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			result.steps = curStep + 1
			Q = append(Q, next)
		}
	}
	result.grids = steps
	return result
}
func main() {
	maze := readMaze()
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println("路径图：")
	for _, v := range steps.grids {
		for _, i := range v {
			fmt.Printf("%3d", i)
		}
		fmt.Println()
	}
	fmt.Printf("一共需要走%d步\n", steps.steps)
}
