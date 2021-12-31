package constants

import "time"

const (
	Path           = "/run/nordvpn/nordvpnd.sock" // Filesystem path to socket
	Address        = "unix://" + Path             // Unix domain socket for nordvpnd
	RequestTimeout = 5 * time.Second              // Maximum time to wait for any request

	Unknown = 0 // Parameter value does not match any predefined constants.
)
