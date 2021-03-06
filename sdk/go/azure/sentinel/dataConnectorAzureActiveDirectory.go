// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Azure Active Directory Data Connector.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sentinel"
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
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sentinel.NewDataConnectorAzureActiveDirectory(ctx, "exampleDataConnectorAzureActiveDirectory", &sentinel.DataConnectorAzureActiveDirectoryArgs{
// 			LogAnalyticsWorkspaceId: exampleAnalyticsWorkspace.ID(),
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
// Azure Active Directory Data Connectors can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sentinel/dataConnectorAzureActiveDirectory:DataConnectorAzureActiveDirectory example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
// ```
type DataConnectorAzureActiveDirectory struct {
	pulumi.CustomResourceState

	// The ID of the Log Analytics Workspace that this Azure Active Directory Data Connector resides in. Changing this forces a new Azure Active Directory Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Azure Active Directory Data Connector. Changing this forces a new Azure Active Directory Data Connector to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the tenant that this Azure Active Directory Data Connector connects to. Changing this forces a new Azure Active Directory Data Connector to be created.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewDataConnectorAzureActiveDirectory registers a new resource with the given unique name, arguments, and options.
func NewDataConnectorAzureActiveDirectory(ctx *pulumi.Context,
	name string, args *DataConnectorAzureActiveDirectoryArgs, opts ...pulumi.ResourceOption) (*DataConnectorAzureActiveDirectory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'LogAnalyticsWorkspaceId'")
	}
	var resource DataConnectorAzureActiveDirectory
	err := ctx.RegisterResource("azure:sentinel/dataConnectorAzureActiveDirectory:DataConnectorAzureActiveDirectory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataConnectorAzureActiveDirectory gets an existing DataConnectorAzureActiveDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataConnectorAzureActiveDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorAzureActiveDirectoryState, opts ...pulumi.ResourceOption) (*DataConnectorAzureActiveDirectory, error) {
	var resource DataConnectorAzureActiveDirectory
	err := ctx.ReadResource("azure:sentinel/dataConnectorAzureActiveDirectory:DataConnectorAzureActiveDirectory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataConnectorAzureActiveDirectory resources.
type dataConnectorAzureActiveDirectoryState struct {
	// The ID of the Log Analytics Workspace that this Azure Active Directory Data Connector resides in. Changing this forces a new Azure Active Directory Data Connector to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Azure Active Directory Data Connector. Changing this forces a new Azure Active Directory Data Connector to be created.
	Name *string `pulumi:"name"`
	// The ID of the tenant that this Azure Active Directory Data Connector connects to. Changing this forces a new Azure Active Directory Data Connector to be created.
	TenantId *string `pulumi:"tenantId"`
}

type DataConnectorAzureActiveDirectoryState struct {
	// The ID of the Log Analytics Workspace that this Azure Active Directory Data Connector resides in. Changing this forces a new Azure Active Directory Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Azure Active Directory Data Connector. Changing this forces a new Azure Active Directory Data Connector to be created.
	Name pulumi.StringPtrInput
	// The ID of the tenant that this Azure Active Directory Data Connector connects to. Changing this forces a new Azure Active Directory Data Connector to be created.
	TenantId pulumi.StringPtrInput
}

func (DataConnectorAzureActiveDirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorAzureActiveDirectoryState)(nil)).Elem()
}

type dataConnectorAzureActiveDirectoryArgs struct {
	// The ID of the Log Analytics Workspace that this Azure Active Directory Data Connector resides in. Changing this forces a new Azure Active Directory Data Connector to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Azure Active Directory Data Connector. Changing this forces a new Azure Active Directory Data Connector to be created.
	Name *string `pulumi:"name"`
	// The ID of the tenant that this Azure Active Directory Data Connector connects to. Changing this forces a new Azure Active Directory Data Connector to be created.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a DataConnectorAzureActiveDirectory resource.
type DataConnectorAzureActiveDirectoryArgs struct {
	// The ID of the Log Analytics Workspace that this Azure Active Directory Data Connector resides in. Changing this forces a new Azure Active Directory Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// The name which should be used for this Azure Active Directory Data Connector. Changing this forces a new Azure Active Directory Data Connector to be created.
	Name pulumi.StringPtrInput
	// The ID of the tenant that this Azure Active Directory Data Connector connects to. Changing this forces a new Azure Active Directory Data Connector to be created.
	TenantId pulumi.StringPtrInput
}

func (DataConnectorAzureActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorAzureActiveDirectoryArgs)(nil)).Elem()
}

type DataConnectorAzureActiveDirectoryInput interface {
	pulumi.Input

	ToDataConnectorAzureActiveDirectoryOutput() DataConnectorAzureActiveDirectoryOutput
	ToDataConnectorAzureActiveDirectoryOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryOutput
}

func (*DataConnectorAzureActiveDirectory) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorAzureActiveDirectory)(nil))
}

