package server

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/thepatrik/hellogrpc/internal/pb/mirror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const serviceName = "hellogrpc"

// Option type
type Option func(*Config)

// Config type
type Config struct {
	Logger *log.Logger
}

// WithLogger sets a logger
func WithLogger(logger *log.Logger) Option {
	return func(cfg *Config) {
		cfg.Logger = logger
	}
}

// Server type
type Server struct {
	cfg    *Config
	server *grpc.Server
}

// New constructs a new gRPC server
func New(options ...Option) *Server {
	cfg := &Config{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
	for _, option := range options {
		option(cfg)
	}
	return &Server{cfg: cfg}
}

// Serve accepts incoming connections.
func (s *Server) Serve(port int) error {
	strport := strconv.Itoa(port)
	listener, err := net.Listen("tcp", ":"+strport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	s.server = grpcServer

	grpc_health_v1.RegisterHealthServer(grpcServer, s)
	pb.RegisterMirrorServer(grpcServer, s)

	log.Printf("running gRPC server on port %s", strport)

	return grpcServer.Serve(listener)
}

// GracefulStop stops the gRPC server gracefully. It stops the server from
// accepting new connections and RPCs and blocks until all the pending RPCs are
// finished.
func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}

// Check checks if the requested service is serving.
func (s *Server) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	resp := &grpc_health_v1.HealthCheckResponse{}
	if len(in.Service) == 0 || in.Service == serviceName {
		resp.Status = grpc_health_v1.HealthCheckResponse_SERVING
	}
	return resp, nil
}

// Watch performs a watch for the serving status of the requested service.
func (s *Server) Watch(in *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	resp := &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}
	return server.Send(resp)
}

// MirrorText mirrors a text
func (s *Server) MirrorText(ctx context.Context, in *pb.MirrorTextRequest) (*pb.MirrorTextResponse, error) {
	input := in.Text
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]

	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	// Convert back to UTF-8.
	out := &pb.MirrorTextResponse{Text: string(rune)}
	return out, nil
}
