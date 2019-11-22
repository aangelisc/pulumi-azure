// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Product Policy
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_product_policy.html.markdown.
type ProductPolicy struct {
	s *pulumi.ResourceState
}

// NewProductPolicy registers a new resource with the given unique name, arguments, and options.
func NewProductPolicy(ctx *pulumi.Context,
	name string, args *ProductPolicyArgs, opts ...pulumi.ResourceOpt) (*ProductPolicy, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["productId"] = nil
		inputs["resourceGroupName"] = nil
		inputs["xmlContent"] = nil
		inputs["xmlLink"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["productId"] = args.ProductId
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["xmlContent"] = args.XmlContent
		inputs["xmlLink"] = args.XmlLink
	}
	s, err := ctx.RegisterResource("azure:apimanagement/productPolicy:ProductPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProductPolicy{s: s}, nil
}

// GetProductPolicy gets an existing ProductPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProductPolicyState, opts ...pulumi.ResourceOpt) (*ProductPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["productId"] = state.ProductId
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["xmlContent"] = state.XmlContent
		inputs["xmlLink"] = state.XmlLink
	}
	s, err := ctx.ReadResource("azure:apimanagement/productPolicy:ProductPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProductPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ProductPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ProductPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the API Management Service. Changing this forces a new resource to be created.
func (r *ProductPolicy) ApiManagementName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
func (r *ProductPolicy) ProductId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["productId"])
}

// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
func (r *ProductPolicy) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The XML Content for this Policy.
func (r *ProductPolicy) XmlContent() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["xmlContent"])
}

// A link to a Policy XML Document, which must be publicly available.
func (r *ProductPolicy) XmlLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["xmlLink"])
}

// Input properties used for looking up and filtering ProductPolicy resources.
type ProductPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId interface{}
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The XML Content for this Policy.
	XmlContent interface{}
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink interface{}
}

// The set of arguments for constructing a ProductPolicy resource.
type ProductPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId interface{}
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The XML Content for this Policy.
	XmlContent interface{}
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink interface{}
}
