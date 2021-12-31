package messages

import (
	"github.com/adamdb5/opennord/pb"
)

// SetWhitelistRequest contains parameters for the SetWhitelist RPC.
type SetWhitelistRequest struct {
	Whitelist Whitelist // The whitelist struct containing the ports and subnets
}

// ToProtoBuffer Converts the SetWhitelistRequest structure into the protobuffer structure.
func (msg SetWhitelistRequest) ToProtoBuffer() *pb.SetWhitelistRequest {
	return &pb.SetWhitelistRequest{
		Whitelist: msg.Whitelist.ToProtoBuffer(),
	}
}
