// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a SQL Azure Database Server.
//
// > **Note:** All arguments including the administrator login and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/sql_server.html.markdown.
type SqlServer struct {
	pulumi.CustomResourceState

	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringOutput `pulumi:"administratorLoginPassword"`
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy SqlServerExtendedAuditingPolicyPtrOutput `pulumi:"extendedAuditingPolicy"`
	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// An `identity` block as defined below.
	Identity SqlServerIdentityPtrOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the SQL Server. This needs to be globally unique within Azure.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSqlServer registers a new resource with the given unique name, arguments, and options.
func NewSqlServer(ctx *pulumi.Context,
	name string, args *SqlServerArgs, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	if args == nil || args.AdministratorLogin == nil {
		return nil, errors.New("missing required argument 'AdministratorLogin'")
	}
	if args == nil || args.AdministratorLoginPassword == nil {
		return nil, errors.New("missing required argument 'AdministratorLoginPassword'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	if args == nil {
		args = &SqlServerArgs{}
	}
	var resource SqlServer
	err := ctx.RegisterResource("azure:sql/sqlServer:SqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlServer gets an existing SqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerState, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	var resource SqlServer
	err := ctx.ReadResource("azure:sql/sqlServer:SqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlServer resources.
type sqlServerState struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy *SqlServerExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// An `identity` block as defined below.
	Identity *SqlServerIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the SQL Server. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version *string `pulumi:"version"`
}

type SqlServerState struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy SqlServerExtendedAuditingPolicyPtrInput
	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity SqlServerIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the SQL Server. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version pulumi.StringPtrInput
}

func (SqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerState)(nil)).Elem()
}

type sqlServerArgs struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword string `pulumi:"administratorLoginPassword"`
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy *SqlServerExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// An `identity` block as defined below.
	Identity *SqlServerIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the SQL Server. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a SqlServer resource.
type SqlServerArgs struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringInput
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringInput
	// A `extendedAuditingPolicy` block as defined below.
	ExtendedAuditingPolicy SqlServerExtendedAuditingPolicyPtrInput
	// An `identity` block as defined below.
	Identity SqlServerIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the SQL Server. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the SQL Server.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version pulumi.StringInput
}

func (SqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerArgs)(nil)).Elem()
}

