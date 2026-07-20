// Video explanation: https://www.youtube.com/watch?v=vE_Vz7vuaOU

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	grid := make([][]uint8, dy)
	for y := range grid {
		row := make([]uint8, dx)
		for x := range row {
			row[x] = uint8(x * y)
		}
		grid[y] = row
	}
	return grid
}

func main() {
	pic.Show(Pic)
}
