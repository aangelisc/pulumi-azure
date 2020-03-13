// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package bot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Directline integration for a Bot Channel
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channel_directline.markdown.
type ChannelDirectLine struct {
	pulumi.CustomResourceState

	BotName pulumi.StringOutput `pulumi:"botName"`
	Location pulumi.StringOutput `pulumi:"location"`
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	Sites ChannelDirectLineSiteArrayOutput `pulumi:"sites"`
}

// NewChannelDirectLine registers a new resource with the given unique name, arguments, and options.
func NewChannelDirectLine(ctx *pulumi.Context,
	name string, args *ChannelDirectLineArgs, opts ...pulumi.ResourceOption) (*ChannelDirectLine, error) {
	if args == nil || args.BotName == nil {
		return nil, errors.New("missing required argument 'BotName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sites == nil {
		return nil, errors.New("missing required argument 'Sites'")
	}
	if args == nil {
		args = &ChannelDirectLineArgs{}
	}
	var resource ChannelDirectLine
	err := ctx.RegisterResource("azure:bot/channelDirectLine:ChannelDirectLine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelDirectLine gets an existing ChannelDirectLine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelDirectLine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelDirectLineState, opts ...pulumi.ResourceOption) (*ChannelDirectLine, error) {
	var resource ChannelDirectLine
	err := ctx.ReadResource("azure:bot/channelDirectLine:ChannelDirectLine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelDirectLine resources.
type channelDirectLineState struct {
	BotName *string `pulumi:"botName"`
	Location *string `pulumi:"location"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	Sites []ChannelDirectLineSite `pulumi:"sites"`
}

type ChannelDirectLineState struct {
	BotName pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
	Sites ChannelDirectLineSiteArrayInput
}

func (ChannelDirectLineState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelDirectLineState)(nil)).Elem()
}

type channelDirectLineArgs struct {
	BotName string `pulumi:"botName"`
	Location *string `pulumi:"location"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Sites []ChannelDirectLineSite `pulumi:"sites"`
}

// The set of arguments for constructing a ChannelDirectLine resource.
type ChannelDirectLineArgs struct {
	BotName pulumi.StringInput
	Location pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sites ChannelDirectLineSiteArrayInput
}

func (ChannelDirectLineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelDirectLineArgs)(nil)).Elem()
}

