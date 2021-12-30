package opennord

import (
	"github.com/adamdb5/opennord/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	address = "unix:///run/nordvpn/nordvpnd.sock"
)

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
