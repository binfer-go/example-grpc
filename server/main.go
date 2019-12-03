package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/proto"
	"io"
	"log"
	"net"
)

var port = ":9900"

type routeGuideServer struct {
	//pb proto.UnimplementedRouteGuideServer
}

/*func (router *routeGuideServer) Getfeature(ctx context.Context, point  *proto.Point) (*proto.Feature, error) {
	fmt.Println("------ getfeature:", point.Longitude, point.Latitude)
	return &proto.Feature{Message:"getfeature"}, nil
}*/

/*func (router *routeGuideServer) ListFeatures(point *proto.Point, server proto.RouteGuide_ListFeaturesServer) error {
	fmt.Println("------ listfeature:", point.Latitude, point.Longitude, server.Context())
	_ = server.Send(&proto.Feature{Message: "listfeature"})
	_ = server.RecvMsg(nil)
	return nil
}*/

/*func (router *routeGuideServer) RecordRoute(stream proto.RouteGuide_RecordRouteServer) error {
	fmt.Println("------ recordRoute:")
	point, err := stream.Recv()
	if err == io.EOF {
		_ = stream.SendAndClose(&proto.Feature{
			Message:              "hello world",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
	}
	fmt.Println("=======", point)
	return nil
}*/

func (router *routeGuideServer) RouteChat(stream proto.RouteGuide_RouteChatServer) error {
	fmt.Println("------ routeChat:")
	point, err := stream.Recv()
	if err != nil {
		return err
	}
	if err == io.EOF {
		_ = stream.Send(&proto.Feature{
			Id:                   1,
			Data:                 byte(1),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
	}
	fmt.Println("=====", point)

	_ = stream.Send(&proto.Feature{
		Message:              "hello ~~~~~",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	return nil
}


func main()  {

	server, err := net.Listen("tcp", "127.0.0.1" + port)
	if err != nil {
		log.Fatalf("tcp listen error: %s \n", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	_ = grpcServer.Serve(server)
}


