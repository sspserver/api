package client

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/context/ctxlogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	archivariuspb "github.com/geniusrabbit/archivarius/internal/server/grpc"
)

// APIClient is a client for the archivarius service
type APIClient struct {
	archivariuspb.ArchivariusServiceClient
	conn *grpc.ClientConn
}

// ConnectAPI creates a new API client connection
func ConnectAPI(ctx context.Context, connect string, opts ...grpc.DialOption) (*APIClient, error) {
	if len(opts) == 0 {
		ctxlogger.Get(ctx).Warn("Current connection is insecure")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	billingCon, err := grpc.NewClient(connect, opts...)
	if err != nil {
		return nil, err
	}
	return &APIClient{
		ArchivariusServiceClient: archivariuspb.NewArchivariusServiceClient(billingCon),
		conn:                     billingCon,
	}, nil
}

// Conn returns the API client connection to GRPC client connection
func (c *APIClient) Conn() *grpc.ClientConn {
	return c.conn
}

// Close closes the API client connection
func (c *APIClient) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}
