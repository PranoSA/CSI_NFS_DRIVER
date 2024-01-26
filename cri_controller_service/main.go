package main

import (
	"log"
	"net"

	"github.com/PranoSA/NFS_API_CSI/controller"
	"github.com/PranoSA/NFS_API_CSI/identity"
	"github.com/PranoSA/NFS_API_CSI/proto"
	"google.golang.org/grpc"
)

/**
 *
 * This is the main.go file for the CSI Service
 */

func main() {
	controllerServer := controller.NFSControllerService{UnimplementedControllerServer: proto.UnimplementedControllerServer{}}
	//nodeService := node.NFSNodeService{UnimplementedNodeServer: proto.UnimplementedNodeServer{}}
	identityService := identity.NFSIdentityService{UnimplementedIdentityServer: proto.UnimplementedIdentityServer{}}

	s := grpc.NewServer()

	proto.RegisterControllerServer(s, &controllerServer)
	proto.RegisterIdentityServer(s, &identityService)
	//proto.RegisterNodeServer(s, &nodeService)

	// Listen on a TCP port
	// Listen on a Unix domain socket
	lis, err := net.Listen(":50051", "0.0.0.0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Serve the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
