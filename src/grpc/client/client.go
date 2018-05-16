package client

import (
	"context"
	"time"

	pb "github.com/a2n/serial/src/grpc/protos"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client 客戶端
type Client struct {
	conn *grpc.ClientConn
	c    pb.SerialServiceeClient
}

// NewClient 創建客戶端
func NewClient() *Client {
	return &Client{}
}

// Dial 撥號
func (c *Client) Dial(a string) error {
	if len(a) == 0 {
		return errors.New("empty address")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var e error
	c.conn, e = grpc.DialContext(ctx, a, grpc.WithInsecure())
	if e != nil {
		return errors.Wrap(e, "")
	}
	c.c = pb.NewSerialServiceeClient(c.conn)
	return nil
}

// Close 關閉連線
func (c *Client) Close() error {
	e := c.Close()
	if e != nil {
		return errors.Wrap(e, "")
	}
	return nil
}

// Get 取得序號
func (c *Client) Get() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	resp, e := c.c.Get(ctx, req)
	if e != nil {
		return 0, errors.Wrap(e, "")
	}
	return resp.GetNo(), nil
}
