package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var output [][]uint8

	for i := 0; i <= dy; i++ {
		var current []uint8
		for j := 0; j <= dx; j++ {
			current = append(current, uint8(j*j+i*i))
		}
		output = append(output, current)
	}
	return output
}

func main() {
	pic.Show(Pic)
}
