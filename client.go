package opennord

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/messages"
	"github.com/adamdb5/opennord/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Client is the public stub for of communicating with the NordVPN daemon.
type Client struct {
	grpcConnection *grpc.ClientConn
	daemonClient   pb.DaemonClient
	config         Config
}

type Config struct {
	technology  constants.Protocol
	protocol    constants.Protocol
	firewall    bool
	killSwitch  bool
	cyberSec    bool
	obfuscate   bool
	notify      bool
	autoConnect *pb.SetAutoConnectRequest
	ipv6        bool
	dns         []string
}

// getContext returns a new context object with a timeout. The value of the timeout is
// defined in the common constants.
func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), constants.RequestTimeout)
	return ctx
}

// Status calls the Status RPC and returns a StatusResponse.
func (c Client) Status() (messages.StatusResponse, error) {
	r, err := c.daemonClient.Status(getContext(), &emptypb.Empty{})
	if err != nil {
		return messages.StatusResponse{}, err
	}
	return messages.FormatStatusResponse(r), nil
}

// AccountInfo calls the AccountInfo RPC and returns an AccountResponse.
func (c Client) AccountInfo() (messages.AccountResponse, error) {
	r, err := c.daemonClient.AccountInfo(getContext(), messages.AccountRequest{}.ToProtoBuffer())
	if err != nil {
		return messages.AccountResponse{}, err
	}
	return messages.FormatAccountResponse(r), nil
}

// Cities calls the Cities RPC and returns a CitiesResponse.
func (c Client) Cities(req messages.CitiesRequest) (messages.CitiesResponse, error) {
	r, err := c.daemonClient.Cities(getContext(), req.ToProtoBuffer())
	if err != nil {
		return messages.CitiesResponse{}, err
	}
	return messages.FormatCitiesResponse(r), nil
}

// Connect does something

// Countries calls the Countries RPC and returns a CountriesResponse.
func (c Client) Countries(req messages.CountriesRequest) (messages.CountriesResponse, error) {
	r, err := c.daemonClient.Countries(getContext(), req.ToProtoBuffer())
	if err != nil {
		return messages.CountriesResponse{}, err
	}
	return messages.FormatCountriesResponse(r), nil
}

// Disconnect calls the Disconnect RPC and terminates the current VPN session.
func (c Client) Disconnect() error {
	_, err := c.daemonClient.Disconnect(getContext(), &pb.DisconnectRequest{Id: 0})
	return err
}

// FrontendCountries calls the FrontendCountries RPC and returns a CountriesResponse.
func (c Client) FrontendCountries(req messages.CountriesRequest) (messages.FrontendCountriesResponse, error) {
	r, err := c.daemonClient.FrontendCountries(getContext(), req.ToProtoBuffer())
	if err != nil {
		return messages.FrontendCountriesResponse{}, err
	}
	return messages.FormatFrontendCountriesResponse(r), nil
}

// Groups calls the Groups RPC and returns a GroupsResponse.
func (c Client) Groups(req messages.GroupsRequest) (messages.GroupsResponse, error) {
	r, err := c.daemonClient.Groups(getContext(), req.ToProtoBuffer())
	if err != nil {
		return messages.GroupsResponse{}, err
	}
	return messages.FormatGroupsResponse(r), nil
}

// IsLoggedIn calls the IsLoggedIn RPC and returns a IsLoggedInResponse.
func (c Client) IsLoggedIn() (messages.IsLoggedInResponse, error) {
	r, err := c.daemonClient.IsLoggedIn(getContext(), &emptypb.Empty{})
	if err != nil {
		return messages.IsLoggedInResponse{}, err
	}
	return messages.FormatIsLoggedInResponse(r), nil
}

// Login calls the Login RPC.
func (c Client) Login(req messages.LoginRequest) error {
	_, err := c.daemonClient.Login(getContext(), req.ToProtoBuffer())
	return err
}

// LoginOAuth2 calls the LoginOAuth2 RPC.
func (c Client) LoginOAuth2() (messages.LoginOAuth2Response, error) {
	r, err := c.daemonClient.LoginOAuth2(getContext(), &emptypb.Empty{})
	if err != nil {
		return messages.LoginOAuth2Response{}, err
	}
	return messages.FormatLoginOAuth2LoginResponse(r), err
}

// Logout calls the Logout RPC.
func (c Client) Logout() error {
	_, err := c.daemonClient.Logout(getContext(), &emptypb.Empty{})
	return err
}

// Plans calls the Plans RPC.
func (c Client) Plans() (messages.PlansResponse, error) {
	r, err := c.daemonClient.Plans(getContext(), &emptypb.Empty{})
	if err != nil {
		return messages.PlansResponse{}, err
	}
	return messages.FormatPlansResponse(r), err
}

