// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.LookupSharedImageVersion(ctx, &compute.LookupSharedImageVersionArgs{
// 			GalleryName:       "my-image-gallery",
// 			ImageName:         "my-image",
// 			Name:              "1.0.0",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSharedImageVersion(ctx *pulumi.Context, args *LookupSharedImageVersionArgs, opts ...pulumi.InvokeOption) (*LookupSharedImageVersionResult, error) {
	var rv LookupSharedImageVersionResult
	err := ctx.Invoke("azure:compute/getSharedImageVersion:getSharedImageVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedImageVersion.
type LookupSharedImageVersionArgs struct {
	// The name of the Shared Image in which the Shared Image exists.
	GalleryName string `pulumi:"galleryName"`
	// The name of the Shared Image in which this Version exists.
	ImageName string `pulumi:"imageName"`
	// The name of the Image Version.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Shared Image Gallery exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSharedImageVersion.
type LookupSharedImageVersionResult struct {
	// Is this Image Version excluded from the `latest` filter?
	ExcludeFromLatest bool   `pulumi:"excludeFromLatest"`
	GalleryName       string `pulumi:"galleryName"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ImageName string `pulumi:"imageName"`
	// The supported Azure location where the Shared Image Gallery exists.
	Location string `pulumi:"location"`
	// The ID of the Managed Image which was the source of this Shared Image Version.
	ManagedImageId string `pulumi:"managedImageId"`
	// The Azure Region in which this Image Version exists.
	Name string `pulumi:"name"`
	// The size of the OS disk snapshot (in Gigabytes) which was the source of this Shared Image Version.
	OsDiskImageSizeGb int `pulumi:"osDiskImageSizeGb"`
	// The ID of the OS disk snapshot which was the source of this Shared Image Version.
	OsDiskSnapshotId  string `pulumi:"osDiskSnapshotId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
	// One or more `targetRegion` blocks as documented below.
	TargetRegions []GetSharedImageVersionTargetRegion `pulumi:"targetRegions"`
}
