// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ModuleModuleLink struct {
	Hash *ModuleModuleLinkHash `pulumi:"hash"`
	// The uri of the module content (zip or nupkg).
	Uri string `pulumi:"uri"`
}

// ModuleModuleLinkInput is an input type that accepts ModuleModuleLinkArgs and ModuleModuleLinkOutput values.
// You can construct a concrete instance of `ModuleModuleLinkInput` via:
//
//          ModuleModuleLinkArgs{...}
type ModuleModuleLinkInput interface {
	pulumi.Input

	ToModuleModuleLinkOutput() ModuleModuleLinkOutput
	ToModuleModuleLinkOutputWithContext(context.Context) ModuleModuleLinkOutput
}

type ModuleModuleLinkArgs struct {
	Hash ModuleModuleLinkHashPtrInput `pulumi:"hash"`
	// The uri of the module content (zip or nupkg).
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (ModuleModuleLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLink)(nil)).Elem()
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkOutput() ModuleModuleLinkOutput {
	return i.ToModuleModuleLinkOutputWithContext(context.Background())
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkOutputWithContext(ctx context.Context) ModuleModuleLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkOutput)
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return i.ToModuleModuleLinkPtrOutputWithContext(context.Background())
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkOutput).ToModuleModuleLinkPtrOutputWithContext(ctx)
}

// ModuleModuleLinkPtrInput is an input type that accepts ModuleModuleLinkArgs, ModuleModuleLinkPtr and ModuleModuleLinkPtrOutput values.
// You can construct a concrete instance of `ModuleModuleLinkPtrInput` via:
//
//          ModuleModuleLinkArgs{...}
//
//  or:
//
//          nil
type ModuleModuleLinkPtrInput interface {
	pulumi.Input

	ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput
	ToModuleModuleLinkPtrOutputWithContext(context.Context) ModuleModuleLinkPtrOutput
}

type moduleModuleLinkPtrType ModuleModuleLinkArgs

func ModuleModuleLinkPtr(v *ModuleModuleLinkArgs) ModuleModuleLinkPtrInput {
	return (*moduleModuleLinkPtrType)(v)
}

func (*moduleModuleLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLink)(nil)).Elem()
}

func (i *moduleModuleLinkPtrType) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return i.ToModuleModuleLinkPtrOutputWithContext(context.Background())
}

func (i *moduleModuleLinkPtrType) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkPtrOutput)
}

type ModuleModuleLinkOutput struct{ *pulumi.OutputState }

func (ModuleModuleLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLink)(nil)).Elem()
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkOutput() ModuleModuleLinkOutput {
	return o
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkOutputWithContext(ctx context.Context) ModuleModuleLinkOutput {
	return o
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return o.ToModuleModuleLinkPtrOutputWithContext(context.Background())
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return o.ApplyT(func(v ModuleModuleLink) *ModuleModuleLink {
		return &v
	}).(ModuleModuleLinkPtrOutput)
}
func (o ModuleModuleLinkOutput) Hash() ModuleModuleLinkHashPtrOutput {
	return o.ApplyT(func(v ModuleModuleLink) *ModuleModuleLinkHash { return v.Hash }).(ModuleModuleLinkHashPtrOutput)
}

// The uri of the module content (zip or nupkg).
func (o ModuleModuleLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ModuleModuleLink) string { return v.Uri }).(pulumi.StringOutput)
}

type ModuleModuleLinkPtrOutput struct{ *pulumi.OutputState }

func (ModuleModuleLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLink)(nil)).Elem()
}

func (o ModuleModuleLinkPtrOutput) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return o
}

func (o ModuleModuleLinkPtrOutput) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return o
}

func (o ModuleModuleLinkPtrOutput) Elem() ModuleModuleLinkOutput {
	return o.ApplyT(func(v *ModuleModuleLink) ModuleModuleLink { return *v }).(ModuleModuleLinkOutput)
}

func (o ModuleModuleLinkPtrOutput) Hash() ModuleModuleLinkHashPtrOutput {
	return o.ApplyT(func(v *ModuleModuleLink) *ModuleModuleLinkHash {
		if v == nil {
			return nil
		}
		return v.Hash
	}).(ModuleModuleLinkHashPtrOutput)
}

