package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// packSizes stores available pack sizes loaded from config.json.
var packSizes []int

func loadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Warning: config.json not found, using default pack sizes")
		packSizes = []int{250, 500, 1000, 2000, 5000}
		return
	}

	var config struct {
		PackSizes []int `json:"pack_sizes"`
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error parsing config.json, using default pack sizes")
		packSizes = []int{250, 500, 1000, 2000, 5000}
		return
	}

	packSizes = config.PackSizes
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	fmt.Println("Loaded pack sizes:", packSizes)
}
