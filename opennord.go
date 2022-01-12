package opennord

import (
	"errors"
	"github.com/adamdb5/opennord/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

// getConnection creates a new gRPC connection.
func getConnection() (*grpc.ClientConn, error) {
	if _, err := os.Stat(SocketPath); errors.Is(err, os.ErrNotExist) {
		return &grpc.ClientConn{}, err
	}

	ctx, _ := context.WithTimeout(context.Background(), RequestTimeout)
	return grpc.DialContext(
		ctx,
		SocketAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
}

// getDaemonClient creates a new protobuffer client.
func getDaemonClient(conn *grpc.ClientConn) pb.DaemonClient {
	return pb.NewDaemonClient(conn)
}

// NewOpenNordClient creates a client for interacting with the NordVPN daemon.
func NewOpenNordClient() (*Client, error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	client := getDaemonClient(conn)
	return &Client{
		grpcConnection: conn,
		daemonClient:   client,
	}, nil
}
