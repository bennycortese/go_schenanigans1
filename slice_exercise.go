package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	zoom := [][] uint8{}
	for i := 0; i < dy; i++ {
		one_chunk := [] uint8{}
		for j := 0; j < dx; j++ {
			if j % 2 == 0 {
				one_chunk = append(one_chunk, 255)
			} else {
				one_chunk = append(one_chunk, 0)
			}
		}
		zoom = append(zoom, one_chunk)
	}
	return zoom
}

func main() {
	pic.Show(Pic)
}
