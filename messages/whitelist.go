package messages

import "github.com/adamdb5/opennord/pb"

// Ports holds lists of UDP and TCP port numbers
type Ports struct {
	Udp []int32 // The UDP ports
	Tcp []int32 // The TCP ports
}

// Whitelist holds a Ports struct and a slice of subnets in CIDR notation.
type Whitelist struct {
	ports   Ports    // The UDP and TCP ports
	subnets []string // The subnets
}

// ToProtoBuffer Converts the Ports structure into the protobuffer structure.
func (p Ports) ToProtoBuffer() *pb.Ports {
	return &pb.Ports{
		Tcp: p.Tcp,
		Udp: p.Udp,
	}
}

// ToProtoBuffer Converts the Whitelist structure into the protobuffer structure.
func (w Whitelist) ToProtoBuffer() *pb.Whitelist {
	return &pb.Whitelist{
		Ports:   w.ports.ToProtoBuffer(),
		Subnets: w.subnets,
	}
}
