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

//StoreHistory method stores history item to memory
func (history *History) StoreHistory(orderId string) {
	ttl, err := strconv.Atoi(config.Config("LOCATION_HISTORY_TTL_SECONDS"))
	if err != nil {
		log.Fatal("Invalid LOCATION_HISTORY_TTL_SECONDS",err)
	}

	heap.Set(orderId, history, int64(ttl))

	heap.Support(History{})

	heap.Save()
}

//GetHistory func retrieves a particular history item based on orderId
func GetHistory(orderId string) *History {

	val, ok := heap.Get(orderId)

	if ok {
		v, k := val.(History)

		if k {
			return &v
		}

	}

	return nil
}
