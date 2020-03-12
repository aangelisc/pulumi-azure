// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Stream Analytics Output to Microsoft SQL Server Database.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_mssql.html.markdown.
type OutputMssql struct {
	pulumi.CustomResourceState

	Database pulumi.StringOutput `pulumi:"database"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	Password pulumi.StringOutput `pulumi:"password"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SQL server url. Changing this forces a new resource to be created.
	Server pulumi.StringOutput `pulumi:"server"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
	// Table in the database that the output points to. Changing this forces a new resource to be created.
	Table pulumi.StringOutput `pulumi:"table"`
	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewOutputMssql registers a new resource with the given unique name, arguments, and options.
func NewOutputMssql(ctx *pulumi.Context,
	name string, args *OutputMssqlArgs, opts ...pulumi.ResourceOption) (*OutputMssql, error) {
	if args == nil || args.Database == nil {
		return nil, errors.New("missing required argument 'Database'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Server == nil {
		return nil, errors.New("missing required argument 'Server'")
	}
	if args == nil || args.StreamAnalyticsJobName == nil {
		return nil, errors.New("missing required argument 'StreamAnalyticsJobName'")
	}
	if args == nil || args.Table == nil {
		return nil, errors.New("missing required argument 'Table'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &OutputMssqlArgs{}
	}
	var resource OutputMssql
	err := ctx.RegisterResource("azure:streamanalytics/outputMssql:OutputMssql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputMssql gets an existing OutputMssql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputMssql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputMssqlState, opts ...pulumi.ResourceOption) (*OutputMssql, error) {
	var resource OutputMssql
	err := ctx.ReadResource("azure:streamanalytics/outputMssql:OutputMssql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputMssql resources.
type outputMssqlState struct {
	Database *string `pulumi:"database"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SQL server url. Changing this forces a new resource to be created.
	Server *string `pulumi:"server"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
	// Table in the database that the output points to. Changing this forces a new resource to be created.
	Table *string `pulumi:"table"`
	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	User *string `pulumi:"user"`
}

type OutputMssqlState struct {
	Database pulumi.StringPtrInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SQL server url. Changing this forces a new resource to be created.
	Server pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
	// Table in the database that the output points to. Changing this forces a new resource to be created.
	Table pulumi.StringPtrInput
	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	User pulumi.StringPtrInput
}

func (OutputMssqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputMssqlState)(nil)).Elem()
}

type outputMssqlArgs struct {
	Database string `pulumi:"database"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	Password string `pulumi:"password"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SQL server url. Changing this forces a new resource to be created.
	Server string `pulumi:"server"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
	// Table in the database that the output points to. Changing this forces a new resource to be created.
	Table string `pulumi:"table"`
	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a OutputMssql resource.
type OutputMssqlArgs struct {
	Database pulumi.StringInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Password used together with username, to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	Password pulumi.StringInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SQL server url. Changing this forces a new resource to be created.
	Server pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
	// Table in the database that the output points to. Changing this forces a new resource to be created.
	Table pulumi.StringInput
	// Username used to login to the Microsoft SQL Server. Changing this forces a new resource to be created.
	User pulumi.StringInput
}

func (OutputMssqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputMssqlArgs)(nil)).Elem()
}

