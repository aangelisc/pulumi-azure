// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Schema within an API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_api_schema.html.markdown.
type ApiSchema struct {
	s *pulumi.ResourceState
}

// NewApiSchema registers a new resource with the given unique name, arguments, and options.
func NewApiSchema(ctx *pulumi.Context,
	name string, args *ApiSchemaArgs, opts ...pulumi.ResourceOpt) (*ApiSchema, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ApiName == nil {
		return nil, errors.New("missing required argument 'ApiName'")
	}
	if args == nil || args.ContentType == nil {
		return nil, errors.New("missing required argument 'ContentType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SchemaId == nil {
		return nil, errors.New("missing required argument 'SchemaId'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["apiName"] = nil
		inputs["contentType"] = nil
		inputs["resourceGroupName"] = nil
		inputs["schemaId"] = nil
		inputs["value"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["apiName"] = args.ApiName
		inputs["contentType"] = args.ContentType
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["schemaId"] = args.SchemaId
		inputs["value"] = args.Value
	}
	s, err := ctx.RegisterResource("azure:apimanagement/apiSchema:ApiSchema", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiSchema{s: s}, nil
}

// GetApiSchema gets an existing ApiSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiSchema(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApiSchemaState, opts ...pulumi.ResourceOpt) (*ApiSchema, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["apiName"] = state.ApiName
		inputs["contentType"] = state.ContentType
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["schemaId"] = state.SchemaId
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("azure:apimanagement/apiSchema:ApiSchema", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiSchema{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApiSchema) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApiSchema) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
func (r *ApiSchema) ApiManagementName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
func (r *ApiSchema) ApiName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiName"])
}

// The content type of the API Schema.
func (r *ApiSchema) ContentType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["contentType"])
}

// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
func (r *ApiSchema) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A unique identifier for this API Schema. Changing this forces a new resource to be created.
func (r *ApiSchema) SchemaId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["schemaId"])
}

// The JSON escaped string defining the document representing the Schema.
func (r *ApiSchema) Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering ApiSchema resources.
type ApiSchemaState struct {
	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
	ApiName interface{}
	// The content type of the API Schema.
	ContentType interface{}
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A unique identifier for this API Schema. Changing this forces a new resource to be created.
	SchemaId interface{}
	// The JSON escaped string defining the document representing the Schema.
	Value interface{}
}

// The set of arguments for constructing a ApiSchema resource.
type ApiSchemaArgs struct {
	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
	ApiName interface{}
	// The content type of the API Schema.
	ContentType interface{}
	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A unique identifier for this API Schema. Changing this forces a new resource to be created.
	SchemaId interface{}
	// The JSON escaped string defining the document representing the Schema.
	Value interface{}
}
