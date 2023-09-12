package main

import (
	"cashier/cashier"
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
	cashier.NewGroup("scores", 2<<10, cashier.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:9999"
	peers := cashier.NewHTTPPool(addr)
	log.Println("cashier is running at", peers)
	log.Fatal(http.ListenAndServe(addr, peers))
}
