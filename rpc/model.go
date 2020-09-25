package data

import (
	"log"

	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

type GrpcConnection struct {
	Conn *grpc.ClientConn
}

var ClientConn *GrpcConnection

// setupConnection set up a connection to the server.
func setupConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can not connect to %s: %s\n", addr, err.Error())
	}
	return conn
}

func (c *GrpcConnection) init() {
	c.Conn = setupConnection(viper.GetString("data_service_url"))
}

func (c *GrpcConnection) close() {
	c.Conn.Close()
}

func (c *GrpcConnection) newClient() DataProviderClient {
	return NewDataProviderClient(ClientConn.Conn)
}

// getState ... 获取连接状态：IDLE/CONNECTING/READY/TRANSIENT_FAILURE/SHUTDOWN
// 正常情况下出于 READY，
// TRANSIENT_FAILURE 时会自动重连
func (c *GrpcConnection) getState() string {
	return c.Conn.GetState().String()
}

// GetClient ... 获取 grpc 客户端
func GetClient() DataProviderClient {
	log.Printf("Grpc client connection state: %s\n", ClientConn.getState())
	return ClientConn.newClient()
}

func CloseConnection() {
	ClientConn.close()
}

func InitConnection() {
	ClientConn = new(GrpcConnection)
	ClientConn.init()
	log.Printf("Grpc client connection state: %s\n", ClientConn.getState())
	log.Println("Grpc connection has set up.")
}
