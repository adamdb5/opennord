package opennord

import (
	"github.com/adamdb5/opennord/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

const (
	address = "unix:///run/nordvpn/nordvpnd.sock"
)

type Client struct {
	grpcConnection *grpc.ClientConn
	daemonClient   pb.DaemonClient
}

func (c Client) GetStatus() *pb.StatusResponse {
	ctx := context.Background()
	r, err := c.daemonClient.Status(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling RPC Status: %s", err)
	}
	return r
}

func getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to socket.")
	}
	return conn
}

func getDaemonClient(conn *grpc.ClientConn) pb.DaemonClient {
	return pb.NewDaemonClient(conn)
}

func NewOpenNordClient() Client {
	conn := getConnection()
	client := getDaemonClient(conn)
	return Client{
		grpcConnection: conn,
		daemonClient:   client,
	}
}

func getStatus(client pb.DaemonClient) *pb.StatusResponse {
	ctx := context.Background()
	r, err := client.Status(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling RPC Status: %s", err)
	}
	return r
}
