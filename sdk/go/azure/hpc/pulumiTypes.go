// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CacheNfsTargetNamespaceJunction struct {
	// The client-facing file path of this NFS target within the HPC Cache NFS Target.
	NamespacePath string `pulumi:"namespacePath"`
	// The NFS export of this NFS target within the HPC Cache NFS Target.
	NfsExport string `pulumi:"nfsExport"`
	// The relative subdirectory path from the `nfsExport` to map to the `namespacePath`. Defaults to `""`, in which case the whole `nfsExport` is exported.
	TargetPath *string `pulumi:"targetPath"`
}

// CacheNfsTargetNamespaceJunctionInput is an input type that accepts CacheNfsTargetNamespaceJunctionArgs and CacheNfsTargetNamespaceJunctionOutput values.
// You can construct a concrete instance of `CacheNfsTargetNamespaceJunctionInput` via:
//
//          CacheNfsTargetNamespaceJunctionArgs{...}
type CacheNfsTargetNamespaceJunctionInput interface {
	pulumi.Input

	ToCacheNfsTargetNamespaceJunctionOutput() CacheNfsTargetNamespaceJunctionOutput
	ToCacheNfsTargetNamespaceJunctionOutputWithContext(context.Context) CacheNfsTargetNamespaceJunctionOutput
}

type CacheNfsTargetNamespaceJunctionArgs struct {
	// The client-facing file path of this NFS target within the HPC Cache NFS Target.
	NamespacePath pulumi.StringInput `pulumi:"namespacePath"`
	// The NFS export of this NFS target within the HPC Cache NFS Target.
	NfsExport pulumi.StringInput `pulumi:"nfsExport"`
	// The relative subdirectory path from the `nfsExport` to map to the `namespacePath`. Defaults to `""`, in which case the whole `nfsExport` is exported.
	TargetPath pulumi.StringPtrInput `pulumi:"targetPath"`
}

func (CacheNfsTargetNamespaceJunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheNfsTargetNamespaceJunction)(nil)).Elem()
}

func (i CacheNfsTargetNamespaceJunctionArgs) ToCacheNfsTargetNamespaceJunctionOutput() CacheNfsTargetNamespaceJunctionOutput {
	return i.ToCacheNfsTargetNamespaceJunctionOutputWithContext(context.Background())
}

func (i CacheNfsTargetNamespaceJunctionArgs) ToCacheNfsTargetNamespaceJunctionOutputWithContext(ctx context.Context) CacheNfsTargetNamespaceJunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNfsTargetNamespaceJunctionOutput)
}

// CacheNfsTargetNamespaceJunctionArrayInput is an input type that accepts CacheNfsTargetNamespaceJunctionArray and CacheNfsTargetNamespaceJunctionArrayOutput values.
// You can construct a concrete instance of `CacheNfsTargetNamespaceJunctionArrayInput` via:
//
//          CacheNfsTargetNamespaceJunctionArray{ CacheNfsTargetNamespaceJunctionArgs{...} }
type CacheNfsTargetNamespaceJunctionArrayInput interface {
	pulumi.Input

	ToCacheNfsTargetNamespaceJunctionArrayOutput() CacheNfsTargetNamespaceJunctionArrayOutput
	ToCacheNfsTargetNamespaceJunctionArrayOutputWithContext(context.Context) CacheNfsTargetNamespaceJunctionArrayOutput
}

type CacheNfsTargetNamespaceJunctionArray []CacheNfsTargetNamespaceJunctionInput

func (CacheNfsTargetNamespaceJunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CacheNfsTargetNamespaceJunction)(nil)).Elem()
}

func (i CacheNfsTargetNamespaceJunctionArray) ToCacheNfsTargetNamespaceJunctionArrayOutput() CacheNfsTargetNamespaceJunctionArrayOutput {
	return i.ToCacheNfsTargetNamespaceJunctionArrayOutputWithContext(context.Background())
}

func (i CacheNfsTargetNamespaceJunctionArray) ToCacheNfsTargetNamespaceJunctionArrayOutputWithContext(ctx context.Context) CacheNfsTargetNamespaceJunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheNfsTargetNamespaceJunctionArrayOutput)
}

type CacheNfsTargetNamespaceJunctionOutput struct{ *pulumi.OutputState }

func (CacheNfsTargetNamespaceJunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheNfsTargetNamespaceJunction)(nil)).Elem()
}

func (o CacheNfsTargetNamespaceJunctionOutput) ToCacheNfsTargetNamespaceJunctionOutput() CacheNfsTargetNamespaceJunctionOutput {
	return o
}

func (o CacheNfsTargetNamespaceJunctionOutput) ToCacheNfsTargetNamespaceJunctionOutputWithContext(ctx context.Context) CacheNfsTargetNamespaceJunctionOutput {
	return o
}

// The client-facing file path of this NFS target within the HPC Cache NFS Target.
func (o CacheNfsTargetNamespaceJunctionOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v CacheNfsTargetNamespaceJunction) string { return v.NamespacePath }).(pulumi.StringOutput)
}

// The NFS export of this NFS target within the HPC Cache NFS Target.
func (o CacheNfsTargetNamespaceJunctionOutput) NfsExport() pulumi.StringOutput {
	return o.ApplyT(func(v CacheNfsTargetNamespaceJunction) string { return v.NfsExport }).(pulumi.StringOutput)
}

// The relative subdirectory path from the `nfsExport` to map to the `namespacePath`. Defaults to `""`, in which case the whole `nfsExport` is exported.
func (o CacheNfsTargetNamespaceJunctionOutput) TargetPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheNfsTargetNamespaceJunction) *string { return v.TargetPath }).(pulumi.StringPtrOutput)
}

type CacheNfsTargetNamespaceJunctionArrayOutput struct{ *pulumi.OutputState }

func (CacheNfsTargetNamespaceJunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CacheNfsTargetNamespaceJunction)(nil)).Elem()
}

func (o CacheNfsTargetNamespaceJunctionArrayOutput) ToCacheNfsTargetNamespaceJunctionArrayOutput() CacheNfsTargetNamespaceJunctionArrayOutput {
	return o
}

func (o CacheNfsTargetNamespaceJunctionArrayOutput) ToCacheNfsTargetNamespaceJunctionArrayOutputWithContext(ctx context.Context) CacheNfsTargetNamespaceJunctionArrayOutput {
	return o
}

func (o CacheNfsTargetNamespaceJunctionArrayOutput) Index(i pulumi.IntInput) CacheNfsTargetNamespaceJunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CacheNfsTargetNamespaceJunction {
		return vs[0].([]CacheNfsTargetNamespaceJunction)[vs[1].(int)]
	}).(CacheNfsTargetNamespaceJunctionOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheNfsTargetNamespaceJunctionOutput{})
	pulumi.RegisterOutputType(CacheNfsTargetNamespaceJunctionArrayOutput{})
}
