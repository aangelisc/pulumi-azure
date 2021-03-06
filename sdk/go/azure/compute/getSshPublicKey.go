// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing SSH Public Key.
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
// 		example, err := compute.LookupSshPublicKey(ctx, &compute.LookupSshPublicKeyArgs{
// 			Name:              "existing",
// 			ResourceGroupName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupSshPublicKey(ctx *pulumi.Context, args *LookupSshPublicKeyArgs, opts ...pulumi.InvokeOption) (*LookupSshPublicKeyResult, error) {
	var rv LookupSshPublicKeyResult
	err := ctx.Invoke("azure:compute/getSshPublicKey:getSshPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSshPublicKey.
type LookupSshPublicKeyArgs struct {
	// The name of this SSH Public Key.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the SSH Public Key exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the SSH Public Key.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getSshPublicKey.
type LookupSshPublicKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The SSH public key used to authenticate to a virtual machine through ssh.
	PublicKey         string            `pulumi:"publicKey"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}
