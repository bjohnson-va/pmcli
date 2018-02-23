// Code generated by mockery v1.0.0
package iam

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"
import iam_v1 "github.com/vendasta/gosdks/pb/iam/v1"
import mock "github.com/stretchr/testify/mock"
import resources "github.com/vendasta/gosdks/iam/resources"

// MockAuthService is an autogenerated mock type for the AuthService type
type MockAuthService struct {
	mock.Mock
}

// AllowPublicMethods provides a mock function with given fields: publicMethods
func (_m *MockAuthService) AllowPublicMethods(publicMethods ...string) {
	_va := make([]interface{}, len(publicMethods))
	for _i := range publicMethods {
		_va[_i] = publicMethods[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// DoesContextHaveValidServiceAccount provides a mock function with given fields: ctx, scopes
func (_m *MockAuthService) DoesContextHaveValidServiceAccount(ctx context.Context, scopes ...AccessScope) error {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...AccessScope) error); ok {
		r0 = rf(ctx, scopes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetServiceAccountFromContext provides a mock function with given fields: ctx
func (_m *MockAuthService) GetServiceAccountFromContext(ctx context.Context) (*ServiceAccountInfo, error) {
	ret := _m.Called(ctx)

	var r0 *ServiceAccountInfo
	if rf, ok := ret.Get(0).(func(context.Context) *ServiceAccountInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ServiceAccountInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionFromContext provides a mock function with given fields: ctx
func (_m *MockAuthService) GetSessionFromContext(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interceptor provides a mock function with given fields:
func (_m *MockAuthService) Interceptor() grpc.UnaryServerInterceptor {
	ret := _m.Called()

	var r0 grpc.UnaryServerInterceptor
	if rf, ok := ret.Get(0).(func() grpc.UnaryServerInterceptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.UnaryServerInterceptor)
		}
	}

	return r0
}

// IsContextAuthorizedToAccessAccountGroup provides a mock function with given fields: ctx, accountGroupID, scopes
func (_m *MockAuthService) IsContextAuthorizedToAccessAccountGroup(ctx context.Context, accountGroupID string, scopes ...AccessScope) error {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, accountGroupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...AccessScope) error); ok {
		r0 = rf(ctx, accountGroupID, scopes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsContextAuthorizedToAccessAccountGroups provides a mock function with given fields: ctx, accountGroupIDs, scopes
func (_m *MockAuthService) IsContextAuthorizedToAccessAccountGroups(ctx context.Context, accountGroupIDs []string, scopes ...AccessScope) error {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, accountGroupIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...AccessScope) error); ok {
		r0 = rf(ctx, accountGroupIDs, scopes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsContextAuthorizedToAccessAccounts provides a mock function with given fields: ctx, accountIdentifiers, scopes
func (_m *MockAuthService) IsContextAuthorizedToAccessAccounts(ctx context.Context, accountIdentifiers []*resources.AccountIdentifier, scopes ...AccessScope) error {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, accountIdentifiers)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*resources.AccountIdentifier, ...AccessScope) error); ok {
		r0 = rf(ctx, accountIdentifiers, scopes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsContextAuthorizedToAccessPartnerMarket provides a mock function with given fields: ctx, partnerID, marketIDs, scopes
func (_m *MockAuthService) IsContextAuthorizedToAccessPartnerMarket(ctx context.Context, partnerID string, marketIDs []string, scopes ...AccessScope) error {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, partnerID, marketIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, ...AccessScope) error); ok {
		r0 = rf(ctx, partnerID, marketIDs, scopes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsContextAuthorizedToAccessResource provides a mock function with given fields: ctx, resourceID, resourceOwnerID, resourceAttributes, scopes
func (_m *MockAuthService) IsContextAuthorizedToAccessResource(ctx context.Context, resourceID string, resourceOwnerID string, resourceAttributes *iam_v1.StructAttribute, scopes ...AccessScope) error {
	_va := make([]interface{}, len(scopes))
	for _i := range scopes {
		_va[_i] = scopes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, resourceID, resourceOwnerID, resourceAttributes)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *iam_v1.StructAttribute, ...AccessScope) error); ok {
		r0 = rf(ctx, resourceID, resourceOwnerID, resourceAttributes, scopes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetServiceAccountOnContext provides a mock function with given fields: ctx, serviceAccountInfo
func (_m *MockAuthService) SetServiceAccountOnContext(ctx context.Context, serviceAccountInfo *ServiceAccountInfo) context.Context {
	ret := _m.Called(ctx, serviceAccountInfo)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, *ServiceAccountInfo) context.Context); ok {
		r0 = rf(ctx, serviceAccountInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}