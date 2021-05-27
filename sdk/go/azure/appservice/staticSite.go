// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service Static Site.
//
// ->**NOTE**: After the Static Site is provisioned, you'll need to associate your target repository, which contains your web app, to the Static Site, by following the [Azure Static Site document](https://docs.microsoft.com/en-us/azure/static-web-apps/github-actions-workflow).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appservice.NewStaticSite(ctx, "example", &appservice.StaticSiteArgs{
// 			Location:          pulumi.String("West Europe"),
// 			ResourceGroupName: pulumi.String("example"),
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
// Static Web Apps can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/staticSite:StaticSite example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1
// ```
type StaticSite struct {
	pulumi.CustomResourceState

	// The API key of this Static Web App, which is used for later interacting with this Static Web App from other clients, e.g. Github Action.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The default host name of the Static Web App.
	DefaultHostName pulumi.StringOutput `pulumi:"defaultHostName"`
	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	ResourceGroupName pulumi.StringOutput    `pulumi:"resourceGroupName"`
	SkuSize           pulumi.StringPtrOutput `pulumi:"skuSize"`
	SkuTier           pulumi.StringPtrOutput `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewStaticSite registers a new resource with the given unique name, arguments, and options.
func NewStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource StaticSite
	err := ctx.RegisterResource("azure:appservice/staticSite:StaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSite gets an existing StaticSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteState, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	var resource StaticSite
	err := ctx.ReadResource("azure:appservice/staticSite:StaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSite resources.
type staticSiteState struct {
	// The API key of this Static Web App, which is used for later interacting with this Static Web App from other clients, e.g. Github Action.
	ApiKey *string `pulumi:"apiKey"`
	// The default host name of the Static Web App.
	DefaultHostName *string `pulumi:"defaultHostName"`
	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SkuSize           *string `pulumi:"skuSize"`
	SkuTier           *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type StaticSiteState struct {
	// The API key of this Static Web App, which is used for later interacting with this Static Web App from other clients, e.g. Github Action.
	ApiKey pulumi.StringPtrInput
	// The default host name of the Static Web App.
	DefaultHostName pulumi.StringPtrInput
	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	ResourceGroupName pulumi.StringPtrInput
	SkuSize           pulumi.StringPtrInput
	SkuTier           pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (StaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteState)(nil)).Elem()
}

type staticSiteArgs struct {
	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SkuSize           *string `pulumi:"skuSize"`
	SkuTier           *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StaticSite resource.
type StaticSiteArgs struct {
	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	ResourceGroupName pulumi.StringInput
	SkuSize           pulumi.StringPtrInput
	SkuTier           pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (StaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteArgs)(nil)).Elem()
}

type StaticSiteInput interface {
	pulumi.Input

	ToStaticSiteOutput() StaticSiteOutput
	ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput
}

func (*StaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSite)(nil))
}

func (i *StaticSite) ToStaticSiteOutput() StaticSiteOutput {
	return i.ToStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSite) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteOutput)
}

func (i *StaticSite) ToStaticSitePtrOutput() StaticSitePtrOutput {
	return i.ToStaticSitePtrOutputWithContext(context.Background())
}

func (i *StaticSite) ToStaticSitePtrOutputWithContext(ctx context.Context) StaticSitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSitePtrOutput)
}

type StaticSitePtrInput interface {
	pulumi.Input

	ToStaticSitePtrOutput() StaticSitePtrOutput
	ToStaticSitePtrOutputWithContext(ctx context.Context) StaticSitePtrOutput
}

type staticSitePtrType StaticSiteArgs

func (*staticSitePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSite)(nil))
}

func (i *staticSitePtrType) ToStaticSitePtrOutput() StaticSitePtrOutput {
	return i.ToStaticSitePtrOutputWithContext(context.Background())
}

func (i *staticSitePtrType) ToStaticSitePtrOutputWithContext(ctx context.Context) StaticSitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSitePtrOutput)
}

// StaticSiteArrayInput is an input type that accepts StaticSiteArray and StaticSiteArrayOutput values.
// You can construct a concrete instance of `StaticSiteArrayInput` via:
//
//          StaticSiteArray{ StaticSiteArgs{...} }
type StaticSiteArrayInput interface {
	pulumi.Input

	ToStaticSiteArrayOutput() StaticSiteArrayOutput
	ToStaticSiteArrayOutputWithContext(context.Context) StaticSiteArrayOutput
}

type StaticSiteArray []StaticSiteInput

func (StaticSiteArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*StaticSite)(nil))
}

func (i StaticSiteArray) ToStaticSiteArrayOutput() StaticSiteArrayOutput {
	return i.ToStaticSiteArrayOutputWithContext(context.Background())
}

func (i StaticSiteArray) ToStaticSiteArrayOutputWithContext(ctx context.Context) StaticSiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteArrayOutput)
}

// StaticSiteMapInput is an input type that accepts StaticSiteMap and StaticSiteMapOutput values.
// You can construct a concrete instance of `StaticSiteMapInput` via:
//
//          StaticSiteMap{ "key": StaticSiteArgs{...} }
type StaticSiteMapInput interface {
	pulumi.Input

	ToStaticSiteMapOutput() StaticSiteMapOutput
	ToStaticSiteMapOutputWithContext(context.Context) StaticSiteMapOutput
}

type StaticSiteMap map[string]StaticSiteInput

func (StaticSiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*StaticSite)(nil))
}

func (i StaticSiteMap) ToStaticSiteMapOutput() StaticSiteMapOutput {
	return i.ToStaticSiteMapOutputWithContext(context.Background())
}

func (i StaticSiteMap) ToStaticSiteMapOutputWithContext(ctx context.Context) StaticSiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteMapOutput)
}

type StaticSiteOutput struct {
	*pulumi.OutputState
}

func (StaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSite)(nil))
}

func (o StaticSiteOutput) ToStaticSiteOutput() StaticSiteOutput {
	return o
}

func (o StaticSiteOutput) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return o
}

func (o StaticSiteOutput) ToStaticSitePtrOutput() StaticSitePtrOutput {
	return o.ToStaticSitePtrOutputWithContext(context.Background())
}

func (o StaticSiteOutput) ToStaticSitePtrOutputWithContext(ctx context.Context) StaticSitePtrOutput {
	return o.ApplyT(func(v StaticSite) *StaticSite {
		return &v
	}).(StaticSitePtrOutput)
}

type StaticSitePtrOutput struct {
	*pulumi.OutputState
}

func (StaticSitePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSite)(nil))
}

func (o StaticSitePtrOutput) ToStaticSitePtrOutput() StaticSitePtrOutput {
	return o
}

func (o StaticSitePtrOutput) ToStaticSitePtrOutputWithContext(ctx context.Context) StaticSitePtrOutput {
	return o
}

type StaticSiteArrayOutput struct{ *pulumi.OutputState }

func (StaticSiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSite)(nil))
}

func (o StaticSiteArrayOutput) ToStaticSiteArrayOutput() StaticSiteArrayOutput {
	return o
}

func (o StaticSiteArrayOutput) ToStaticSiteArrayOutputWithContext(ctx context.Context) StaticSiteArrayOutput {
	return o
}

func (o StaticSiteArrayOutput) Index(i pulumi.IntInput) StaticSiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticSite {
		return vs[0].([]StaticSite)[vs[1].(int)]
	}).(StaticSiteOutput)
}

type StaticSiteMapOutput struct{ *pulumi.OutputState }

func (StaticSiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StaticSite)(nil))
}

func (o StaticSiteMapOutput) ToStaticSiteMapOutput() StaticSiteMapOutput {
	return o
}

func (o StaticSiteMapOutput) ToStaticSiteMapOutputWithContext(ctx context.Context) StaticSiteMapOutput {
	return o
}

func (o StaticSiteMapOutput) MapIndex(k pulumi.StringInput) StaticSiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StaticSite {
		return vs[0].(map[string]StaticSite)[vs[1].(string)]
	}).(StaticSiteOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteOutput{})
	pulumi.RegisterOutputType(StaticSitePtrOutput{})
	pulumi.RegisterOutputType(StaticSiteArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteMapOutput{})
}
