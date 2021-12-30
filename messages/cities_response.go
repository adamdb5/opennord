package messages

import "github.com/adamdb5/opennord/pb"

// CitiesResponse holds the response from a Cities RPC call.
type CitiesResponse struct {
	names []string
}

// FormatCitiesResponse converts the protobuffer struct to a CitiesResponse.
func FormatCitiesResponse(response *pb.CitiesResponse) CitiesResponse {
	return CitiesResponse{
		response.GetNames(),
	}
}

func (msg CitiesResponse) Names() []string {
	return msg.names
}
