// Code generated by sdkgen. DO NOT EDIT.

//nolint
package iam

import (
	"context"

	"google.golang.org/grpc"

	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

//revive:disable

// UserAccountServiceClient is a iam.UserAccountServiceClient with
// lazy GRPC connection initialization.
type UserAccountServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements iam.UserAccountServiceClient
func (c *UserAccountServiceClient) Get(ctx context.Context, in *iam.GetUserAccountRequest, opts ...grpc.CallOption) (*iam.UserAccount, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewUserAccountServiceClient(conn).Get(ctx, in, opts...)
}
