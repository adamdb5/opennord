package opennord

import (
	"github.com/adamdb5/opennord/messages"
	"github.com/adamdb5/opennord/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

// Client is the public stub for of communicating with the NordVPN daemon.
type Client struct {
	grpcConnection *grpc.ClientConn
	daemonClient   pb.DaemonClient
}

// Status calls the Status RPC and returns a StatusResponse.
func (c Client) Status() messages.StatusResponse {
	ctx := context.Background()
	r, err := c.daemonClient.Status(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling RPC Status: %s", err)
	}
	return messages.FormatStatusResponse(r)
}

// AccountInfo calls the AccountInfo RPC and returns an AccountResponse.
func (c Client) AccountInfo() messages.AccountResponse {
	ctx := context.Background()
	r, err := c.daemonClient.AccountInfo(ctx, messages.AccountRequest{}.ToProtoBuffer())
	if err != nil {
		log.Fatalf("Error calling RPC AccountInfo: %s", err)
	}
	return messages.FormatAccountResponse(r)
}

func (c Client) Cities(req messages.CitiesRequest) messages.CitiesResponse {
	ctx := context.Background()
	r, err := c.daemonClient.Cities(ctx, req.ToProtoBuffer())
	if err != nil {
		log.Fatalf("Error calling RPC Cities: %s", err)
	}
	return messages.FormatCitiesResponse(r)
}
