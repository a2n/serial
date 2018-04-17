package server

import (
	"context"
	"flag"
	"net"

	"github.com/a2n/serial/src"
	pb "github.com/a2n/serial/src/grpc/protos"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// SerialServer 序號服務
type SerialServer struct {
	ids *serial.IdService
}

// NewSerialServer 創建序號服務
func NewSerialServer() *SerialServer {
	return &SerialServer{
		ids: serial.NewIdService(),
	}
}

// Port GRPC 埠號
const Port string = ":57888"

// Start 啟動
func (ss *SerialServer) Start() {
	flag.Parse()
	l, e := net.Listen("tcp", Port)
	if e != nil {
		glog.Fatalf("%+v", e)
	}

	srv := grpc.NewServer()
	pb.RegisterSerialServiceeServer(srv, ss)
	glog.Infof("Starting grpc server %s", Port)
	glog.Fatalf("%+v", srv.Serve(l))
}

// Get 取得
func (ss *SerialServer) Get(ctx context.Context, ept *pb.Empty) (*pb.Response, error) {
	return &pb.Response{No: ss.ids.Increase()}, nil
}
