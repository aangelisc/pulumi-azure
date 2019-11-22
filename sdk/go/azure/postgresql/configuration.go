// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Sets a PostgreSQL Configuration value on a PostgreSQL Server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/postgresql_configuration.html.markdown.
type Configuration struct {
	s *pulumi.ResourceState
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOpt) (*Configuration, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serverName"] = nil
		inputs["value"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serverName"] = args.ServerName
		inputs["value"] = args.Value
	}
	s, err := ctx.RegisterResource("azure:postgresql/configuration:Configuration", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Configuration{s: s}, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigurationState, opts ...pulumi.ResourceOpt) (*Configuration, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serverName"] = state.ServerName
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("azure:postgresql/configuration:Configuration", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Configuration{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Configuration) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Configuration) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the name of the PostgreSQL Configuration, which needs [to be a valid PostgreSQL configuration name](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIER). Changing this forces a new resource to be created.
func (r *Configuration) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
func (r *Configuration) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
func (r *Configuration) ServerName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serverName"])
}

// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
func (r *Configuration) Value() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering Configuration resources.
type ConfigurationState struct {
	// Specifies the name of the PostgreSQL Configuration, which needs [to be a valid PostgreSQL configuration name](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIER). Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName interface{}
	// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
	Value interface{}
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Specifies the name of the PostgreSQL Configuration, which needs [to be a valid PostgreSQL configuration name](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIER). Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName interface{}
	// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
	Value interface{}
}
