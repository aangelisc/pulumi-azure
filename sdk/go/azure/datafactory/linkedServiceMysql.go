// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Linked Service (connection) between MySQL and Azure Data Factory.
// 
// > **Note:** All arguments including the connectionString will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_factory_linked_service_mysql.html.markdown.
type LinkedServiceMysql struct {
	s *pulumi.ResourceState
}

// NewLinkedServiceMysql registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceMysql(ctx *pulumi.Context,
	name string, args *LinkedServiceMysqlArgs, opts ...pulumi.ResourceOpt) (*LinkedServiceMysql, error) {
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["additionalProperties"] = nil
		inputs["annotations"] = nil
		inputs["connectionString"] = nil
		inputs["dataFactoryName"] = nil
		inputs["description"] = nil
		inputs["integrationRuntimeName"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["additionalProperties"] = args.AdditionalProperties
		inputs["annotations"] = args.Annotations
		inputs["connectionString"] = args.ConnectionString
		inputs["dataFactoryName"] = args.DataFactoryName
		inputs["description"] = args.Description
		inputs["integrationRuntimeName"] = args.IntegrationRuntimeName
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:datafactory/linkedServiceMysql:LinkedServiceMysql", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinkedServiceMysql{s: s}, nil
}

// GetLinkedServiceMysql gets an existing LinkedServiceMysql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceMysql(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LinkedServiceMysqlState, opts ...pulumi.ResourceOpt) (*LinkedServiceMysql, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["additionalProperties"] = state.AdditionalProperties
		inputs["annotations"] = state.Annotations
		inputs["connectionString"] = state.ConnectionString
		inputs["dataFactoryName"] = state.DataFactoryName
		inputs["description"] = state.Description
		inputs["integrationRuntimeName"] = state.IntegrationRuntimeName
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:datafactory/linkedServiceMysql:LinkedServiceMysql", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinkedServiceMysql{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LinkedServiceMysql) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LinkedServiceMysql) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A map of additional properties to associate with the Data Factory Linked Service MySQL.
func (r *LinkedServiceMysql) AdditionalProperties() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["additionalProperties"])
}

// List of tags that can be used for describing the Data Factory Linked Service MySQL.
func (r *LinkedServiceMysql) Annotations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["annotations"])
}

// The connection string in which to authenticate with MySQL.
func (r *LinkedServiceMysql) ConnectionString() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["connectionString"])
}

// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
func (r *LinkedServiceMysql) DataFactoryName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dataFactoryName"])
}

// The description for the Data Factory Linked Service MySQL.
func (r *LinkedServiceMysql) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
func (r *LinkedServiceMysql) IntegrationRuntimeName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["integrationRuntimeName"])
}

// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
func (r *LinkedServiceMysql) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A map of parameters to associate with the Data Factory Linked Service MySQL.
func (r *LinkedServiceMysql) Parameters() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["parameters"])
}

// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
func (r *LinkedServiceMysql) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering LinkedServiceMysql resources.
type LinkedServiceMysqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties interface{}
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations interface{}
	// The connection string in which to authenticate with MySQL.
	ConnectionString interface{}
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName interface{}
	// The description for the Data Factory Linked Service MySQL.
	Description interface{}
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName interface{}
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name interface{}
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters interface{}
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName interface{}
}

// The set of arguments for constructing a LinkedServiceMysql resource.
type LinkedServiceMysqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	AdditionalProperties interface{}
	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations interface{}
	// The connection string in which to authenticate with MySQL.
	ConnectionString interface{}
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName interface{}
	// The description for the Data Factory Linked Service MySQL.
	Description interface{}
	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName interface{}
	// Specifies the name of the Data Factory Linked Service MySQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name interface{}
	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	Parameters interface{}
	// The name of the resource group in which to create the Data Factory Linked Service MySQL. Changing this forces a new resource
	ResourceGroupName interface{}
}
