package messages

import "github.com/adamdb5/opennord/pb"

// FrontendCountry holds a country's name and code.
type FrontendCountry struct {
	name string // The country's name
	code string // The country's code, e.g. us, uk
}

// FormatFrontendCountry converts the protobuffer struct to a FrontendCountry.
func FormatFrontendCountry(response *pb.FrontendCountry) FrontendCountry {
	return FrontendCountry{
		name: response.GetName(),
		code: response.GetCode(),
	}
}

// Name returns the name of the country.
func (msg FrontendCountry) Name() string {
	return msg.name
}

// Code returns the name of the country.
func (msg FrontendCountry) Code() string {
	return msg.code
}
