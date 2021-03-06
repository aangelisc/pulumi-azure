// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionHttpRunAfter struct {
	// Specifies the name of the precedent HTTP Action.
	ActionName string `pulumi:"actionName"`
	// Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
	ActionResult string `pulumi:"actionResult"`
}

// ActionHttpRunAfterInput is an input type that accepts ActionHttpRunAfterArgs and ActionHttpRunAfterOutput values.
// You can construct a concrete instance of `ActionHttpRunAfterInput` via:
//
//          ActionHttpRunAfterArgs{...}
type ActionHttpRunAfterInput interface {
	pulumi.Input

	ToActionHttpRunAfterOutput() ActionHttpRunAfterOutput
	ToActionHttpRunAfterOutputWithContext(context.Context) ActionHttpRunAfterOutput
}

type ActionHttpRunAfterArgs struct {
	// Specifies the name of the precedent HTTP Action.
	ActionName pulumi.StringInput `pulumi:"actionName"`
	// Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
	ActionResult pulumi.StringInput `pulumi:"actionResult"`
}

func (ActionHttpRunAfterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionHttpRunAfter)(nil)).Elem()
}

func (i ActionHttpRunAfterArgs) ToActionHttpRunAfterOutput() ActionHttpRunAfterOutput {
	return i.ToActionHttpRunAfterOutputWithContext(context.Background())
}

func (i ActionHttpRunAfterArgs) ToActionHttpRunAfterOutputWithContext(ctx context.Context) ActionHttpRunAfterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpRunAfterOutput)
}

// ActionHttpRunAfterArrayInput is an input type that accepts ActionHttpRunAfterArray and ActionHttpRunAfterArrayOutput values.
// You can construct a concrete instance of `ActionHttpRunAfterArrayInput` via:
//
//          ActionHttpRunAfterArray{ ActionHttpRunAfterArgs{...} }
type ActionHttpRunAfterArrayInput interface {
	pulumi.Input

	ToActionHttpRunAfterArrayOutput() ActionHttpRunAfterArrayOutput
	ToActionHttpRunAfterArrayOutputWithContext(context.Context) ActionHttpRunAfterArrayOutput
}

type ActionHttpRunAfterArray []ActionHttpRunAfterInput

func (ActionHttpRunAfterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionHttpRunAfter)(nil)).Elem()
}

func (i ActionHttpRunAfterArray) ToActionHttpRunAfterArrayOutput() ActionHttpRunAfterArrayOutput {
	return i.ToActionHttpRunAfterArrayOutputWithContext(context.Background())
}

func (i ActionHttpRunAfterArray) ToActionHttpRunAfterArrayOutputWithContext(ctx context.Context) ActionHttpRunAfterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpRunAfterArrayOutput)
}

type ActionHttpRunAfterOutput struct{ *pulumi.OutputState }

func (ActionHttpRunAfterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionHttpRunAfter)(nil)).Elem()
}

func (o ActionHttpRunAfterOutput) ToActionHttpRunAfterOutput() ActionHttpRunAfterOutput {
	return o
}

func (o ActionHttpRunAfterOutput) ToActionHttpRunAfterOutputWithContext(ctx context.Context) ActionHttpRunAfterOutput {
	return o
}

// Specifies the name of the precedent HTTP Action.
func (o ActionHttpRunAfterOutput) ActionName() pulumi.StringOutput {
	return o.ApplyT(func(v ActionHttpRunAfter) string { return v.ActionName }).(pulumi.StringOutput)
}

// Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include `Succeeded`, `Failed`, `Skipped` and `TimedOut`.
func (o ActionHttpRunAfterOutput) ActionResult() pulumi.StringOutput {
	return o.ApplyT(func(v ActionHttpRunAfter) string { return v.ActionResult }).(pulumi.StringOutput)
}

type ActionHttpRunAfterArrayOutput struct{ *pulumi.OutputState }

func (ActionHttpRunAfterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionHttpRunAfter)(nil)).Elem()
}