func (i *DataConnectorAzureActiveDirectory) ToDataConnectorAzureActiveDirectoryOutput() DataConnectorAzureActiveDirectoryOutput {
	return i.ToDataConnectorAzureActiveDirectoryOutputWithContext(context.Background())
}

func (i *DataConnectorAzureActiveDirectory) ToDataConnectorAzureActiveDirectoryOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureActiveDirectoryOutput)
}

func (i *DataConnectorAzureActiveDirectory) ToDataConnectorAzureActiveDirectoryPtrOutput() DataConnectorAzureActiveDirectoryPtrOutput {
	return i.ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *DataConnectorAzureActiveDirectory) ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureActiveDirectoryPtrOutput)
}

type DataConnectorAzureActiveDirectoryPtrInput interface {
	pulumi.Input

	ToDataConnectorAzureActiveDirectoryPtrOutput() DataConnectorAzureActiveDirectoryPtrOutput
	ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryPtrOutput
}

type dataConnectorAzureActiveDirectoryPtrType DataConnectorAzureActiveDirectoryArgs

func (*dataConnectorAzureActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorAzureActiveDirectory)(nil))
}

func (i *dataConnectorAzureActiveDirectoryPtrType) ToDataConnectorAzureActiveDirectoryPtrOutput() DataConnectorAzureActiveDirectoryPtrOutput {
	return i.ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *dataConnectorAzureActiveDirectoryPtrType) ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureActiveDirectoryPtrOutput)
}

// DataConnectorAzureActiveDirectoryArrayInput is an input type that accepts DataConnectorAzureActiveDirectoryArray and DataConnectorAzureActiveDirectoryArrayOutput values.
// You can construct a concrete instance of `DataConnectorAzureActiveDirectoryArrayInput` via:
//
//          DataConnectorAzureActiveDirectoryArray{ DataConnectorAzureActiveDirectoryArgs{...} }
type DataConnectorAzureActiveDirectoryArrayInput interface {
	pulumi.Input

	ToDataConnectorAzureActiveDirectoryArrayOutput() DataConnectorAzureActiveDirectoryArrayOutput
	ToDataConnectorAzureActiveDirectoryArrayOutputWithContext(context.Context) DataConnectorAzureActiveDirectoryArrayOutput
}

type DataConnectorAzureActiveDirectoryArray []DataConnectorAzureActiveDirectoryInput

func (DataConnectorAzureActiveDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DataConnectorAzureActiveDirectory)(nil))
}

func (i DataConnectorAzureActiveDirectoryArray) ToDataConnectorAzureActiveDirectoryArrayOutput() DataConnectorAzureActiveDirectoryArrayOutput {
	return i.ToDataConnectorAzureActiveDirectoryArrayOutputWithContext(context.Background())
}

func (i DataConnectorAzureActiveDirectoryArray) ToDataConnectorAzureActiveDirectoryArrayOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureActiveDirectoryArrayOutput)
}

// DataConnectorAzureActiveDirectoryMapInput is an input type that accepts DataConnectorAzureActiveDirectoryMap and DataConnectorAzureActiveDirectoryMapOutput values.
// You can construct a concrete instance of `DataConnectorAzureActiveDirectoryMapInput` via:
//
//          DataConnectorAzureActiveDirectoryMap{ "key": DataConnectorAzureActiveDirectoryArgs{...} }
type DataConnectorAzureActiveDirectoryMapInput interface {
	pulumi.Input

	ToDataConnectorAzureActiveDirectoryMapOutput() DataConnectorAzureActiveDirectoryMapOutput
	ToDataConnectorAzureActiveDirectoryMapOutputWithContext(context.Context) DataConnectorAzureActiveDirectoryMapOutput
}

