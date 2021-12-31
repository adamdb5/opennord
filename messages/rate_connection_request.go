package messages

import (
	"github.com/adamdb5/opennord/pb"
)

// RateConnectionRequest contains parameters for the RateConnection RPC.
type RateConnectionRequest struct {
	Rating int32 // The rating for the connection
}

// ToProtoBuffer Converts the RateConnectionRequest structure into the protobuffer structure.
func (msg RateConnectionRequest) ToProtoBuffer() *pb.RateConnectionRequest {
	return &pb.RateConnectionRequest{
		Rating: msg.Rating,
	}
}