// Ping checks that the daemon is alive via the Ping RPC.
// If the value returned is nil, the ping succeeded.
func (c Client) Ping() error {
	_, err := c.daemonClient.Ping(getContext(), &emptypb.Empty{})
	return err
}

// RateConnection calls the RateConnection RPC.
func (c Client) RateConnection(req messages.RateConnectionRequest) error {
	_, err := c.daemonClient.RateConnection(getContext(), req.ToProtoBuffer())
	return err
}

// SetAutoConnect calls the SetAutoConnect RPC.
// TODO: fix
func (c Client) SetAutoConnect(req messages.SetAutoConnectRequest) (*pb.Payload, error) {
	r, err := c.daemonClient.SetAutoConnect(getContext(), req.ToProtoBuffer())
	if err != nil {
		return &pb.Payload{}, err
	}
	return r, err
}

// SetWhitelist calls the SetWhitelist RPC.
func (c Client) SetWhitelist(req messages.SetWhitelistRequest) error {
	_, err := c.daemonClient.SetWhitelist(getContext(), req.ToProtoBuffer())
	return err
}

// SetCyberSec calls the SetCyberSec RPC.
func (c Client) SetCyberSec(enabled bool) error {
	_, err := c.daemonClient.SetCyberSec(getContext(), &pb.SetCyberSecRequest{
		CyberSec: enabled},
	)
	return err
}

// SetDefaults calls the SetDefaults RPC.
func (c Client) SetDefaults() error {
	_, err := c.daemonClient.SetDefaults(getContext(), &emptypb.Empty{})
	return err
}

// SetKillSwitch calls the SetKilLSwitch RPC.
func (c Client) SetKillSwitch(enabled bool) error {
	_, err := c.daemonClient.SetKillSwitch(getContext(), &pb.SetKillSwitchRequest{
		Enabled: enabled,
	})
	return err
}

// SetFirewall calls the SetFirewall RPC.
func (c Client) SetFirewall(enabled bool) error {
	_, err := c.daemonClient.SetFirewall(getContext(), &pb.SetGenericRequest{
		Enabled: enabled,
	})
	return err
}

// SetNotify calls the SetNotify RPC.
func (c Client) SetNotify(enabled bool) error {
	_, err := c.daemonClient.SetNotify(getContext(), &pb.SetNotifyRequest{
		Uid:    1000, // Seems to be a magic number
		Notify: enabled,
	})
	return err
}

// SetObfuscate calls the SetObfuscate RPC.
func (c Client) SetObfuscate(enabled bool) error {
	_, err := c.daemonClient.SetObfuscate(getContext(), &pb.SetGenericRequest{
		Enabled: enabled,
	})
	return err
}

// SetProtocol calls the SetProtocol RPC.
func (c Client) SetProtocol(protocol constants.Protocol) error {
	_, err := c.daemonClient.SetProtocol(getContext(), &pb.SetProtocolRequest{
		Protocol: pb.ProtocolEnum(protocol),
	})
	return err
}

// SetTechnology calls the SetTechnology RPC.
func (c Client) SetTechnology(technology constants.Technology) error {
	_, err := c.daemonClient.SetTechnology(getContext(), &pb.SetTechnologyRequest{
		Technology: pb.TechnologyEnum(technology),
	})
	return err
}

// SetIpv6 calls the SetIPv6 RPC.
func (c Client) SetIpv6(enabled bool) error {
	_, err := c.daemonClient.SetIpv6(getContext(), &pb.SetGenericRequest{
		Enabled: enabled,
	})
	return err
}

// SetDns calls the SetDns RPC.
func (c Client) SetDns(servers []string, cyberSec bool) error {
	_, err := c.daemonClient.SetDns(getContext(), &pb.SetDNSRequest{
		Dns:      servers,
		CyberSec: cyberSec,
	})
	return err
}

// Settings calls the Settings RPC. (but the settings RPC doesn't do anything anyway)
func (c Client) Settings() error {
	_, err := c.daemonClient.Settings(getContext(), &pb.SettingsRequest{})
	return err
}

// SettingsProtocols calls the SettingsProtocols RPC.
func (c Client) SettingsProtocols() (*pb.ProtocolsResponse, error) {
	r, err := c.daemonClient.SettingsProtocols(getContext(), &pb.SettingsRequest{})
	if err != nil {
		return &pb.ProtocolsResponse{}, err
	}
	return r, err
}

// SettingsTechnologies calls the SettingsTechnologies RPC.
func (c Client) SettingsTechnologies() (*pb.TechnologyResponse, error) {
	r, err := c.daemonClient.SettingsTechnologies(getContext(), &pb.SettingsRequest{})
	if err != nil {
		return &pb.TechnologyResponse{}, err
	}
	return r, err
}
