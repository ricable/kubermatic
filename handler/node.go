package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/kubermatic/api"
	"github.com/kubermatic/api/provider"
	"golang.org/x/net/context"
	"k8s.io/client-go/pkg/api/v1"
	metav1 "k8s.io/client-go/pkg/apis/meta/v1"
)

func nodesEndpoint(
	kps map[string]provider.KubernetesProvider,
	cps map[string]provider.CloudProvider,
) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(nodesReq)

		kp, found := kps[req.dc]
		if !found {
			return nil, NewBadRequest("unknown kubernetes datacenter %q", req.dc)
		}

		c, err := kp.Cluster(req.user, req.cluster)
		if err != nil {
			return nil, err
		}

		_, cp, err := provider.ClusterCloudProvider(cps, c)
		if err != nil {
			return nil, err
		}
		if cp == nil {
			return []*api.Node{}, nil
		}

		return cp.Nodes(ctx, c)
	}
}

func kubernetesNodesEndpoint(kps map[string]provider.KubernetesProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(nodesReq)

		kp, found := kps[req.dc]
		if !found {
			return nil, NewBadRequest("unknown kubernetes datacenter %q", req.dc)
		}

		c, err := kp.Cluster(req.user, req.cluster)
		if err != nil {
			return nil, err
		}

		client, err := c.GetClient()
		if err != nil {
			return nil, err
		}

		return client.Nodes().List(v1.ListOptions{})
	}
}

func kubernetesNodeInfoEndpoint(kps map[string]provider.KubernetesProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(nodeReq)

		kp, found := kps[req.dc]
		if !found {
			return nil, NewBadRequest("unknown kubernetes datacenter %q", req.dc)
		}

		c, err := kp.Cluster(req.user, req.cluster)
		if err != nil {
			return nil, err
		}

		client, err := c.GetClient()
		if err != nil {
			return nil, err
		}

		return client.Nodes().Get(req.uid, metav1.GetOptions{})
	}
}

func deleteNodeEndpoint(
	kps map[string]provider.KubernetesProvider,
	cps map[string]provider.CloudProvider,
) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(nodeReq)

		kp, found := kps[req.dc]
		if !found {
			return nil, NewBadRequest("unknown kubernetes datacenter %q", req.dc)
		}

		c, err := kp.Cluster(req.user, req.cluster)
		if err != nil {
			return nil, err
		}

		_, cp, err := provider.ClusterCloudProvider(cps, c)
		if err != nil {
			return nil, err
		}

		if cp == nil {
			return []*api.Node{}, nil
		}

		client, err := c.GetClient()
		if err != nil {
			return nil, err
		}

		nodes, err := cp.Nodes(ctx, c)
		if err != nil {
			return nil, err
		}

		// HACK: This is dirty. We should correlate the Kubermatic UID to the Kubernetes Name somewhere...
		for _, node := range nodes {
			if node.Metadata.UID == req.uid {
				err = client.Nodes().Delete(node.Status.Addresses["public"], &v1.DeleteOptions{})
				if err != nil {
					return nil, err
				}
			}
		}

		return nil, cp.DeleteNodes(ctx, c, []string{req.uid})
	}
}

func createNodesEndpoint(
	kps map[string]provider.KubernetesProvider,
	cps map[string]provider.CloudProvider,
) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createNodesReq)

		glog.V(1).Infoln("Get k8s provider")
		kp, found := kps[req.dc]
		if !found {
			return nil, NewBadRequest("unknown kubernetes datacenter %q", req.dc)
		}

		glog.V(1).Infoln("Get Cluster from username and clustername")
		c, err := kp.Cluster(req.user, req.cluster)
		if err != nil {
			return nil, err
		}

		glog.V(1).Infoln("Get ClusterCloudProvider")
		cpName, cp, err := provider.ClusterCloudProvider(cps, c)
		if err != nil {
			return nil, err
		}
		if cp == nil {
			return nil, NewBadRequest("cannot create nodes without cloud provider")
		}

		glog.V(1).Infoln("Get NodeCloudProviderName")
		npName, err := provider.NodeCloudProviderName(&req.Spec)
		if err != nil {
			return nil, err
		}
		if npName != cpName {
			return nil, NewBadRequest("cluster cloud provider %q and node cloud provider %q do not match",
				cpName, npName)
		}

		glog.V(1).Infoln("Call create Nodes")
		return cp.CreateNodes(ctx, c, &req.Spec, req.Instances)
	}
}

type nodesReq struct {
	clusterReq
}

func decodeNodesReq(c context.Context, r *http.Request) (interface{}, error) {
	var req nodesReq

	cr, err := decodeClusterReq(c, r)
	if err != nil {
		return nil, err
	}
	req.clusterReq = cr.(clusterReq)

	return req, nil
}

type createNodesReq struct {
	clusterReq
	Instances int          `json:"instances"`
	Spec      api.NodeSpec `json:"spec"`
}

func decodeCreateNodesReq(c context.Context, r *http.Request) (interface{}, error) {
	var req createNodesReq

	cr, err := decodeClusterReq(c, r)
	if err != nil {
		return nil, err
	}
	req.clusterReq = cr.(clusterReq)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

type nodeReq struct {
	nodesReq
	uid string
}

func decodeNodeReq(c context.Context, r *http.Request) (interface{}, error) {
	var req nodeReq

	cr, err := decodeNodesReq(c, r)
	if err != nil {
		return nil, err
	}
	req.nodesReq = cr.(nodesReq)
	req.uid = mux.Vars(r)["node"]

	return req, nil
}
