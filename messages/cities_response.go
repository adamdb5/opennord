package messages

import "github.com/adamdb5/opennord/pb"

// CitiesResponse holds the response from a Cities RPC call.
type CitiesResponse struct {
	names []string
	//protocol  string // The protocol supported by the server
	//obfuscate bool   // If the server supports obfuscation
	//country   string // The country that the server is located in
}

// FormatCitiesResponse converts the protobuffer struct to a CitiesResponse.
func FormatCitiesResponse(response *pb.CitiesResponse) CitiesResponse {
	return CitiesResponse{
		response.GetName(),
	}
}

//// Protocol returns the protocol supported by the server.
//func (msg CitiesResponse) Protocol() string {
//	return msg.protocol
//}
//
//// Obfuscate returns if the server supports obfuscation.
//func (msg CitiesResponse) Obfuscate() bool {
//	return msg.obfuscate
//}
//
//// Country returns the country that the server is located in.
//func (msg CitiesResponse) Country() string {
//	return msg.country
//}

func (msg CitiesResponse) Names() []string {
	return msg.names
}
