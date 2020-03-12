// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package operationalinsights

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Workspace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/log_analytics_workspace.html.markdown.
type AnalyticsWorkspace struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Portal URL for the Log Analytics Workspace.
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey pulumi.StringOutput `pulumi:"primarySharedKey"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays pulumi.IntOutput `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey pulumi.StringOutput `pulumi:"secondarySharedKey"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewAnalyticsWorkspace registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsWorkspace(ctx *pulumi.Context,
	name string, args *AnalyticsWorkspaceArgs, opts ...pulumi.ResourceOption) (*AnalyticsWorkspace, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &AnalyticsWorkspaceArgs{}
	}
	var resource AnalyticsWorkspace
	err := ctx.RegisterResource("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsWorkspace gets an existing AnalyticsWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsWorkspaceState, opts ...pulumi.ResourceOption) (*AnalyticsWorkspace, error) {
	var resource AnalyticsWorkspace
	err := ctx.ReadResource("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsWorkspace resources.
type analyticsWorkspaceState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Portal URL for the Log Analytics Workspace.
	PortalUrl *string `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey *string `pulumi:"primarySharedKey"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type AnalyticsWorkspaceState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Portal URL for the Log Analytics Workspace.
	PortalUrl pulumi.StringPtrInput
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey pulumi.StringPtrInput
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays pulumi.IntPtrInput
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey pulumi.StringPtrInput
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId pulumi.StringPtrInput
}

func (AnalyticsWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsWorkspaceState)(nil)).Elem()
}

type analyticsWorkspaceArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AnalyticsWorkspace resource.
type AnalyticsWorkspaceArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays pulumi.IntPtrInput
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AnalyticsWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsWorkspaceArgs)(nil)).Elem()
}

