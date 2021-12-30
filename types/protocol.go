package types

import "github.com/adamdb5/opennord/pb"

type Protocol int32

const (
	TCP = Protocol(pb.ProtocolEnum_TCP)
	UDP = Protocol(pb.ProtocolEnum_UDP)
)