// The uri of the module content (zip or nupkg).
func (o ModuleModuleLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleModuleLink) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type ModuleModuleLinkHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}

// ModuleModuleLinkHashInput is an input type that accepts ModuleModuleLinkHashArgs and ModuleModuleLinkHashOutput values.
// You can construct a concrete instance of `ModuleModuleLinkHashInput` via:
//
//          ModuleModuleLinkHashArgs{...}
type ModuleModuleLinkHashInput interface {
	pulumi.Input

	ToModuleModuleLinkHashOutput() ModuleModuleLinkHashOutput
	ToModuleModuleLinkHashOutputWithContext(context.Context) ModuleModuleLinkHashOutput
}

type ModuleModuleLinkHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (ModuleModuleLinkHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLinkHash)(nil)).Elem()
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashOutput() ModuleModuleLinkHashOutput {
	return i.ToModuleModuleLinkHashOutputWithContext(context.Background())
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashOutputWithContext(ctx context.Context) ModuleModuleLinkHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkHashOutput)
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return i.ToModuleModuleLinkHashPtrOutputWithContext(context.Background())
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkHashOutput).ToModuleModuleLinkHashPtrOutputWithContext(ctx)
}

// ModuleModuleLinkHashPtrInput is an input type that accepts ModuleModuleLinkHashArgs, ModuleModuleLinkHashPtr and ModuleModuleLinkHashPtrOutput values.
// You can construct a concrete instance of `ModuleModuleLinkHashPtrInput` via:
//
//          ModuleModuleLinkHashArgs{...}
//
//  or:
//
//          nil
type ModuleModuleLinkHashPtrInput interface {
	pulumi.Input

	ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput
	ToModuleModuleLinkHashPtrOutputWithContext(context.Context) ModuleModuleLinkHashPtrOutput
}

type moduleModuleLinkHashPtrType ModuleModuleLinkHashArgs

func ModuleModuleLinkHashPtr(v *ModuleModuleLinkHashArgs) ModuleModuleLinkHashPtrInput {
	return (*moduleModuleLinkHashPtrType)(v)
}

func (*moduleModuleLinkHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLinkHash)(nil)).Elem()
}

func (i *moduleModuleLinkHashPtrType) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return i.ToModuleModuleLinkHashPtrOutputWithContext(context.Background())
}

func (i *moduleModuleLinkHashPtrType) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkHashPtrOutput)
}

type ModuleModuleLinkHashOutput struct{ *pulumi.OutputState }

func (ModuleModuleLinkHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLinkHash)(nil)).Elem()
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashOutput() ModuleModuleLinkHashOutput {
	return o
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashOutputWithContext(ctx context.Context) ModuleModuleLinkHashOutput {
	return o
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return o.ToModuleModuleLinkHashPtrOutputWithContext(context.Background())
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return o.ApplyT(func(v ModuleModuleLinkHash) *ModuleModuleLinkHash {
		return &v
	}).(ModuleModuleLinkHashPtrOutput)
}
func (o ModuleModuleLinkHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ModuleModuleLinkHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ModuleModuleLinkHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ModuleModuleLinkHash) string { return v.Value }).(pulumi.StringOutput)
}

type ModuleModuleLinkHashPtrOutput struct{ *pulumi.OutputState }

func (ModuleModuleLinkHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLinkHash)(nil)).Elem()
}

func (o ModuleModuleLinkHashPtrOutput) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return o
}

func (o ModuleModuleLinkHashPtrOutput) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return o
}

func (o ModuleModuleLinkHashPtrOutput) Elem() ModuleModuleLinkHashOutput {
	return o.ApplyT(func(v *ModuleModuleLinkHash) ModuleModuleLinkHash { return *v }).(ModuleModuleLinkHashOutput)
}

func (o ModuleModuleLinkHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleModuleLinkHash) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ModuleModuleLinkHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleModuleLinkHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type RunBookJobSchedule struct {
	JobScheduleId *string           `pulumi:"jobScheduleId"`
	Parameters    map[string]string `pulumi:"parameters"`
	RunOn         *string           `pulumi:"runOn"`
	ScheduleName  string            `pulumi:"scheduleName"`
}

