// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ad

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/azuread_service_principal_password.html.markdown.
type ServicePrincipalPassword struct {
	s *pulumi.ResourceState
}

// NewServicePrincipalPassword registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalPassword(ctx *pulumi.Context,
	name string, args *ServicePrincipalPasswordArgs, opts ...pulumi.ResourceOpt) (*ServicePrincipalPassword, error) {
	if args == nil || args.EndDate == nil {
		return nil, errors.New("missing required argument 'EndDate'")
	}
	if args == nil || args.ServicePrincipalId == nil {
		return nil, errors.New("missing required argument 'ServicePrincipalId'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["endDate"] = nil
		inputs["keyId"] = nil
		inputs["servicePrincipalId"] = nil
		inputs["startDate"] = nil
		inputs["value"] = nil
	} else {
		inputs["endDate"] = args.EndDate
		inputs["keyId"] = args.KeyId
		inputs["servicePrincipalId"] = args.ServicePrincipalId
		inputs["startDate"] = args.StartDate
		inputs["value"] = args.Value
	}
	s, err := ctx.RegisterResource("azure:ad/servicePrincipalPassword:ServicePrincipalPassword", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServicePrincipalPassword{s: s}, nil
}

// GetServicePrincipalPassword gets an existing ServicePrincipalPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalPassword(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServicePrincipalPasswordState, opts ...pulumi.ResourceOpt) (*ServicePrincipalPassword, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["endDate"] = state.EndDate
		inputs["keyId"] = state.KeyId
		inputs["servicePrincipalId"] = state.ServicePrincipalId
		inputs["startDate"] = state.StartDate
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("azure:ad/servicePrincipalPassword:ServicePrincipalPassword", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServicePrincipalPassword{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ServicePrincipalPassword) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ServicePrincipalPassword) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
func (r *ServicePrincipalPassword) EndDate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["endDate"])
}

// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
func (r *ServicePrincipalPassword) KeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyId"])
}

// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
func (r *ServicePrincipalPassword) ServicePrincipalId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["servicePrincipalId"])
}

// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
func (r *ServicePrincipalPassword) StartDate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["startDate"])
}

// The Password for this Service Principal.
func (r *ServicePrincipalPassword) Value() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering ServicePrincipalPassword resources.
type ServicePrincipalPasswordState struct {
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate interface{}
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId interface{}
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId interface{}
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate interface{}
	// The Password for this Service Principal.
	Value interface{}
}

// The set of arguments for constructing a ServicePrincipalPassword resource.
type ServicePrincipalPasswordArgs struct {
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate interface{}
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId interface{}
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId interface{}
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate interface{}
	// The Password for this Service Principal.
	Value interface{}
}
