package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file ,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row,col int

	fmt.Fscanf(file,"%d %d",&row,&col)
	fmt.Printf("row:%d,col:%d\n",row,col)

	//二维切片（数组）的行、列是人们抽象出来的，便于好理解，本质就是一个大slice里面有很多小slice，每个小slice里面有很多数
	maze := make([][]int,row)//定义二维slice或数组，先定义大slice里有多少组小slice，即抽象的行数
	for i := range maze{
		maze[i] = make([]int,col)//再定义每个小slice里有多少个元素，即抽象的列数
		for j := range maze[i] {
			fmt.Fscanf(file,"%d",&maze[i][j])//因为maze[i][j]需要传入fmt.Fscanf函数进行赋值，所以需要取地址
			fmt.Printf("maze[%d][%d]=%d\n",i,j,maze[i][j])
		}
	}
	return maze
}

type point struct{
	i ,j int
}

var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1},
}
func (p point) add(r point) point {
	return point{p.i + r.i,p.j + r.j}
}

func walk(maze [6][5]int,start,end point){
	steps := make([][]int,len(maze))
	for i := range steps {
		steps[i] = make([]int,len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0{
		cur := Q[0]
		Q = Q[1:]

		for _,dir := range dirs{
			next := cur.add(dir)
		}
	}
}

func main() {
	//maze := readMaze("maze/maze.txt")
	maze := [6][5]int{ {0,1,0,0,0},
					   {0,0,0,1,0},
					   {0,1,0,1,0},
					   {1,1,1,0,0},
					   {0,1,0,0,1},
					   {0,1,0,0,0}}//由于readMaze读出来的数据不对，所以先直接赋值

	//fmt.Println("maze:",maze)
	for _,row := range maze{
		for _,value := range row {
			fmt.Printf("%d ",value)
		}
		fmt.Println()
	}
	//len(maze)计算大slice里面有多少个小slice，len(maze[0])计算小slice里面有多少个元素
	walk(maze,point{0,0},point{len(maze) -1,len(maze[0]) - 1})
}
