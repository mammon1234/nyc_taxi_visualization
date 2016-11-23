package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"database/sql"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"github.com/yy9415/nyctaxi/util"
	"time"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

const (
	url = "http://data.fcc.gov/api/block/find?showall=false&format=json"
)

func main() {
	//error := log.New(os.Stderr,
	//	"ERROR: ",
	//	log.Ldate|log.Ltime|log.Lshortfile)
	//error.Println("error")

	db, err := sql.Open("mysql", "martin:martin@/taxi?charset=utf8")
	util.CheckErr(err)
	defer db.Close()

	// dat, err := os.Open("/Volumes/YuDriver/data/april/1-3/xaa")
	dat, err := os.Open("/Volumes/YuDriver/data/temp/data.sorted.csv")
	util.CheckErr(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)

	var count int = 0
	var record_number int = 0;

	// months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	months := []string{"01"}
	stmt := make(map[string]*sql.Stmt)

	for _, v := range months {
		st, err := db.Prepare("INSERT trip SET pickup_time=?, dropoff_time=?, pickup_longitude=?, " +
		"pickup_latitude=?, dropoff_longitude=?, dropoff_latitude=?, route=?, distance=?, passenger=?, " +
		"fare=?, pickup_geohash=?, dropoff_geohash=?, pickup_census=?, dropoff_census=?")
		util.CheckErr(err)
		stmt[v] = st
	}

	t := time.Now()
	fmt.Print("Start: ")
    fmt.Println(t.Format(time.RFC3339))

	for scanner.Scan() {
		record_number++

		if record_number == 1 {
			continue
		}

		if record_number % 100000 == 0 {
			fmt.Printf("Scaned: %d\n", record_number)
		}

		str := scanner.Text()
		arr := strings.Split(str, ",")
		if len(arr) > 13 {
			pickup_time := arr[1]
			dropoff_time := arr[2]
			passenger := arr[3]
			distance := arr[4]
			if arr[5] == "" || arr[6] == "" || arr[9] == "" || arr[10] == "" {
				continue
			}
			if pickup_time >= "2014-02-01 00:00:00" {
				break;
			}

			pickup_longitude, err := strconv.ParseFloat(arr[5], 64)
			if err != nil {
				continue
			}
			if pickup_longitude == 0 {
				continue;
			}
			pickup_latitude, err := strconv.ParseFloat(arr[6], 64)
			if err != nil {
				continue
			}
			dropoff_longitude, err := strconv.ParseFloat(arr[9], 64)
			if err != nil {
				continue
			}
			dropoff_latitude, err := strconv.ParseFloat(arr[10], 64)
			if err != nil {
				continue
			}
			fare_amount := arr[12]
			//route, err := routeQuery(pickup_latitude, pickup_longitude, dropoff_latitude, dropoff_longitude)
			pickup_geohash := util.Encode(pickup_latitude, pickup_longitude)
			dropoff_geohash := util.Encode(dropoff_latitude, dropoff_longitude)
			pickup_census := ""
			dropoff_census := ""
			
			// m := strings.Split(strings.Split(pickup_time, " ")[0], "-")[1]
			st := stmt["01"]

			_, err = st.Exec(pickup_time, dropoff_time, pickup_longitude, pickup_latitude, dropoff_longitude,
				dropoff_latitude, "", distance, passenger, fare_amount, pickup_geohash, dropoff_geohash, pickup_census, dropoff_census)
			if err != nil {
				continue
			}
			count++
		}

		if count % 100000 == 0 {
			fmt.Printf("Saved: %d\n", count)
		} 
	}
	fmt.Printf("Finished: %d\n", count)
	t = time.Now()
    fmt.Println(t.Format(time.RFC3339))
	//https://maps.googleapis.com/maps/api/directions/json?origin=40.727271999999999,-74.000607000000002&destination=40.725496999999997,-73.993049999999997
	//routeQuery("40.727271999999999", "-74.000607000000002", "40.725496999999997", "-73.993049999999997")
}

func queryCensusCode(latitude string, longitude string) string {
	tmpURL := url + "&latitude=" + latitude + "&longitude=" + longitude
	// fmt.Println(tmpURL)
	resp, err := http.Get(tmpURL)
	defer resp.Body.Close()
	if err != nil {
		return ""
	}
	var census Census
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(jsonDataFromHttp), &census)
	if census.Status == "OK" {
		return census.County.FIPS
	} else {
		return ""
	}
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

type Census struct {
	Block Block
	County County 
	State State
	Status string
	ExecutionTime string
}

type Block struct {
	FIPS string
}

type County struct {
	FIPS string
	Name string
}

type State struct {
	FIPS string
	Code string
	Name string
}
