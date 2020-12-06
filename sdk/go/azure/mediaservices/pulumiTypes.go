// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediaservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AccountIdentity struct {
	// The Principal ID associated with this Managed Service Identity.
	PrincipalId *string `pulumi:"principalId"`
	// The Tenant ID associated with this Managed Service Identity.
	TenantId *string `pulumi:"tenantId"`
	// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible value is  `SystemAssigned`.
	Type *string `pulumi:"type"`
}

// AccountIdentityInput is an input type that accepts AccountIdentityArgs and AccountIdentityOutput values.
// You can construct a concrete instance of `AccountIdentityInput` via:
//
//          AccountIdentityArgs{...}
type AccountIdentityInput interface {
	pulumi.Input

	ToAccountIdentityOutput() AccountIdentityOutput
	ToAccountIdentityOutputWithContext(context.Context) AccountIdentityOutput
}

type AccountIdentityArgs struct {
	// The Principal ID associated with this Managed Service Identity.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The Tenant ID associated with this Managed Service Identity.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible value is  `SystemAssigned`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (AccountIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentity)(nil)).Elem()
}

func (i AccountIdentityArgs) ToAccountIdentityOutput() AccountIdentityOutput {
	return i.ToAccountIdentityOutputWithContext(context.Background())
}

func (i AccountIdentityArgs) ToAccountIdentityOutputWithContext(ctx context.Context) AccountIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityOutput)
}

func (i AccountIdentityArgs) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return i.ToAccountIdentityPtrOutputWithContext(context.Background())
}

func (i AccountIdentityArgs) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityOutput).ToAccountIdentityPtrOutputWithContext(ctx)
}

// AccountIdentityPtrInput is an input type that accepts AccountIdentityArgs, AccountIdentityPtr and AccountIdentityPtrOutput values.
// You can construct a concrete instance of `AccountIdentityPtrInput` via:
//
//          AccountIdentityArgs{...}
//
//  or:
//
//          nil
type AccountIdentityPtrInput interface {
	pulumi.Input

	ToAccountIdentityPtrOutput() AccountIdentityPtrOutput
	ToAccountIdentityPtrOutputWithContext(context.Context) AccountIdentityPtrOutput
}

type accountIdentityPtrType AccountIdentityArgs

func AccountIdentityPtr(v *AccountIdentityArgs) AccountIdentityPtrInput {
	return (*accountIdentityPtrType)(v)
}

func (*accountIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIdentity)(nil)).Elem()
}

func (i *accountIdentityPtrType) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return i.ToAccountIdentityPtrOutputWithContext(context.Background())
}

func (i *accountIdentityPtrType) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIdentityPtrOutput)
}

type AccountIdentityOutput struct{ *pulumi.OutputState }

func (AccountIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIdentity)(nil)).Elem()
}

func (o AccountIdentityOutput) ToAccountIdentityOutput() AccountIdentityOutput {
	return o
}

func (o AccountIdentityOutput) ToAccountIdentityOutputWithContext(ctx context.Context) AccountIdentityOutput {
	return o
}

func (o AccountIdentityOutput) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return o.ToAccountIdentityPtrOutputWithContext(context.Background())
}

func (o AccountIdentityOutput) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *AccountIdentity {
		return &v
	}).(AccountIdentityPtrOutput)
}

