package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"database/sql"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yy9415/nyctaxi/util"
	"github.com/yy9415/nyctaxi/model"
)

func LoadPickUpData(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	start := params.Get("start")
	end := params.Get("end")
	// fmt.Println("S: " + start)
	// fmt.Println("E: " + end)

	db, err := sql.Open("mysql", "martin:martin@/taxi?charset=utf8")
	defer db.Close()
	util.CheckErr(err)

	points := model.Geojson {
		Type:"FeatureCollection",
		Features: []model.Feature{
		},
	}

	// query data
	rows, err := db.Query("SELECT * FROM trip where pickup_time > '" + start +
	"' and pickup_time < '" + end + "'")
	util.CheckErr(err)
	defer rows.Close()

	var row_count = 0
	var passenger_count = 0
	var distance_count = 0.0

	for rows.Next() {
		var id int32
		var pickup_time string
		var dropoff_time string
		var passenger int
		var distance float64
		var pickup_longitude string
		var pickup_latitude string
		var dropoff_longitude string
		var dropoff_latitude string
		var fare float64
		var route string
		var pickup_geohash string
		var dropoff_geohash string
		var dropoff_census string

		err = rows.Scan(&id, &pickup_time, &dropoff_time, &pickup_longitude, &pickup_latitude,
			&dropoff_longitude, &dropoff_latitude, &distance, &passenger, &fare,
			&route, &pickup_geohash, &dropoff_geohash, &dropoff_census)
		util.CheckErr(err)
		p_longitude, _ := strconv.ParseFloat(pickup_longitude, 64)
		p_latitude, _ := strconv.ParseFloat(pickup_latitude, 64)
		d_longitude, _ := strconv.ParseFloat(dropoff_longitude, 64)
		d_latitude, _ := strconv.ParseFloat(dropoff_latitude, 64)


		row_count++
		passenger_count += passenger
		distance_count += distance

		p_feature := model.Feature {
			Type: "point",
			Id: id,
			Property: model.Property {
				Type: 0,
			},
			Geometry: model.Geometry {
				Type: "point",
				Coordinates: []float64 {
					p_longitude, p_latitude,
				},
			},
		}

		d_feature := model.Feature {
			Type: "point",
			Id: id,
			Property: model.Property {
				Type: 1,
			},
			Geometry: model.Geometry {
				Type: "point",
				Coordinates: []float64 {
					d_longitude, d_latitude,
				},
			},
		}
		points.Features = append(points.Features, p_feature)
		points.Features = append(points.Features, d_feature)
	}

	trips := model.Trips {
		Result: "success",
		HasNext: true,
		Count: row_count,
		Passenger: passenger_count,
		Distance: distance_count,
		Geojson: points,
	}

	if row_count == 0 {
		trips.HasNext = false
	}

	js, err := json.Marshal(trips)
	util.CheckErr(err)
	fmt.Println(row_count)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}