package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// IsLoggedInResponse holds the response from an IsLoggedIn RPC call.
type IsLoggedInResponse struct {
	isLoggedIn constants.Boolean // If the user is currently logged in
}

// FormatIsLoggedInResponse converts the protobuffer struct to a IsLoggedInResponse.
func FormatIsLoggedInResponse(response *pb.IsLoggedInResponse) IsLoggedInResponse {
	return IsLoggedInResponse{
		isLoggedIn: constants.Boolean(response.GetIsLoggedIn()),
	}
}

// IsLoggedIn returns whether the user is currently logged in.
func (msg IsLoggedInResponse) IsLoggedIn() constants.Boolean {
	return msg.isLoggedIn
}
