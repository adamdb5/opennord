package opennord

import (
	"time"
)

type Boolean uint32
type Protocol uint32
type Technology uint32

const (
	SocketPath     = "/run/nordvpn/nordvpnd.sock"
	SocketAddress  = "unix://" + SocketPath
	RequestTimeout = 5 * time.Second

	ErrOk = 1000
)
