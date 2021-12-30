package types

import "github.com/adamdb5/opennord/pb"

type Technology int32

const (
	OpenVPN  = Technology(pb.TechnologyEnum_OPENVPN)
	NordLynx = Technology(pb.TechnologyEnum_NORDLYNX)
	Skylark  = Technology(pb.TechnologyEnum_SKYLARK)
)
