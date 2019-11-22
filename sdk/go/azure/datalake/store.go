// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Data Lake Store.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_lake_store.html.markdown.
type Store struct {
	s *pulumi.ResourceState
}

// NewStore registers a new resource with the given unique name, arguments, and options.
func NewStore(ctx *pulumi.Context,
	name string, args *StoreArgs, opts ...pulumi.ResourceOpt) (*Store, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["encryptionState"] = nil
		inputs["encryptionType"] = nil
		inputs["firewallAllowAzureIps"] = nil
		inputs["firewallState"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["tier"] = nil
	} else {
		inputs["encryptionState"] = args.EncryptionState
		inputs["encryptionType"] = args.EncryptionType
		inputs["firewallAllowAzureIps"] = args.FirewallAllowAzureIps
		inputs["firewallState"] = args.FirewallState
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["tier"] = args.Tier
	}
	inputs["endpoint"] = nil
	s, err := ctx.RegisterResource("azure:datalake/store:Store", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Store{s: s}, nil
}

// GetStore gets an existing Store resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStore(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StoreState, opts ...pulumi.ResourceOpt) (*Store, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["encryptionState"] = state.EncryptionState
		inputs["encryptionType"] = state.EncryptionType
		inputs["endpoint"] = state.Endpoint
		inputs["firewallAllowAzureIps"] = state.FirewallAllowAzureIps
		inputs["firewallState"] = state.FirewallState
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["tier"] = state.Tier
	}
	s, err := ctx.ReadResource("azure:datalake/store:Store", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Store{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Store) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Store) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
func (r *Store) EncryptionState() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["encryptionState"])
}

// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
func (r *Store) EncryptionType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["encryptionType"])
}

// The Endpoint for the Data Lake Store.
func (r *Store) Endpoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["endpoint"])
}

// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
func (r *Store) FirewallAllowAzureIps() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["firewallAllowAzureIps"])
}

// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
func (r *Store) FirewallState() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["firewallState"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *Store) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
func (r *Store) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the Data Lake Store.
func (r *Store) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *Store) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
func (r *Store) Tier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tier"])
}

// Input properties used for looking up and filtering Store resources.
type StoreState struct {
	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState interface{}
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType interface{}
	// The Endpoint for the Data Lake Store.
	Endpoint interface{}
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps interface{}
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name interface{}
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier interface{}
}

// The set of arguments for constructing a Store resource.
type StoreArgs struct {
	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState interface{}
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType interface{}
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps interface{}
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name interface{}
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier interface{}
}
