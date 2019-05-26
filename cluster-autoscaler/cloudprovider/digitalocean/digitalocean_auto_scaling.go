package digitalocean

import (
	"context"
	"github.com/digitalocean/godo"
	"k8s.io/klog"
)

func main(clusterID, poolID string) {

	ctx := context.TODO()
	pool, _, err := godo.KubernetesService.GetNodePool(nil, ctx, clusterID, poolID)
	newCount := pool.Count + 1
	poolUpdateReq := &godo.KubernetesNodePoolUpdateRequest{
		Name:  pool.Name,
		Count: newCount,
	}
	pool, _, err = godo.KubernetesService.UpdateNodePool(nil, ctx, clusterID, poolID, poolUpdateReq)
	if err != nil {
		klog.Error(err)
	}

}
