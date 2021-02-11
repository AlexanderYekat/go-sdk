// Code generated by sdkgen. DO NOT EDIT.

package elasticsearch

import (
	"context"

	"google.golang.org/grpc"
)

// ElasticSearch provides access to "elasticsearch" component of Yandex.Cloud
type ElasticSearch struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewElasticSearch creates instance of ElasticSearch
func NewElasticSearch(g func(ctx context.Context) (*grpc.ClientConn, error)) *ElasticSearch {
	return &ElasticSearch{g}
}

// Cluster gets ClusterService client
func (e *ElasticSearch) Cluster() *ClusterServiceClient {
	return &ClusterServiceClient{getConn: e.getConn}
}

// ResourcePreset gets ResourcePresetService client
func (e *ElasticSearch) ResourcePreset() *ResourcePresetServiceClient {
	return &ResourcePresetServiceClient{getConn: e.getConn}
}

// User gets UserService client
func (e *ElasticSearch) User() *UserServiceClient {
	return &UserServiceClient{getConn: e.getConn}
}