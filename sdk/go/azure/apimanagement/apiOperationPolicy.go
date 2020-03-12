// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management API Operation Policy
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_api_operation_policy.html.markdown.
type ApiOperationPolicy struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent pulumi.StringOutput `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrOutput `pulumi:"xmlLink"`
}

// NewApiOperationPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiOperationPolicy(ctx *pulumi.Context,
	name string, args *ApiOperationPolicyArgs, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ApiName == nil {
		return nil, errors.New("missing required argument 'ApiName'")
	}
	if args == nil || args.OperationId == nil {
		return nil, errors.New("missing required argument 'OperationId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApiOperationPolicyArgs{}
	}
	var resource ApiOperationPolicy
	err := ctx.RegisterResource("azure:apimanagement/apiOperationPolicy:ApiOperationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperationPolicy gets an existing ApiOperationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationPolicyState, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	var resource ApiOperationPolicy
	err := ctx.ReadResource("azure:apimanagement/apiOperationPolicy:ApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperationPolicy resources.
type apiOperationPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName *string `pulumi:"apiName"`
	OperationId *string `pulumi:"operationId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

type ApiOperationPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringPtrInput
	OperationId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The XML Content for this Policy.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyState)(nil)).Elem()
}

type apiOperationPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName string `pulumi:"apiName"`
	OperationId string `pulumi:"operationId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

// The set of arguments for constructing a ApiOperationPolicy resource.
type ApiOperationPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringInput
	OperationId pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The XML Content for this Policy.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyArgs)(nil)).Elem()
}

