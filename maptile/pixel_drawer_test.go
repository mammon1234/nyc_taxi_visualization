package maptile_draw

import (
	// "fmt"
	// "time"
	// "io/ioutil"
	"testing"
)

func TestDraw(t *testing.T) {
	draw("/Users/yangyu/Desktop/test1.png", 60, 216, []uint8{0, 255, 204, 50})
}

func TestTileDraw(t *testing.T) {
	/*
	pixels := [][]int{
		[]int{40, 216},
		[]int{30, 216},
		[]int{20, 216},
		[]int{10, 216},
	}
	TileDraw("/Users/yangyu/Desktop/test2.png", pixels, []uint8{0, 255, 204, 150})
	*/
}