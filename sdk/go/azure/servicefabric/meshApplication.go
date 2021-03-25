// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Service Fabric Mesh Application can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicefabric/meshApplication:MeshApplication application1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabricMesh/applications/application1
// ```
type MeshApplication struct {
	pulumi.CustomResourceState

	// Specifies the Azure Region where the Service Fabric Mesh Application should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Service Fabric Mesh Application. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Application exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Any number of `service` block as defined below.
	Services MeshApplicationServiceArrayOutput `pulumi:"services"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewMeshApplication registers a new resource with the given unique name, arguments, and options.
func NewMeshApplication(ctx *pulumi.Context,
	name string, args *MeshApplicationArgs, opts ...pulumi.ResourceOption) (*MeshApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Services == nil {
		return nil, errors.New("invalid value for required argument 'Services'")
	}
	var resource MeshApplication
	err := ctx.RegisterResource("azure:servicefabric/meshApplication:MeshApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeshApplication gets an existing MeshApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeshApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeshApplicationState, opts ...pulumi.ResourceOption) (*MeshApplication, error) {
	var resource MeshApplication
	err := ctx.ReadResource("azure:servicefabric/meshApplication:MeshApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeshApplication resources.
type meshApplicationState struct {
	// Specifies the Azure Region where the Service Fabric Mesh Application should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Application. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Application exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Any number of `service` block as defined below.
	Services []MeshApplicationService `pulumi:"services"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type MeshApplicationState struct {
	// Specifies the Azure Region where the Service Fabric Mesh Application should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Application. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Application exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Any number of `service` block as defined below.
	Services MeshApplicationServiceArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*meshApplicationState)(nil)).Elem()
}

