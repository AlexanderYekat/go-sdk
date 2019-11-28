// Code generated by sdkgen. DO NOT EDIT.

//nolint
package clickhouse

import (
	"context"

	"google.golang.org/grpc"

	clickhouse "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ClusterServiceClient is a clickhouse.ClusterServiceClient with
// lazy GRPC connection initialization.
type ClusterServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddHosts implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) AddHosts(ctx context.Context, in *clickhouse.AddClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).AddHosts(ctx, in, opts...)
}

// AddShard implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) AddShard(ctx context.Context, in *clickhouse.AddClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).AddShard(ctx, in, opts...)
}

// AddZookeeper implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) AddZookeeper(ctx context.Context, in *clickhouse.AddClusterZookeeperRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).AddZookeeper(ctx, in, opts...)
}

// Backup implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Backup(ctx context.Context, in *clickhouse.BackupClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Backup(ctx, in, opts...)
}

// Create implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Create(ctx context.Context, in *clickhouse.CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Create(ctx, in, opts...)
}

// CreateExternalDictionary implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) CreateExternalDictionary(ctx context.Context, in *clickhouse.CreateClusterExternalDictionaryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).CreateExternalDictionary(ctx, in, opts...)
}

// Delete implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Delete(ctx context.Context, in *clickhouse.DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteExternalDictionary implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) DeleteExternalDictionary(ctx context.Context, in *clickhouse.DeleteClusterExternalDictionaryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).DeleteExternalDictionary(ctx, in, opts...)
}

// DeleteHosts implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) DeleteHosts(ctx context.Context, in *clickhouse.DeleteClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).DeleteHosts(ctx, in, opts...)
}

// DeleteShard implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) DeleteShard(ctx context.Context, in *clickhouse.DeleteClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).DeleteShard(ctx, in, opts...)
}

// Get implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Get(ctx context.Context, in *clickhouse.GetClusterRequest, opts ...grpc.CallOption) (*clickhouse.Cluster, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Get(ctx, in, opts...)
}

// GetShard implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) GetShard(ctx context.Context, in *clickhouse.GetClusterShardRequest, opts ...grpc.CallOption) (*clickhouse.Shard, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).GetShard(ctx, in, opts...)
}

// List implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) List(ctx context.Context, in *clickhouse.ListClustersRequest, opts ...grpc.CallOption) (*clickhouse.ListClustersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).List(ctx, in, opts...)
}

type ClusterIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ClusterServiceClient
	request *clickhouse.ListClustersRequest

	items []*clickhouse.Cluster
}

func (c *ClusterServiceClient) ClusterIterator(ctx context.Context, folderId string, opts ...grpc.CallOption) *ClusterIterator {
	return &ClusterIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListClustersRequest{
			FolderId: folderId,
			PageSize: 1000,
		},
	}
}

func (it *ClusterIterator) Next() bool {
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

	it.items = response.Clusters
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterIterator) Value() *clickhouse.Cluster {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterIterator) Error() error {
	return it.err
}

// ListBackups implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) ListBackups(ctx context.Context, in *clickhouse.ListClusterBackupsRequest, opts ...grpc.CallOption) (*clickhouse.ListClusterBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).ListBackups(ctx, in, opts...)
}

type ClusterBackupsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ClusterServiceClient
	request *clickhouse.ListClusterBackupsRequest

	items []*clickhouse.Backup
}

func (c *ClusterServiceClient) ClusterBackupsIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *ClusterBackupsIterator {
	return &ClusterBackupsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListClusterBackupsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *ClusterBackupsIterator) Next() bool {
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

	response, err := it.client.ListBackups(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Backups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterBackupsIterator) Value() *clickhouse.Backup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterBackupsIterator) Error() error {
	return it.err
}

// ListHosts implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) ListHosts(ctx context.Context, in *clickhouse.ListClusterHostsRequest, opts ...grpc.CallOption) (*clickhouse.ListClusterHostsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).ListHosts(ctx, in, opts...)
}

type ClusterHostsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ClusterServiceClient
	request *clickhouse.ListClusterHostsRequest

	items []*clickhouse.Host
}

func (c *ClusterServiceClient) ClusterHostsIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *ClusterHostsIterator {
	return &ClusterHostsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListClusterHostsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *ClusterHostsIterator) Next() bool {
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

	response, err := it.client.ListHosts(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Hosts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterHostsIterator) Value() *clickhouse.Host {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterHostsIterator) Error() error {
	return it.err
}

// ListLogs implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) ListLogs(ctx context.Context, in *clickhouse.ListClusterLogsRequest, opts ...grpc.CallOption) (*clickhouse.ListClusterLogsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).ListLogs(ctx, in, opts...)
}

type ClusterLogsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ClusterServiceClient
	request *clickhouse.ListClusterLogsRequest

	items []*clickhouse.LogRecord
}

func (c *ClusterServiceClient) ClusterLogsIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *ClusterLogsIterator {
	return &ClusterLogsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListClusterLogsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *ClusterLogsIterator) Next() bool {
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

	response, err := it.client.ListLogs(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Logs
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterLogsIterator) Value() *clickhouse.LogRecord {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterLogsIterator) Error() error {
	return it.err
}

// ListOperations implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) ListOperations(ctx context.Context, in *clickhouse.ListClusterOperationsRequest, opts ...grpc.CallOption) (*clickhouse.ListClusterOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).ListOperations(ctx, in, opts...)
}

type ClusterOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ClusterServiceClient
	request *clickhouse.ListClusterOperationsRequest

	items []*operation.Operation
}

func (c *ClusterServiceClient) ClusterOperationsIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *ClusterOperationsIterator {
	return &ClusterOperationsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListClusterOperationsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *ClusterOperationsIterator) Next() bool {
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

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterOperationsIterator) Error() error {
	return it.err
}

// ListShards implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) ListShards(ctx context.Context, in *clickhouse.ListClusterShardsRequest, opts ...grpc.CallOption) (*clickhouse.ListClusterShardsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).ListShards(ctx, in, opts...)
}

type ClusterShardsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ClusterServiceClient
	request *clickhouse.ListClusterShardsRequest

	items []*clickhouse.Shard
}

func (c *ClusterServiceClient) ClusterShardsIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *ClusterShardsIterator {
	return &ClusterShardsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListClusterShardsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *ClusterShardsIterator) Next() bool {
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

	response, err := it.client.ListShards(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Shards
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterShardsIterator) Value() *clickhouse.Shard {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterShardsIterator) Error() error {
	return it.err
}

// Move implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Move(ctx context.Context, in *clickhouse.MoveClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Move(ctx, in, opts...)
}

// Restore implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Restore(ctx context.Context, in *clickhouse.RestoreClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Restore(ctx, in, opts...)
}

// Start implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Start(ctx context.Context, in *clickhouse.StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Start(ctx, in, opts...)
}

// Stop implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Stop(ctx context.Context, in *clickhouse.StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) Update(ctx context.Context, in *clickhouse.UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateShard implements clickhouse.ClusterServiceClient
func (c *ClusterServiceClient) UpdateShard(ctx context.Context, in *clickhouse.UpdateClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewClusterServiceClient(conn).UpdateShard(ctx, in, opts...)
}
