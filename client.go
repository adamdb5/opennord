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

// Disconnect calls the Disconnect RPC and terminates the current VPN session.
func (c Client) Disconnect() error {
	_, err := c.daemonClient.Disconnect(getContext(), &pb.DisconnectRequest{Id: 0})
	return err
}

// Groups calls the Groups RPC and returns a GroupsResponse.
func (c Client) Groups(req messages.GroupsRequest) (messages.GroupsResponse, error) {
	r, err := c.daemonClient.Groups(getContext(), req.ToProtoBuffer())
	if err != nil {
		return messages.GroupsResponse{}, err
	}
	return messages.FormatGroupsResponse(r), nil
}
