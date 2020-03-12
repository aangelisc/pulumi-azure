// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Property.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_property.html.markdown.
type Property struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The display name of this API Management Property.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrOutput `pulumi:"secret"`
	// A list of tags to be applied to the API Management Property.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The value of this API Management Property.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewProperty registers a new resource with the given unique name, arguments, and options.
func NewProperty(ctx *pulumi.Context,
	name string, args *PropertyArgs, opts ...pulumi.ResourceOption) (*Property, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &PropertyArgs{}
	}
	var resource Property
	err := ctx.RegisterResource("azure:apimanagement/property:Property", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProperty gets an existing Property resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyState, opts ...pulumi.ResourceOption) (*Property, error) {
	var resource Property
	err := ctx.ReadResource("azure:apimanagement/property:Property", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Property resources.
type propertyState struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The display name of this API Management Property.
	DisplayName *string `pulumi:"displayName"`
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Property.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Property.
	Value *string `pulumi:"value"`
}

type PropertyState struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The display name of this API Management Property.
	DisplayName pulumi.StringPtrInput
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Property.
	Tags pulumi.StringArrayInput
	// The value of this API Management Property.
	Value pulumi.StringPtrInput
}

func (PropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyState)(nil)).Elem()
}

type propertyArgs struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The display name of this API Management Property.
	DisplayName string `pulumi:"displayName"`
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Property.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Property.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Property resource.
type PropertyArgs struct {
	// The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The display name of this API Management Property.
	DisplayName pulumi.StringInput
	// The name of the API Management Property. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Property.
	Tags pulumi.StringArrayInput
	// The value of this API Management Property.
	Value pulumi.StringInput
}

func (PropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyArgs)(nil)).Elem()
}