type DataConnectorAzureActiveDirectoryMap map[string]DataConnectorAzureActiveDirectoryInput

func (DataConnectorAzureActiveDirectoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DataConnectorAzureActiveDirectory)(nil))
}

func (i DataConnectorAzureActiveDirectoryMap) ToDataConnectorAzureActiveDirectoryMapOutput() DataConnectorAzureActiveDirectoryMapOutput {
	return i.ToDataConnectorAzureActiveDirectoryMapOutputWithContext(context.Background())
}

func (i DataConnectorAzureActiveDirectoryMap) ToDataConnectorAzureActiveDirectoryMapOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureActiveDirectoryMapOutput)
}

type DataConnectorAzureActiveDirectoryOutput struct {
	*pulumi.OutputState
}

func (DataConnectorAzureActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorAzureActiveDirectory)(nil))
}

func (o DataConnectorAzureActiveDirectoryOutput) ToDataConnectorAzureActiveDirectoryOutput() DataConnectorAzureActiveDirectoryOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryOutput) ToDataConnectorAzureActiveDirectoryOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryOutput) ToDataConnectorAzureActiveDirectoryPtrOutput() DataConnectorAzureActiveDirectoryPtrOutput {
	return o.ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o DataConnectorAzureActiveDirectoryOutput) ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryPtrOutput {
	return o.ApplyT(func(v DataConnectorAzureActiveDirectory) *DataConnectorAzureActiveDirectory {
		return &v
	}).(DataConnectorAzureActiveDirectoryPtrOutput)
}

type DataConnectorAzureActiveDirectoryPtrOutput struct {
	*pulumi.OutputState
}

func (DataConnectorAzureActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorAzureActiveDirectory)(nil))
}

func (o DataConnectorAzureActiveDirectoryPtrOutput) ToDataConnectorAzureActiveDirectoryPtrOutput() DataConnectorAzureActiveDirectoryPtrOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryPtrOutput) ToDataConnectorAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryPtrOutput {
	return o
}

type DataConnectorAzureActiveDirectoryArrayOutput struct{ *pulumi.OutputState }

func (DataConnectorAzureActiveDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataConnectorAzureActiveDirectory)(nil))
}

func (o DataConnectorAzureActiveDirectoryArrayOutput) ToDataConnectorAzureActiveDirectoryArrayOutput() DataConnectorAzureActiveDirectoryArrayOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryArrayOutput) ToDataConnectorAzureActiveDirectoryArrayOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryArrayOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryArrayOutput) Index(i pulumi.IntInput) DataConnectorAzureActiveDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataConnectorAzureActiveDirectory {
		return vs[0].([]DataConnectorAzureActiveDirectory)[vs[1].(int)]
	}).(DataConnectorAzureActiveDirectoryOutput)
}

type DataConnectorAzureActiveDirectoryMapOutput struct{ *pulumi.OutputState }

func (DataConnectorAzureActiveDirectoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataConnectorAzureActiveDirectory)(nil))
}

func (o DataConnectorAzureActiveDirectoryMapOutput) ToDataConnectorAzureActiveDirectoryMapOutput() DataConnectorAzureActiveDirectoryMapOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryMapOutput) ToDataConnectorAzureActiveDirectoryMapOutputWithContext(ctx context.Context) DataConnectorAzureActiveDirectoryMapOutput {
	return o
}

func (o DataConnectorAzureActiveDirectoryMapOutput) MapIndex(k pulumi.StringInput) DataConnectorAzureActiveDirectoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataConnectorAzureActiveDirectory {
		return vs[0].(map[string]DataConnectorAzureActiveDirectory)[vs[1].(string)]
	}).(DataConnectorAzureActiveDirectoryOutput)
}

func init() {
	pulumi.RegisterOutputType(DataConnectorAzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(DataConnectorAzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(DataConnectorAzureActiveDirectoryArrayOutput{})
	pulumi.RegisterOutputType(DataConnectorAzureActiveDirectoryMapOutput{})
}
