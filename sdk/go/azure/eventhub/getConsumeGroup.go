// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Event Hubs Consumer Group within an Event Hub.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/eventhub_consumer_group.html.markdown.
func GetConsumeGroup(ctx *pulumi.Context, args *GetConsumeGroupArgs, opts ...pulumi.InvokeOption) (*GetConsumeGroupResult, error) {
	var rv GetConsumeGroupResult
	err := ctx.Invoke("azure:eventhub/getConsumeGroup:getConsumeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConsumeGroup.
type GetConsumeGroupArgs struct {
	// Specifies the name of the EventHub.
	EventhubName string `pulumi:"eventhubName"`
	// Specifies the name of the EventHub Consumer Group resource.
	Name string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getConsumeGroup.
type GetConsumeGroupResult struct {
	EventhubName string `pulumi:"eventhubName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	NamespaceName string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the user metadata.
	UserMetadata string `pulumi:"userMetadata"`
}

