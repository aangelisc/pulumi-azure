// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Dev Test Lab.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/devtest"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := devtest.LookupLab(ctx, &devtest.LookupLabArgs{
// 			Name:              "example-lab",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("uniqueIdentifier", example.UniqueIdentifier)
// 		return nil
// 	})
// }
// ```
func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure:devtest/getLab:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLab.
type LookupLabArgs struct {
	// The name of the Dev Test Lab.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the Dev Test Lab exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getLab.
type LookupLabResult struct {
	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId string `pulumi:"artifactsStorageAccountId"`
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId string `pulumi:"defaultPremiumStorageAccountId"`
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId string `pulumi:"defaultStorageAccountId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The Azure location where the Dev Test Lab exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId string `pulumi:"premiumDataDiskStorageAccountId"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab.
	StorageType string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
}
