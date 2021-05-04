// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Redis Enterprise Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/redis"
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
// 		_, err = redis.NewEnterpriseCluster(ctx, "exampleEnterpriseCluster", &redis.EnterpriseClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("EnterpriseFlash_F300-3"),
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
// Redis Enterprise Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:redis/enterpriseCluster:EnterpriseCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1
// ```
type EnterpriseCluster struct {
	pulumi.CustomResourceState

	// DNS name of the cluster endpoint.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The minimum TLS version.  Defaults to `1.2`. Changing this forces a new Redis Enterprise Cluster to be created.
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `skuName`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Version of redis the cluster supports, e.g. '6'.
	//
	// Deprecated: This field currently is not yet being returned from the service API, please see https://github.com/Azure/azure-sdk-for-go/issues/14420 for more information
	Version pulumi.StringOutput `pulumi:"version"`
	// A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewEnterpriseCluster registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseCluster(ctx *pulumi.Context,
	name string, args *EnterpriseClusterArgs, opts ...pulumi.ResourceOption) (*EnterpriseCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource EnterpriseCluster
	err := ctx.RegisterResource("azure:redis/enterpriseCluster:EnterpriseCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseCluster gets an existing EnterpriseCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseClusterState, opts ...pulumi.ResourceOption) (*EnterpriseCluster, error) {
	var resource EnterpriseCluster
	err := ctx.ReadResource("azure:redis/enterpriseCluster:EnterpriseCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseCluster resources.
type enterpriseClusterState struct {
	// DNS name of the cluster endpoint.
	Hostname *string `pulumi:"hostname"`
	// The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	Location *string `pulumi:"location"`
	// The minimum TLS version.  Defaults to `1.2`. Changing this forces a new Redis Enterprise Cluster to be created.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `skuName`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Tags map[string]string `pulumi:"tags"`
	// Version of redis the cluster supports, e.g. '6'.
	//
	// Deprecated: This field currently is not yet being returned from the service API, please see https://github.com/Azure/azure-sdk-for-go/issues/14420 for more information
	Version *string `pulumi:"version"`
	// A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
	Zones []string `pulumi:"zones"`
}

type EnterpriseClusterState struct {
	// DNS name of the cluster endpoint.
	Hostname pulumi.StringPtrInput
	// The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	Location pulumi.StringPtrInput
	// The minimum TLS version.  Defaults to `1.2`. Changing this forces a new Redis Enterprise Cluster to be created.
	MinimumTlsVersion pulumi.StringPtrInput
	// The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `skuName`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
	SkuName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Tags pulumi.StringMapInput
	// Version of redis the cluster supports, e.g. '6'.
	//
	// Deprecated: This field currently is not yet being returned from the service API, please see https://github.com/Azure/azure-sdk-for-go/issues/14420 for more information
	Version pulumi.StringPtrInput
	// A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
	Zones pulumi.StringArrayInput
}

func (EnterpriseClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseClusterState)(nil)).Elem()
}

type enterpriseClusterArgs struct {
	// The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	Location *string `pulumi:"location"`
	// The minimum TLS version.  Defaults to `1.2`. Changing this forces a new Redis Enterprise Cluster to be created.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `skuName`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Tags map[string]string `pulumi:"tags"`
	// A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a EnterpriseCluster resource.
type EnterpriseClusterArgs struct {
	// The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	Location pulumi.StringPtrInput
	// The minimum TLS version.  Defaults to `1.2`. Changing this forces a new Redis Enterprise Cluster to be created.
	MinimumTlsVersion pulumi.StringPtrInput
	// The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
	ResourceGroupName pulumi.StringInput
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `skuName`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
	SkuName pulumi.StringInput
	// A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
	Tags pulumi.StringMapInput
	// A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
	Zones pulumi.StringArrayInput
}

func (EnterpriseClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseClusterArgs)(nil)).Elem()
}

type EnterpriseClusterInput interface {
	pulumi.Input

	ToEnterpriseClusterOutput() EnterpriseClusterOutput
	ToEnterpriseClusterOutputWithContext(ctx context.Context) EnterpriseClusterOutput
}

func (*EnterpriseCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseCluster)(nil))
}

func (i *EnterpriseCluster) ToEnterpriseClusterOutput() EnterpriseClusterOutput {
	return i.ToEnterpriseClusterOutputWithContext(context.Background())
}

func (i *EnterpriseCluster) ToEnterpriseClusterOutputWithContext(ctx context.Context) EnterpriseClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseClusterOutput)
}

func (i *EnterpriseCluster) ToEnterpriseClusterPtrOutput() EnterpriseClusterPtrOutput {
	return i.ToEnterpriseClusterPtrOutputWithContext(context.Background())
}

func (i *EnterpriseCluster) ToEnterpriseClusterPtrOutputWithContext(ctx context.Context) EnterpriseClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseClusterPtrOutput)
}

type EnterpriseClusterPtrInput interface {
	pulumi.Input

	ToEnterpriseClusterPtrOutput() EnterpriseClusterPtrOutput
	ToEnterpriseClusterPtrOutputWithContext(ctx context.Context) EnterpriseClusterPtrOutput
}

type enterpriseClusterPtrType EnterpriseClusterArgs

func (*enterpriseClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseCluster)(nil))
}

func (i *enterpriseClusterPtrType) ToEnterpriseClusterPtrOutput() EnterpriseClusterPtrOutput {
	return i.ToEnterpriseClusterPtrOutputWithContext(context.Background())
}

