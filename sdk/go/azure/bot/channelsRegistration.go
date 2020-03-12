// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package bot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Bot Channels Registration.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channels_registration.markdown.
type ChannelsRegistration struct {
	pulumi.CustomResourceState

	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringOutput `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringOutput `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringOutput `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringOutput `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewChannelsRegistration registers a new resource with the given unique name, arguments, and options.
func NewChannelsRegistration(ctx *pulumi.Context,
	name string, args *ChannelsRegistrationArgs, opts ...pulumi.ResourceOption) (*ChannelsRegistration, error) {
	if args == nil || args.MicrosoftAppId == nil {
		return nil, errors.New("missing required argument 'MicrosoftAppId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ChannelsRegistrationArgs{}
	}
	var resource ChannelsRegistration
	err := ctx.RegisterResource("azure:bot/channelsRegistration:ChannelsRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelsRegistration gets an existing ChannelsRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelsRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelsRegistrationState, opts ...pulumi.ResourceOption) (*ChannelsRegistration, error) {
	var resource ChannelsRegistration
	err := ctx.ReadResource("azure:bot/channelsRegistration:ChannelsRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelsRegistration resources.
type channelsRegistrationState struct {
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId *string `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ChannelsRegistrationState struct {
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringPtrInput
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ChannelsRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelsRegistrationState)(nil)).Elem()
}

type channelsRegistrationArgs struct {
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId string `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ChannelsRegistration resource.
type ChannelsRegistrationArgs struct {
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringInput
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ChannelsRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelsRegistrationArgs)(nil)).Elem()
}

