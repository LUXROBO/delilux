package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TrackParcelPayload struct {
	TKey     string `json:"t_key"`
	TCode    string `json:"t_code"`
	TInvoice string `json:"t_invoice"`
}

func main() {
	payload := TrackParcelPayload{
		TKey:     "wS6QdxNImv6DaPt0or1X4g",
		TCode:    "04",
		TInvoice: "648428990916",
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"GET",
		"http://info.sweettracker.co.kr"+"/api/v1/trackingInfo",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	fmt.Println(data)
}
