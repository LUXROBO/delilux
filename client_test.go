package parcelux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpParceluxClient() *ParceluxClient {
	return NewClient("wS6QdxNImv6DaPt0or1X4g")
}

func TestTrackParcel(t *testing.T) {
	parceluxClient := setUpParceluxClient()
	trackResult := parceluxClient.TrackParcel("04", "648428990916")
	assert.Equal(t, true, trackResult.(*TrackResp).Complete)
}
