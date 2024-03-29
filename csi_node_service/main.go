package main

import (
	"log"
	"net"
	"os"

	"github.com/PranoSA/NFS_API_CSI/identity"
	"github.com/PranoSA/NFS_API_CSI/node"
	"github.com/PranoSA/NFS_API_CSI/proto"
	"google.golang.org/grpc"
)

/**
 *
 * This is the main.go file for the CSI Service
 */

func main() {
	//controllerServer := controller.NFSControllerService{UnimplementedControllerServer: proto.UnimplementedControllerServer{}}
	nodeService := node.NFSNodeService{UnimplementedNodeServer: proto.UnimplementedNodeServer{}}
	identityService := identity.NFSIdentityService{UnimplementedIdentityServer: proto.UnimplementedIdentityServer{}}

	s := grpc.NewServer()

	// Get CSI Driver Socket From Env
	sock := os.Getenv("CSI_ENDPOINT")

	//proto.RegisterControllerServer(s, &controllerServer)
	proto.RegisterIdentityServer(s, &identityService)
	proto.RegisterNodeServer(s, &nodeService)

	// Listen on a TCP port
	// Listen on a Unix domain socket
	//"/var/lib/kubelet/plugins/my-driver/csi.sock"
	lis, err := net.Listen("unix", sock)
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
