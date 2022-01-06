package opennord

import (
	"time"
)

const (
	SocketPath     = "/run/nordvpn/nordvpnd.sock"
	SocketAddress  = "unix://" + SocketPath
	RequestTimeout = 5 * time.Second

	ErrOk              = 1000
	StatusOk           = 1000
	StatusConnecting   = 1001
	StatusConnected    = 1002
	StatusGenericError = 2000
	StatusAlreadyRated = 3019
)
