// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Application Security Group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/application_security_group.html.markdown.
type ApplicationSecurityGroup struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Application Security Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Application Security Group.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewApplicationSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationSecurityGroup(ctx *pulumi.Context,
	name string, args *ApplicationSecurityGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationSecurityGroup, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationSecurityGroupArgs{}
	}
	var resource ApplicationSecurityGroup
	err := ctx.RegisterResource("azure:network/applicationSecurityGroup:ApplicationSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSecurityGroup gets an existing ApplicationSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSecurityGroupState, opts ...pulumi.ResourceOption) (*ApplicationSecurityGroup, error) {
	var resource ApplicationSecurityGroup
	err := ctx.ReadResource("azure:network/applicationSecurityGroup:ApplicationSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSecurityGroup resources.
type applicationSecurityGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Security Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Application Security Group.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ApplicationSecurityGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Security Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Application Security Group.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ApplicationSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSecurityGroupState)(nil)).Elem()
}

type applicationSecurityGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Security Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Application Security Group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationSecurityGroup resource.
type ApplicationSecurityGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Security Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Application Security Group.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ApplicationSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSecurityGroupArgs)(nil)).Elem()
}

