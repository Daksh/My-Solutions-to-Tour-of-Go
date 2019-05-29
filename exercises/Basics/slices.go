package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for y:=0; y < dy; y++{
		ans[y] = make([]uint8, dx)
		for x:=0; x < dx; x++ {
			ans[y][x] = (uint8(x)+uint8(y))/uint8(2)
			ans[y][x] = uint8(x)*uint8(y)
			ans[y][x] = (uint8(x) ^ uint8(y)) * (uint8(x) ^ uint8(y))
		}
	}
	return ans
}

func main() {
	pic.Show(Pic)
}
