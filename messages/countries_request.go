package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// CountriesRequest contains parameters for the Countries RPC.
type CountriesRequest struct {
	Protocol  constants.Protocol // The protocol supported by the server
	Obfuscate constants.Boolean  // If the server supports obfuscation
}

// ToProtoBuffer Converts the CountriesRequest structure into the protobuffer structure.
func (msg CountriesRequest) ToProtoBuffer() *pb.CountriesRequest {
	return &pb.CountriesRequest{
		Protocol:  pb.ProtocolEnum(msg.Protocol),
		Obfuscate: pb.BooleanEnum(msg.Obfuscate),
	}
}
