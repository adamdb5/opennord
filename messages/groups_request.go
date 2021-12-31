package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// GroupsRequest contains parameters for the Groups RPC.
type GroupsRequest struct {
	Protocol  constants.Protocol // Protocol of the desired server groups
	Obfuscate constants.Boolean  // If the desired servers need to support obfuscation
}

// ToProtoBuffer Converts the GroupsRequest structure into the protobuffer structure.
func (msg GroupsRequest) ToProtoBuffer() *pb.GroupsRequest {
	return &pb.GroupsRequest{
		Protocol:  pb.ProtocolEnum(msg.Protocol),
		Obfuscate: pb.BooleanEnum(msg.Obfuscate),
	}
}
