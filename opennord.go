package opennord

import (
	"github.com/adamdb5/opennord/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	address = "unix:///run/nordvpn/nordvpnd.sock" // Unix domain socket for nordvpnd
)

// getConnection creates a new gRPC connection.
func getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to socket.")
	}
	return conn
}

// getDaemonClient creates a new protobuffer client.
func getDaemonClient(conn *grpc.ClientConn) pb.DaemonClient {
	return pb.NewDaemonClient(conn)
}

// NewOpenNordClient creates a client for interacting with the NordVPN daemon.
func NewOpenNordClient() Client {
	conn := getConnection()
	client := getDaemonClient(conn)
	return Client{
		grpcConnection: conn,
		daemonClient:   client,
	}
}
