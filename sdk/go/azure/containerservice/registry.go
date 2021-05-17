// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Container Registry.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = containerservice.NewRegistry(ctx, "acr", &containerservice.RegistryArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			Sku:               pulumi.String("Premium"),
// 			AdminEnabled:      pulumi.Bool(false),
// 			GeoreplicationLocations: pulumi.StringArray{
// 				pulumi.String("East US"),
// 				pulumi.String("West Europe"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Encryption)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleUserAssignedIdentity, err := authorization.NewUserAssignedIdentity(ctx, "exampleUserAssignedIdentity", &authorization.UserAssignedIdentityArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:          pulumi.Any(azurerm_resource_group.Example.Location),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKey, err := keyvault.LookupKey(ctx, &keyvault.LookupKeyArgs{
// 			Name:       "super-secret",
// 			KeyVaultId: data.Azurerm_key_vault.Existing.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = containerservice.NewRegistry(ctx, "acr", &containerservice.RegistryArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			Sku:               pulumi.String("Premium"),
// 			Identity: &containerservice.RegistryIdentityArgs{
// 				Type: pulumi.String("UserAssigned"),
// 				IdentityIds: pulumi.StringArray{
// 					exampleUserAssignedIdentity.ID(),
// 				},
// 			},
// 			Encryption: &containerservice.RegistryEncryptionArgs{
// 				Enabled:          pulumi.Bool(true),
// 				KeyVaultKeyId:    pulumi.String(exampleKey.Id),
// 				IdentityClientId: exampleUserAssignedIdentity.ClientId,
// 			},
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
// Container Registries can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:containerservice/registry:Registry example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1
// ```
type Registry struct {
	pulumi.CustomResourceState

	// Specifies whether the admin user is enabled. Defaults to `false`.
	AdminEnabled pulumi.BoolPtrOutput `pulumi:"adminEnabled"`
	// The Password associated with the Container Registry Admin account - if the admin account is enabled.
	AdminPassword pulumi.StringOutput `pulumi:"adminPassword"`
	// The Username associated with the Container Registry Admin account - if the admin account is enabled.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// An `encryption` block as documented below.
	Encryption RegistryEncryptionOutput `pulumi:"encryption"`
	// A list of Azure locations where the container registry should be geo-replicated.
	//
	// Deprecated: Deprecated in favour of `georeplications`
	GeoreplicationLocations pulumi.StringArrayOutput `pulumi:"georeplicationLocations"`
	// A `georeplications` block as documented below.
	Georeplications RegistryGeoreplicationArrayOutput `pulumi:"georeplications"`
	// An `identity` block as documented below.
	Identity RegistryIdentityOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer pulumi.StringOutput `pulumi:"loginServer"`
	// Specifies the name of the Container Registry. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkRuleSet` block as documented below.
	NetworkRuleSet RegistryNetworkRuleSetOutput `pulumi:"networkRuleSet"`
	// Whether public network access is allowed for the container registry. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
	QuarantinePolicyEnabled pulumi.BoolPtrOutput `pulumi:"quarantinePolicyEnabled"`
	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy RegistryRetentionPolicyOutput `pulumi:"retentionPolicy"`
	// The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `trustPolicy` block as documented below.
	TrustPolicy RegistryTrustPolicyOutput `pulumi:"trustPolicy"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Registry
	err := ctx.RegisterResource("azure:containerservice/registry:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azure:containerservice/registry:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// Specifies whether the admin user is enabled. Defaults to `false`.
	AdminEnabled *bool `pulumi:"adminEnabled"`
	// The Password associated with the Container Registry Admin account - if the admin account is enabled.
	AdminPassword *string `pulumi:"adminPassword"`
	// The Username associated with the Container Registry Admin account - if the admin account is enabled.
	AdminUsername *string `pulumi:"adminUsername"`
	// An `encryption` block as documented below.
	Encryption *RegistryEncryption `pulumi:"encryption"`
	// A list of Azure locations where the container registry should be geo-replicated.
	//
	// Deprecated: Deprecated in favour of `georeplications`
	GeoreplicationLocations []string `pulumi:"georeplicationLocations"`
	// A `georeplications` block as documented below.
	Georeplications []RegistryGeoreplication `pulumi:"georeplications"`
	// An `identity` block as documented below.
	Identity *RegistryIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer *string `pulumi:"loginServer"`
	// Specifies the name of the Container Registry. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkRuleSet` block as documented below.
	NetworkRuleSet *RegistryNetworkRuleSet `pulumi:"networkRuleSet"`
	// Whether public network access is allowed for the container registry. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
	QuarantinePolicyEnabled *bool `pulumi:"quarantinePolicyEnabled"`
	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy *RegistryRetentionPolicy `pulumi:"retentionPolicy"`
	// The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
	Sku *string `pulumi:"sku"`
	// The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.  Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `trustPolicy` block as documented below.
	TrustPolicy *RegistryTrustPolicy `pulumi:"trustPolicy"`
}

