package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/jirawan-chuapradit/grpc-gateway-example/pkg/example"
	"google.golang.org/grpc"
)

func main() {

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

	restPort := "8095"
	log.Printf("Serving gRPC-Gateway on http://localhost:%s", restPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", restPort), mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
