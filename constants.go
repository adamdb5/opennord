package opennord

import (
	"github.com/adamdb5/opennord/pb"
	"time"
)

type AccountType int32
type Boolean int32
type Protocol int32
type Technology int32

const (
	SocketPath     = "/run/nordvpn/nordvpnd.sock"
	SocketAddress  = "unix://" + SocketPath
	RequestTimeout = 5 * time.Second
)

const (
	Inactive = AccountType(pb.AccountTypeEnum_INACTIVE)
	Active   = AccountType(pb.AccountTypeEnum_ACTIVE)
)

const (
	UDP = Protocol(pb.ProtocolEnum_UDP)
	TCP = Protocol(pb.ProtocolEnum_TCP)
)

const (
	OpenVPN  = Technology(pb.TechnologyEnum_OPENVPN)
	NordLynx = Technology(pb.TechnologyEnum_NORDLYNX)
)
