// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 			RetentionInDays:   pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Log Analytics Workspaces can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1
// ```
type AnalyticsWorkspace struct {
	pulumi.CustomResourceState

	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb             pulumi.Float64PtrOutput `pulumi:"dailyQuotaGb"`
	InternetIngestionEnabled pulumi.BoolPtrOutput    `pulumi:"internetIngestionEnabled"`
	// Should the Log Analytics Workflow support querying over the Public Internet? Defaults to `true`.
	InternetQueryEnabled pulumi.BoolPtrOutput `pulumi:"internetQueryEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey pulumi.StringOutput `pulumi:"primarySharedKey"`
	// The capacity reservation level in GB for this workspace.  Must be in increments of 100  between 100 and 5000.
	ReservationCapcityInGbPerDay pulumi.IntPtrOutput `pulumi:"reservationCapcityInGbPerDay"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays pulumi.IntOutput `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey pulumi.StringOutput `pulumi:"secondarySharedKey"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, `CapacityReservation`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewAnalyticsWorkspace registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsWorkspace(ctx *pulumi.Context,
	name string, args *AnalyticsWorkspaceArgs, opts ...pulumi.ResourceOption) (*AnalyticsWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
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
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb             *float64 `pulumi:"dailyQuotaGb"`
	InternetIngestionEnabled *bool    `pulumi:"internetIngestionEnabled"`
	// Should the Log Analytics Workflow support querying over the Public Internet? Defaults to `true`.
	InternetQueryEnabled *bool `pulumi:"internetQueryEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	PortalUrl *string `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey *string `pulumi:"primarySharedKey"`
	// The capacity reservation level in GB for this workspace.  Must be in increments of 100  between 100 and 5000.
	ReservationCapcityInGbPerDay *int `pulumi:"reservationCapcityInGbPerDay"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, `CapacityReservation`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type AnalyticsWorkspaceState struct {
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb             pulumi.Float64PtrInput
	InternetIngestionEnabled pulumi.BoolPtrInput
	// Should the Log Analytics Workflow support querying over the Public Internet? Defaults to `true`.
	InternetQueryEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	PortalUrl pulumi.StringPtrInput
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey pulumi.StringPtrInput
	// The capacity reservation level in GB for this workspace.  Must be in increments of 100  between 100 and 5000.
	ReservationCapcityInGbPerDay pulumi.IntPtrInput
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays pulumi.IntPtrInput
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey pulumi.StringPtrInput
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, `CapacityReservation`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
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
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb             *float64 `pulumi:"dailyQuotaGb"`
	InternetIngestionEnabled *bool    `pulumi:"internetIngestionEnabled"`
	// Should the Log Analytics Workflow support querying over the Public Internet? Defaults to `true`.
	InternetQueryEnabled *bool `pulumi:"internetQueryEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The capacity reservation level in GB for this workspace.  Must be in increments of 100  between 100 and 5000.
	ReservationCapcityInGbPerDay *int `pulumi:"reservationCapcityInGbPerDay"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, `CapacityReservation`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AnalyticsWorkspace resource.
type AnalyticsWorkspaceArgs struct {
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb             pulumi.Float64PtrInput
	InternetIngestionEnabled pulumi.BoolPtrInput
	// Should the Log Analytics Workflow support querying over the Public Internet? Defaults to `true`.
	InternetQueryEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The capacity reservation level in GB for this workspace.  Must be in increments of 100  between 100 and 5000.
	ReservationCapcityInGbPerDay pulumi.IntPtrInput
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays pulumi.IntPtrInput
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, `CapacityReservation`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AnalyticsWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsWorkspaceArgs)(nil)).Elem()
}

type AnalyticsWorkspaceInput interface {
	pulumi.Input

	ToAnalyticsWorkspaceOutput() AnalyticsWorkspaceOutput
	ToAnalyticsWorkspaceOutputWithContext(ctx context.Context) AnalyticsWorkspaceOutput
}

func (*AnalyticsWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsWorkspace)(nil))
}

func (i *AnalyticsWorkspace) ToAnalyticsWorkspaceOutput() AnalyticsWorkspaceOutput {
	return i.ToAnalyticsWorkspaceOutputWithContext(context.Background())
}

func (i *AnalyticsWorkspace) ToAnalyticsWorkspaceOutputWithContext(ctx context.Context) AnalyticsWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsWorkspaceOutput)
}

func (i *AnalyticsWorkspace) ToAnalyticsWorkspacePtrOutput() AnalyticsWorkspacePtrOutput {
	return i.ToAnalyticsWorkspacePtrOutputWithContext(context.Background())
}

func (i *AnalyticsWorkspace) ToAnalyticsWorkspacePtrOutputWithContext(ctx context.Context) AnalyticsWorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsWorkspacePtrOutput)
}

type AnalyticsWorkspacePtrInput interface {
	pulumi.Input

	ToAnalyticsWorkspacePtrOutput() AnalyticsWorkspacePtrOutput
	ToAnalyticsWorkspacePtrOutputWithContext(ctx context.Context) AnalyticsWorkspacePtrOutput
}

type analyticsWorkspacePtrType AnalyticsWorkspaceArgs

func (*analyticsWorkspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsWorkspace)(nil))
}

func (i *analyticsWorkspacePtrType) ToAnalyticsWorkspacePtrOutput() AnalyticsWorkspacePtrOutput {
	return i.ToAnalyticsWorkspacePtrOutputWithContext(context.Background())
}

func (i *analyticsWorkspacePtrType) ToAnalyticsWorkspacePtrOutputWithContext(ctx context.Context) AnalyticsWorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsWorkspacePtrOutput)
}

// AnalyticsWorkspaceArrayInput is an input type that accepts AnalyticsWorkspaceArray and AnalyticsWorkspaceArrayOutput values.
// You can construct a concrete instance of `AnalyticsWorkspaceArrayInput` via:
//
//          AnalyticsWorkspaceArray{ AnalyticsWorkspaceArgs{...} }
type AnalyticsWorkspaceArrayInput interface {
	pulumi.Input

	ToAnalyticsWorkspaceArrayOutput() AnalyticsWorkspaceArrayOutput
	ToAnalyticsWorkspaceArrayOutputWithContext(context.Context) AnalyticsWorkspaceArrayOutput
}

type AnalyticsWorkspaceArray []AnalyticsWorkspaceInput

func (AnalyticsWorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AnalyticsWorkspace)(nil))
}

func (i AnalyticsWorkspaceArray) ToAnalyticsWorkspaceArrayOutput() AnalyticsWorkspaceArrayOutput {
	return i.ToAnalyticsWorkspaceArrayOutputWithContext(context.Background())
}

func (i AnalyticsWorkspaceArray) ToAnalyticsWorkspaceArrayOutputWithContext(ctx context.Context) AnalyticsWorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsWorkspaceArrayOutput)
}

// AnalyticsWorkspaceMapInput is an input type that accepts AnalyticsWorkspaceMap and AnalyticsWorkspaceMapOutput values.
// You can construct a concrete instance of `AnalyticsWorkspaceMapInput` via:
//
//          AnalyticsWorkspaceMap{ "key": AnalyticsWorkspaceArgs{...} }
type AnalyticsWorkspaceMapInput interface {
	pulumi.Input

	ToAnalyticsWorkspaceMapOutput() AnalyticsWorkspaceMapOutput
	ToAnalyticsWorkspaceMapOutputWithContext(context.Context) AnalyticsWorkspaceMapOutput
}

type AnalyticsWorkspaceMap map[string]AnalyticsWorkspaceInput

func (AnalyticsWorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AnalyticsWorkspace)(nil))
}

func (i AnalyticsWorkspaceMap) ToAnalyticsWorkspaceMapOutput() AnalyticsWorkspaceMapOutput {
	return i.ToAnalyticsWorkspaceMapOutputWithContext(context.Background())
}

func (i AnalyticsWorkspaceMap) ToAnalyticsWorkspaceMapOutputWithContext(ctx context.Context) AnalyticsWorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsWorkspaceMapOutput)
}

type AnalyticsWorkspaceOutput struct {
	*pulumi.OutputState
}

func (AnalyticsWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsWorkspace)(nil))
}

func (o AnalyticsWorkspaceOutput) ToAnalyticsWorkspaceOutput() AnalyticsWorkspaceOutput {
	return o
}

func (o AnalyticsWorkspaceOutput) ToAnalyticsWorkspaceOutputWithContext(ctx context.Context) AnalyticsWorkspaceOutput {
	return o
}

func (o AnalyticsWorkspaceOutput) ToAnalyticsWorkspacePtrOutput() AnalyticsWorkspacePtrOutput {
	return o.ToAnalyticsWorkspacePtrOutputWithContext(context.Background())
}

func (o AnalyticsWorkspaceOutput) ToAnalyticsWorkspacePtrOutputWithContext(ctx context.Context) AnalyticsWorkspacePtrOutput {
	return o.ApplyT(func(v AnalyticsWorkspace) *AnalyticsWorkspace {
		return &v
	}).(AnalyticsWorkspacePtrOutput)
}

type AnalyticsWorkspacePtrOutput struct {
	*pulumi.OutputState
}

func (AnalyticsWorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsWorkspace)(nil))
}

func (o AnalyticsWorkspacePtrOutput) ToAnalyticsWorkspacePtrOutput() AnalyticsWorkspacePtrOutput {
	return o
}

func (o AnalyticsWorkspacePtrOutput) ToAnalyticsWorkspacePtrOutputWithContext(ctx context.Context) AnalyticsWorkspacePtrOutput {
	return o
}

type AnalyticsWorkspaceArrayOutput struct{ *pulumi.OutputState }

func (AnalyticsWorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsWorkspace)(nil))
}

func (o AnalyticsWorkspaceArrayOutput) ToAnalyticsWorkspaceArrayOutput() AnalyticsWorkspaceArrayOutput {
	return o
}

func (o AnalyticsWorkspaceArrayOutput) ToAnalyticsWorkspaceArrayOutputWithContext(ctx context.Context) AnalyticsWorkspaceArrayOutput {
	return o
}

func (o AnalyticsWorkspaceArrayOutput) Index(i pulumi.IntInput) AnalyticsWorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AnalyticsWorkspace {
		return vs[0].([]AnalyticsWorkspace)[vs[1].(int)]
	}).(AnalyticsWorkspaceOutput)
}

type AnalyticsWorkspaceMapOutput struct{ *pulumi.OutputState }

func (AnalyticsWorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AnalyticsWorkspace)(nil))
}

func (o AnalyticsWorkspaceMapOutput) ToAnalyticsWorkspaceMapOutput() AnalyticsWorkspaceMapOutput {
	return o
}

func (o AnalyticsWorkspaceMapOutput) ToAnalyticsWorkspaceMapOutputWithContext(ctx context.Context) AnalyticsWorkspaceMapOutput {
	return o
}

func (o AnalyticsWorkspaceMapOutput) MapIndex(k pulumi.StringInput) AnalyticsWorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AnalyticsWorkspace {
		return vs[0].(map[string]AnalyticsWorkspace)[vs[1].(string)]
	}).(AnalyticsWorkspaceOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsWorkspaceOutput{})
	pulumi.RegisterOutputType(AnalyticsWorkspacePtrOutput{})
	pulumi.RegisterOutputType(AnalyticsWorkspaceArrayOutput{})
	pulumi.RegisterOutputType(AnalyticsWorkspaceMapOutput{})
}
