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
	SOCKET_PATH     = "/run/nordvpn/nordvpnd.sock"
	SOCKET_ADDRESS  = "unix://" + SOCKET_PATH
	REQUEST_TIMEOUT = 5 * time.Second
)

const (
	INACTIVE = AccountType(pb.AccountTypeEnum_INACTIVE)
	ACTIVE   = AccountType(pb.AccountTypeEnum_ACTIVE)
)

const (
	FALSE = Boolean(pb.BooleanEnum_FALSE)
	TRUE  = Boolean(pb.BooleanEnum_TRUE)
)

const (
	UDP = Protocol(pb.ProtocolEnum_UDP)
	TCP = Protocol(pb.ProtocolEnum_TCP)
)

const (
	OPENVPN  = Technology(pb.TechnologyEnum_OPENVPN)
	NORDLYNX = Technology(pb.TechnologyEnum_NORDLYNX)
)
