// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ChannelDirectLineSite struct {
	// Enables/Disables this site. Enabled by default
	Enabled *bool `pulumi:"enabled"`
	// Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
	EnhancedAuthenticationEnabled *bool `pulumi:"enhancedAuthenticationEnabled"`
	// Id for the site
	Id *string `pulumi:"id"`
	// Primary key for accessing this site
	Key *string `pulumi:"key"`
	// Secondary key for accessing this site
	Key2 *string `pulumi:"key2"`
	// The name of the site
	Name string `pulumi:"name"`
	// This field is required when `isSecureSiteEnabled` is enabled. Determines which origins can establish a Directline conversation for this site.
	TrustedOrigins []string `pulumi:"trustedOrigins"`
	// Enables v1 of the Directline protocol for this site. Enabled by default
	V1Allowed *bool `pulumi:"v1Allowed"`
	// Enables v3 of the Directline protocol for this site. Enabled by default
	V3Allowed *bool `pulumi:"v3Allowed"`
}

// ChannelDirectLineSiteInput is an input type that accepts ChannelDirectLineSiteArgs and ChannelDirectLineSiteOutput values.
// You can construct a concrete instance of `ChannelDirectLineSiteInput` via:
//
//          ChannelDirectLineSiteArgs{...}
type ChannelDirectLineSiteInput interface {
	pulumi.Input

	ToChannelDirectLineSiteOutput() ChannelDirectLineSiteOutput
	ToChannelDirectLineSiteOutputWithContext(context.Context) ChannelDirectLineSiteOutput
}

type ChannelDirectLineSiteArgs struct {
	// Enables/Disables this site. Enabled by default
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
	EnhancedAuthenticationEnabled pulumi.BoolPtrInput `pulumi:"enhancedAuthenticationEnabled"`
	// Id for the site
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Primary key for accessing this site
	Key pulumi.StringPtrInput `pulumi:"key"`
	// Secondary key for accessing this site
	Key2 pulumi.StringPtrInput `pulumi:"key2"`
	// The name of the site
	Name pulumi.StringInput `pulumi:"name"`
	// This field is required when `isSecureSiteEnabled` is enabled. Determines which origins can establish a Directline conversation for this site.
	TrustedOrigins pulumi.StringArrayInput `pulumi:"trustedOrigins"`
	// Enables v1 of the Directline protocol for this site. Enabled by default
	V1Allowed pulumi.BoolPtrInput `pulumi:"v1Allowed"`
	// Enables v3 of the Directline protocol for this site. Enabled by default
	V3Allowed pulumi.BoolPtrInput `pulumi:"v3Allowed"`
}

func (ChannelDirectLineSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelDirectLineSite)(nil)).Elem()
}

func (i ChannelDirectLineSiteArgs) ToChannelDirectLineSiteOutput() ChannelDirectLineSiteOutput {
	return i.ToChannelDirectLineSiteOutputWithContext(context.Background())
}

func (i ChannelDirectLineSiteArgs) ToChannelDirectLineSiteOutputWithContext(ctx context.Context) ChannelDirectLineSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSiteOutput)
}

// ChannelDirectLineSiteArrayInput is an input type that accepts ChannelDirectLineSiteArray and ChannelDirectLineSiteArrayOutput values.
// You can construct a concrete instance of `ChannelDirectLineSiteArrayInput` via:
//
//          ChannelDirectLineSiteArray{ ChannelDirectLineSiteArgs{...} }
type ChannelDirectLineSiteArrayInput interface {
	pulumi.Input

	ToChannelDirectLineSiteArrayOutput() ChannelDirectLineSiteArrayOutput
	ToChannelDirectLineSiteArrayOutputWithContext(context.Context) ChannelDirectLineSiteArrayOutput
}

type ChannelDirectLineSiteArray []ChannelDirectLineSiteInput

func (ChannelDirectLineSiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelDirectLineSite)(nil)).Elem()
}

func (i ChannelDirectLineSiteArray) ToChannelDirectLineSiteArrayOutput() ChannelDirectLineSiteArrayOutput {
	return i.ToChannelDirectLineSiteArrayOutputWithContext(context.Background())
}

func (i ChannelDirectLineSiteArray) ToChannelDirectLineSiteArrayOutputWithContext(ctx context.Context) ChannelDirectLineSiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSiteArrayOutput)
}

type ChannelDirectLineSiteOutput struct{ *pulumi.OutputState }

func (ChannelDirectLineSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelDirectLineSite)(nil)).Elem()
}

func (o ChannelDirectLineSiteOutput) ToChannelDirectLineSiteOutput() ChannelDirectLineSiteOutput {
	return o
}

func (o ChannelDirectLineSiteOutput) ToChannelDirectLineSiteOutputWithContext(ctx context.Context) ChannelDirectLineSiteOutput {
	return o
}

// Enables/Disables this site. Enabled by default
func (o ChannelDirectLineSiteOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
func (o ChannelDirectLineSiteOutput) EnhancedAuthenticationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *bool { return v.EnhancedAuthenticationEnabled }).(pulumi.BoolPtrOutput)
}

// Id for the site
func (o ChannelDirectLineSiteOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Primary key for accessing this site
func (o ChannelDirectLineSiteOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// Secondary key for accessing this site
func (o ChannelDirectLineSiteOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

// The name of the site
func (o ChannelDirectLineSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) string { return v.Name }).(pulumi.StringOutput)
}

// This field is required when `isSecureSiteEnabled` is enabled. Determines which origins can establish a Directline conversation for this site.
func (o ChannelDirectLineSiteOutput) TrustedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) []string { return v.TrustedOrigins }).(pulumi.StringArrayOutput)
}

// Enables v1 of the Directline protocol for this site. Enabled by default
func (o ChannelDirectLineSiteOutput) V1Allowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *bool { return v.V1Allowed }).(pulumi.BoolPtrOutput)
}

// Enables v3 of the Directline protocol for this site. Enabled by default
func (o ChannelDirectLineSiteOutput) V3Allowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSite) *bool { return v.V3Allowed }).(pulumi.BoolPtrOutput)
}

type ChannelDirectLineSiteArrayOutput struct{ *pulumi.OutputState }

func (ChannelDirectLineSiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelDirectLineSite)(nil)).Elem()
}

func (o ChannelDirectLineSiteArrayOutput) ToChannelDirectLineSiteArrayOutput() ChannelDirectLineSiteArrayOutput {
	return o
}

func (o ChannelDirectLineSiteArrayOutput) ToChannelDirectLineSiteArrayOutputWithContext(ctx context.Context) ChannelDirectLineSiteArrayOutput {
	return o
}

func (o ChannelDirectLineSiteArrayOutput) Index(i pulumi.IntInput) ChannelDirectLineSiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelDirectLineSite {
		return vs[0].([]ChannelDirectLineSite)[vs[1].(int)]
	}).(ChannelDirectLineSiteOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelDirectLineSiteOutput{})
	pulumi.RegisterOutputType(ChannelDirectLineSiteArrayOutput{})
}
