package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// A Pack represents a individual SKU with a given count of products
type Pack struct {
	Count int `json:"count"`
}

// Packs available
var packs []Pack

func GetPacks() []Pack {
	return packs
}

func AddPack(pack Pack) {
	for _, p := range packs {
		if p.Count == pack.Count {
			return
		}
	}
	packs = append(packs, pack)
}

// Handler for route to add a pack
func addPackHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	var pack Pack
	err = json.Unmarshal(bodyBytes, &pack)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	if pack.Count < 1 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	AddPack(pack)

	data, _ := json.Marshal(pack)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Handler for route to show all packs
func getPacksHandler(w http.ResponseWriter, r *http.Request) {
	output, _ := json.Marshal(packs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
