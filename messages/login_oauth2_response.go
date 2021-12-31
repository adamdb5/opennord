package messages

import "github.com/adamdb5/opennord/pb"

// LoginOAuth2LoginResponse holds the response from a LoginOAuth2 RPC call.
type LoginOAuth2Response struct {
	url string // OAuth URL
}

// FormatLoginOAuth2LoginResponse converts the protobuffer struct to a FormatLoginOAuth2Response.
func FormatLoginOAuth2LoginResponse(response *pb.LoginOAuth2Response) LoginOAuth2Response {
	return LoginOAuth2Response{
		url: response.GetUrl(),
	}
}

// Url returns OAuth2 URL.
func (msg LoginOAuth2Response) Url() string {
	return msg.url
}
