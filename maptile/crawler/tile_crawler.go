package maptile_crawler

import (
	"fmt"
	"strconv"
	"os"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"github.com/yy9415/nyctaxi/util"
	"github.com/oleiade/lane"
)

const (
	url = "https://api.mapbox.com/v4/mapbox.light/"
	token = "pk.eyJ1IjoieWFuZ3l1OTQxNSIsImEiOiJjaWp1c2pyeWQwZ2ZvdHdrczk3ZGl3b2JvIn0.Khjkqpomtr2iC7-XxSHbDw"
	dir = "/Users/yangyu/Developer/go/lighttiles/"
	//http://a.tiles.mapbox.com/v4/yangyu9415.1f39fa18/{z}/{x}/{y}.png?access_token=pk.eyJ1IjoieWFuZ3l1OTQxNSIsImEiOiJjaWp1c2pyeWQwZ2ZvdHdrczk3ZGl3b2JvIn0.Khjkqpomtr2iC7-XxSHbDw
)

type Tile struct {
	Z int
	X int
	Y int
}

/**
 * z zoom level
 * x x level
 * y y level
 */
func FetchTile(z int, x int, y int) string {

	tmpURL := url + strconv.Itoa(z) + "/" + strconv.Itoa(x) + "/" + strconv.Itoa(y) + ".png?access_token=" + token

	resp, err := http.Get(tmpURL)
	util.CheckErr(err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	util.CheckErr(err)
	dir := dir + strconv.Itoa(z) + "/" + strconv.Itoa(x) + "/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// path does not exist
		os.MkdirAll(dir, 0777)
	}
	file := dir + strconv.Itoa(y) + ".png"
	ioutil.WriteFile(file, data, 0777)

	abs, _ := filepath.Abs(file)
	return abs
}

func BatchFetchTile(z_1 int, z_2 int, x_1 int, x_2 int, y_1 int, y_2 int) {
	queue := lane.NewQueue()
	for x := x_1; x <= x_2; x++ {
		for y := y_1; y <= y_2; y++ {
			tile := Tile{z_1, x, y}
			queue.Enqueue(tile)
		}
	}
	for !queue.Empty() {
		tile := queue.Dequeue().(Tile)
		if tile.Z >	 z_2 {
			break
		}
		if !isExist(tile) {
			fmt.Println(tile)
			FetchTile(tile.Z, tile.X, tile.Y)
		}
		tiles := getChildTiles(tile.Z, tile.X, tile.Y)
		for _, v := range tiles {
			queue.Enqueue(v)
		}
	}
	fmt.Println(queue.Size())
}

func isExist(tile Tile) bool {
	file, _ := ioutil.ReadFile(dir + strconv.Itoa(tile.Z) + "/" + strconv.Itoa(tile.X) + "/" + strconv.Itoa(tile.Y) + ".png")
	return file != nil
}

func getChildTiles(z int, x int, y int) []Tile {
	tiles := []Tile {}
	x = x*2
	y = y*2
	tiles = append(tiles, Tile{z+1, x, y})
	tiles = append(tiles, Tile{z+1, x+1, y})
	tiles = append(tiles, Tile{z+1, x, y+1})
	tiles = append(tiles, Tile{z+1, x+1, y+1})
	return tiles
}