type meshApplicationArgs struct {
	// Specifies the Azure Region where the Service Fabric Mesh Application should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Application. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Application exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Any number of `service` block as defined below.
	Services []MeshApplicationService `pulumi:"services"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MeshApplication resource.
type MeshApplicationArgs struct {
	// Specifies the Azure Region where the Service Fabric Mesh Application should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Application. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Application exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Any number of `service` block as defined below.
	Services MeshApplicationServiceArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meshApplicationArgs)(nil)).Elem()
}

type MeshApplicationInput interface {
	pulumi.Input

	ToMeshApplicationOutput() MeshApplicationOutput
	ToMeshApplicationOutputWithContext(ctx context.Context) MeshApplicationOutput
}

func (*MeshApplication) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshApplication)(nil))
}

func (i *MeshApplication) ToMeshApplicationOutput() MeshApplicationOutput {
	return i.ToMeshApplicationOutputWithContext(context.Background())
}

func (i *MeshApplication) ToMeshApplicationOutputWithContext(ctx context.Context) MeshApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshApplicationOutput)
}

func (i *MeshApplication) ToMeshApplicationPtrOutput() MeshApplicationPtrOutput {
	return i.ToMeshApplicationPtrOutputWithContext(context.Background())
}

func (i *MeshApplication) ToMeshApplicationPtrOutputWithContext(ctx context.Context) MeshApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshApplicationPtrOutput)
}

type MeshApplicationPtrInput interface {
	pulumi.Input

	ToMeshApplicationPtrOutput() MeshApplicationPtrOutput
	ToMeshApplicationPtrOutputWithContext(ctx context.Context) MeshApplicationPtrOutput
}

type meshApplicationPtrType MeshApplicationArgs

func (*meshApplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MeshApplication)(nil))
}

func (i *meshApplicationPtrType) ToMeshApplicationPtrOutput() MeshApplicationPtrOutput {
	return i.ToMeshApplicationPtrOutputWithContext(context.Background())
}

func (i *meshApplicationPtrType) ToMeshApplicationPtrOutputWithContext(ctx context.Context) MeshApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshApplicationPtrOutput)
}

// MeshApplicationArrayInput is an input type that accepts MeshApplicationArray and MeshApplicationArrayOutput values.
// You can construct a concrete instance of `MeshApplicationArrayInput` via:
//
//          MeshApplicationArray{ MeshApplicationArgs{...} }
type MeshApplicationArrayInput interface {
	pulumi.Input

	ToMeshApplicationArrayOutput() MeshApplicationArrayOutput
	ToMeshApplicationArrayOutputWithContext(context.Context) MeshApplicationArrayOutput
}

type MeshApplicationArray []MeshApplicationInput

func (MeshApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MeshApplication)(nil))
}

func (i MeshApplicationArray) ToMeshApplicationArrayOutput() MeshApplicationArrayOutput {
	return i.ToMeshApplicationArrayOutputWithContext(context.Background())
}

func (i MeshApplicationArray) ToMeshApplicationArrayOutputWithContext(ctx context.Context) MeshApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshApplicationArrayOutput)
}

// MeshApplicationMapInput is an input type that accepts MeshApplicationMap and MeshApplicationMapOutput values.
// You can construct a concrete instance of `MeshApplicationMapInput` via:
//
//          MeshApplicationMap{ "key": MeshApplicationArgs{...} }
type MeshApplicationMapInput interface {
	pulumi.Input

	ToMeshApplicationMapOutput() MeshApplicationMapOutput
	ToMeshApplicationMapOutputWithContext(context.Context) MeshApplicationMapOutput
}

type MeshApplicationMap map[string]MeshApplicationInput

func (MeshApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MeshApplication)(nil))
}

func (i MeshApplicationMap) ToMeshApplicationMapOutput() MeshApplicationMapOutput {
	return i.ToMeshApplicationMapOutputWithContext(context.Background())
}

func (i MeshApplicationMap) ToMeshApplicationMapOutputWithContext(ctx context.Context) MeshApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshApplicationMapOutput)
}

type MeshApplicationOutput struct {
	*pulumi.OutputState
}

func (MeshApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshApplication)(nil))
}

func (o MeshApplicationOutput) ToMeshApplicationOutput() MeshApplicationOutput {
	return o
}

func (o MeshApplicationOutput) ToMeshApplicationOutputWithContext(ctx context.Context) MeshApplicationOutput {
	return o
}

func (o MeshApplicationOutput) ToMeshApplicationPtrOutput() MeshApplicationPtrOutput {
	return o.ToMeshApplicationPtrOutputWithContext(context.Background())
}

func (o MeshApplicationOutput) ToMeshApplicationPtrOutputWithContext(ctx context.Context) MeshApplicationPtrOutput {
	return o.ApplyT(func(v MeshApplication) *MeshApplication {
		return &v
	}).(MeshApplicationPtrOutput)
}

type MeshApplicationPtrOutput struct {
	*pulumi.OutputState
}

func (MeshApplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MeshApplication)(nil))
}

func (o MeshApplicationPtrOutput) ToMeshApplicationPtrOutput() MeshApplicationPtrOutput {
	return o
}

func (o MeshApplicationPtrOutput) ToMeshApplicationPtrOutputWithContext(ctx context.Context) MeshApplicationPtrOutput {
	return o
}

type MeshApplicationArrayOutput struct{ *pulumi.OutputState }

func (MeshApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MeshApplication)(nil))
}

func (o MeshApplicationArrayOutput) ToMeshApplicationArrayOutput() MeshApplicationArrayOutput {
	return o
}

func (o MeshApplicationArrayOutput) ToMeshApplicationArrayOutputWithContext(ctx context.Context) MeshApplicationArrayOutput {
	return o
}

func (o MeshApplicationArrayOutput) Index(i pulumi.IntInput) MeshApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MeshApplication {
		return vs[0].([]MeshApplication)[vs[1].(int)]
	}).(MeshApplicationOutput)
}

type MeshApplicationMapOutput struct{ *pulumi.OutputState }

func (MeshApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MeshApplication)(nil))
}

func (o MeshApplicationMapOutput) ToMeshApplicationMapOutput() MeshApplicationMapOutput {
	return o
}

func (o MeshApplicationMapOutput) ToMeshApplicationMapOutputWithContext(ctx context.Context) MeshApplicationMapOutput {
	return o
}

func (o MeshApplicationMapOutput) MapIndex(k pulumi.StringInput) MeshApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MeshApplication {
		return vs[0].(map[string]MeshApplication)[vs[1].(string)]
	}).(MeshApplicationOutput)
}

func init() {
	pulumi.RegisterOutputType(MeshApplicationOutput{})
	pulumi.RegisterOutputType(MeshApplicationPtrOutput{})
	pulumi.RegisterOutputType(MeshApplicationArrayOutput{})
	pulumi.RegisterOutputType(MeshApplicationMapOutput{})
}
