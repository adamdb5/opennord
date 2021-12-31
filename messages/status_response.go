package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// StatusResponse holds the response from a Status RPC call.
type StatusResponse struct {
	state      string               // The state of the VPN
	technology constants.Technology // The technology currently being used
	protocol   constants.Protocol   // The protocol currently being used
	ip         string               // The remote server's IP address
	hostname   string               // The remote server's hostname
	country    string               // The remote server's country
	city       string               // The remote server's city
	download   int64                // The number of bytes downloaded in this session
	upload     int64                // The number of bytes uploaded in this session
	uptime     int64                // The time elapsed for this session in nanoseconds
}

// FormatStatusResponse converts the protobuffer struct to an StatusResponse.
func FormatStatusResponse(response *pb.StatusResponse) StatusResponse {
	var technology constants.Technology
	var protocol constants.Protocol

	switch response.GetTechnology() {
	case pb.TechnologyEnum_OPENVPN:
		technology = constants.OpenVPN
	case pb.TechnologyEnum_NORDLYNX:
		technology = constants.NordLynx
	case pb.TechnologyEnum_SKYLARK:
		technology = constants.Skylark
	default:
		technology = constants.Unknown
	}

	switch response.GetProtocol() {
	case pb.ProtocolEnum_UDP:
		protocol = constants.UDP
	case pb.ProtocolEnum_TCP:
		protocol = constants.TCP
	default:
		protocol = constants.Unknown
	}

	return StatusResponse{
		state:      response.GetState(),
		technology: technology,
		protocol:   protocol,
		ip:         response.GetIp(),
		hostname:   response.GetHostname(),
		country:    response.GetCountry(),
		city:       response.GetCity(),
		download:   response.GetDownload(),
		upload:     response.GetUpload(),
		uptime:     response.GetUptime(),
	}
}

// State returns the current VPN state.
func (msg StatusResponse) State() string {
	return msg.state
}

// Technology returns the current technology being used.
// Possible values are OpenVPN(1), NordLynx(2), Skylark(3)
func (msg StatusResponse) Technology() constants.Technology {
	return msg.technology
}

// Protocol returns the current protocol being used.
// Possible values are UDP(1), TCP(2)
func (msg StatusResponse) Protocol() constants.Protocol {
	return msg.protocol
}

// Ip returns the IP address of the remote server.
func (msg StatusResponse) Ip() string {
	return msg.ip
}

// Hostname returns the hostname of the remote server.
func (msg StatusResponse) Hostname() string {
	return msg.hostname
}

// Country returns the country of the remote server.
func (msg StatusResponse) Country() string {
	return msg.country
}

// City returns the city of the remote server.
func (msg StatusResponse) City() string {
	return msg.city
}

// Downloaded returns the number of bytes downloaded in this session.
func (msg StatusResponse) Downloaded() int64 {
	return msg.download
}

// Uploaded returns the number of bytes uploaded in this session.
func (msg StatusResponse) Uploaded() int64 {
	return msg.download
}

// Uptime returns the time elapsed for this session in nanoseconds
func (msg StatusResponse) Uptime() int64 {
	return msg.uptime
}
