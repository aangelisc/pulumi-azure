// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing EventHub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := eventhub.LookupCluster(ctx, &eventhub.LookupClusterArgs{
// 			Name:              "search-eventhub",
// 			ResourceGroupName: "search-service",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eventhubId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure:eventhub/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The name of this EventHub Cluster.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the EventHub Cluster exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Location of the EventHub Cluster.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU name of the EventHub Cluster.
	SkuName string `pulumi:"skuName"`
}