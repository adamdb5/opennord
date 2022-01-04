package opennord

import (
	"errors"
	"github.com/adamdb5/opennord/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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

// AccountInfo calls the AccountInfo RPC and returns an AccountResponse.
func (c Client) AccountInfo() (*pb.AccountResponse, error) {
	r, err := c.daemonClient.AccountInfo(getContext(), &emptypb.Empty{})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return r, err
}

// Cities calls the Cities RPC and returns a CitiesResponse.
func (c Client) Cities(country string) (*pb.CitiesResponse, error) {
	r, err := c.daemonClient.Cities(getContext(), &pb.CitiesRequest{
		Protocol:  pb.ProtocolEnum_UDP,
		Obfuscate: false,
		Country:   country,
	})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	if r.GetType() != ErrOk {
		return nil, errors.New("unknown error")
	}
	return r, err
}

// Connect calls the Connect RPC and returns a stream of ConnectResponse.
func (c Client) Connect(req *pb.ConnectRequest) (pb.Daemon_ConnectClient, error) {
	r, err := c.daemonClient.Connect(getContext(), req)
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return r, err
}

// Countries calls the Countries RPC and returns a CountriesResponse.
func (c Client) Countries() (*pb.CountriesResponse, error) {
	r, err := c.daemonClient.Countries(getContext(), &pb.CountriesRequest{
		Protocol:  pb.ProtocolEnum_UDP,
		Obfuscate: false,
	})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return r, err
}

// Disconnect calls the Disconnect RPC and terminates the current VPN session.
func (c Client) Disconnect() error {
	_, err := c.daemonClient.Disconnect(getContext(), &pb.DisconnectRequest{Id: 0})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	return err
}

// FrontendCountries calls the FrontendCountries RPC and returns a CountriesResponse.
func (c Client) FrontendCountries() (*pb.FrontendCountriesResponse, error) {
	r, err := c.daemonClient.FrontendCountries(getContext(), &pb.CountriesRequest{
		Protocol:  pb.ProtocolEnum_UDP,
		Obfuscate: false,
	})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return r, err
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
	r, err := c.daemonClient.Login(getContext(), req)
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() != ErrOk {
		return errors.New("unknown error")
	}
	return err
}

// LoginOAuth2 calls the LoginOAuth2 RPC.
func (c Client) LoginOAuth2() (*pb.LoginOAuth2Response, error) {
	r, err := c.daemonClient.LoginOAuth2(getContext(), &emptypb.Empty{})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return r, err
}

// Logout calls the Logout RPC.
func (c Client) Logout() error {
	_, err := c.daemonClient.Logout(getContext(), &emptypb.Empty{})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	return err
}

// Plans calls the Plans RPC.
func (c Client) Plans() (*pb.PlansResponse, error) {
	r, err := c.daemonClient.Plans(getContext(), &emptypb.Empty{})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	if r.GetType() != ErrOk {
		return nil, errors.New("unknown error")
	}
	return r, err
}

// Ping checks that the daemon is alive via the Ping RPC.
// If the value returned is nil, the ping succeeded.
func (c Client) Ping() error {
	r, err := c.daemonClient.Ping(getContext(), &emptypb.Empty{})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() != ErrOk {
		return errors.New("unknown error")
	}
	return err
}

// RateConnection calls the RateConnection RPC.
func (c Client) RateConnection(req *pb.RateConnectionRequest) error {
	r, err := c.daemonClient.RateConnection(getContext(), req)
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() != ErrOk {
		return errors.New("unknown error")
	}
	return err
}

// SetAutoConnect calls the SetAutoConnect RPC.
func (c Client) SetAutoConnect(req *pb.SetAutoConnectRequest) (*pb.Payload, error) {
	r, err := c.daemonClient.SetAutoConnect(getContext(), req)
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	if r.GetType() == StatusGenericError {
		return nil, errors.New("auto-connect already enabled / disabled")
	}
	if r.GetType() != StatusOk {
		return nil, errors.New("unknown error")
	}
	return r, err
}

// SetCyberSec calls the SetCyberSec RPC.
func (c Client) SetCyberSec(enabled bool) error {
	r, err := c.daemonClient.SetCyberSec(getContext(), &pb.SetCyberSecRequest{
		CyberSec: enabled,
	})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() != StatusOk {
		return errors.New("unknown error")
	}
	return err
}

// SetDefaults calls the SetDefaults RPC.
func (c Client) SetDefaults() error {
	r, err := c.daemonClient.SetDefaults(getContext(), &emptypb.Empty{})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() != StatusOk {
		return errors.New("unknown error")
	}
	return err
}

// SetDns calls the SetDns RPC.
func (c Client) SetDns(servers []string, cyberSec bool) error {
	r, err := c.daemonClient.SetDns(getContext(), &pb.SetDNSRequest{
		Dns:      servers,
		CyberSec: cyberSec,
	})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() != StatusOk {
		return errors.New("unknown error")
	}
	return err
}

// SetFirewall calls the SetFirewall RPC.
func (c Client) SetFirewall(enabled bool) error {
	r, err := c.daemonClient.SetFirewall(getContext(), &pb.SetGenericRequest{
		Enabled: enabled,
	})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() == StatusGenericError {
		return errors.New("firewall already enabled / disabled")
	}
	if r.GetType() != StatusOk {
		return errors.New("unknown error")
	}
	return err
}

// SetIpv6 calls the SetIPv6 RPC.
func (c Client) SetIpv6(enabled bool) error {
	r, err := c.daemonClient.SetIpv6(getContext(), &pb.SetGenericRequest{
		Enabled: enabled,
	})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() == StatusGenericError {
		return errors.New("IPv6 already enabled / disabled")
	}
	if r.GetType() != StatusOk {
		return errors.New("unknown error")
	}
	return err
}

// SetKillSwitch calls the SetKilLSwitch RPC.
func (c Client) SetKillSwitch(enabled bool) error {
	r, err := c.daemonClient.SetKillSwitch(getContext(), &pb.SetKillSwitchRequest{
		Enabled: enabled,
	})
	if err != nil {
		return errors.New(status.Convert(err).Message())
	}
	if r.GetType() == StatusGenericError {
		return errors.New("kill switch already enabled / disabled")
	}
	if r.GetType() != StatusOk {
		return errors.New("unknown error")
	}
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

// SetWhitelist calls the SetWhitelist RPC.
func (c Client) SetWhitelist(req *pb.SetWhitelistRequest) error {
	_, err := c.daemonClient.SetWhitelist(getContext(), req)
	return err
}

// Settings calls the Settings RPC. (but the settings RPC doesn't do anything anyway)
func (c Client) Settings() error {
	_, err := c.daemonClient.Settings(getContext(), &emptypb.Empty{})
	return err
}

// SettingsProtocols calls the SettingsProtocols RPC.
func (c Client) SettingsProtocols() (*pb.ProtocolsResponse, error) {
	r, err := c.daemonClient.SettingsProtocols(getContext(), &emptypb.Empty{})
	if err != nil {
		return &pb.ProtocolsResponse{}, err
	}
	return r, err
}

// SettingsTechnologies calls the SettingsTechnologies RPC.
func (c Client) SettingsTechnologies() (*pb.TechnologyResponse, error) {
	r, err := c.daemonClient.SettingsTechnologies(getContext(), &emptypb.Empty{})
	if err != nil {
		return &pb.TechnologyResponse{}, err
	}
	return r, err
}

// Status calls the Status RPC and returns a StatusResponse.
func (c Client) Status() (*pb.StatusResponse, error) {
	return c.daemonClient.Status(getContext(), &emptypb.Empty{})
}
