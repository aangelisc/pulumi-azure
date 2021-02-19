// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Kubernetes Cluster Node Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/containerservice"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := containerservice.GetClusterNodePool(ctx, &containerservice.GetClusterNodePoolArgs{
// 			Name:                  "existing",
// 			KubernetesClusterName: "existing-cluster",
// 			ResourceGroupName:     "existing-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func GetClusterNodePool(ctx *pulumi.Context, args *GetClusterNodePoolArgs, opts ...pulumi.InvokeOption) (*GetClusterNodePoolResult, error) {
	var rv GetClusterNodePoolResult
	err := ctx.Invoke("azure:containerservice/getClusterNodePool:getClusterNodePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterNodePool.
type GetClusterNodePoolArgs struct {
	// The Name of the Kubernetes Cluster where this Node Pool is located.
	KubernetesClusterName string `pulumi:"kubernetesClusterName"`
	// The name of this Kubernetes Cluster Node Pool.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Kubernetes Cluster exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getClusterNodePool.
type GetClusterNodePoolResult struct {
	// A list of Availability Zones in which the Nodes in this Node Pool exists.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Does this Node Pool have Auto-Scaling enabled?
	EnableAutoScaling bool `pulumi:"enableAutoScaling"`
	// Do nodes in this Node Pool have a Public IP Address?
	EnableNodePublicIp bool `pulumi:"enableNodePublicIp"`
	// The eviction policy used for Virtual Machines in the Virtual Machine Scale Set, when `priority` is set to `Spot`.
	EvictionPolicy string `pulumi:"evictionPolicy"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string `pulumi:"id"`
	KubernetesClusterName string `pulumi:"kubernetesClusterName"`
	// The maximum number of Nodes allowed when auto-scaling is enabled.
	MaxCount int `pulumi:"maxCount"`
	// The maximum number of Pods allowed on each Node in this Node Pool.
	MaxPods int `pulumi:"maxPods"`
	// The minimum number of Nodes allowed when auto-scaling is enabled.
	MinCount int `pulumi:"minCount"`
	// The Mode for this Node Pool, specifying how these Nodes should be used (for either System or User resources).
	Mode string `pulumi:"mode"`
	Name string `pulumi:"name"`
	// The current number of Nodes in the Node Pool.
	NodeCount int `pulumi:"nodeCount"`
	// A map of Kubernetes Labels applied to each Node in this Node Pool.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// A map of Kubernetes Taints applied to each Node in this Node Pool.
	NodeTaints []string `pulumi:"nodeTaints"`
	// The version of Kubernetes configured on each Node in this Node Pool.
	OrchestratorVersion string `pulumi:"orchestratorVersion"`
	// The size of the OS Disk on each Node in this Node Pool.
	OsDiskSizeGb int `pulumi:"osDiskSizeGb"`
	// The type of the OS Disk on each Node in this Node Pool.
	OsDiskType string `pulumi:"osDiskType"`
	// The operating system used on each Node in this Node Pool.
	OsType string `pulumi:"osType"`
	// The priority of the Virtual Machines in the Virtual Machine Scale Set backing this Node Pool.
	Priority string `pulumi:"priority"`
	// The ID of the Proximity Placement Group where the Virtual Machine Scale Set backing this Node Pool will be placed.
	ProximityPlacementGroupId string `pulumi:"proximityPlacementGroupId"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	// The maximum price being paid for Virtual Machines in this Scale Set. `-1` means the current on-demand price for a Virtual Machine.
	SpotMaxPrice float64 `pulumi:"spotMaxPrice"`
	// A mapping of tags assigned to the Kubernetes Cluster Node Pool.
	Tags map[string]string `pulumi:"tags"`
	// A `upgradeSettings` block as documented below.
	UpgradeSettings []GetClusterNodePoolUpgradeSetting `pulumi:"upgradeSettings"`
	// The size of the Virtual Machines used in the Virtual Machine Scale Set backing this Node Pool.
	VmSize string `pulumi:"vmSize"`
	// The ID of the Subnet in which this Node Pool exists.
	VnetSubnetId string `pulumi:"vnetSubnetId"`
}