// The Principal ID associated with this Managed Service Identity.
func (o AccountIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The Tenant ID associated with this Managed Service Identity.
func (o AccountIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible value is  `SystemAssigned`.
func (o AccountIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AccountIdentityPtrOutput struct{ *pulumi.OutputState }

func (AccountIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIdentity)(nil)).Elem()
}

func (o AccountIdentityPtrOutput) ToAccountIdentityPtrOutput() AccountIdentityPtrOutput {
	return o
}

func (o AccountIdentityPtrOutput) ToAccountIdentityPtrOutputWithContext(ctx context.Context) AccountIdentityPtrOutput {
	return o
}

func (o AccountIdentityPtrOutput) Elem() AccountIdentityOutput {
	return o.ApplyT(func(v *AccountIdentity) AccountIdentity { return *v }).(AccountIdentityOutput)
}

// The Principal ID associated with this Managed Service Identity.
func (o AccountIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The Tenant ID associated with this Managed Service Identity.
func (o AccountIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible value is  `SystemAssigned`.
func (o AccountIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type AccountStorageAccount struct {
	// Specifies the ID of the Storage Account that will be associated with the Media Services instance.
	Id string `pulumi:"id"`
	// Specifies whether the storage account should be the primary account or not. Defaults to `false`.
	IsPrimary *bool `pulumi:"isPrimary"`
}

// AccountStorageAccountInput is an input type that accepts AccountStorageAccountArgs and AccountStorageAccountOutput values.
// You can construct a concrete instance of `AccountStorageAccountInput` via:
//
//          AccountStorageAccountArgs{...}
type AccountStorageAccountInput interface {
	pulumi.Input

	ToAccountStorageAccountOutput() AccountStorageAccountOutput
	ToAccountStorageAccountOutputWithContext(context.Context) AccountStorageAccountOutput
}

type AccountStorageAccountArgs struct {
	// Specifies the ID of the Storage Account that will be associated with the Media Services instance.
	Id pulumi.StringInput `pulumi:"id"`
	// Specifies whether the storage account should be the primary account or not. Defaults to `false`.
	IsPrimary pulumi.BoolPtrInput `pulumi:"isPrimary"`
}

func (AccountStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountStorageAccount)(nil)).Elem()
}

func (i AccountStorageAccountArgs) ToAccountStorageAccountOutput() AccountStorageAccountOutput {
	return i.ToAccountStorageAccountOutputWithContext(context.Background())
}

func (i AccountStorageAccountArgs) ToAccountStorageAccountOutputWithContext(ctx context.Context) AccountStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountStorageAccountOutput)
}

// AccountStorageAccountArrayInput is an input type that accepts AccountStorageAccountArray and AccountStorageAccountArrayOutput values.
// You can construct a concrete instance of `AccountStorageAccountArrayInput` via:
//
//          AccountStorageAccountArray{ AccountStorageAccountArgs{...} }
type AccountStorageAccountArrayInput interface {
	pulumi.Input

	ToAccountStorageAccountArrayOutput() AccountStorageAccountArrayOutput
	ToAccountStorageAccountArrayOutputWithContext(context.Context) AccountStorageAccountArrayOutput
}

type AccountStorageAccountArray []AccountStorageAccountInput

func (AccountStorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountStorageAccount)(nil)).Elem()
}

func (i AccountStorageAccountArray) ToAccountStorageAccountArrayOutput() AccountStorageAccountArrayOutput {
	return i.ToAccountStorageAccountArrayOutputWithContext(context.Background())
}

func (i AccountStorageAccountArray) ToAccountStorageAccountArrayOutputWithContext(ctx context.Context) AccountStorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountStorageAccountArrayOutput)
}

type AccountStorageAccountOutput struct{ *pulumi.OutputState }

func (AccountStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountStorageAccount)(nil)).Elem()
}

func (o AccountStorageAccountOutput) ToAccountStorageAccountOutput() AccountStorageAccountOutput {
	return o
}

func (o AccountStorageAccountOutput) ToAccountStorageAccountOutputWithContext(ctx context.Context) AccountStorageAccountOutput {
	return o
}

// Specifies the ID of the Storage Account that will be associated with the Media Services instance.
func (o AccountStorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccountStorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies whether the storage account should be the primary account or not. Defaults to `false`.
func (o AccountStorageAccountOutput) IsPrimary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AccountStorageAccount) *bool { return v.IsPrimary }).(pulumi.BoolPtrOutput)
}

type AccountStorageAccountArrayOutput struct{ *pulumi.OutputState }

func (AccountStorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountStorageAccount)(nil)).Elem()
}

func (o AccountStorageAccountArrayOutput) ToAccountStorageAccountArrayOutput() AccountStorageAccountArrayOutput {
	return o
}

func (o AccountStorageAccountArrayOutput) ToAccountStorageAccountArrayOutputWithContext(ctx context.Context) AccountStorageAccountArrayOutput {
	return o
}

func (o AccountStorageAccountArrayOutput) Index(i pulumi.IntInput) AccountStorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountStorageAccount {
		return vs[0].([]AccountStorageAccount)[vs[1].(int)]
	}).(AccountStorageAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountIdentityOutput{})
	pulumi.RegisterOutputType(AccountIdentityPtrOutput{})
	pulumi.RegisterOutputType(AccountStorageAccountOutput{})
	pulumi.RegisterOutputType(AccountStorageAccountArrayOutput{})
}
