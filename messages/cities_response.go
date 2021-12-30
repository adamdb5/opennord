package messages

import "github.com/adamdb5/opennord/pb"

// CitiesResponse holds the response from a Cities RPC call.
type CitiesResponse struct {
	names []string // Array of cities
}

// FormatCitiesResponse converts the protobuffer struct to a CitiesResponse.
func FormatCitiesResponse(response *pb.CitiesResponse) CitiesResponse {
	return CitiesResponse{
		response.GetNames(),
	}
}

// Names returns the names of all cities matching the criteria provided in the initial request.
func (msg CitiesResponse) Names() []string {
	return msg.names
}
