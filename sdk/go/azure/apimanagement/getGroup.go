// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing API Management Group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/api_management_group.html.markdown.
func LookupGroup(ctx *pulumi.Context, args *GetGroupArgs) (*GetGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:apimanagement/getGroup:getGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetGroupResult{
		ApiManagementName: outputs["apiManagementName"],
		Description: outputs["description"],
		DisplayName: outputs["displayName"],
		ExternalId: outputs["externalId"],
		Name: outputs["name"],
		ResourceGroupName: outputs["resourceGroupName"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getGroup.
type GetGroupArgs struct {
	// The Name of the API Management Service in which this Group exists.
	ApiManagementName interface{}
	// The Name of the API Management Group.
	Name interface{}
	// The Name of the Resource Group in which the API Management Service exists.
	ResourceGroupName interface{}
}

// A collection of values returned by getGroup.
type GetGroupResult struct {
	ApiManagementName interface{}
	// The description of this API Management Group.
	Description interface{}
	// The display name of this API Management Group.
	DisplayName interface{}
	// The identifier of the external Group.
	ExternalId interface{}
	Name interface{}
	ResourceGroupName interface{}
	// The type of this API Management Group, such as `custom` or `external`.
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
