// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package securitycenter

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a resources Advanced Threat Protection setting.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/advanced_threat_protection.html.markdown.
type AdvancedThreatProtection struct {
	pulumi.CustomResourceState

	// Should Advanced Threat Protection be enabled on this resource?
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewAdvancedThreatProtection registers a new resource with the given unique name, arguments, and options.
func NewAdvancedThreatProtection(ctx *pulumi.Context,
	name string, args *AdvancedThreatProtectionArgs, opts ...pulumi.ResourceOption) (*AdvancedThreatProtection, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.TargetResourceId == nil {
		return nil, errors.New("missing required argument 'TargetResourceId'")
	}
	if args == nil {
		args = &AdvancedThreatProtectionArgs{}
	}
	var resource AdvancedThreatProtection
	err := ctx.RegisterResource("azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdvancedThreatProtection gets an existing AdvancedThreatProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdvancedThreatProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdvancedThreatProtectionState, opts ...pulumi.ResourceOption) (*AdvancedThreatProtection, error) {
	var resource AdvancedThreatProtection
	err := ctx.ReadResource("azure:securitycenter/advancedThreatProtection:AdvancedThreatProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdvancedThreatProtection resources.
type advancedThreatProtectionState struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type AdvancedThreatProtectionState struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled pulumi.BoolPtrInput
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (AdvancedThreatProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*advancedThreatProtectionState)(nil)).Elem()
}

type advancedThreatProtectionArgs struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled bool `pulumi:"enabled"`
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a AdvancedThreatProtection resource.
type AdvancedThreatProtectionArgs struct {
	// Should Advanced Threat Protection be enabled on this resource?
	Enabled pulumi.BoolInput
	// The ID of the Azure Resource which to enable Advanced Threat Protection on. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (AdvancedThreatProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*advancedThreatProtectionArgs)(nil)).Elem()
}

