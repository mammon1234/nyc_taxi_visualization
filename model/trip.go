package model

type Trip struct {
	Id int32 `json:"id"`
	Pickup_time string `json:pickup_time`
	Dropoff_time string `json:dropoff_time`
	Passenger int `json:passenger`
	Distance float64 `json:distance`
	Pickup_longitude string `json:pickup_longitude`
	Pickup_latitude string `json:pickup_latitude`
	Dropoff_longitude string `json:dropoff_longitude`
	Dropoff_latitude string `json:dropoff_latitude`
	Fare float64 `json:fare`
	Route string `json:route`
	Pickup_geohash string `json:pickup_geohash`
	Dropoff_geohash string `json:dropoff_geohash`
}

// type Trips []Trip