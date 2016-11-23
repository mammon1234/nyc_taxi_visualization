package model

import (
	"testing"
)

func TestPush(t *testing.T) {
	var q Queue
	trip := Trip{1, "2014-01-01 10:10:10", "2014-01-01 10:10:12", 1, 1.2, 
		"123.234459", "45.124455", "123.234459", "45.124455", 10.3, "1", "abc", "bca"}
	q.Push(&trip)
	if q.size != 1 || len(q.trips) != 1 {
		t.Errorf("Push failed")
	}
}

func TestPop(t *testing.T) {
	var q Queue
	trip := Trip{1, "2014-01-01 10:10:10", "2014-01-01 10:10:12", 1, 1.2, 
		"123.234459", "45.124455", "123.234459", "45.124455", 10.3, "1", "abc", "bca"}
	q.Push(&trip)
	item := q.Pop()
	if item == nil {
		t.Errorf("Nil pop out")
	}
	if q.size != 0 || len(q.trips) != 1 {
		t.Errorf("Item remained: size %d, len %d", q.size, len(q.trips))
	}
}