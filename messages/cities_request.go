package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// CitiesRequest contains parameters for the Cities RPC.
type CitiesRequest struct {
	Protocol  constants.Protocol // The protocol supported by the server
	Obfuscate constants.Boolean  // If the server supports obfuscation
	Country   string             // The country that the server is located in
}

// ToProtoBuffer Converts the CitiesRequest structure into the protobuffer structure.
func (msg CitiesRequest) ToProtoBuffer() *pb.CitiesRequest {
	return &pb.CitiesRequest{
		Protocol:  pb.ProtocolEnum(msg.Protocol),
		Obfuscate: pb.BooleanEnum(msg.Obfuscate),
		Country:   msg.Country,
	}
}
