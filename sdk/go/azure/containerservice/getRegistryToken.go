// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Container Registry token.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := containerservice.LookupRegistryToken(ctx, &containerservice.LookupRegistryTokenArgs{
// 			Name:                  "exampletoken",
// 			ResourceGroupName:     "example-resource-group",
// 			ContainerRegistryName: "example-registry",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("scopeMapId", example.ScopeMapId)
// 		return nil
// 	})
// }
// ```
func LookupRegistryToken(ctx *pulumi.Context, args *LookupRegistryTokenArgs, opts ...pulumi.InvokeOption) (*LookupRegistryTokenResult, error) {
	var rv LookupRegistryTokenResult
	err := ctx.Invoke("azure:containerservice/getRegistryToken:getRegistryToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryToken.
type LookupRegistryTokenArgs struct {
	// The Name of the Container Registry where the token exists.
	ContainerRegistryName string `pulumi:"containerRegistryName"`
	// The name of the Container Registry token.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where this Container Registry token exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getRegistryToken.
type LookupRegistryTokenResult struct {
	ContainerRegistryName string `pulumi:"containerRegistryName"`
	// Whether this Token is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Scope Map ID used by the token.
	ScopeMapId string `pulumi:"scopeMapId"`
}
