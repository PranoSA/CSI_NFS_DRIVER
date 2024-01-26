package node

import (
	"context"
	"fmt"
	"os"
	"syscall"

	csi "github.com/PranoSA/NFS_API_CSI/proto"
	proto "github.com/PranoSA/NFS_API_CSI/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
 *
 * 	Node Service Implimentation For CSI Drivers
 *  Define all the functionality in this package -> Maybe Split it Off Later
 */

var _ proto.NodeServer = &NFSNodeService{}

type NFSNodeService struct {
	csi.UnimplementedNodeServer
}

func mount(source string, target string, fstype string, flags uintptr, data string) error {
	// Mount the NFS share at the staging target path
	err := os.MkdirAll(target, 0755)

	if err != nil {
		return err
	}

	// Mount the NFS share
	err = syscall.Mount(source, target, fstype, flags, data)

	if err != nil {
		return err
	}

	return nil
}

func mountNfsShare(server string, share string, targetPath string) error {
	// Mount the NFS share at the staging target path

	// Create the target path if it doesn't exist
	err := os.MkdirAll(targetPath, 0755)

	if err != nil {
		return err
	}

	// Mount the NFS share
	err = mount(server+":"+share, targetPath, "nfs", 0, "")

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

/**
 *
 * Attach Volume to Node
 *	Mount Volume to Node
 * 	Bind Volume to Node
 */
func (s *NFSNodeService) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	// Implement your logic here
	nfsServer, ok := req.GetVolumeContext()["server"]
	if !ok {
		return nil, fmt.Errorf("server not found in volume context")
	}

	nfsPath, ok := req.GetVolumeContext()["path"]

	if !ok {
		return nil, fmt.Errorf("path not found in volume context")
	}

	// Mount the NFS share at the staging target path
	err := mountNfsShare(nfsServer, nfsPath, req.GetStagingTargetPath())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.NodeStageVolumeResponse{}, nil

}

func (s *NFSNodeService) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	// Implement your logic here
	return nil, nil
}

/**
 *
 * What Does This Do
 */
func bindMount(source string, target string) error {
	// Bind-mount the staging target path to the publish target path
	err := os.MkdirAll(target, 0755)

	if err != nil {
		return err
	}

	// Bind-mount the staging target path to the publish target path
	err = syscall.Mount(source, target, "", syscall.MS_BIND, "")

	if err != nil {
		return err
	}

	return nil
}

func (s *NFSNodeService) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// Implement your logic here
	// Bind-mount the staging target path to the publish target path
	err := bindMount(req.GetStagingTargetPath(), req.GetTargetPath())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

func (s *NFSNodeService) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	// Implement your logic here
	return nil, nil
}

func (s *NFSNodeService) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	// Implement your logic here
	return nil, nil
}

func (s *NFSNodeService) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	// Implement your logic here
	return nil, nil
}

func (s *NFSNodeService) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	// Implement your logic here

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
		},
	}, nil

}

func (s *NFSNodeService) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {

	return &csi.NodeGetInfoResponse{
		NodeId:            "nfs_node",
		MaxVolumesPerNode: 256,
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				"topology.kubernetes.io/region": "us-east-1",
				"topology.kubernetes.io/zone":   "us-east-1a",
			},
		},
	}, nil
}
