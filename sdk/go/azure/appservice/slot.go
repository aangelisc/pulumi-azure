// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an App Service Slot (within an App Service).
//
// > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_service_slot.html.markdown.
type Slot struct {
	pulumi.CustomResourceState

	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringOutput `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapOutput `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings SlotAuthSettingsOutput `pulumi:"authSettings"`
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolOutput `pulumi:"clientAffinityEnabled"`
	// An `connectionString` block as defined below.
	ConnectionStrings SlotConnectionStringArrayOutput `pulumi:"connectionStrings"`
	// The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringOutput `pulumi:"defaultSiteHostname"`
	// Is the App Service Slot Enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity SlotIdentityOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	Logs SlotLogsOutput `pulumi:"logs"`
	// The name of the Connection String.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig SlotSiteConfigOutput `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials SlotSiteCredentialArrayOutput `pulumi:"siteCredentials"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSlot registers a new resource with the given unique name, arguments, and options.
func NewSlot(ctx *pulumi.Context,
	name string, args *SlotArgs, opts ...pulumi.ResourceOption) (*Slot, error) {
	if args == nil || args.AppServiceName == nil {
		return nil, errors.New("missing required argument 'AppServiceName'")
	}
	if args == nil || args.AppServicePlanId == nil {
		return nil, errors.New("missing required argument 'AppServicePlanId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &SlotArgs{}
	}
	var resource Slot
	err := ctx.RegisterResource("azure:appservice/slot:Slot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlot gets an existing Slot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SlotState, opts ...pulumi.ResourceOption) (*Slot, error) {
	var resource Slot
	err := ctx.ReadResource("azure:appservice/slot:Slot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slot resources.
type slotState struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *SlotAuthSettings `pulumi:"authSettings"`
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// An `connectionString` block as defined below.
	ConnectionStrings []SlotConnectionString `pulumi:"connectionStrings"`
	// The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
	DefaultSiteHostname *string `pulumi:"defaultSiteHostname"`
	// Is the App Service Slot Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity *SlotIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	Logs *SlotLogs `pulumi:"logs"`
	// The name of the Connection String.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *SlotSiteConfig `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials []SlotSiteCredential `pulumi:"siteCredentials"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SlotState struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringPtrInput
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringPtrInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings SlotAuthSettingsPtrInput
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// An `connectionString` block as defined below.
	ConnectionStrings SlotConnectionStringArrayInput
	// The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringPtrInput
	// Is the App Service Slot Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// A Managed Service Identity block as defined below.
	Identity SlotIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	Logs SlotLogsPtrInput
	// The name of the Connection String.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName pulumi.StringPtrInput
	// A `siteConfig` object as defined below.
	SiteConfig SlotSiteConfigPtrInput
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials SlotSiteCredentialArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*slotState)(nil)).Elem()
}

type slotArgs struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *SlotAuthSettings `pulumi:"authSettings"`
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// An `connectionString` block as defined below.
	ConnectionStrings []SlotConnectionString `pulumi:"connectionStrings"`
	// Is the App Service Slot Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity *SlotIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	Logs *SlotLogs `pulumi:"logs"`
	// The name of the Connection String.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *SlotSiteConfig `pulumi:"siteConfig"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Slot resource.
type SlotArgs struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringInput
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings SlotAuthSettingsPtrInput
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// An `connectionString` block as defined below.
	ConnectionStrings SlotConnectionStringArrayInput
	// Is the App Service Slot Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// A Managed Service Identity block as defined below.
	Identity SlotIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	Logs SlotLogsPtrInput
	// The name of the Connection String.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName pulumi.StringInput
	// A `siteConfig` object as defined below.
	SiteConfig SlotSiteConfigPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*slotArgs)(nil)).Elem()
}