// RunBookJobScheduleInput is an input type that accepts RunBookJobScheduleArgs and RunBookJobScheduleOutput values.
// You can construct a concrete instance of `RunBookJobScheduleInput` via:
//
//          RunBookJobScheduleArgs{...}
type RunBookJobScheduleInput interface {
	pulumi.Input

	ToRunBookJobScheduleOutput() RunBookJobScheduleOutput
	ToRunBookJobScheduleOutputWithContext(context.Context) RunBookJobScheduleOutput
}

type RunBookJobScheduleArgs struct {
	JobScheduleId pulumi.StringPtrInput `pulumi:"jobScheduleId"`
	Parameters    pulumi.StringMapInput `pulumi:"parameters"`
	RunOn         pulumi.StringPtrInput `pulumi:"runOn"`
	ScheduleName  pulumi.StringInput    `pulumi:"scheduleName"`
}

func (RunBookJobScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookJobSchedule)(nil)).Elem()
}

func (i RunBookJobScheduleArgs) ToRunBookJobScheduleOutput() RunBookJobScheduleOutput {
	return i.ToRunBookJobScheduleOutputWithContext(context.Background())
}

func (i RunBookJobScheduleArgs) ToRunBookJobScheduleOutputWithContext(ctx context.Context) RunBookJobScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookJobScheduleOutput)
}

// RunBookJobScheduleArrayInput is an input type that accepts RunBookJobScheduleArray and RunBookJobScheduleArrayOutput values.
// You can construct a concrete instance of `RunBookJobScheduleArrayInput` via:
//
//          RunBookJobScheduleArray{ RunBookJobScheduleArgs{...} }
type RunBookJobScheduleArrayInput interface {
	pulumi.Input

	ToRunBookJobScheduleArrayOutput() RunBookJobScheduleArrayOutput
	ToRunBookJobScheduleArrayOutputWithContext(context.Context) RunBookJobScheduleArrayOutput
}

type RunBookJobScheduleArray []RunBookJobScheduleInput

func (RunBookJobScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunBookJobSchedule)(nil)).Elem()
}

func (i RunBookJobScheduleArray) ToRunBookJobScheduleArrayOutput() RunBookJobScheduleArrayOutput {
	return i.ToRunBookJobScheduleArrayOutputWithContext(context.Background())
}

func (i RunBookJobScheduleArray) ToRunBookJobScheduleArrayOutputWithContext(ctx context.Context) RunBookJobScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookJobScheduleArrayOutput)
}

type RunBookJobScheduleOutput struct{ *pulumi.OutputState }

func (RunBookJobScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookJobSchedule)(nil)).Elem()
}

func (o RunBookJobScheduleOutput) ToRunBookJobScheduleOutput() RunBookJobScheduleOutput {
	return o
}

func (o RunBookJobScheduleOutput) ToRunBookJobScheduleOutputWithContext(ctx context.Context) RunBookJobScheduleOutput {
	return o
}

func (o RunBookJobScheduleOutput) JobScheduleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunBookJobSchedule) *string { return v.JobScheduleId }).(pulumi.StringPtrOutput)
}

func (o RunBookJobScheduleOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v RunBookJobSchedule) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o RunBookJobScheduleOutput) RunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunBookJobSchedule) *string { return v.RunOn }).(pulumi.StringPtrOutput)
}

func (o RunBookJobScheduleOutput) ScheduleName() pulumi.StringOutput {
	return o.ApplyT(func(v RunBookJobSchedule) string { return v.ScheduleName }).(pulumi.StringOutput)
}

type RunBookJobScheduleArrayOutput struct{ *pulumi.OutputState }

func (RunBookJobScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunBookJobSchedule)(nil)).Elem()
}

func (o RunBookJobScheduleArrayOutput) ToRunBookJobScheduleArrayOutput() RunBookJobScheduleArrayOutput {
	return o
}

func (o RunBookJobScheduleArrayOutput) ToRunBookJobScheduleArrayOutputWithContext(ctx context.Context) RunBookJobScheduleArrayOutput {
	return o
}

func (o RunBookJobScheduleArrayOutput) Index(i pulumi.IntInput) RunBookJobScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RunBookJobSchedule {
		return vs[0].([]RunBookJobSchedule)[vs[1].(int)]
	}).(RunBookJobScheduleOutput)
}

