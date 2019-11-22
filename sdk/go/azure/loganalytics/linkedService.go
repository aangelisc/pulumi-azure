// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Links a Log Analytics (formally Operational Insights) Workspace to another resource. The (currently) only linkable service is an Azure Automation Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/log_analytics_linked_service.html.markdown.
type LinkedService struct {
	s *pulumi.ResourceState
}

// NewLinkedService registers a new resource with the given unique name, arguments, and options.
func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOpt) (*LinkedService, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["linkedServiceName"] = nil
		inputs["linkedServiceProperties"] = nil
		inputs["resourceGroupName"] = nil
		inputs["resourceId"] = nil
		inputs["tags"] = nil
		inputs["workspaceName"] = nil
	} else {
		inputs["linkedServiceName"] = args.LinkedServiceName
		inputs["linkedServiceProperties"] = args.LinkedServiceProperties
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["resourceId"] = args.ResourceId
		inputs["tags"] = args.Tags
		inputs["workspaceName"] = args.WorkspaceName
	}
	inputs["name"] = nil
	s, err := ctx.RegisterResource("azure:loganalytics/linkedService:LinkedService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinkedService{s: s}, nil
}

// GetLinkedService gets an existing LinkedService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LinkedServiceState, opts ...pulumi.ResourceOpt) (*LinkedService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["linkedServiceName"] = state.LinkedServiceName
		inputs["linkedServiceProperties"] = state.LinkedServiceProperties
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["resourceId"] = state.ResourceId
		inputs["tags"] = state.Tags
		inputs["workspaceName"] = state.WorkspaceName
	}
	s, err := ctx.ReadResource("azure:loganalytics/linkedService:LinkedService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinkedService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LinkedService) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LinkedService) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
func (r *LinkedService) LinkedServiceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["linkedServiceName"])
}

// A `linkedServiceProperties` block as defined below.
func (r *LinkedService) LinkedServiceProperties() pulumi.Output {
	return r.s.State["linkedServiceProperties"]
}

// The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
func (r *LinkedService) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
func (r *LinkedService) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
func (r *LinkedService) ResourceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceId"])
}

// A mapping of tags to assign to the resource.
func (r *LinkedService) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
func (r *LinkedService) WorkspaceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["workspaceName"])
}

// Input properties used for looking up and filtering LinkedService resources.
type LinkedServiceState struct {
	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName interface{}
	// A `linkedServiceProperties` block as defined below.
	LinkedServiceProperties interface{}
	// The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
	Name interface{}
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
	ResourceId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName interface{}
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName interface{}
	// A `linkedServiceProperties` block as defined below.
	LinkedServiceProperties interface{}
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
	ResourceId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName interface{}
}
