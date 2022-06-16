package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ParceluxClient is a wrapper class for sweet-tracker client
type ParceluxClient interface {
	trackParcel(trackParcelPayload interface{}) interface{}
}

// HTTPInfo includes http information of request instance
type HTTPInfo struct {
	Method string
	URL    string
	Header struct {
		Authorization string
		ContentType   string
	}
}

// RequestWithPayload makes request with a given payload
func RequestWithPayload(
	payload interface{},
	response interface{},
	httpInfo HTTPInfo,
) interface{} {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		httpInfo.Method,
		httpInfo.URL,
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	httpHeader := httpInfo.Header
	req.Header.Add("Authorization", httpHeader.Authorization)
	req.Header.Add("Content-Type", httpHeader.ContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(response)
	return response
}
