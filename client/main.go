package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/proto"
	"log"
	"time"
)

const (
	address     = "localhost:9900"
)

func main() {

	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client server error: %s \n", err)
	}
	defer conn.Close()

	routeServe := proto.NewRouteGuideClient(conn)
	feature, err := routeServe.RouteChat(context.Background())
	if err != nil {
		log.Printf("client feature error: %s \n", err)
	}
	_ = feature.Send(&proto.Point{
		Latitude:             111,
		Longitude:            231,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	fmt.Println(feature.Recv())

}

