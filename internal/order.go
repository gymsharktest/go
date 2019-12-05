package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

// OrderRequest is a request for products from the user
type OrderRequest struct {
	Count int `json:"count"`
}

// Order is a map [pack size] => count
type Order struct {
	Lines map[int]int `json:"order"`
}

// Add some packs to the order
func (o *Order) addPacks(pack Pack, qty int) {
	if o.Lines == nil {
		o.Lines = make(map[int]int)
	}
	o.Lines[pack.Count] = qty
}

// Handle a request to make an order
func orderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling order request")

	var orderReq OrderRequest
	var err = json.NewDecoder(r.Body).Decode(&orderReq)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	if orderReq.Count < 1 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	count := orderReq.Count

	// sort the array of packs in a descending order
	sort.Slice(packs, func(i, j int) bool {
		return packs[i].Count > packs[j].Count
	})

	order := makeOrder(count)

	data, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Create an order for the requested amout of products
func makeOrder(count int) Order {
	var order Order
	for _, pack := range packs {
		if pack.Count <= count {
			order.addPacks(pack, count/pack.Count)
			count = count % pack.Count
		}
	}
	// edge case, send minimum
	if len(order.Lines) == 0 || count > 0 {
		order.addPacks(packs[len(packs)-1], 1)
	}
	return order
}
