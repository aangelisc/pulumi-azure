// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Workspace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/log_analytics_workspace.html.markdown.
type AnalyticsWorkspace struct {
	s *pulumi.ResourceState
}

// NewAnalyticsWorkspace registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsWorkspace(ctx *pulumi.Context,
	name string, args *AnalyticsWorkspaceArgs, opts ...pulumi.ResourceOpt) (*AnalyticsWorkspace, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["retentionInDays"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["retentionInDays"] = args.RetentionInDays
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["portalUrl"] = nil
	inputs["primarySharedKey"] = nil
	inputs["secondarySharedKey"] = nil
	inputs["workspaceId"] = nil
	s, err := ctx.RegisterResource("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsWorkspace{s: s}, nil
}

// GetAnalyticsWorkspace gets an existing AnalyticsWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsWorkspace(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AnalyticsWorkspaceState, opts ...pulumi.ResourceOpt) (*AnalyticsWorkspace, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["portalUrl"] = state.PortalUrl
		inputs["primarySharedKey"] = state.PrimarySharedKey
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["retentionInDays"] = state.RetentionInDays
		inputs["secondarySharedKey"] = state.SecondarySharedKey
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
		inputs["workspaceId"] = state.WorkspaceId
	}
	s, err := ctx.ReadResource("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsWorkspace{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AnalyticsWorkspace) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AnalyticsWorkspace) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *AnalyticsWorkspace) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
func (r *AnalyticsWorkspace) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The Portal URL for the Log Analytics Workspace.
func (r *AnalyticsWorkspace) PortalUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portalUrl"])
}

// The Primary shared key for the Log Analytics Workspace.
func (r *AnalyticsWorkspace) PrimarySharedKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["primarySharedKey"])
}

// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
func (r *AnalyticsWorkspace) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The workspace data retention in days. Possible values range between 30 and 730.
func (r *AnalyticsWorkspace) RetentionInDays() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["retentionInDays"])
}

// The Secondary shared key for the Log Analytics Workspace.
func (r *AnalyticsWorkspace) SecondarySharedKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["secondarySharedKey"])
}

// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
func (r *AnalyticsWorkspace) Sku() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sku"])
}

// A mapping of tags to assign to the resource.
func (r *AnalyticsWorkspace) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The Workspace (or Customer) ID for the Log Analytics Workspace.
func (r *AnalyticsWorkspace) WorkspaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["workspaceId"])
}

// Input properties used for looking up and filtering AnalyticsWorkspace resources.
type AnalyticsWorkspaceState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name interface{}
	// The Portal URL for the Log Analytics Workspace.
	PortalUrl interface{}
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey interface{}
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays interface{}
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey interface{}
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId interface{}
}

// The set of arguments for constructing a AnalyticsWorkspace resource.
type AnalyticsWorkspaceArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The workspace data retention in days. Possible values range between 30 and 730.
	RetentionInDays interface{}
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`).
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
