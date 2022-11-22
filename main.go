package main

import (
	"encoding/json"
	"fmt"

	// "lcorequest/go/pkg/mod/github.com/gorilla/mux@v1.8.0"
	"log"
	"net/http"

	// "lcorequest/go/pkg/mod/github.com/gorilla/mux@v1.8.0"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json :"UID"`
	Name  string  `json :"Name"`
	Desc  string  `json : "Desc"`
	Price float64 `json : "Price"`
}

var Inventory []Item

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoints called: homepage()")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(Inventory)
}

func _deleteItemUid(uid string) {
	for index, Item := range Inventory {
		if Item.UID == uid {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
			break

		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item

	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	deleteItemAtUid(params["uid"])
}

func deleteItemAtUid(s string) {
	panic("unimplemented")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function called: getInventory()")

	json.NewEncoder(w).Encode(Inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	Inventory = append(Inventory, item)

	json.NewEncoder(w).Encode(item)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Inventory = append(Inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "A fine block of cheese",
		Price: 4.99,
	})

	Inventory = append(Inventory, Item{
		UID:   "1",
		Name:  "Milk",
		Desc:  "A jug of milk",
		Price: 3.25,
	})

	handleRequests()
}
