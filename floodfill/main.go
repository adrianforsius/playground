package main

import (
	"fmt"
)

func floodFill(s [][]int, color int, p point) [][]int {
	return nil
}

type point struct {
	x int
	y int
}

func main() {
	out := floodFill([][]int{
		{1, 1, 0},
		{1, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
	}, 2, point{0, 2})
	fmt.Println("out", out)
}
