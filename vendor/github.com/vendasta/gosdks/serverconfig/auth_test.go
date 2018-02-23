package serverconfig

import (
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/vendasta/gosdks/logging"

	"golang.org/x/net/context"

	"encoding/base64"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApiKeyAuthorization(t *testing.T) {
	Convey("Given an ApiKey Authorization Interceptor", t, func() {
		ti := NewAPIKeyAuthInterceptor(map[string][]APIKeyUser{
			"test": []APIKeyUser{
				APIKeyUser{Key: "test-key1", UID: "uid-1"},
			},
			"demo": []APIKeyUser{
				APIKeyUser{Key: "demo-key1", UID: "uid-2"},
				APIKeyUser{Key: "demo-key2", UID: "uid-3"},
			},
		})

		logging.InitSilentLogger()

		interceptor := ti.Interceptor()
		ctx := context.Background()
		os.Setenv("ENVIRONMENT", "test")

		Convey("API Key is required", func() {
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})

		Convey("Authorization header must include \"Bearer\"", func() {
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{"authorization": []string{"test-key1"}})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})
		Convey("Invalid token refused", func() {
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{"authorization": []string{"Bearer invalid-key"}})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})
		Convey("Valid token accepted", func() {
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{"authorization": []string{"Bearer test-key1"}})
			testHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return nil, nil
			}
			_, err := interceptor(ctx, nil, nil, testHandler)
			So(err, ShouldEqual, nil)
		})
		Convey("Environment respected", func() {
			os.Setenv("ENVIRONMENT", "demo")
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{"authorization": []string{"Bearer test-key1"}})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})
	})
}

func TestJwtAuthorization(t *testing.T) {
	Convey("Given a JWT Authorization Interceptor", t, func() {
		ti := NewJwtAuthInterceptor(map[string][]string{
			"test": []string{
				"uid-1@example.com",
			},
			"demo": []string{
				"uid-2@example.com",
				"uid-3@example.com",
			},
		})
		logging.InitSilentLogger()

		jwtKey := "x-endpoint-api-userinfo"

		interceptor := ti.Interceptor()
		ctx := context.Background()
		os.Setenv("ENVIRONMENT", "test")

		Convey("API Key is required", func() {
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})
		Convey("JWT must be base64 encoded", func() {
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{jwtKey: []string{"not base64"}})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})

		Convey("Invalid user refused", func() {
			token := `{
				"aud":"some-audience",
				"email":"invalid-email@example.com",
				"email_verified":true,
				"iss":"some-issuer",
				"iat":1491329072,
				"exp":1491332672}`
			encodedToken := base64.RawURLEncoding.EncodeToString([]byte(token))
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{jwtKey: []string{encodedToken}})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})
		Convey("Valid user accepted", func() {
			token := `{
				"aud":"some-audience",
				"email":"uid-1@example.com",
				"email_verified":true,
				"iss":"some-issuer",
				"iat":1491329072,
				"exp":1491332672}`
			encodedToken := base64.URLEncoding.EncodeToString([]byte(token))
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{jwtKey: []string{encodedToken}})
			testHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return nil, nil
			}
			_, err := interceptor(ctx, nil, nil, testHandler)
			So(err, ShouldEqual, nil)
		})
		Convey("Environment respected", func() {
			os.Setenv("ENVIRONMENT", "demo")
			token := `{
				"aud":"some-audience",
				"email":"uid-1@example.com",
				"email_verified":true,
				"iss":"some-issuer",
				"iat":1491329072,
				"exp":1491332672}`
			encodedToken := base64.RawURLEncoding.EncodeToString([]byte(token))
			ctx = metadata.NewIncomingContext(ctx, metadata.MD{jwtKey: []string{encodedToken}})
			_, err := interceptor(ctx, nil, nil, nil)
			code := grpc.Code(err)
			So(code, ShouldEqual, codes.PermissionDenied)
		})
	})
}
