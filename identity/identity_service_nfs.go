package identity

import (
	"context"

	"github.com/PranoSA/NFS_API_CSI/proto"
)

var _ proto.IdentityServer = &NFSIdentityService{}

type NFSIdentityService struct {
	proto.UnimplementedIdentityServer
}

func (is *NFSIdentityService) GetPluginInfo(ctx context.Context, req *proto.GetPluginInfoRequest) (*proto.GetPluginInfoResponse, error) {
	return &proto.GetPluginInfoResponse{
		Name:          "NFS-CSI",
		VendorVersion: "0.0.1",
	}, nil
}

func (is *NFSIdentityService) GetPluginCapabilities(ctx context.Context, req *proto.GetPluginCapabilitiesRequest) (*proto.GetPluginCapabilitiesResponse, error) {
	return &proto.GetPluginCapabilitiesResponse{
		Capabilities: []*proto.PluginCapability{
			{
				Type: &proto.PluginCapability_Service_{
					Service: &proto.PluginCapability_Service{
						Type: proto.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
			{
				Type: &proto.PluginCapability_Service_{
					Service: &proto.PluginCapability_Service{
						Type: proto.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
					},
				},
			},
		},
	}, nil
}
