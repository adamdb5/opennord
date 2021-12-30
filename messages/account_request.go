package messages

import "github.com/adamdb5/opennord/pb"

// AccountRequest is a placeholder for an empty request.
// The AccountInfo RPC call does not require any parameters.
type AccountRequest struct{}

// ToProtoBuffer Converts the AccountRequest structure into the protobuffer structure.
func (msg AccountRequest) ToProtoBuffer() *pb.AccountRequest {
	return &pb.AccountRequest{}
}