type RunBookPublishContentLink struct {
	Hash *RunBookPublishContentLinkHash `pulumi:"hash"`
	// The uri of the runbook content.
	Uri     string  `pulumi:"uri"`
	Version *string `pulumi:"version"`
}

// RunBookPublishContentLinkInput is an input type that accepts RunBookPublishContentLinkArgs and RunBookPublishContentLinkOutput values.
// You can construct a concrete instance of `RunBookPublishContentLinkInput` via:
//
//          RunBookPublishContentLinkArgs{...}
type RunBookPublishContentLinkInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkOutput() RunBookPublishContentLinkOutput
	ToRunBookPublishContentLinkOutputWithContext(context.Context) RunBookPublishContentLinkOutput
}

type RunBookPublishContentLinkArgs struct {
	Hash RunBookPublishContentLinkHashPtrInput `pulumi:"hash"`
	// The uri of the runbook content.
	Uri     pulumi.StringInput    `pulumi:"uri"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (RunBookPublishContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLink)(nil)).Elem()
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkOutput() RunBookPublishContentLinkOutput {
	return i.ToRunBookPublishContentLinkOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkOutputWithContext(ctx context.Context) RunBookPublishContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkOutput)
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return i.ToRunBookPublishContentLinkPtrOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkOutput).ToRunBookPublishContentLinkPtrOutputWithContext(ctx)
}

// RunBookPublishContentLinkPtrInput is an input type that accepts RunBookPublishContentLinkArgs, RunBookPublishContentLinkPtr and RunBookPublishContentLinkPtrOutput values.
// You can construct a concrete instance of `RunBookPublishContentLinkPtrInput` via:
//
//          RunBookPublishContentLinkArgs{...}
//
//  or:
//
//          nil
type RunBookPublishContentLinkPtrInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput
	ToRunBookPublishContentLinkPtrOutputWithContext(context.Context) RunBookPublishContentLinkPtrOutput
}

type runBookPublishContentLinkPtrType RunBookPublishContentLinkArgs

func RunBookPublishContentLinkPtr(v *RunBookPublishContentLinkArgs) RunBookPublishContentLinkPtrInput {
	return (*runBookPublishContentLinkPtrType)(v)
}

func (*runBookPublishContentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLink)(nil)).Elem()
}

func (i *runBookPublishContentLinkPtrType) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return i.ToRunBookPublishContentLinkPtrOutputWithContext(context.Background())
}

func (i *runBookPublishContentLinkPtrType) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkPtrOutput)
}

type RunBookPublishContentLinkOutput struct{ *pulumi.OutputState }

func (RunBookPublishContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLink)(nil)).Elem()
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkOutput() RunBookPublishContentLinkOutput {
	return o
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkOutputWithContext(ctx context.Context) RunBookPublishContentLinkOutput {
	return o
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return o.ToRunBookPublishContentLinkPtrOutputWithContext(context.Background())
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return o.ApplyT(func(v RunBookPublishContentLink) *RunBookPublishContentLink {
		return &v
	}).(RunBookPublishContentLinkPtrOutput)
}
func (o RunBookPublishContentLinkOutput) Hash() RunBookPublishContentLinkHashPtrOutput {
	return o.ApplyT(func(v RunBookPublishContentLink) *RunBookPublishContentLinkHash { return v.Hash }).(RunBookPublishContentLinkHashPtrOutput)
}

// The uri of the runbook content.
func (o RunBookPublishContentLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v RunBookPublishContentLink) string { return v.Uri }).(pulumi.StringOutput)
}

func (o RunBookPublishContentLinkOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunBookPublishContentLink) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type RunBookPublishContentLinkPtrOutput struct{ *pulumi.OutputState }

func (RunBookPublishContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLink)(nil)).Elem()
}

func (o RunBookPublishContentLinkPtrOutput) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return o
}

func (o RunBookPublishContentLinkPtrOutput) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return o
}

func (o RunBookPublishContentLinkPtrOutput) Elem() RunBookPublishContentLinkOutput {
	return o.ApplyT(func(v *RunBookPublishContentLink) RunBookPublishContentLink { return *v }).(RunBookPublishContentLinkOutput)
}

func (o RunBookPublishContentLinkPtrOutput) Hash() RunBookPublishContentLinkHashPtrOutput {
	return o.ApplyT(func(v *RunBookPublishContentLink) *RunBookPublishContentLinkHash {
		if v == nil {
			return nil
		}
		return v.Hash
	}).(RunBookPublishContentLinkHashPtrOutput)
}

// The uri of the runbook content.
func (o RunBookPublishContentLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunBookPublishContentLink) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o RunBookPublishContentLinkPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunBookPublishContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type RunBookPublishContentLinkHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}

// RunBookPublishContentLinkHashInput is an input type that accepts RunBookPublishContentLinkHashArgs and RunBookPublishContentLinkHashOutput values.
// You can construct a concrete instance of `RunBookPublishContentLinkHashInput` via:
//
//          RunBookPublishContentLinkHashArgs{...}
type RunBookPublishContentLinkHashInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkHashOutput() RunBookPublishContentLinkHashOutput
	ToRunBookPublishContentLinkHashOutputWithContext(context.Context) RunBookPublishContentLinkHashOutput
}

type RunBookPublishContentLinkHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (RunBookPublishContentLinkHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLinkHash)(nil)).Elem()
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashOutput() RunBookPublishContentLinkHashOutput {
	return i.ToRunBookPublishContentLinkHashOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkHashOutput)
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return i.ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkHashOutput).ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx)
}

// RunBookPublishContentLinkHashPtrInput is an input type that accepts RunBookPublishContentLinkHashArgs, RunBookPublishContentLinkHashPtr and RunBookPublishContentLinkHashPtrOutput values.
// You can construct a concrete instance of `RunBookPublishContentLinkHashPtrInput` via:
//
//          RunBookPublishContentLinkHashArgs{...}
//
//  or:
//
//          nil
type RunBookPublishContentLinkHashPtrInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput
	ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Context) RunBookPublishContentLinkHashPtrOutput
}

type runBookPublishContentLinkHashPtrType RunBookPublishContentLinkHashArgs

func RunBookPublishContentLinkHashPtr(v *RunBookPublishContentLinkHashArgs) RunBookPublishContentLinkHashPtrInput {
	return (*runBookPublishContentLinkHashPtrType)(v)
}

func (*runBookPublishContentLinkHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLinkHash)(nil)).Elem()
}

func (i *runBookPublishContentLinkHashPtrType) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return i.ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Background())
}

func (i *runBookPublishContentLinkHashPtrType) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkHashPtrOutput)
}

type RunBookPublishContentLinkHashOutput struct{ *pulumi.OutputState }

func (RunBookPublishContentLinkHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLinkHash)(nil)).Elem()
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashOutput() RunBookPublishContentLinkHashOutput {
	return o
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashOutput {
	return o
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return o.ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Background())
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return o.ApplyT(func(v RunBookPublishContentLinkHash) *RunBookPublishContentLinkHash {
		return &v
	}).(RunBookPublishContentLinkHashPtrOutput)
}
func (o RunBookPublishContentLinkHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v RunBookPublishContentLinkHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o RunBookPublishContentLinkHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v RunBookPublishContentLinkHash) string { return v.Value }).(pulumi.StringOutput)
}

type RunBookPublishContentLinkHashPtrOutput struct{ *pulumi.OutputState }

func (RunBookPublishContentLinkHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLinkHash)(nil)).Elem()
}

func (o RunBookPublishContentLinkHashPtrOutput) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return o
}

func (o RunBookPublishContentLinkHashPtrOutput) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return o
}

func (o RunBookPublishContentLinkHashPtrOutput) Elem() RunBookPublishContentLinkHashOutput {
	return o.ApplyT(func(v *RunBookPublishContentLinkHash) RunBookPublishContentLinkHash { return *v }).(RunBookPublishContentLinkHashOutput)
}

func (o RunBookPublishContentLinkHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunBookPublishContentLinkHash) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o RunBookPublishContentLinkHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunBookPublishContentLinkHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ScheduleMonthlyOccurrence struct {
	// Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	Day string `pulumi:"day"`
	// Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
	Occurrence int `pulumi:"occurrence"`
}

// ScheduleMonthlyOccurrenceInput is an input type that accepts ScheduleMonthlyOccurrenceArgs and ScheduleMonthlyOccurrenceOutput values.
// You can construct a concrete instance of `ScheduleMonthlyOccurrenceInput` via:
//
//          ScheduleMonthlyOccurrenceArgs{...}
type ScheduleMonthlyOccurrenceInput interface {
	pulumi.Input

	ToScheduleMonthlyOccurrenceOutput() ScheduleMonthlyOccurrenceOutput
	ToScheduleMonthlyOccurrenceOutputWithContext(context.Context) ScheduleMonthlyOccurrenceOutput
}

type ScheduleMonthlyOccurrenceArgs struct {
	// Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	Day pulumi.StringInput `pulumi:"day"`
	// Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
	Occurrence pulumi.IntInput `pulumi:"occurrence"`
}

func (ScheduleMonthlyOccurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i ScheduleMonthlyOccurrenceArgs) ToScheduleMonthlyOccurrenceOutput() ScheduleMonthlyOccurrenceOutput {
	return i.ToScheduleMonthlyOccurrenceOutputWithContext(context.Background())
}

func (i ScheduleMonthlyOccurrenceArgs) ToScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMonthlyOccurrenceOutput)
}

// ScheduleMonthlyOccurrenceArrayInput is an input type that accepts ScheduleMonthlyOccurrenceArray and ScheduleMonthlyOccurrenceArrayOutput values.
// You can construct a concrete instance of `ScheduleMonthlyOccurrenceArrayInput` via:
//
//          ScheduleMonthlyOccurrenceArray{ ScheduleMonthlyOccurrenceArgs{...} }
type ScheduleMonthlyOccurrenceArrayInput interface {
	pulumi.Input

	ToScheduleMonthlyOccurrenceArrayOutput() ScheduleMonthlyOccurrenceArrayOutput
	ToScheduleMonthlyOccurrenceArrayOutputWithContext(context.Context) ScheduleMonthlyOccurrenceArrayOutput
}

type ScheduleMonthlyOccurrenceArray []ScheduleMonthlyOccurrenceInput

func (ScheduleMonthlyOccurrenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i ScheduleMonthlyOccurrenceArray) ToScheduleMonthlyOccurrenceArrayOutput() ScheduleMonthlyOccurrenceArrayOutput {
	return i.ToScheduleMonthlyOccurrenceArrayOutputWithContext(context.Background())
}

func (i ScheduleMonthlyOccurrenceArray) ToScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMonthlyOccurrenceArrayOutput)
}

type ScheduleMonthlyOccurrenceOutput struct{ *pulumi.OutputState }

func (ScheduleMonthlyOccurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o ScheduleMonthlyOccurrenceOutput) ToScheduleMonthlyOccurrenceOutput() ScheduleMonthlyOccurrenceOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceOutput) ToScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceOutput {
	return o
}

// Day of the occurrence. Must be one of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
func (o ScheduleMonthlyOccurrenceOutput) Day() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleMonthlyOccurrence) string { return v.Day }).(pulumi.StringOutput)
}

// Occurrence of the week within the month. Must be between `1` and `5`. `-1` for last week within the month.
func (o ScheduleMonthlyOccurrenceOutput) Occurrence() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleMonthlyOccurrence) int { return v.Occurrence }).(pulumi.IntOutput)
}

type ScheduleMonthlyOccurrenceArrayOutput struct{ *pulumi.OutputState }

func (ScheduleMonthlyOccurrenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o ScheduleMonthlyOccurrenceArrayOutput) ToScheduleMonthlyOccurrenceArrayOutput() ScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceArrayOutput) ToScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceArrayOutput) Index(i pulumi.IntInput) ScheduleMonthlyOccurrenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleMonthlyOccurrence {
		return vs[0].([]ScheduleMonthlyOccurrence)[vs[1].(int)]
	}).(ScheduleMonthlyOccurrenceOutput)
}

func init() {
	pulumi.RegisterOutputType(ModuleModuleLinkOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkPtrOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkHashOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkHashPtrOutput{})
	pulumi.RegisterOutputType(RunBookJobScheduleOutput{})
	pulumi.RegisterOutputType(RunBookJobScheduleArrayOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkPtrOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkHashOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkHashPtrOutput{})
	pulumi.RegisterOutputType(ScheduleMonthlyOccurrenceOutput{})
	pulumi.RegisterOutputType(ScheduleMonthlyOccurrenceArrayOutput{})
}
