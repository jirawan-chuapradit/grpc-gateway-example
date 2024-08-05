package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/jirawan-chuapradit/grpc-gateway-example/pkg/example"
	"google.golang.org/grpc"
	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	// Start the DataDog tracer
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("example-gateway"),
		tracer.WithGlobalTag("component", "example-gateway"),
		tracer.WithServiceVersion("v1.3.0"),
		// tc.WithAgentAddr(t.Config.AgentHost),
		tracer.WithAnalytics(true),
		tracer.WithRuntimeMetrics(),
	)

	defer tracer.Stop()

	grpcPort := "50051"
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterExampleServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", grpcPort), opts)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	tracedMux := ddhttp.WrapHandler(mux, "example-gateway", "http.router")
	restPort := "8095"
	log.Printf("Serving gRPC-Gateway on http://localhost:%s", restPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", restPort), tracedMux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
