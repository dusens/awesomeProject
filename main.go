package main

import (
	"code.dss.com/awesomeProject/geecahe"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	geecahe.NewGroup("scores", 2<<10, geecahe.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDb] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exits", key)
		}))
	addr := "localhost:9999"
	peers := geecahe.NewHTTPPool(addr)
	log.Println("geecache is ruuning at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