func (o ActionHttpRunAfterArrayOutput) ToActionHttpRunAfterArrayOutput() ActionHttpRunAfterArrayOutput {
	return o
}

func (o ActionHttpRunAfterArrayOutput) ToActionHttpRunAfterArrayOutputWithContext(ctx context.Context) ActionHttpRunAfterArrayOutput {
	return o
}

func (o ActionHttpRunAfterArrayOutput) Index(i pulumi.IntInput) ActionHttpRunAfterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionHttpRunAfter {
		return vs[0].([]ActionHttpRunAfter)[vs[1].(int)]
	}).(ActionHttpRunAfterOutput)
}

type TriggerRecurrenceSchedule struct {
	// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
	AtTheseHours []int `pulumi:"atTheseHours"`
	// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
	AtTheseMinutes []int `pulumi:"atTheseMinutes"`
	// Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
	OnTheseDays []string `pulumi:"onTheseDays"`
}

// TriggerRecurrenceScheduleInput is an input type that accepts TriggerRecurrenceScheduleArgs and TriggerRecurrenceScheduleOutput values.
// You can construct a concrete instance of `TriggerRecurrenceScheduleInput` via:
//
//          TriggerRecurrenceScheduleArgs{...}
type TriggerRecurrenceScheduleInput interface {
	pulumi.Input

	ToTriggerRecurrenceScheduleOutput() TriggerRecurrenceScheduleOutput
	ToTriggerRecurrenceScheduleOutputWithContext(context.Context) TriggerRecurrenceScheduleOutput
}

type TriggerRecurrenceScheduleArgs struct {
	// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
	AtTheseHours pulumi.IntArrayInput `pulumi:"atTheseHours"`
	// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
	AtTheseMinutes pulumi.IntArrayInput `pulumi:"atTheseMinutes"`
	// Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
	OnTheseDays pulumi.StringArrayInput `pulumi:"onTheseDays"`
}

func (TriggerRecurrenceScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerRecurrenceSchedule)(nil)).Elem()
}

func (i TriggerRecurrenceScheduleArgs) ToTriggerRecurrenceScheduleOutput() TriggerRecurrenceScheduleOutput {
	return i.ToTriggerRecurrenceScheduleOutputWithContext(context.Background())
}

func (i TriggerRecurrenceScheduleArgs) ToTriggerRecurrenceScheduleOutputWithContext(ctx context.Context) TriggerRecurrenceScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerRecurrenceScheduleOutput)
}

func (i TriggerRecurrenceScheduleArgs) ToTriggerRecurrenceSchedulePtrOutput() TriggerRecurrenceSchedulePtrOutput {
	return i.ToTriggerRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (i TriggerRecurrenceScheduleArgs) ToTriggerRecurrenceSchedulePtrOutputWithContext(ctx context.Context) TriggerRecurrenceSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerRecurrenceScheduleOutput).ToTriggerRecurrenceSchedulePtrOutputWithContext(ctx)
}

// TriggerRecurrenceSchedulePtrInput is an input type that accepts TriggerRecurrenceScheduleArgs, TriggerRecurrenceSchedulePtr and TriggerRecurrenceSchedulePtrOutput values.
// You can construct a concrete instance of `TriggerRecurrenceSchedulePtrInput` via:
//
//          TriggerRecurrenceScheduleArgs{...}
//
//  or:
//
//          nil
type TriggerRecurrenceSchedulePtrInput interface {
	pulumi.Input

	ToTriggerRecurrenceSchedulePtrOutput() TriggerRecurrenceSchedulePtrOutput
	ToTriggerRecurrenceSchedulePtrOutputWithContext(context.Context) TriggerRecurrenceSchedulePtrOutput
}

type triggerRecurrenceSchedulePtrType TriggerRecurrenceScheduleArgs

func TriggerRecurrenceSchedulePtr(v *TriggerRecurrenceScheduleArgs) TriggerRecurrenceSchedulePtrInput {
	return (*triggerRecurrenceSchedulePtrType)(v)
}

