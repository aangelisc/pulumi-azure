// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Subscription within a API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_subscription.html.markdown.
type Subscription struct {
	pulumi.CustomResourceState

	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The display name of this Subscription.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil {
		args = &SubscriptionArgs{}
	}
	var resource Subscription
	err := ctx.RegisterResource("azure:apimanagement/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure:apimanagement/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The display name of this Subscription.
	DisplayName *string `pulumi:"displayName"`
	PrimaryKey *string `pulumi:"primaryKey"`
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State *string `pulumi:"state"`
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId *string `pulumi:"userId"`
}

type SubscriptionState struct {
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The display name of this Subscription.
	DisplayName pulumi.StringPtrInput
	PrimaryKey pulumi.StringPtrInput
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	SecondaryKey pulumi.StringPtrInput
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State pulumi.StringPtrInput
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The display name of this Subscription.
	DisplayName string `pulumi:"displayName"`
	PrimaryKey *string `pulumi:"primaryKey"`
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State *string `pulumi:"state"`
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The display name of this Subscription.
	DisplayName pulumi.StringInput
	PrimaryKey pulumi.StringPtrInput
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	SecondaryKey pulumi.StringPtrInput
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State pulumi.StringPtrInput
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

