package controller

import (
	"context"

	"github.com/PranoSA/NFS_API_CSI/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.ControllerServer = &NFSControllerService{}

type NFSControllerService struct {
	proto.UnimplementedControllerServer
}

type NfsShare struct {
	Id     string
	Server string
	Path   string
	Size   int64
}

func allocateNfsShare(name string, required_bytes int64) (*NfsShare, error) {

	//What is the API or DOMAIN NAME ????

	// HELM CHART OR NAME SPACE???

	return &NfsShare{
		Id:     "1",
		Server: "",
		Path:   "",
		Size:   required_bytes,
	}, nil
}

/**
 * Actualy ALlocate Volume, What is the "state" ????
 *
 *
 */
func (cs *NFSControllerService) CreateVolume(ctx context.Context, req *proto.CreateVolumeRequest) (*proto.CreateVolumeResponse, error) {
	nfsShare, err := allocateNfsShare(req.GetName(), req.GetCapacityRange().GetRequiredBytes())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Create a VolumeContext with the NFS server and share
	volumeContext := map[string]string{
		"nfsServer": nfsShare.Server,
		"nfsShare":  nfsShare.Path,
	}

	return &proto.CreateVolumeResponse{
		Volume: &proto.Volume{
			VolumeId:      nfsShare.Id,
			CapacityBytes: nfsShare.Size,
			VolumeContext: volumeContext,
		},
	}, nil
	return nil, nil
}

func (cs *NFSControllerService) DeleteVolume(ctx context.Context, req *proto.DeleteVolumeRequest) (*proto.DeleteVolumeResponse, error) {
	return nil, nil
}

/**
 *
 */

func (cs *NFSControllerService) ControllerPublishVolume(ctx context.Context, req *proto.ControllerPublishVolumeRequest) (*proto.ControllerPublishVolumeResponse, error) {
	return nil, nil
}

func (cs *NFSControllerService) ControllerUnpublishVolume(ctx context.Context, req *proto.ControllerUnpublishVolumeRequest) (*proto.ControllerUnpublishVolumeResponse, error) {
	return nil, nil
}

func (cs *NFSControllerService) ValidateVolumeCapabilities(ctx context.Context, req *proto.ValidateVolumeCapabilitiesRequest) (*proto.ValidateVolumeCapabilitiesResponse, error) {
	return nil, nil
}
