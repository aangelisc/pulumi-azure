// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about a set of existing Public IP Addresses.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/public_ips.html.markdown.
func GetPublicIPs(ctx *pulumi.Context, args *GetPublicIPsArgs, opts ...pulumi.InvokeOption) (*GetPublicIPsResult, error) {
	var rv GetPublicIPsResult
	err := ctx.Invoke("azure:network/getPublicIPs:getPublicIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicIPs.
type GetPublicIPsArgs struct {
	// The Allocation Type for the Public IP Address. Possible values include `Static` or `Dynamic`.
	AllocationType *string `pulumi:"allocationType"`
	// Filter to include IP Addresses which are attached to a device, such as a VM/LB (`true`) or unattached (`false`).
	Attached *bool `pulumi:"attached"`
	// A prefix match used for the IP Addresses `name` field, case sensitive.
	NamePrefix *string `pulumi:"namePrefix"`
	// Specifies the name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getPublicIPs.
type GetPublicIPsResult struct {
	AllocationType *string `pulumi:"allocationType"`
	Attached *bool `pulumi:"attached"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	NamePrefix *string `pulumi:"namePrefix"`
	// A List of `publicIps` blocks as defined below filtered by the criteria above.
	PublicIps []GetPublicIPsPublicIp `pulumi:"publicIps"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

