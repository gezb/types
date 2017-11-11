package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterType            = "cluster"
	ClusterFieldAPIVersion = "apiVersion"
	ClusterFieldKind       = "kind"
	ClusterFieldObjectMeta = "objectMeta"
	ClusterFieldSpec       = "spec"
	ClusterFieldStatus     = "status"
)

type Cluster struct {
	types.Resource
	APIVersion string         `json:"apiVersion,omitempty"`
	Kind       string         `json:"kind,omitempty"`
	ObjectMeta ObjectMeta     `json:"objectMeta,omitempty"`
	Spec       ClusterSpec    `json:"spec,omitempty"`
	Status     *ClusterStatus `json:"status,omitempty"`
}
type ClusterCollection struct {
	types.Collection
	Data   []Cluster `json:"data,omitempty"`
	client *ClusterClient
}

type ClusterClient struct {
	apiClient *Client
}

type ClusterOperations interface {
	List(opts *types.ListOpts) (*ClusterCollection, error)
	Create(opts *Cluster) (*Cluster, error)
	Update(existing *Cluster, updates interface{}) (*Cluster, error)
	ById(id string) (*Cluster, error)
	Delete(container *Cluster) error
}

func newClusterClient(apiClient *Client) *ClusterClient {
	return &ClusterClient{
		apiClient: apiClient,
	}
}

func (c *ClusterClient) Create(container *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoCreate(ClusterType, container, resp)
	return resp, err
}

func (c *ClusterClient) Update(existing *Cluster, updates interface{}) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoUpdate(ClusterType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterClient) List(opts *types.ListOpts) (*ClusterCollection, error) {
	resp := &ClusterCollection{}
	err := c.apiClient.Ops.DoList(ClusterType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterCollection) Next() (*ClusterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterClient) ById(id string) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoById(ClusterType, id, resp)
	return resp, err
}

func (c *ClusterClient) Delete(container *Cluster) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterType, &container.Resource)
}
