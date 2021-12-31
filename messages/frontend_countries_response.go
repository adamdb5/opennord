package messages

import "github.com/adamdb5/opennord/pb"

// FrontendCountriesResponse holds the response from a FrontendCountries RPC call.
type FrontendCountriesResponse struct {
	countries []FrontendCountry // Array of countries
}

// FormatFrontendCountriesResponse converts the protobuffer struct to a FrontendCountriesResponse.
func FormatFrontendCountriesResponse(response *pb.FrontendCountriesResponse) FrontendCountriesResponse {
	var countries []FrontendCountry
	for i := 0; i < len(response.GetCountries()); i++ {
		countries[i] = FrontendCountry{
			name: response.GetCountries()[i].GetName(),
			code: response.GetCountries()[i].GetCode(),
		}
	}
	return FrontendCountriesResponse{
		countries: countries,
	}
}

// Countries returns the countries matching the criteria provided in the initial request.
func (msg FrontendCountriesResponse) Countries() []FrontendCountry {
	return msg.countries
}
