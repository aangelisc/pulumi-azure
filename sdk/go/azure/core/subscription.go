// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Alias for a Subscription - which adds an Alias to an existing Subscription, allowing it to be managed in the provider - or create a new Subscription with a new Alias.
//
// > **NOTE:** Destroying a Subscription controlled by this resource will place the Subscription into a cancelled state. It is possible to re-activate a subscription within 90-days of cancellation, after which time the Subscription is irrevocably deleted, and the Subscription ID cannot be re-used. For further information see [here](https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/cancel-azure-subscription#what-happens-after-subscription-cancellation). Users can optionally delete a Subscription once 72 hours have passed, however, this functionality is not suitable for this provider. A `Deleted` subscription cannot be reactivated.
//
// > **NOTE:** It is not possible to destroy (cancel) a subscription if it contains resources. If resources are present that are not managed by the provider then these will need to be removed before the Subscription can be destroyed.
//
// > **NOTE:** Azure supports Multiple Aliases per Subscription, however, to reliably manage this resource in this provider only a single Alias is supported.
//
// ## Example Usage
// ### Creating A New Alias And Subscription For An Enrollment Account
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/billing"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleEnrollmentAccountScope, err := billing.GetEnrollmentAccountScope(ctx, &billing.GetEnrollmentAccountScopeArgs{
// 			BillingAccountName:    "1234567890",
// 			EnrollmentAccountName: "0123456",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewSubscription(ctx, "exampleSubscription", &core.SubscriptionArgs{
// 			SubscriptionName: pulumi.String("My Example EA Subscription"),
// 			BillingScopeId:   pulumi.String(exampleEnrollmentAccountScope.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Creating A New Alias And Subscription For A Microsoft Customer Account
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/billing"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleMcaAccountScope, err := billing.GetMcaAccountScope(ctx, &billing.GetMcaAccountScopeArgs{
// 			BillingAccountName: "e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
// 			BillingProfileName: "PE2Q-NOIT-BG7-TGB",
// 			InvoiceSectionName: "MTT4-OBS7-PJA-TGB",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewSubscription(ctx, "exampleSubscription", &core.SubscriptionArgs{
// 			SubscriptionName: pulumi.String("My Example MCA Subscription"),
// 			BillingScopeId:   pulumi.String(exampleMcaAccountScope.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Adding An Alias To An Existing Subscription
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.NewSubscription(ctx, "example", &core.SubscriptionArgs{
// 			Alias:            pulumi.String("examplesub"),
// 			SubscriptionId:   pulumi.String("12345678-12234-5678-9012-123456789012"),
// 			SubscriptionName: pulumi.String("My Example Subscription"),
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
// Subscriptions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/subscription:Subscription example "/providers/Microsoft.Subscription/aliases/subscription1"
// ```
//
//  In this scenario, the `subscription_id` property can be completed and this provider will assume control of the existing subscription by creating an Alias. See the `adding an Alias to an existing Subscription` above. This provider requires an alias to correctly manage Subscription resources due to Azure Subscription API design.
type Subscription struct {
	pulumi.CustomResourceState

	// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
	Alias          pulumi.StringOutput    `pulumi:"alias"`
	BillingScopeId pulumi.StringPtrOutput `pulumi:"billingScopeId"`
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName pulumi.StringOutput    `pulumi:"subscriptionName"`
	Tags             pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Tenant to which the subscription belongs.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload pulumi.StringPtrOutput `pulumi:"workload"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	var resource Subscription
	err := ctx.RegisterResource("azure:core/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure:core/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
	Alias          *string `pulumi:"alias"`
	BillingScopeId *string `pulumi:"billingScopeId"`
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName *string           `pulumi:"subscriptionName"`
	Tags             map[string]string `pulumi:"tags"`
	// The ID of the Tenant to which the subscription belongs.
	TenantId *string `pulumi:"tenantId"`
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload *string `pulumi:"workload"`
}

type SubscriptionState struct {
	// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
	Alias          pulumi.StringPtrInput
	BillingScopeId pulumi.StringPtrInput
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId pulumi.StringPtrInput
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
	// The ID of the Tenant to which the subscription belongs.
	TenantId pulumi.StringPtrInput
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
	Alias          *string `pulumi:"alias"`
	BillingScopeId *string `pulumi:"billingScopeId"`
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload *string `pulumi:"workload"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
	Alias          pulumi.StringPtrInput
	BillingScopeId pulumi.StringPtrInput
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId pulumi.StringPtrInput
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName pulumi.StringInput
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload pulumi.StringPtrInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

func (i *Subscription) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return i.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPtrOutput)
}

type SubscriptionPtrInput interface {
	pulumi.Input

	ToSubscriptionPtrOutput() SubscriptionPtrOutput
	ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput
}

type subscriptionPtrType SubscriptionArgs

func (*subscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil))
}

func (i *subscriptionPtrType) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return i.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (i *subscriptionPtrType) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPtrOutput)
}

// SubscriptionArrayInput is an input type that accepts SubscriptionArray and SubscriptionArrayOutput values.
// You can construct a concrete instance of `SubscriptionArrayInput` via:
//
//          SubscriptionArray{ SubscriptionArgs{...} }
type SubscriptionArrayInput interface {
	pulumi.Input

	ToSubscriptionArrayOutput() SubscriptionArrayOutput
	ToSubscriptionArrayOutputWithContext(context.Context) SubscriptionArrayOutput
}

type SubscriptionArray []SubscriptionInput

func (SubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Subscription)(nil))
}

func (i SubscriptionArray) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return i.ToSubscriptionArrayOutputWithContext(context.Background())
}

func (i SubscriptionArray) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionArrayOutput)
}

// SubscriptionMapInput is an input type that accepts SubscriptionMap and SubscriptionMapOutput values.
// You can construct a concrete instance of `SubscriptionMapInput` via:
//
//          SubscriptionMap{ "key": SubscriptionArgs{...} }
type SubscriptionMapInput interface {
	pulumi.Input

	ToSubscriptionMapOutput() SubscriptionMapOutput
	ToSubscriptionMapOutputWithContext(context.Context) SubscriptionMapOutput
}

type SubscriptionMap map[string]SubscriptionInput

func (SubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Subscription)(nil))
}

func (i SubscriptionMap) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return i.ToSubscriptionMapOutputWithContext(context.Background())
}

func (i SubscriptionMap) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionMapOutput)
}

type SubscriptionOutput struct {
	*pulumi.OutputState
}

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return o.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (o SubscriptionOutput) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return o.ApplyT(func(v Subscription) *Subscription {
		return &v
	}).(SubscriptionPtrOutput)
}

type SubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (SubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil))
}

func (o SubscriptionPtrOutput) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return o
}

func (o SubscriptionPtrOutput) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return o
}

type SubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subscription)(nil))
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) Index(i pulumi.IntInput) SubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subscription {
		return vs[0].([]Subscription)[vs[1].(int)]
	}).(SubscriptionOutput)
}

type SubscriptionMapOutput struct{ *pulumi.OutputState }

func (SubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Subscription)(nil))
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) MapIndex(k pulumi.StringInput) SubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Subscription {
		return vs[0].(map[string]Subscription)[vs[1].(string)]
	}).(SubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
	pulumi.RegisterOutputType(SubscriptionPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionMapOutput{})
}