func (*triggerRecurrenceSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerRecurrenceSchedule)(nil)).Elem()
}

func (i *triggerRecurrenceSchedulePtrType) ToTriggerRecurrenceSchedulePtrOutput() TriggerRecurrenceSchedulePtrOutput {
	return i.ToTriggerRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (i *triggerRecurrenceSchedulePtrType) ToTriggerRecurrenceSchedulePtrOutputWithContext(ctx context.Context) TriggerRecurrenceSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerRecurrenceSchedulePtrOutput)
}

type TriggerRecurrenceScheduleOutput struct{ *pulumi.OutputState }

func (TriggerRecurrenceScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerRecurrenceSchedule)(nil)).Elem()
}

func (o TriggerRecurrenceScheduleOutput) ToTriggerRecurrenceScheduleOutput() TriggerRecurrenceScheduleOutput {
	return o
}

func (o TriggerRecurrenceScheduleOutput) ToTriggerRecurrenceScheduleOutputWithContext(ctx context.Context) TriggerRecurrenceScheduleOutput {
	return o
}

func (o TriggerRecurrenceScheduleOutput) ToTriggerRecurrenceSchedulePtrOutput() TriggerRecurrenceSchedulePtrOutput {
	return o.ToTriggerRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (o TriggerRecurrenceScheduleOutput) ToTriggerRecurrenceSchedulePtrOutputWithContext(ctx context.Context) TriggerRecurrenceSchedulePtrOutput {
	return o.ApplyT(func(v TriggerRecurrenceSchedule) *TriggerRecurrenceSchedule {
		return &v
	}).(TriggerRecurrenceSchedulePtrOutput)
}

// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
func (o TriggerRecurrenceScheduleOutput) AtTheseHours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v TriggerRecurrenceSchedule) []int { return v.AtTheseHours }).(pulumi.IntArrayOutput)
}

// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
func (o TriggerRecurrenceScheduleOutput) AtTheseMinutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v TriggerRecurrenceSchedule) []int { return v.AtTheseMinutes }).(pulumi.IntArrayOutput)
}

// Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
func (o TriggerRecurrenceScheduleOutput) OnTheseDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerRecurrenceSchedule) []string { return v.OnTheseDays }).(pulumi.StringArrayOutput)
}

type TriggerRecurrenceSchedulePtrOutput struct{ *pulumi.OutputState }

func (TriggerRecurrenceSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerRecurrenceSchedule)(nil)).Elem()
}

func (o TriggerRecurrenceSchedulePtrOutput) ToTriggerRecurrenceSchedulePtrOutput() TriggerRecurrenceSchedulePtrOutput {
	return o
}

func (o TriggerRecurrenceSchedulePtrOutput) ToTriggerRecurrenceSchedulePtrOutputWithContext(ctx context.Context) TriggerRecurrenceSchedulePtrOutput {
	return o
}

func (o TriggerRecurrenceSchedulePtrOutput) Elem() TriggerRecurrenceScheduleOutput {
	return o.ApplyT(func(v *TriggerRecurrenceSchedule) TriggerRecurrenceSchedule { return *v }).(TriggerRecurrenceScheduleOutput)
}

// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
func (o TriggerRecurrenceSchedulePtrOutput) AtTheseHours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *TriggerRecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.AtTheseHours
	}).(pulumi.IntArrayOutput)
}

// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
func (o TriggerRecurrenceSchedulePtrOutput) AtTheseMinutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *TriggerRecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.AtTheseMinutes
	}).(pulumi.IntArrayOutput)
}

// Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
func (o TriggerRecurrenceSchedulePtrOutput) OnTheseDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TriggerRecurrenceSchedule) []string {
		if v == nil {
			return nil
		}
		return v.OnTheseDays
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionHttpRunAfterOutput{})
	pulumi.RegisterOutputType(ActionHttpRunAfterArrayOutput{})
	pulumi.RegisterOutputType(TriggerRecurrenceScheduleOutput{})
	pulumi.RegisterOutputType(TriggerRecurrenceSchedulePtrOutput{})
}
