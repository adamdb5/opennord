package opennord

import (
	"github.com/adamdb5/opennord/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "unix:///run/nordvpn/nordvpnd.sock" // Unix domain socket for nordvpnd
)

// getConnection creates a new gRPC connection.
func getConnection() (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
}

// getDaemonClient creates a new protobuffer client.
func getDaemonClient(conn *grpc.ClientConn) pb.DaemonClient {
	return pb.NewDaemonClient(conn)
}

// NewOpenNordClient creates a client for interacting with the NordVPN daemon.
func NewOpenNordClient() (Client, error) {
	conn, err := getConnection()
	if err != nil {
		return Client{}, err
	}
	client := getDaemonClient(conn)
	return Client{
		grpcConnection: conn,
		daemonClient:   client,
	}, nil
}
