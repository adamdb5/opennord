package messages

import "github.com/adamdb5/opennord/pb"

// CountriesResponse holds the response from a Countries RPC call.
type CountriesResponse struct {
	names []string // Array of countries
}

// FormatCountriesResponse converts the protobuffer struct to a CountriesResponse.
func FormatCountriesResponse(response *pb.CountriesResponse) CountriesResponse {
	return CountriesResponse{
		response.GetNames(),
	}
}

// Names returns the names of all countries matching the criteria provided in the initial request.
func (msg CountriesResponse) Names() []string {
	return msg.names
}
