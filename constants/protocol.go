package constants

import "github.com/adamdb5/opennord/pb"

type Protocol int32

const (
	UDP = Protocol(pb.ProtocolEnum_UDP) // UDP protocol
	TCP = Protocol(pb.ProtocolEnum_TCP) // TCP protocol
)
