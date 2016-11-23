package model

type Trips  struct {
	Result string `json:"result"`
	HasNext bool `json:"hasNext"`
	Count int `json:"count"`
	Passenger int `json:"passenger"`
	Distance float64 `json:"distance"` 
	Geojson Geojson `json:"geojson"`
}

type Geojson struct {
	Type string `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type string `json:"type"`
	Id int32 `json:"id"`
	Property Property `json:"properties"`
	Geometry Geometry `json:"geometry"`
}

type Property struct {
	Type int32 `json:"type"`
}

type Geometry struct {
	Type string `json:"type"`
	Coordinates []float64  `json:"coordinates"`
}