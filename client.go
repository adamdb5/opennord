package opennord

import (
	"github.com/adamdb5/opennord/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Client is the public stub for of communicating with the NordVPN daemon.
type Client struct {
	grpcConnection *grpc.ClientConn
	daemonClient   pb.DaemonClient
}

// getContext returns a new context object with a timeout. The value of the timeout is defined in the constants.
func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), RequestTimeout)
	return ctx
}

// Status calls the Status RPC and returns a StatusResponse.
func (c Client) Status() (*pb.StatusResponse, error) {
	return c.daemonClient.Status(getContext(), &emptypb.Empty{})
}

// AccountInfo calls the AccountInfo RPC and returns an AccountResponse.
func (c Client) AccountInfo() (*pb.AccountResponse, error) {
	return c.daemonClient.AccountInfo(getContext(), &pb.AccountRequest{})
}

// Cities calls the Cities RPC and returns a CitiesResponse.
func (c Client) Cities(req *pb.CitiesRequest) (*pb.CitiesResponse, error) {
	return c.daemonClient.Cities(getContext(), req)
}

// Connect does something

// Countries calls the Countries RPC and returns a CountriesResponse.
func (c Client) Countries(req *pb.CountriesRequest) (*pb.CountriesResponse, error) {
	return c.daemonClient.Countries(getContext(), req)
}

// Disconnect calls the Disconnect RPC and terminates the current VPN session.
func (c Client) Disconnect() error {
	_, err := c.daemonClient.Disconnect(getContext(), &pb.DisconnectRequest{Id: 0})
	return err
}

// FrontendCountries calls the FrontendCountries RPC and returns a CountriesResponse.
func (c Client) FrontendCountries(req *pb.CountriesRequest) (*pb.FrontendCountriesResponse, error) {
	return c.daemonClient.FrontendCountries(getContext(), req)
}

// Groups calls the Groups RPC and returns a GroupsResponse.
func (c Client) Groups(req *pb.GroupsRequest) (*pb.GroupsResponse, error) {
	return c.daemonClient.Groups(getContext(), req)
}

// IsLoggedIn calls the IsLoggedIn RPC and returns a IsLoggedInResponse.
func (c Client) IsLoggedIn() (*pb.IsLoggedInResponse, error) {
	return c.daemonClient.IsLoggedIn(getContext(), &emptypb.Empty{})
}

// Login calls the Login RPC.
func (c Client) Login(req *pb.LoginRequest) error {
	_, err := c.daemonClient.Login(getContext(), req)
	return err
}

// LoginOAuth2 calls the LoginOAuth2 RPC.
func (c Client) LoginOAuth2() (*pb.LoginOAuth2Response, error) {
	return c.daemonClient.LoginOAuth2(getContext(), &emptypb.Empty{})
}

// Logout calls the Logout RPC.
func (c Client) Logout() error {
	_, err := c.daemonClient.Logout(getContext(), &emptypb.Empty{})
	return err
}

// Plans calls the Plans RPC.
func (c Client) Plans() (*pb.PlansResponse, error) {
	return c.daemonClient.Plans(getContext(), &emptypb.Empty{})
}

// Ping checks that the daemon is alive via the Ping RPC.
// If the value returned is nil, the ping succeeded.
func (c Client) Ping() error {
	_, err := c.daemonClient.Ping(getContext(), &emptypb.Empty{})
	return err
}

// RateConnection calls the RateConnection RPC.
func (c Client) RateConnection(req *pb.RateConnectionRequest) error {
	_, err := c.daemonClient.RateConnection(getContext(), req)
	return err
}

// SetAutoConnect calls the SetAutoConnect RPC.
// TODO: fix
func (c Client) SetAutoConnect(req *pb.SetAutoConnectRequest) (*pb.Payload, error) {
	return c.daemonClient.SetAutoConnect(getContext(), req)
}

// SetWhitelist calls the SetWhitelist RPC.
func (c Client) SetWhitelist(req *pb.SetWhitelistRequest) error {
	_, err := c.daemonClient.SetWhitelist(getContext(), req)
	return err
}

// SetCyberSec calls the SetCyberSec RPC.
func (c Client) SetCyberSec(enabled bool) error {
	_, err := c.daemonClient.SetCyberSec(getContext(), &pb.SetCyberSecRequest{
		CyberSec: enabled,
	})
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
func (c Client) SetProtocol(protocol Protocol) error {
	_, err := c.daemonClient.SetProtocol(getContext(), &pb.SetProtocolRequest{
		Protocol: pb.ProtocolEnum(protocol),
	})
	return err
}

// SetTechnology calls the SetTechnology RPC.
func (c Client) SetTechnology(technology Technology) error {
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
