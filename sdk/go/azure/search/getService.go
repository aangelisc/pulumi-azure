// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Search Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/search"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure:search/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// The Name of the Search Service.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Search Service exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identity` block as defined below.
	Identities []GetServiceIdentity `pulumi:"identities"`
	// The name of this Query Key.
	Name string `pulumi:"name"`
	// The number of partitions which have been created.
	PartitionCount int `pulumi:"partitionCount"`
	// The Primary Key used for Search Service Administration.
	PrimaryKey string `pulumi:"primaryKey"`
	// Whether or not public network access is enabled for this resource.
	PublicNetworkAccessEnabled bool `pulumi:"publicNetworkAccessEnabled"`
	// A `queryKeys` block as defined below.
	QueryKeys []GetServiceQueryKey `pulumi:"queryKeys"`
	// The number of replica's which have been created.
	ReplicaCount      int    `pulumi:"replicaCount"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary Key used for Search Service Administration.
	SecondaryKey string `pulumi:"secondaryKey"`
}
