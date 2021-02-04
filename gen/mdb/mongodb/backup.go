// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mongodb

import (
	"context"

	"google.golang.org/grpc"

	mongodb "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// BackupServiceClient is a mongodb.BackupServiceClient with
// lazy GRPC connection initialization.
type BackupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Delete implements mongodb.BackupServiceClient
func (c *BackupServiceClient) Delete(ctx context.Context, in *mongodb.DeleteBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewBackupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements mongodb.BackupServiceClient
func (c *BackupServiceClient) Get(ctx context.Context, in *mongodb.GetBackupRequest, opts ...grpc.CallOption) (*mongodb.Backup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewBackupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements mongodb.BackupServiceClient
func (c *BackupServiceClient) List(ctx context.Context, in *mongodb.ListBackupsRequest, opts ...grpc.CallOption) (*mongodb.ListBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewBackupServiceClient(conn).List(ctx, in, opts...)
}

type BackupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *BackupServiceClient
	request *mongodb.ListBackupsRequest

	items []*mongodb.Backup
}

func (c *BackupServiceClient) BackupIterator(ctx context.Context, folderId string, opts ...grpc.CallOption) *BackupIterator {
	return &BackupIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &mongodb.ListBackupsRequest{
			FolderId: folderId,
			PageSize: 1000,
		},
	}
}

func (it *BackupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Backups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BackupIterator) Value() *mongodb.Backup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BackupIterator) Error() error {
	return it.err
}
