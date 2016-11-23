package maptile_crawler

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFetchTile(t *testing.T) {
	f := FetchTile(11, 1204, 1537)
	file, err := ioutil.ReadFile(f)
	if  file == nil || err != nil{
		t.Errorf("Fetch failed")
	}
}

func TestBatchFetchTile(t *testing.T) {
	// BatchFetchTile(12, 17, 1204, 1208, 1537, 1541)
	BatchFetchTile(12, 17, 1203, 1209, 1536, 1542)
}

func TestGetChildTile(t *testing.T) {
	tiles := getChildTiles(12, 100, 1000)

	for _, v := range tiles {
		fmt.Println(v)
	}
	if len(tiles) != 4 {
		t.Errorf("Wrong child tiles")
	}
}

func TestIsExist(t *testing.T) {
	tile := Tile{11, 1204, 1537}
	if !isExist(tile) {
		t.Errorf("TestIsExist failed")
	}
}

