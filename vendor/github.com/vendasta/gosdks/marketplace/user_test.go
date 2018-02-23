package marketplace

import (
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/net/context"

	"testing"
)

func Test_GetUserAccountPermission(t *testing.T) {
	type testCase struct {
		name      string
		userID    string
		accountID string
		getError  error

		expectedPermission bool
		expectedError      error
	}

	testCases := []*testCase{
		{
			name:               "returns false with error when userID empty",
			userID:             "",
			accountID:          "AG-123",
			expectedPermission: false,
			expectedError:      util.Error(util.InvalidArgument, "User ID cannot be empty"),
		},
		{
			name:               "returns false with error when accountID empty",
			userID:             "UID-123",
			accountID:          "",
			expectedPermission: false,
			expectedError:      util.Error(util.InvalidArgument, "Account ID cannot be empty"),
		},
		{
			name:               "returns false with error when Get returns error",
			userID:             "UID-123",
			accountID:          "AG-123",
			getError:           util.Error(util.Unauthenticated, "unauthed"),
			expectedPermission: false,
			expectedError:      util.Error(util.Unauthenticated, "unauthed"),
		},
		{
			name:               "returns true with no error when Get returns no error",
			userID:             "UID-123",
			accountID:          "AG-123",
			getError:           nil,
			expectedPermission: true,
			expectedError:      nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			client := &userClient{
				&basesdk.BaseClientMock{
					Error: c.getError,
				},
			}

			permission, err := client.GetUserAccountPermission(ctx, c.userID, c.accountID)

			assert.Equal(t, c.expectedPermission, permission)
			assert.Equal(t, c.expectedError, err)
		})
	}
}
