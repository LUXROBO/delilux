package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client is a wrapper class for sweet-tracker client
type Client interface {
	trackParcel(trackParcelPayload interface{}) interface{}
}

// ParceluxClient ...
type ParceluxClient struct {
	apiURL string
	header struct {
		Authorization string
		ContentType   string
	}
}

// NewParceluxClient returns a new instance of ParceluxClient
func NewParceluxClient(apiKey string) *ParceluxClient {
	plClient := &ParceluxClient{
		apiURL: "http://info.sweettracker.co.kr/tracking/5",
		header: Header{
			Authorization: "Basic " + apiKey,
			ContentType:   "application/json",
		},
	}
	return plClient
}

// Header includes header information of request instance
type Header struct {
	Authorization string
	ContentType   string
}

// HTTPInfo includes http information of request instance
type HTTPInfo struct {
	Method string
	URL    string
	Header Header
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

// TrackParcel ...
func (c ParceluxClient) TrackParcel(
	trackParcelPayload interface{},
) interface{} {
	var trackParcelResp trackParcelResp
	payload := trackParcelPayload.(TrackParcelPayload)
	httpInfo := HTTPInfo{
		Method: "GET",
		URL:    c.apiURL + "/api/v1/trackingInfo" + payload.CustomerUID,
		Header: c.header,
	}
	client.RequestWithPayload(
		Payload,
		&trackParcelResp,
		httpInfo,
	)
	return trackParcelResp
}
