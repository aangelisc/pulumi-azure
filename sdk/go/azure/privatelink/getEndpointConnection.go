// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access the connection status information about an existing Private Endpoint Connection.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatelink"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := privatelink.GetEndpointConnection(ctx, &privatelink.GetEndpointConnectionArgs{
// 			Name:              "example-private-endpoint",
// 			ResourceGroupName: "example-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("privateEndpointStatus", example.PrivateServiceConnections[0].Status)
// 		return nil
// 	})
// }
// ```
func GetEndpointConnection(ctx *pulumi.Context, args *GetEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*GetEndpointConnectionResult, error) {
	var rv GetEndpointConnectionResult
	err := ctx.Invoke("azure:privatelink/getEndpointConnection:getEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpointConnection.
type GetEndpointConnectionArgs struct {
	// Specifies the Name of the private endpoint.
	Name string `pulumi:"name"`
	// Specifies the Name of the Resource Group within which the private endpoint exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getEndpointConnection.
type GetEndpointConnectionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The supported Azure location where the resource exists.
	Location string `pulumi:"location"`
	// The name of the private endpoint.
	Name                      string                                          `pulumi:"name"`
	PrivateServiceConnections []GetEndpointConnectionPrivateServiceConnection `pulumi:"privateServiceConnections"`
	ResourceGroupName         string                                          `pulumi:"resourceGroupName"`
}
