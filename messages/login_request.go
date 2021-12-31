package messages

import (
	"github.com/adamdb5/opennord/pb"
)

// LoginRequest contains parameters for the Login RPC.
type LoginRequest struct {
	Username string // Username to use for logging in
	Password string // Plaintext password to use for logging in
}

// ToProtoBuffer Converts the LoginRequest structure into the protobuffer structure.
func (msg LoginRequest) ToProtoBuffer() *pb.LoginRequest {
	return &pb.LoginRequest{
		Username: msg.Username,
		Password: msg.Password,
	}
}
