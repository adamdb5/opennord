package messages

import "github.com/adamdb5/opennord/pb"

// CitiesRequest contains parameters for the Cities RPC.
type CitiesRequest struct {
	protocol  string // The protocol supported by the server
	obfuscate bool   // If the server supports obfuscation
	country   string // The country that the server is located in
}

// ToProtoBuffer Converts the CitiesRequest structure into the protobuffer structure.
func (msg CitiesRequest) ToProtoBuffer() *pb.CitiesRequest {
	return &pb.CitiesRequest{
		Protocol:  msg.protocol,
		Obfuscate: msg.obfuscate,
		Country:   msg.country,
	}
}
