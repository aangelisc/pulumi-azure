// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows you to set a user or group as the AD administrator for an Azure SQL server
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/sql_active_directory_administrator.html.markdown.
type ActiveDirectoryAdministrator struct {
	s *pulumi.ResourceState
}

// NewActiveDirectoryAdministrator registers a new resource with the given unique name, arguments, and options.
func NewActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, args *ActiveDirectoryAdministratorArgs, opts ...pulumi.ResourceOpt) (*ActiveDirectoryAdministrator, error) {
	if args == nil || args.Login == nil {
		return nil, errors.New("missing required argument 'Login'")
	}
	if args == nil || args.ObjectId == nil {
		return nil, errors.New("missing required argument 'ObjectId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["login"] = nil
		inputs["objectId"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serverName"] = nil
		inputs["tenantId"] = nil
	} else {
		inputs["login"] = args.Login
		inputs["objectId"] = args.ObjectId
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serverName"] = args.ServerName
		inputs["tenantId"] = args.TenantId
	}
	s, err := ctx.RegisterResource("azure:sql/activeDirectoryAdministrator:ActiveDirectoryAdministrator", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ActiveDirectoryAdministrator{s: s}, nil
}

// GetActiveDirectoryAdministrator gets an existing ActiveDirectoryAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ActiveDirectoryAdministratorState, opts ...pulumi.ResourceOpt) (*ActiveDirectoryAdministrator, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["login"] = state.Login
		inputs["objectId"] = state.ObjectId
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serverName"] = state.ServerName
		inputs["tenantId"] = state.TenantId
	}
	s, err := ctx.ReadResource("azure:sql/activeDirectoryAdministrator:ActiveDirectoryAdministrator", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ActiveDirectoryAdministrator{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ActiveDirectoryAdministrator) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ActiveDirectoryAdministrator) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The login name of the principal to set as the server administrator
func (r *ActiveDirectoryAdministrator) Login() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["login"])
}

// The ID of the principal to set as the server administrator
func (r *ActiveDirectoryAdministrator) ObjectId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["objectId"])
}

// The name of the resource group for the SQL server. Changing this forces a new resource to be created.
func (r *ActiveDirectoryAdministrator) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The name of the SQL Server on which to set the administrator. Changing this forces a new resource to be created.
func (r *ActiveDirectoryAdministrator) ServerName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serverName"])
}

// The Azure Tenant ID
func (r *ActiveDirectoryAdministrator) TenantId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tenantId"])
}

// Input properties used for looking up and filtering ActiveDirectoryAdministrator resources.
type ActiveDirectoryAdministratorState struct {
	// The login name of the principal to set as the server administrator
	Login interface{}
	// The ID of the principal to set as the server administrator
	ObjectId interface{}
	// The name of the resource group for the SQL server. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the SQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName interface{}
	// The Azure Tenant ID
	TenantId interface{}
}

// The set of arguments for constructing a ActiveDirectoryAdministrator resource.
type ActiveDirectoryAdministratorArgs struct {
	// The login name of the principal to set as the server administrator
	Login interface{}
	// The ID of the principal to set as the server administrator
	ObjectId interface{}
	// The name of the resource group for the SQL server. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the SQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName interface{}
	// The Azure Tenant ID
	TenantId interface{}
}
