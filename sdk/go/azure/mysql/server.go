// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package mysql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a MySQL Server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/mysql_server.html.markdown.
type Server struct {
	pulumi.CustomResourceState

	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword pulumi.StringOutput `pulumi:"administratorLoginPassword"`
	// The FQDN of the MySQL Server.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement pulumi.StringOutput `pulumi:"sslEnforcement"`
	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfileOutput `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil || args.AdministratorLogin == nil {
		return nil, errors.New("missing required argument 'AdministratorLogin'")
	}
	if args == nil || args.AdministratorLoginPassword == nil {
		return nil, errors.New("missing required argument 'AdministratorLoginPassword'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil || args.SslEnforcement == nil {
		return nil, errors.New("missing required argument 'SslEnforcement'")
	}
	if args == nil || args.StorageProfile == nil {
		return nil, errors.New("missing required argument 'StorageProfile'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	if args == nil {
		args = &ServerArgs{}
	}
	var resource Server
	err := ctx.RegisterResource("azure:mysql/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure:mysql/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The FQDN of the MySQL Server.
	Fqdn *string `pulumi:"fqdn"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName *string `pulumi:"skuName"`
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// A `storageProfile` block as defined below.
	StorageProfile *ServerStorageProfile `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version *string `pulumi:"version"`
}

type ServerState struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword pulumi.StringPtrInput
	// The FQDN of the MySQL Server.
	Fqdn pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName pulumi.StringPtrInput
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement pulumi.StringPtrInput
	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfilePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword string `pulumi:"administratorLoginPassword"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName string `pulumi:"skuName"`
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement string `pulumi:"sslEnforcement"`
	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfile `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Administrator Login for the MySQL Server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringInput
	// The Password associated with the `administratorLogin` for the MySQL Server.
	AdministratorLoginPassword pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the MySQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#sku).
	SkuName pulumi.StringInput
	// Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
	SslEnforcement pulumi.StringInput
	// A `storageProfile` block as defined below.
	StorageProfile ServerStorageProfileInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the version of MySQL to use. Valid values are `5.6`, `5.7`, and `8.0`. Changing this forces a new resource to be created.
	Version pulumi.StringInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