type RegistryState struct {
	// Specifies whether the admin user is enabled. Defaults to `false`.
	AdminEnabled pulumi.BoolPtrInput
	// The Password associated with the Container Registry Admin account - if the admin account is enabled.
	AdminPassword pulumi.StringPtrInput
	// The Username associated with the Container Registry Admin account - if the admin account is enabled.
	AdminUsername pulumi.StringPtrInput
	// An `encryption` block as documented below.
	Encryption RegistryEncryptionPtrInput
	// A list of Azure locations where the container registry should be geo-replicated.
	//
	// Deprecated: Deprecated in favour of `georeplications`
	GeoreplicationLocations pulumi.StringArrayInput
	// A `georeplications` block as documented below.
	Georeplications RegistryGeoreplicationArrayInput
	// An `identity` block as documented below.
	Identity RegistryIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The URL that can be used to log into the container registry.
	LoginServer pulumi.StringPtrInput
	// Specifies the name of the Container Registry. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkRuleSet` block as documented below.
	NetworkRuleSet RegistryNetworkRuleSetPtrInput
	// Whether public network access is allowed for the container registry. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
	QuarantinePolicyEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `retentionPolicy` block as documented below.
	RetentionPolicy RegistryRetentionPolicyPtrInput
	// The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
	Sku pulumi.StringPtrInput
	// The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `trustPolicy` block as documented below.
	TrustPolicy RegistryTrustPolicyPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// Specifies whether the admin user is enabled. Defaults to `false`.
	AdminEnabled *bool `pulumi:"adminEnabled"`
	// An `encryption` block as documented below.
	Encryption *RegistryEncryption `pulumi:"encryption"`
	// A list of Azure locations where the container registry should be geo-replicated.
	//
	// Deprecated: Deprecated in favour of `georeplications`
	GeoreplicationLocations []string `pulumi:"georeplicationLocations"`
	// A `georeplications` block as documented below.
	Georeplications []RegistryGeoreplication `pulumi:"georeplications"`
	// An `identity` block as documented below.
	Identity *RegistryIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Registry. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkRuleSet` block as documented below.
	NetworkRuleSet *RegistryNetworkRuleSet `pulumi:"networkRuleSet"`
	// Whether public network access is allowed for the container registry. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
	QuarantinePolicyEnabled *bool `pulumi:"quarantinePolicyEnabled"`
	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy *RegistryRetentionPolicy `pulumi:"retentionPolicy"`
	// The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
	Sku *string `pulumi:"sku"`
	// The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.  Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `trustPolicy` block as documented below.
	TrustPolicy *RegistryTrustPolicy `pulumi:"trustPolicy"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// Specifies whether the admin user is enabled. Defaults to `false`.
	AdminEnabled pulumi.BoolPtrInput
	// An `encryption` block as documented below.
	Encryption RegistryEncryptionPtrInput
	// A list of Azure locations where the container registry should be geo-replicated.
	//
	// Deprecated: Deprecated in favour of `georeplications`
	GeoreplicationLocations pulumi.StringArrayInput
	// A `georeplications` block as documented below.
	Georeplications RegistryGeoreplicationArrayInput
	// An `identity` block as documented below.
	Identity RegistryIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Registry. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkRuleSet` block as documented below.
	NetworkRuleSet RegistryNetworkRuleSetPtrInput
	// Whether public network access is allowed for the container registry. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
	QuarantinePolicyEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `retentionPolicy` block as documented below.
	RetentionPolicy RegistryRetentionPolicyPtrInput
	// The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
	Sku pulumi.StringPtrInput
	// The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `trustPolicy` block as documented below.
	TrustPolicy RegistryTrustPolicyPtrInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil))
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

func (i *Registry) ToRegistryPtrOutput() RegistryPtrOutput {
	return i.ToRegistryPtrOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPtrOutput)
}

type RegistryPtrInput interface {
	pulumi.Input

	ToRegistryPtrOutput() RegistryPtrOutput
	ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput
}

type registryPtrType RegistryArgs

func (*registryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil))
}

func (i *registryPtrType) ToRegistryPtrOutput() RegistryPtrOutput {
	return i.ToRegistryPtrOutputWithContext(context.Background())
}

func (i *registryPtrType) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPtrOutput)
}

// RegistryArrayInput is an input type that accepts RegistryArray and RegistryArrayOutput values.
// You can construct a concrete instance of `RegistryArrayInput` via:
//
//          RegistryArray{ RegistryArgs{...} }
type RegistryArrayInput interface {
	pulumi.Input

	ToRegistryArrayOutput() RegistryArrayOutput
	ToRegistryArrayOutputWithContext(context.Context) RegistryArrayOutput
}

type RegistryArray []RegistryInput

func (RegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Registry)(nil))
}

func (i RegistryArray) ToRegistryArrayOutput() RegistryArrayOutput {
	return i.ToRegistryArrayOutputWithContext(context.Background())
}

func (i RegistryArray) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryArrayOutput)
}

// RegistryMapInput is an input type that accepts RegistryMap and RegistryMapOutput values.
// You can construct a concrete instance of `RegistryMapInput` via:
//
//          RegistryMap{ "key": RegistryArgs{...} }
type RegistryMapInput interface {
	pulumi.Input

	ToRegistryMapOutput() RegistryMapOutput
	ToRegistryMapOutputWithContext(context.Context) RegistryMapOutput
}

type RegistryMap map[string]RegistryInput

func (RegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Registry)(nil))
}

func (i RegistryMap) ToRegistryMapOutput() RegistryMapOutput {
	return i.ToRegistryMapOutputWithContext(context.Background())
}

func (i RegistryMap) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMapOutput)
}

type RegistryOutput struct {
	*pulumi.OutputState
}

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil))
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryPtrOutput() RegistryPtrOutput {
	return o.ToRegistryPtrOutputWithContext(context.Background())
}

func (o RegistryOutput) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return o.ApplyT(func(v Registry) *Registry {
		return &v
	}).(RegistryPtrOutput)
}

type RegistryPtrOutput struct {
	*pulumi.OutputState
}

func (RegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil))
}

func (o RegistryPtrOutput) ToRegistryPtrOutput() RegistryPtrOutput {
	return o
}

func (o RegistryPtrOutput) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return o
}

type RegistryArrayOutput struct{ *pulumi.OutputState }

func (RegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Registry)(nil))
}

func (o RegistryArrayOutput) ToRegistryArrayOutput() RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) Index(i pulumi.IntInput) RegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Registry {
		return vs[0].([]Registry)[vs[1].(int)]
	}).(RegistryOutput)
}

type RegistryMapOutput struct{ *pulumi.OutputState }

func (RegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Registry)(nil))
}

func (o RegistryMapOutput) ToRegistryMapOutput() RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) MapIndex(k pulumi.StringInput) RegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Registry {
		return vs[0].(map[string]Registry)[vs[1].(string)]
	}).(RegistryOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryOutput{})
	pulumi.RegisterOutputType(RegistryPtrOutput{})
	pulumi.RegisterOutputType(RegistryArrayOutput{})
	pulumi.RegisterOutputType(RegistryMapOutput{})
}
