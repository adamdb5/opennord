package messages

import (
	"github.com/adamdb5/opennord/pb"
	"github.com/adamdb5/opennord/types"
)

// CitiesRequest contains parameters for the Cities RPC.
type CitiesRequest struct {
	Protocol  types.Protocol // The protocol supported by the server
	Obfuscate types.Boolean  // If the server supports obfuscation
	Country   string         // The country that the server is located in
}

// ToProtoBuffer Converts the CitiesRequest structure into the protobuffer structure.
func (msg CitiesRequest) ToProtoBuffer() *pb.CitiesRequest {
	return &pb.CitiesRequest{
		Protocol:  pb.ProtocolEnum(msg.Protocol),
		Obfuscate: pb.BooleanEnum(msg.Obfuscate),
		Country:   msg.Country,
	}
}
