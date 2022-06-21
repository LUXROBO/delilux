package main

import (
	"fmt"

	"github.com/luxrobo/parcelux/client"
)

func main() {
	// req, err := http.NewRequest(
	// 	"GET",
	// 	"http://info.sweettracker.co.kr"+"/api/v1/trackingInfo",
	// 	nil,
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// req.Header.Add("Content-Type", "application/json")

	// q := req.URL.Query()
	// q.Add("t_key", "wS6QdxNImv6DaPt0or1X4g")
	// q.Add("t_code", "04")
	// q.Add("t_invoice", "648428990916")
	// req.URL.RawQuery = q.Encode()

	// httpClient := &http.Client{}
	// resp, err := httpClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// var data interface{}
	// json.NewDecoder(resp.Body).Decode(&data)
	// fmt.Println(data)

	plClient := client.NewParceluxClient("wS6QdxNImv6DaPt0or1X4g")
	plResult := plClient.TrackParcel("04", "648428990916")
	fmt.Println(plResult)
}
