package types

import "github.com/adamdb5/opennord/pb"

type Technology int32

const (
	OpenVPN  = Technology(pb.TechnologyEnum_OPENVPN)  // OpenVPN technology
	NordLynx = Technology(pb.TechnologyEnum_NORDLYNX) // NordLynx technology
	Skylark  = Technology(pb.TechnologyEnum_SKYLARK)  // Skylark technology
)
