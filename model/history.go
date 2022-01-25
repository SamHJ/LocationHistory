package model

import (
	"locationhistory/config"
	"log"
	"strconv"

	ttl_map "github.com/leprosus/golang-ttl-map"
)

type History struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

var heap *ttl_map.Heap

//initialize the in-memory storage module
func init() {
	heap = ttl_map.New()
	heap.Path("./history/history.tsv")
}

//StoreHistory method stores history item to memory with time to live (ttl)
func (history History) StoreHistory(orderId string) {
	ttl, err := strconv.Atoi(config.Config("LOCATION_HISTORY_TTL_SECONDS"))
	if err != nil {
		log.Fatal("Invalid LOCATION_HISTORY_TTL_SECONDS", err)
	}

	var histories []History

	oldHistory := GetHistory(orderId)

	if oldHistory == nil {
		histories = append(histories, history)

	} else {
		histories = append(histories, (oldHistory.([]History))...)
	}

	heap.Set(orderId, histories, int64(ttl))

	heap.Support([]History{})

	heap.Save()
}

//GetHistory func retrieves a particular history item based on orderId
func GetHistory(orderId string) interface{} {
	val, ok := heap.Get(orderId)
	if ok{
		return val
	}
	return nil
}

//DeleteHistory func deletes a history item
func DeleteHistory(orderId string) {
	heap.Del(orderId)
}
