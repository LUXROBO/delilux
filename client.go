package parcelux

import (
	"encoding/json"
	"net/http"
)

// IParceluxClient is an interface for sweet-tracker api
type IParceluxClient interface {
	TrackParcel(trackCode string, trackInvoice string) TrackResp
}

// ParceluxClient is a wrapper class for sweet-tracker api
type ParceluxClient struct {
	apiURL string
	apiKey string
	header struct {
		ContentType string
	}
}

// NewClient returns a new instance of ParceluxClient
func NewClient(apiKey string) *ParceluxClient {
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

// requestToTrack makes a request for tracking parcel with a given params
func (pc ParceluxClient) requestToTrack(
	trackParams TrackParams,
	response *TrackResp,
	httpInfo HTTPInfo,
) *TrackResp {
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
) TrackResp {
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

	pc.requestToTrack(
		trackParams,
		&trackResp,
		httpInfo,
	)

	return trackResp
}
