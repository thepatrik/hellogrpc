package server_test

import (
	"context"
	"os"
	"testing"

	pb "github.com/thepatrik/hellogrpc/internal/pb/mirror"

	"github.com/stretchr/testify/require"
	"github.com/thepatrik/hellogrpc/pkg/server"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var grpcServer *server.Server

func TestMain(m *testing.M) {
	grpcServer = server.New()
	code := m.Run()
	os.Exit(code)
}

func TestHealthCheckServing(t *testing.T) {
	req := &grpc_health_v1.HealthCheckRequest{}
	actual, err := grpcServer.Check(context.Background(), req)
	require.Nil(t, err, "Error should be nil")
	expected := grpc_health_v1.HealthCheckResponse_SERVING
	require.Equal(t, expected, actual.Status)
}

func TestHealthCheckServiceName(t *testing.T) {
	req := &grpc_health_v1.HealthCheckRequest{Service: "hellogrpc"}
	actual, err := grpcServer.Check(context.Background(), req)
	require.Nil(t, err, "Error should be nil")
	expected := grpc_health_v1.HealthCheckResponse_SERVING
	require.Equal(t, expected, actual.Status)
}

func TestHealthCheckUnknown(t *testing.T) {
	req := &grpc_health_v1.HealthCheckRequest{Service: "foo"}
	actual, err := grpcServer.Check(context.Background(), req)
	require.Nil(t, err, "Error should be nil")
	expected := grpc_health_v1.HealthCheckResponse_UNKNOWN
	require.Equal(t, expected, actual.Status)
}

func TestGetItemPrices(t *testing.T) {
	req := &pb.MirrorTextRequest{Text: "Hello World!"}
	res, err := grpcServer.MirrorText(context.Background(), req)
	require.Nil(t, err, "err should be nil")
	require.NotNil(t, res, "res should not be nil")
	require.Equal(t, "!dlroW olleH", res.Text)
}
