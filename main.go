package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AutoGenerated struct {
	Type  string `json:"type"`
	Valid bool   `json:"valid"`
	Value string `json:"value"`
}

func main() {
	for {
		time.Sleep(1 * time.Second)
		go req()
	}
}
func req() {
	req, _ := http.NewRequest("GET", "https://api.marleyspoon.com/promotions/46BDD2?brand=ms&country=us&product_type=web", nil)
	req.Header.Add("authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJtcyIsImNvdW50cnkiOiJ1cyIsImJyYW5kIjoibXMiLCJ0cyI6MTY0NzA4OTA0MywicmFuZG9tX2lkIjoiYWFlMGExIn0.rWekfowPYSuWufVjEVaL80EqfSCROmqiIS_VV0xsRy8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data AutoGenerated
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON error")
		panic(err)
	}
	fmt.Println(data)

	if data.Valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
