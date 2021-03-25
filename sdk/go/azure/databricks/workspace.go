// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Databricks Workspace
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/databricks"
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
// 		_, err = databricks.NewWorkspace(ctx, "exampleWorkspace", &databricks.WorkspaceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku:               pulumi.String("standard"),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
// 			},
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
// Databrick Workspaces can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:databricks/workspace:Workspace workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// A `customParameters` block as documented below.
	CustomParameters WorkspaceCustomParametersOutput `pulumi:"customParameters"`
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID of the Managed Resource Group created by the Databricks Workspace.
	ManagedResourceGroupId pulumi.StringOutput `pulumi:"managedResourceGroupId"`
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName pulumi.StringOutput `pulumi:"managedResourceGroupName"`
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique identifier of the databricks workspace in Databricks control plane.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
	// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceUrl pulumi.StringOutput `pulumi:"workspaceUrl"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Workspace
	err := ctx.RegisterResource("azure:databricks/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure:databricks/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// A `customParameters` block as documented below.
	CustomParameters *WorkspaceCustomParameters `pulumi:"customParameters"`
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The ID of the Managed Resource Group created by the Databricks Workspace.
	ManagedResourceGroupId *string `pulumi:"managedResourceGroupId"`
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique identifier of the databricks workspace in Databricks control plane.
	WorkspaceId *string `pulumi:"workspaceId"`
	// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceUrl *string `pulumi:"workspaceUrl"`
}

type WorkspaceState struct {
	// A `customParameters` block as documented below.
	CustomParameters WorkspaceCustomParametersPtrInput
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The ID of the Managed Resource Group created by the Databricks Workspace.
	ManagedResourceGroupId pulumi.StringPtrInput
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique identifier of the databricks workspace in Databricks control plane.
	WorkspaceId pulumi.StringPtrInput
	// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceUrl pulumi.StringPtrInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// A `customParameters` block as documented below.
	CustomParameters *WorkspaceCustomParameters `pulumi:"customParameters"`
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// A `customParameters` block as documented below.
	CustomParameters WorkspaceCustomParametersPtrInput
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The `sku` to use for the Databricks Workspace. Possible values are `standard`, `premium`, or `trial`. Changing this can force a new resource to be created in some circumstances.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i *Workspace) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
}

type WorkspacePtrInput interface {
	pulumi.Input

	ToWorkspacePtrOutput() WorkspacePtrOutput
	ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput
}

type workspacePtrType WorkspaceArgs

func (*workspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil))
}

func (i *workspacePtrType) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i *workspacePtrType) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
}

// WorkspaceArrayInput is an input type that accepts WorkspaceArray and WorkspaceArrayOutput values.
// You can construct a concrete instance of `WorkspaceArrayInput` via:
//
//          WorkspaceArray{ WorkspaceArgs{...} }
type WorkspaceArrayInput interface {
	pulumi.Input

	ToWorkspaceArrayOutput() WorkspaceArrayOutput
	ToWorkspaceArrayOutputWithContext(context.Context) WorkspaceArrayOutput
}

type WorkspaceArray []WorkspaceInput

func (WorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Workspace)(nil))
}

func (i WorkspaceArray) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return i.ToWorkspaceArrayOutputWithContext(context.Background())
}

func (i WorkspaceArray) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceArrayOutput)
}

// WorkspaceMapInput is an input type that accepts WorkspaceMap and WorkspaceMapOutput values.
// You can construct a concrete instance of `WorkspaceMapInput` via:
//
//          WorkspaceMap{ "key": WorkspaceArgs{...} }
type WorkspaceMapInput interface {
	pulumi.Input

	ToWorkspaceMapOutput() WorkspaceMapOutput
	ToWorkspaceMapOutputWithContext(context.Context) WorkspaceMapOutput
}

type WorkspaceMap map[string]WorkspaceInput

func (WorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Workspace)(nil))
}

func (i WorkspaceMap) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return i.ToWorkspaceMapOutputWithContext(context.Background())
}

func (i WorkspaceMap) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceMapOutput)
}

type WorkspaceOutput struct {
	*pulumi.OutputState
}

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o.ToWorkspacePtrOutputWithContext(context.Background())
}

func (o WorkspaceOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o.ApplyT(func(v Workspace) *Workspace {
		return &v
	}).(WorkspacePtrOutput)
}

type WorkspacePtrOutput struct {
	*pulumi.OutputState
}

func (WorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil))
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o
}

type WorkspaceArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Workspace)(nil))
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) Index(i pulumi.IntInput) WorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Workspace {
		return vs[0].([]Workspace)[vs[1].(int)]
	}).(WorkspaceOutput)
}

type WorkspaceMapOutput struct{ *pulumi.OutputState }

func (WorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Workspace)(nil))
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) MapIndex(k pulumi.StringInput) WorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Workspace {
		return vs[0].(map[string]Workspace)[vs[1].(string)]
	}).(WorkspaceOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspacePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceMapOutput{})
}
