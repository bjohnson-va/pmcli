package cssdk

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

var geoResponse = `{
  "version": "1.0",
  "data": {
    "city": "SASKATOON",
    "country": "CA",
    "longitude": -106.647656,
    "state": null,
    "address": null,
    "latitude": 52.134369900000003
  },
  "requestId": "594ae62c00ff028e617c958dd10001737e726570636f72652d70726f6400016170693a3630372d656e2d6765742d676d622d6c6f632d61732d6469637400010102",
  "responseTime": 53,
  "statusCode": 200
}`

func Test_InferGeolocationReturnGeoPointOn200(t *testing.T) {
	client := BuildGeoClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: geoResponse}

	result, err := client.InferGeolocation(context.Background(), "CA", WithCity("saskatoon"))

	assert.Nil(t, err)
	expected := &GeoPoint{
		Latitude:  52.1343699,
		Longitude: -106.647656,
	}
	assert.Equal(t, expected, result)
}

func Test_InferGeolocationUsesProvidedOptions(t *testing.T) {
	client := BuildGeoClient("user", "key", config.Local)
	m := &basesdk.BaseClientMock{JSONBody: geoResponse}
	client.SDKClient = m

	_, _ = client.InferGeolocation(context.Background(), "CA", WithCity("saskatoon"), WithState("SK"), WithAddress("123 cowboy lane"))

	expected := map[string]interface{}{
		"country": "CA",
		"state":   "SK",
		"city":    "saskatoon",
		"address": "123 cowboy lane",
	}
	assert.Equal(t, expected, m.ParamsSent)
}

func Test_InferGeolocationReturnsErrorWhenCoreReturnsError(t *testing.T) {
	client := BuildGeoClient("user", "key", config.Local)
	expectedError := errors.New("New error")
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: geoResponse, Error: expectedError}

	_, err := client.InferGeolocation(context.Background(), "CA", WithCity("saskatoon"))

	assert.Equal(t, expectedError, err)
}

func Test_InferGeolocationReturnsErrorWhenInflatingResponseHasError(t *testing.T) {
	client := BuildGeoClient("user", "key", config.Local)
	client.SDKClient = &basesdk.BaseClientMock{JSONBody: `{"data":"garbage"}`}

	_, err := client.InferGeolocation(context.Background(), "CA", WithCity("saskatoon"))

	assert.Error(t, err)
}
