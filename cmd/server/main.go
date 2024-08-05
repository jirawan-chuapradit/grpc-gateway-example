package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pb "github.com/jirawan-chuapradit/grpc-gateway-example/pkg/example"
	ddgrpc "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Start a new span for the SayHello method
	span, _ := tracer.StartSpanFromContext(ctx, "server.SayHello", tracer.SpanType(ext.SpanTypeWeb))
	defer span.Finish()
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	// Start the DataDog tracer
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("example-grpc-server"),
		tracer.WithGlobalTag("component", "example-grpc-server"),
		tracer.WithServiceVersion("v1.3.0"),
		// tc.WithAgentAddr(t.Config.AgentHost),
		tracer.WithAnalytics(true),
		tracer.WithRuntimeMetrics(),
	)

	defer tracer.Stop()

	grpcPort := "50051"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			ddgrpc.UnaryServerInterceptor(ddgrpc.WithServiceName("example-grpc-server")),
		),
	)
	pb.RegisterExampleServiceServer(s, &server{})
	log.Printf("Serving gRPC on :%s\n", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
