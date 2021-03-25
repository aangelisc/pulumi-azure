// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Shared Image within a Shared Image Gallery.
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
// 		_, err := compute.LookupSharedImage(ctx, &compute.LookupSharedImageArgs{
// 			GalleryName:       "my-image-gallery",
// 			Name:              "my-image",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSharedImage(ctx *pulumi.Context, args *LookupSharedImageArgs, opts ...pulumi.InvokeOption) (*LookupSharedImageResult, error) {
	var rv LookupSharedImageResult
	err := ctx.Invoke("azure:compute/getSharedImage:getSharedImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedImage.
type LookupSharedImageArgs struct {
	// The name of the Shared Image Gallery in which the Shared Image exists.
	GalleryName string `pulumi:"galleryName"`
	// The name of the Shared Image.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Shared Image Gallery exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSharedImage.
type LookupSharedImageResult struct {
	// The description of this Shared Image.
	Description string `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula        string `pulumi:"eula"`
	GalleryName string `pulumi:"galleryName"`
	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on.
	HyperVGeneration string `pulumi:"hyperVGeneration"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identifier` block as defined below.
	Identifiers []GetSharedImageIdentifier `pulumi:"identifiers"`
	// The supported Azure location where the Shared Image Gallery exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The type of Operating System present in this Shared Image.
	OsType string `pulumi:"osType"`
	// The URI containing the Privacy Statement for this Shared Image.
	PrivacyStatementUri string `pulumi:"privacyStatementUri"`
	// The URI containing the Release Notes for this Shared Image.
	ReleaseNoteUri    string `pulumi:"releaseNoteUri"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run).
	Specialized bool `pulumi:"specialized"`
	// A mapping of tags assigned to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
}
