package maptile_draw

import (
	"fmt"
	"math"
	"os"
	"image"  
    "image/color"  
    "image/png"  
	"github.com/yy9415/nyctaxi/util"
)

const (
	TILE_SIZE = 256
)

func LayerDraw() {

}

func TileDraw(file string, pixels [][]int, c []uint8) {
	imgfile, _ := os.Create(file)
  	defer imgfile.Close()
	img := image.NewNRGBA(image.Rect(0, 0, TILE_SIZE, TILE_SIZE))
	color := color.RGBA{c[0], c[1], c[2], c[3]}

	for i := 0; i < len(pixels); i++ {
		pixel_x := pixels[i][0]
		pixel_y := pixels[i][1]

		img.Set(pixel_x, pixel_y, color)
		img.Set(pixel_x, pixel_y+1, color)
		img.Set(pixel_x, pixel_y+2, color)
		img.Set(pixel_x+1, pixel_y, color)
		img.Set(pixel_x+1, pixel_y+1, color)
		img.Set(pixel_x+1, pixel_y+2, color)
		img.Set(pixel_x+2, pixel_y, color)
		img.Set(pixel_x+2, pixel_y+1, color)
		img.Set(pixel_x+2, pixel_y+2, color)

		img.Set(pixel_x+1, pixel_y-1, color)
		img.Set(pixel_x+3, pixel_y+1, color)
		img.Set(pixel_x+1, pixel_y+3, color)
		img.Set(pixel_x-1, pixel_y+1, color)
	}

	err := png.Encode(imgfile, img)
	util.CheckErr(err)
}

func draw(file string, pixel_x int, pixel_y int, c []uint8) {

	oldImgFile, err := os.Open(file)
  	defer oldImgFile.Close()
  	util.CheckErr(err)
  	oldImg, err := png.Decode(oldImgFile)
  	util.CheckErr(err)

  	newImageFile, _ := os.Create(file)
  	defer newImageFile.Close()
	newImg := image.NewNRGBA(image.Rect(0, 0, TILE_SIZE, TILE_SIZE))
	color := color.RGBA{c[0], c[1], c[2], c[3]}

	bound := oldImg.Bounds()
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			col := oldImg.At(x, y)
			newImg.Set(x, y, col)
		}
	}
	
	newImg.Set(pixel_x, pixel_y, color)
	newImg.Set(pixel_x, pixel_y+1, color)
	newImg.Set(pixel_x, pixel_y+2, color)
	newImg.Set(pixel_x+1, pixel_y, color)
	newImg.Set(pixel_x+1, pixel_y+1, color)
	newImg.Set(pixel_x+1, pixel_y+2, color)
	newImg.Set(pixel_x+2, pixel_y, color)
	newImg.Set(pixel_x+2, pixel_y+1, color)
	newImg.Set(pixel_x+2, pixel_y+2, color)

	newImg.Set(pixel_x+1, pixel_y-1, color)
	newImg.Set(pixel_x+3, pixel_y+1, color)
	newImg.Set(pixel_x+1, pixel_y+3, color)
	newImg.Set(pixel_x-1, pixel_y+1, color)

	err = png.Encode(newImageFile, newImg)
	util.CheckErr(err)
}

func pointToPixel(lat float64, lng float64, zoom uint64) (uint64, uint64, uint64, uint64) {
	var factor uint64
	factor = 1 << zoom
	scale := float64(factor)

	worldCoordinateX, worldCoordinateY := project(lat, lng)
	pixel_x := uint64(math.Floor(worldCoordinateX * scale))
    pixel_y := uint64(math.Floor(worldCoordinateY * scale))
    fmt.Println(pixel_x, pixel_y)
	tileX := uint64(math.Floor(worldCoordinateX * scale / TILE_SIZE))
	tileY := uint64(math.Floor(worldCoordinateY * scale / TILE_SIZE))
	fmt.Println(tileX, tileY)
	return tileX, tileY, pixel_x - (TILE_SIZE*tileX), pixel_y - (TILE_SIZE*tileY)
}

func project(lat, lng float64) (float64, float64) {
    siny := math.Sin(lat * math.Pi / 180);
    siny = math.Min(math.Max(siny, -0.9999), 0.9999);
    x := TILE_SIZE * (0.5 + lng / 360)
    y := TILE_SIZE * (0.5 - math.Log((1 + siny) / (1 - siny)) / (4 * math.Pi))
    return x, y
}