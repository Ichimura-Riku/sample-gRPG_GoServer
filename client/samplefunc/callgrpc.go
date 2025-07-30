package samplefunc

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	weatherv1 "buf-demo/example/gen/go/jyapp/weather/v1"
)

const grpcServerAddr = "localhost:8080"

func CallGRPC() {
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	client := weatherv1.NewWeatherServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// サンプル値でリクエスト
	req := &weatherv1.GetWeatherRequest{
		Latitude:  35.0,
		Longitude: 139.0,
	}
	resp, err := client.GetWeather(ctx, req)
	if err != nil {
		log.Printf("gRPC error: %v", err)
		return
	}
	fmt.Printf("gRPC Response: Temperature=%.2f, Condition=%v\n", resp.Temperature, resp.Conditions)
}
