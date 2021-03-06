// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing App Service Environment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := appservice.GetAppServiceEnvironment(ctx, &appservice.GetAppServiceEnvironmentArgs{
// 			Name:              "existing-ase",
// 			ResourceGroupName: "existing-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func GetAppServiceEnvironment(ctx *pulumi.Context, args *GetAppServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*GetAppServiceEnvironmentResult, error) {
	var rv GetAppServiceEnvironmentResult
	err := ctx.Invoke("azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppServiceEnvironment.
type GetAppServiceEnvironmentArgs struct {
	// The name of this App Service Environment.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the App Service Environment exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAppServiceEnvironment.
type GetAppServiceEnvironmentResult struct {
	// Zero or more `clusterSetting` blocks as defined below.
	ClusterSettings []GetAppServiceEnvironmentClusterSetting `pulumi:"clusterSettings"`
	// The number of app instances per App Service Environment Front End.
	FrontEndScaleFactor int `pulumi:"frontEndScaleFactor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP address of internal load balancer of the App Service Environment.
	InternalIpAddress string `pulumi:"internalIpAddress"`
	// The Azure Region where the App Service Environment exists.
	Location string `pulumi:"location"`
	// The name of the Cluster Setting.
	Name string `pulumi:"name"`
	// List of outbound IP addresses of the App Service Environment.
	OutboundIpAddresses []string `pulumi:"outboundIpAddresses"`
	// The Pricing Tier (Isolated SKU) of the App Service Environment.
	PricingTier       string `pulumi:"pricingTier"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// IP address of service endpoint of the App Service Environment.
	ServiceIpAddress string `pulumi:"serviceIpAddress"`
	// A mapping of tags assigned to the App Service Environment.
	Tags map[string]string `pulumi:"tags"`
}