func (i *enterpriseClusterPtrType) ToEnterpriseClusterPtrOutputWithContext(ctx context.Context) EnterpriseClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseClusterPtrOutput)
}

// EnterpriseClusterArrayInput is an input type that accepts EnterpriseClusterArray and EnterpriseClusterArrayOutput values.
// You can construct a concrete instance of `EnterpriseClusterArrayInput` via:
//
//          EnterpriseClusterArray{ EnterpriseClusterArgs{...} }
type EnterpriseClusterArrayInput interface {
	pulumi.Input

	ToEnterpriseClusterArrayOutput() EnterpriseClusterArrayOutput
	ToEnterpriseClusterArrayOutputWithContext(context.Context) EnterpriseClusterArrayOutput
}

type EnterpriseClusterArray []EnterpriseClusterInput

func (EnterpriseClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EnterpriseCluster)(nil))
}

func (i EnterpriseClusterArray) ToEnterpriseClusterArrayOutput() EnterpriseClusterArrayOutput {
	return i.ToEnterpriseClusterArrayOutputWithContext(context.Background())
}

func (i EnterpriseClusterArray) ToEnterpriseClusterArrayOutputWithContext(ctx context.Context) EnterpriseClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseClusterArrayOutput)
}

// EnterpriseClusterMapInput is an input type that accepts EnterpriseClusterMap and EnterpriseClusterMapOutput values.
// You can construct a concrete instance of `EnterpriseClusterMapInput` via:
//
//          EnterpriseClusterMap{ "key": EnterpriseClusterArgs{...} }
type EnterpriseClusterMapInput interface {
	pulumi.Input

	ToEnterpriseClusterMapOutput() EnterpriseClusterMapOutput
	ToEnterpriseClusterMapOutputWithContext(context.Context) EnterpriseClusterMapOutput
}

type EnterpriseClusterMap map[string]EnterpriseClusterInput

func (EnterpriseClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EnterpriseCluster)(nil))
}

func (i EnterpriseClusterMap) ToEnterpriseClusterMapOutput() EnterpriseClusterMapOutput {
	return i.ToEnterpriseClusterMapOutputWithContext(context.Background())
}

func (i EnterpriseClusterMap) ToEnterpriseClusterMapOutputWithContext(ctx context.Context) EnterpriseClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseClusterMapOutput)
}

type EnterpriseClusterOutput struct {
	*pulumi.OutputState
}

func (EnterpriseClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseCluster)(nil))
}

func (o EnterpriseClusterOutput) ToEnterpriseClusterOutput() EnterpriseClusterOutput {
	return o
}

func (o EnterpriseClusterOutput) ToEnterpriseClusterOutputWithContext(ctx context.Context) EnterpriseClusterOutput {
	return o
}

func (o EnterpriseClusterOutput) ToEnterpriseClusterPtrOutput() EnterpriseClusterPtrOutput {
	return o.ToEnterpriseClusterPtrOutputWithContext(context.Background())
}

func (o EnterpriseClusterOutput) ToEnterpriseClusterPtrOutputWithContext(ctx context.Context) EnterpriseClusterPtrOutput {
	return o.ApplyT(func(v EnterpriseCluster) *EnterpriseCluster {
		return &v
	}).(EnterpriseClusterPtrOutput)
}

type EnterpriseClusterPtrOutput struct {
	*pulumi.OutputState
}

func (EnterpriseClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseCluster)(nil))
}

func (o EnterpriseClusterPtrOutput) ToEnterpriseClusterPtrOutput() EnterpriseClusterPtrOutput {
	return o
}

func (o EnterpriseClusterPtrOutput) ToEnterpriseClusterPtrOutputWithContext(ctx context.Context) EnterpriseClusterPtrOutput {
	return o
}

type EnterpriseClusterArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnterpriseCluster)(nil))
}

func (o EnterpriseClusterArrayOutput) ToEnterpriseClusterArrayOutput() EnterpriseClusterArrayOutput {
	return o
}

func (o EnterpriseClusterArrayOutput) ToEnterpriseClusterArrayOutputWithContext(ctx context.Context) EnterpriseClusterArrayOutput {
	return o
}

func (o EnterpriseClusterArrayOutput) Index(i pulumi.IntInput) EnterpriseClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnterpriseCluster {
		return vs[0].([]EnterpriseCluster)[vs[1].(int)]
	}).(EnterpriseClusterOutput)
}

type EnterpriseClusterMapOutput struct{ *pulumi.OutputState }

func (EnterpriseClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EnterpriseCluster)(nil))
}

func (o EnterpriseClusterMapOutput) ToEnterpriseClusterMapOutput() EnterpriseClusterMapOutput {
	return o
}

func (o EnterpriseClusterMapOutput) ToEnterpriseClusterMapOutputWithContext(ctx context.Context) EnterpriseClusterMapOutput {
	return o
}

func (o EnterpriseClusterMapOutput) MapIndex(k pulumi.StringInput) EnterpriseClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EnterpriseCluster {
		return vs[0].(map[string]EnterpriseCluster)[vs[1].(string)]
	}).(EnterpriseClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterpriseClusterOutput{})
	pulumi.RegisterOutputType(EnterpriseClusterPtrOutput{})
	pulumi.RegisterOutputType(EnterpriseClusterArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseClusterMapOutput{})
}
