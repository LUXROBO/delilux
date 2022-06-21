package client

import (
	"encoding/json"
	"net/http"
)

// Client is a wrapper class for sweet-tracker client
type IParceluxClient interface {
	trackParcel(trackCode string, trackInvoice string) interface{}
}

// ParceluxClient is a wrapper class for sweet-tracker client
type ParceluxClient struct {
	apiURL string
	apiKey string
	header struct {
		ContentType string
	}
}

// NewParceluxClient returns a new instance of ParceluxClient
func NewParceluxClient(apiKey string) *ParceluxClient {
	parceluxClient := &ParceluxClient{
		apiURL: "http://info.sweettracker.co.kr",
		apiKey: apiKey,
		header: Header{
			ContentType: "application/json",
		},
	}
	return parceluxClient
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
func (pc ParceluxClient) requestWithParams(
	trackParams TrackParams,
	response *TrackResp,
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

	q := req.URL.Query()
	q.Add("t_key", pc.apiKey)
	q.Add("t_code", trackParams.TCode)
	q.Add("t_invoice", trackParams.TInvoice)
	req.URL.RawQuery = q.Encode()

	httpHeader := httpInfo.Header
	req.Header.Add("Content-Type", httpHeader.ContentType)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(response)
	return response
}

// TrackParcel returns a response of tracking parcel
func (pc ParceluxClient) TrackParcel(
	trackCode string, trackInvoice string,
) interface{} {
	trackParams := TrackParams{
		TCode:    trackCode,
		TInvoice: trackInvoice,
	}

	var trackResp TrackResp

	httpInfo := HTTPInfo{
		Method: "GET",
		URL:    pc.apiURL + "/api/v1/trackingInfo",
		Header: pc.header,
	}

	pc.requestWithParams(
		trackParams,
		&trackResp,
		httpInfo,
	)

	return trackResp
}
