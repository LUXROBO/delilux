package client

import (
	"encoding/json"
	"net/http"
)

// Client is a wrapper class for sweet-tracker client
type Client interface {
	trackParcel(trackParcelParams interface{}) interface{}
}

// ParceluxClient ...
type ParceluxClient struct {
	apiURL string
	header struct {
		ContentType string
	}
}

// NewParceluxClient returns a new instance of ParceluxClient
func NewParceluxClient(apiKey string) *ParceluxClient {
	plClient := &ParceluxClient{
		apiURL: "http://info.sweettracker.co.kr",
		header: Header{
			ContentType: "application/json",
		},
	}
	return plClient
}

// Header includes header information of request instance
type Header struct {
	ContentType string
}

// HTTPInfo includes http information of request instance
type HTTPInfo struct {
	Method string
	URL    string
	Header Header
}

// RequestWithPayload makes request with a given payload
func RequestWithQueryParams(
	params interface{},
	response interface{},
	httpInfo HTTPInfo,
) interface{} {
	req, err := http.NewRequest(
		httpInfo.Method,
		httpInfo.URL,
		nil,
	)
	if err != nil {
		panic(err)
	}

	trackParcelParams := params.(TrackParcelParams)

	q := req.URL.Query()
	q.Add("t_key", trackParcelParams.TKey)
	q.Add("t_code", trackParcelParams.TCode)
	q.Add("t_invoice", trackParcelParams.TInvoice)
	req.URL.RawQuery = q.Encode()

	httpHeader := httpInfo.Header
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
	trackParcelParams interface{},
) interface{} {
	var trackParcelResp TrackParcelResp
	queryParams := trackParcelParams.(TrackParcelParams)
	httpInfo := HTTPInfo{
		Method: "GET",
		URL:    c.apiURL + "/api/v1/trackingInfo",
		Header: c.header,
	}
	RequestWithQueryParams(
		queryParams,
		&trackParcelResp,
		httpInfo,
	)
	return trackParcelResp
}
