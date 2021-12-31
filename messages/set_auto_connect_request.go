package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// SetAutoConnectRequest contains parameters for the SetAutoConnect RPC.
type SetAutoConnectRequest struct {
	ServerTag   string             // The server to use when autoconnecting
	Protocol    constants.Protocol // The protocol to use when autoconnecting
	CyberSec    constants.Boolean  // Whether cyberSec should be used when autoconnecting
	Obfuscate   constants.Boolean  // Whether obfuscation should be used when autoconnecting
	AutoConnect constants.Boolean  // Whether the NordVPN daemon should autoconnect
	Dns         constants.Boolean  // Whether custom dns should be used when autoconnecting
	Whitelist   constants.Boolean  // Whether whitelist should be used when autoconnecting
}

// ToProtoBuffer Converts the SetAutoConnectRequest structure into the protobuffer structure.
func (msg SetAutoConnectRequest) ToProtoBuffer() *pb.SetAutoConnectRequest {
	return &pb.SetAutoConnectRequest{
		ServerTag:   msg.ServerTag,
		Protocol:    pb.ProtocolEnum(msg.Protocol),
		CyberSec:    pb.BooleanEnum(msg.CyberSec),
		Obfuscate:   pb.BooleanEnum(msg.Obfuscate),
		AutoConnect: pb.BooleanEnum(msg.AutoConnect),
		Dns:         pb.BooleanEnum(msg.Dns),
		Whitelist:   pb.BooleanEnum(msg.Whitelist),
	}
}
