package parcelux

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpApiKey(t *testing.T) {
	t.Setenv("API_KEY", "wS6QdxNImv6DaPt0or1X4g")
}

func setUpParceluxClient(apiKey string) *ParceluxClient {
	return NewClient(apiKey)
}

func TestTrackParcel(t *testing.T) {
	setUpApiKey(t)
	API_KEY := os.Getenv("API_KEY")
	parceluxClient := setUpParceluxClient(API_KEY)

	trackCode, trackInvoice := "04", "648428990916"
	trackResult := parceluxClient.TrackParcel(trackCode, trackInvoice)
	fmt.Println("trackResult:", trackResult)
	assert.Equal(t, true, trackResult.Complete)
}
