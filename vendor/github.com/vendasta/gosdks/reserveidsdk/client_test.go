package reserveidsdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"testing"
)

func Test_ReserveID_StatusCodeIsString(t *testing.T) {
	responseData := `{"data":{"status":"200","internal_id":"AC-123","reason":"test reason"}}`

	baseClient := &basesdk.BaseClientMock{}
	baseClient.JSONBody = responseData

	client := ReserveIDClient{baseClient}

	_, err := client.ReserveID(context.Background(), "VUNI", "AG-1234", "", "MP-987")

	assert.EqualError(t, err, "rpc error: code = Internal desc = Failed to convert response to ReserveIDResult: json: cannot unmarshal string into Go struct field ReserveIDResponse.status of type int,")
}

func Test_ReserveID_StatusCodeIsInt(t *testing.T) {
	responseData := `{"data":{"status":200,"internal_id":"AC-123","reason":"test reason"}}`

	baseClient := &basesdk.BaseClientMock{}
	baseClient.JSONBody = responseData

	client := ReserveIDClient{baseClient}

	reserveIDResponse, err := client.ReserveID(context.Background(), "VUNI", "AG-1234", "", "MP-987")

	assert.NoError(t, err)
	assert.Equal(t, 200, reserveIDResponse.Status)
	assert.Equal(t, "AC-123", reserveIDResponse.VendorInternalID)
	assert.Equal(t, "test reason", reserveIDResponse.Reason)
}
