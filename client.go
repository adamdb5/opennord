package opennord

import (
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

// Status calls the Status RPC and returns a StatusResponse.
func (c Client) Status() (messages.StatusResponse, error) {
	ctx := context.Background()
	r, err := c.daemonClient.Status(ctx, &emptypb.Empty{})
	if err != nil {
		return messages.StatusResponse{}, err
	}
	return messages.FormatStatusResponse(r), nil
}

// AccountInfo calls the AccountInfo RPC and returns an AccountResponse.
func (c Client) AccountInfo() (messages.AccountResponse, error) {
	ctx := context.Background()
	r, err := c.daemonClient.AccountInfo(ctx, messages.AccountRequest{}.ToProtoBuffer())
	if err != nil {
		return messages.AccountResponse{}, err
	}
	return messages.FormatAccountResponse(r), nil
}

func (c Client) Cities(req messages.CitiesRequest) (messages.CitiesResponse, error) {
	ctx := context.Background()
	r, err := c.daemonClient.Cities(ctx, req.ToProtoBuffer())
	if err != nil {
		return messages.CitiesResponse{}, err
	}
	return messages.FormatCitiesResponse(r), nil
}
