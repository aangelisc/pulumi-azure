// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a backend within an API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_backend.html.markdown.
type Backend struct {
	s *pulumi.ResourceState
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOpt) (*Backend, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["credentials"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["protocol"] = nil
		inputs["proxy"] = nil
		inputs["resourceGroupName"] = nil
		inputs["resourceId"] = nil
		inputs["serviceFabricCluster"] = nil
		inputs["title"] = nil
		inputs["tls"] = nil
		inputs["url"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["credentials"] = args.Credentials
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["protocol"] = args.Protocol
		inputs["proxy"] = args.Proxy
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["resourceId"] = args.ResourceId
		inputs["serviceFabricCluster"] = args.ServiceFabricCluster
		inputs["title"] = args.Title
		inputs["tls"] = args.Tls
		inputs["url"] = args.Url
	}
	s, err := ctx.RegisterResource("azure:apimanagement/backend:Backend", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Backend{s: s}, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BackendState, opts ...pulumi.ResourceOpt) (*Backend, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["credentials"] = state.Credentials
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["protocol"] = state.Protocol
		inputs["proxy"] = state.Proxy
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["resourceId"] = state.ResourceId
		inputs["serviceFabricCluster"] = state.ServiceFabricCluster
		inputs["title"] = state.Title
		inputs["tls"] = state.Tls
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("azure:apimanagement/backend:Backend", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Backend{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Backend) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Backend) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
func (r *Backend) ApiManagementName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// A `credentials` block as documented below.
func (r *Backend) Credentials() pulumi.Output {
	return r.s.State["credentials"]
}

// The description of the backend.
func (r *Backend) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The name of the API Management backend. Changing this forces a new resource to be created.
func (r *Backend) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The protocol used by the backend host. Possible values are `http` or `soap`.
func (r *Backend) Protocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocol"])
}

// A `proxy` block as documented below.
func (r *Backend) Proxy() pulumi.Output {
	return r.s.State["proxy"]
}

// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
func (r *Backend) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
func (r *Backend) ResourceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceId"])
}

// A `serviceFabricCluster` block as documented below.
func (r *Backend) ServiceFabricCluster() pulumi.Output {
	return r.s.State["serviceFabricCluster"]
}

// The title of the backend.
func (r *Backend) Title() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["title"])
}

// A `tls` block as documented below.
func (r *Backend) Tls() pulumi.Output {
	return r.s.State["tls"]
}

// The URL of the backend host.
func (r *Backend) Url() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering Backend resources.
type BackendState struct {
	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// A `credentials` block as documented below.
	Credentials interface{}
	// The description of the backend.
	Description interface{}
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name interface{}
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol interface{}
	// A `proxy` block as documented below.
	Proxy interface{}
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId interface{}
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster interface{}
	// The title of the backend.
	Title interface{}
	// A `tls` block as documented below.
	Tls interface{}
	// The URL of the backend host.
	Url interface{}
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// A `credentials` block as documented below.
	Credentials interface{}
	// The description of the backend.
	Description interface{}
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name interface{}
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol interface{}
	// A `proxy` block as documented below.
	Proxy interface{}
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId interface{}
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster interface{}
	// The title of the backend.
	Title interface{}
	// A `tls` block as documented below.
	Tls interface{}
	// The URL of the backend host.
	Url interface{}
}
