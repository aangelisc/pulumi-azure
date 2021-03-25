// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Dev Test Lab.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/devtest"
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
// 		_, err = devtest.NewLab(ctx, "exampleLab", &devtest.LabArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Tags: pulumi.StringMap{
// 				"Sydney": pulumi.String("Australia"),
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
// Dev Test Labs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:devtest/lab:Lab lab1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1
// ```
type Lab struct {
	pulumi.CustomResourceState

	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId pulumi.StringOutput `pulumi:"artifactsStorageAccountId"`
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId pulumi.StringOutput `pulumi:"defaultPremiumStorageAccountId"`
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId pulumi.StringOutput `pulumi:"defaultStorageAccountId"`
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId pulumi.StringOutput `pulumi:"premiumDataDiskStorageAccountId"`
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
}

// NewLab registers a new resource with the given unique name, arguments, and options.
func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Lab
	err := ctx.RegisterResource("azure:devtest/lab:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLab gets an existing Lab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure:devtest/lab:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lab resources.
type labState struct {
	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId *string `pulumi:"artifactsStorageAccountId"`
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId *string `pulumi:"defaultPremiumStorageAccountId"`
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId *string `pulumi:"defaultStorageAccountId"`
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId *string `pulumi:"premiumDataDiskStorageAccountId"`
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

type LabState struct {
	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountId pulumi.StringPtrInput
	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountId pulumi.StringPtrInput
	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountId pulumi.StringPtrInput
	// The ID of the Key used for this Dev Test Lab.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountId pulumi.StringPtrInput
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier pulumi.StringPtrInput
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Lab resource.
type LabArgs struct {
	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((*Lab)(nil))
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

func (i *Lab) ToLabPtrOutput() LabPtrOutput {
	return i.ToLabPtrOutputWithContext(context.Background())
}

func (i *Lab) ToLabPtrOutputWithContext(ctx context.Context) LabPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPtrOutput)
}

type LabPtrInput interface {
	pulumi.Input

	ToLabPtrOutput() LabPtrOutput
	ToLabPtrOutputWithContext(ctx context.Context) LabPtrOutput
}

type labPtrType LabArgs

func (*labPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil))
}

func (i *labPtrType) ToLabPtrOutput() LabPtrOutput {
	return i.ToLabPtrOutputWithContext(context.Background())
}

func (i *labPtrType) ToLabPtrOutputWithContext(ctx context.Context) LabPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPtrOutput)
}

// LabArrayInput is an input type that accepts LabArray and LabArrayOutput values.
// You can construct a concrete instance of `LabArrayInput` via:
//
//          LabArray{ LabArgs{...} }
type LabArrayInput interface {
	pulumi.Input

	ToLabArrayOutput() LabArrayOutput
	ToLabArrayOutputWithContext(context.Context) LabArrayOutput
}

type LabArray []LabInput

func (LabArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Lab)(nil))
}

func (i LabArray) ToLabArrayOutput() LabArrayOutput {
	return i.ToLabArrayOutputWithContext(context.Background())
}

func (i LabArray) ToLabArrayOutputWithContext(ctx context.Context) LabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabArrayOutput)
}

// LabMapInput is an input type that accepts LabMap and LabMapOutput values.
// You can construct a concrete instance of `LabMapInput` via:
//
//          LabMap{ "key": LabArgs{...} }
type LabMapInput interface {
	pulumi.Input

	ToLabMapOutput() LabMapOutput
	ToLabMapOutputWithContext(context.Context) LabMapOutput
}

type LabMap map[string]LabInput

func (LabMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Lab)(nil))
}

func (i LabMap) ToLabMapOutput() LabMapOutput {
	return i.ToLabMapOutputWithContext(context.Background())
}

func (i LabMap) ToLabMapOutputWithContext(ctx context.Context) LabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabMapOutput)
}

type LabOutput struct {
	*pulumi.OutputState
}

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lab)(nil))
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func (o LabOutput) ToLabPtrOutput() LabPtrOutput {
	return o.ToLabPtrOutputWithContext(context.Background())
}

func (o LabOutput) ToLabPtrOutputWithContext(ctx context.Context) LabPtrOutput {
	return o.ApplyT(func(v Lab) *Lab {
		return &v
	}).(LabPtrOutput)
}

type LabPtrOutput struct {
	*pulumi.OutputState
}

func (LabPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil))
}

func (o LabPtrOutput) ToLabPtrOutput() LabPtrOutput {
	return o
}

func (o LabPtrOutput) ToLabPtrOutputWithContext(ctx context.Context) LabPtrOutput {
	return o
}

type LabArrayOutput struct{ *pulumi.OutputState }

func (LabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Lab)(nil))
}

func (o LabArrayOutput) ToLabArrayOutput() LabArrayOutput {
	return o
}

func (o LabArrayOutput) ToLabArrayOutputWithContext(ctx context.Context) LabArrayOutput {
	return o
}

func (o LabArrayOutput) Index(i pulumi.IntInput) LabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Lab {
		return vs[0].([]Lab)[vs[1].(int)]
	}).(LabOutput)
}

type LabMapOutput struct{ *pulumi.OutputState }

func (LabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Lab)(nil))
}

func (o LabMapOutput) ToLabMapOutput() LabMapOutput {
	return o
}

func (o LabMapOutput) ToLabMapOutputWithContext(ctx context.Context) LabMapOutput {
	return o
}

func (o LabMapOutput) MapIndex(k pulumi.StringInput) LabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Lab {
		return vs[0].(map[string]Lab)[vs[1].(string)]
	}).(LabOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
	pulumi.RegisterOutputType(LabPtrOutput{})
	pulumi.RegisterOutputType(LabArrayOutput{})
	pulumi.RegisterOutputType(LabMapOutput{})
}
