package main

import (
	"os"
	"database/sql"
	"bufio"
	"fmt"
	"strings"
)

const (
	MongoDBHosts = "mongodb://127.0.0.1:27017"
	AuthDatabase = "nyctaxi"
	AuthUserName = "martin"
	AuthPassword = "123456"
	//TestDatabase = "goinggo"
)

func main() {

	db, err := sql.Open("mysql", "martin:martin@/nyctaxi?charset=utf8")
	checkErr(err)

	dat, err := os.Open("/Volumes/YuDriver/data/april/1-3/xaa")
	checkErr(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)

	var count int = 0
	var record_number int = 0;

	stmt, err := db.Prepare("INSERT trip_april SET pickup_time=?, dropoff_time=?, pickup_longitude=?, " +
		"pickup_latitude=?, dropoff_longitude=?, dropoff_latitude=?, route=?, distance=?, passenger=?, fare=?")
	checkErr(err)

	for scanner.Scan() {
		record_number++
		if record_number == 1 {
			continue
		}
		if record_number % 60 != 0 {
			//fmt.Println(count)
			continue
		}
		if record_number % 10000 == 0 {
			fmt.Println("Scan: ", record_number)
		}

		if count % 1000 == 0 {
			fmt.Println("Save: ", count)
		}
		str := scanner.Text()
		arr := strings.Split(str, ",")
		if len(arr) > 13 {
			pickup_time := arr[1]
			dropoff_time := arr[2]
			passenger := arr[3]
			distance := arr[4]
			pickup_longitude := arr[5]
			pickup_latitude := arr[6]
			dropoff_longitude := arr[9]
			dropoff_latitude := arr[10]
			fare_amount := arr[12]
			route, err := routeQuery(pickup_latitude, pickup_longitude, dropoff_latitude, dropoff_longitude)
			checkErr(err)
			if route != "" {
				_, err = stmt.Exec(pickup_time, dropoff_time, pickup_longitude, pickup_latitude, dropoff_longitude,
					dropoff_latitude, route, distance, passenger, fare_amount)
				checkErr(err)
				count++
			}
		}
	}
	db.Close()
	fmt.Println(count)
	//https://maps.googleapis.com/maps/api/directions/json?origin=40.727271999999999,-74.000607000000002&destination=40.725496999999997,-73.993049999999997
	//routeQuery("40.727271999999999", "-74.000607000000002", "40.725496999999997", "-73.993049999999997")
}

type Direction struct {
	Geocoded_waypoints []interface{}
	Routes []Routes
	Status string
}

type Routes struct {
	Bounds interface{}
	Copyrights string
	Legs []interface{}
	Overview_polyline Overview_polyline
	Summary string
	Warnings []interface{}
	waypoint_order []interface{}
}

type Overview_polyline struct {
	Points string
}