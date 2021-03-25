// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SQL Stored Procedure within a Cosmos DB Account SQL Database.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := cosmosdb.LookupAccount(ctx, &cosmosdb.LookupAccountArgs{
// 			Name:              "tfex-cosmosdb-account",
// 			ResourceGroupName: "tfex-cosmosdb-account-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleSqlDatabase, err := cosmosdb.NewSqlDatabase(ctx, "exampleSqlDatabase", &cosmosdb.SqlDatabaseArgs{
// 			ResourceGroupName: pulumi.String(exampleAccount.ResourceGroupName),
// 			AccountName:       pulumi.String(exampleAccount.Name),
// 			Throughput:        pulumi.Int(400),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSqlContainer, err := cosmosdb.NewSqlContainer(ctx, "exampleSqlContainer", &cosmosdb.SqlContainerArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_cosmosdb_account.Example.Resource_group_name),
// 			AccountName:       pulumi.Any(azurerm_cosmosdb_account.Example.Name),
// 			DatabaseName:      exampleSqlDatabase.Name,
// 			PartitionKeyPath:  pulumi.String("/id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewSqlStoredProcedure(ctx, "exampleSqlStoredProcedure", &cosmosdb.SqlStoredProcedureArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_cosmosdb_account.Example.Resource_group_name),
// 			AccountName:       pulumi.Any(azurerm_cosmosdb_account.Example.Name),
// 			DatabaseName:      exampleSqlDatabase.Name,
// 			ContainerName:     exampleSqlContainer.Name,
// 			Body: pulumi.String("  	function () { var context = getContext(); var response = context.getResponse(); response.setBody('Hello, World'); }\n"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// CosmosDB SQL Stored Procedures can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/sqlStoredProcedure:SqlStoredProcedure db1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1/containers/c1/storedProcedures/sp1
// ```
type SqlStoredProcedure struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The body of the stored procedure.
	Body pulumi.StringOutput `pulumi:"body"`
	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Specifies the name of the Cosmos DB SQL Stored Procedure. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewSqlStoredProcedure registers a new resource with the given unique name, arguments, and options.
func NewSqlStoredProcedure(ctx *pulumi.Context,
	name string, args *SqlStoredProcedureArgs, opts ...pulumi.ResourceOption) (*SqlStoredProcedure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SqlStoredProcedure
	err := ctx.RegisterResource("azure:cosmosdb/sqlStoredProcedure:SqlStoredProcedure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlStoredProcedure gets an existing SqlStoredProcedure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlStoredProcedure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlStoredProcedureState, opts ...pulumi.ResourceOption) (*SqlStoredProcedure, error) {
	var resource SqlStoredProcedure
	err := ctx.ReadResource("azure:cosmosdb/sqlStoredProcedure:SqlStoredProcedure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlStoredProcedure resources.
type sqlStoredProcedureState struct {
	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// The body of the stored procedure.
	Body *string `pulumi:"body"`
	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	ContainerName *string `pulumi:"containerName"`
	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// Specifies the name of the Cosmos DB SQL Stored Procedure. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type SqlStoredProcedureState struct {
	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// The body of the stored procedure.
	Body pulumi.StringPtrInput
	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	ContainerName pulumi.StringPtrInput
	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// Specifies the name of the Cosmos DB SQL Stored Procedure. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (SqlStoredProcedureState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlStoredProcedureState)(nil)).Elem()
}

type sqlStoredProcedureArgs struct {
	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// The body of the stored procedure.
	Body string `pulumi:"body"`
	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	ContainerName string `pulumi:"containerName"`
	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies the name of the Cosmos DB SQL Stored Procedure. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a SqlStoredProcedure resource.
type SqlStoredProcedureArgs struct {
	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// The body of the stored procedure.
	Body pulumi.StringInput
	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	ContainerName pulumi.StringInput
	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// Specifies the name of the Cosmos DB SQL Stored Procedure. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (SqlStoredProcedureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlStoredProcedureArgs)(nil)).Elem()
}

type SqlStoredProcedureInput interface {
	pulumi.Input

	ToSqlStoredProcedureOutput() SqlStoredProcedureOutput
	ToSqlStoredProcedureOutputWithContext(ctx context.Context) SqlStoredProcedureOutput
}

func (*SqlStoredProcedure) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedure)(nil))
}

func (i *SqlStoredProcedure) ToSqlStoredProcedureOutput() SqlStoredProcedureOutput {
	return i.ToSqlStoredProcedureOutputWithContext(context.Background())
}

func (i *SqlStoredProcedure) ToSqlStoredProcedureOutputWithContext(ctx context.Context) SqlStoredProcedureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureOutput)
}

func (i *SqlStoredProcedure) ToSqlStoredProcedurePtrOutput() SqlStoredProcedurePtrOutput {
	return i.ToSqlStoredProcedurePtrOutputWithContext(context.Background())
}

func (i *SqlStoredProcedure) ToSqlStoredProcedurePtrOutputWithContext(ctx context.Context) SqlStoredProcedurePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedurePtrOutput)
}

type SqlStoredProcedurePtrInput interface {
	pulumi.Input

	ToSqlStoredProcedurePtrOutput() SqlStoredProcedurePtrOutput
	ToSqlStoredProcedurePtrOutputWithContext(ctx context.Context) SqlStoredProcedurePtrOutput
}

type sqlStoredProcedurePtrType SqlStoredProcedureArgs

func (*sqlStoredProcedurePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedure)(nil))
}

func (i *sqlStoredProcedurePtrType) ToSqlStoredProcedurePtrOutput() SqlStoredProcedurePtrOutput {
	return i.ToSqlStoredProcedurePtrOutputWithContext(context.Background())
}

func (i *sqlStoredProcedurePtrType) ToSqlStoredProcedurePtrOutputWithContext(ctx context.Context) SqlStoredProcedurePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedurePtrOutput)
}

// SqlStoredProcedureArrayInput is an input type that accepts SqlStoredProcedureArray and SqlStoredProcedureArrayOutput values.
// You can construct a concrete instance of `SqlStoredProcedureArrayInput` via:
//
//          SqlStoredProcedureArray{ SqlStoredProcedureArgs{...} }
type SqlStoredProcedureArrayInput interface {
	pulumi.Input

	ToSqlStoredProcedureArrayOutput() SqlStoredProcedureArrayOutput
	ToSqlStoredProcedureArrayOutputWithContext(context.Context) SqlStoredProcedureArrayOutput
}

type SqlStoredProcedureArray []SqlStoredProcedureInput

func (SqlStoredProcedureArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SqlStoredProcedure)(nil))
}

func (i SqlStoredProcedureArray) ToSqlStoredProcedureArrayOutput() SqlStoredProcedureArrayOutput {
	return i.ToSqlStoredProcedureArrayOutputWithContext(context.Background())
}

func (i SqlStoredProcedureArray) ToSqlStoredProcedureArrayOutputWithContext(ctx context.Context) SqlStoredProcedureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureArrayOutput)
}

// SqlStoredProcedureMapInput is an input type that accepts SqlStoredProcedureMap and SqlStoredProcedureMapOutput values.
// You can construct a concrete instance of `SqlStoredProcedureMapInput` via:
//
//          SqlStoredProcedureMap{ "key": SqlStoredProcedureArgs{...} }
type SqlStoredProcedureMapInput interface {
	pulumi.Input

	ToSqlStoredProcedureMapOutput() SqlStoredProcedureMapOutput
	ToSqlStoredProcedureMapOutputWithContext(context.Context) SqlStoredProcedureMapOutput
}

type SqlStoredProcedureMap map[string]SqlStoredProcedureInput

func (SqlStoredProcedureMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SqlStoredProcedure)(nil))
}

func (i SqlStoredProcedureMap) ToSqlStoredProcedureMapOutput() SqlStoredProcedureMapOutput {
	return i.ToSqlStoredProcedureMapOutputWithContext(context.Background())
}

func (i SqlStoredProcedureMap) ToSqlStoredProcedureMapOutputWithContext(ctx context.Context) SqlStoredProcedureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureMapOutput)
}

type SqlStoredProcedureOutput struct {
	*pulumi.OutputState
}

func (SqlStoredProcedureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedure)(nil))
}

func (o SqlStoredProcedureOutput) ToSqlStoredProcedureOutput() SqlStoredProcedureOutput {
	return o
}

func (o SqlStoredProcedureOutput) ToSqlStoredProcedureOutputWithContext(ctx context.Context) SqlStoredProcedureOutput {
	return o
}

func (o SqlStoredProcedureOutput) ToSqlStoredProcedurePtrOutput() SqlStoredProcedurePtrOutput {
	return o.ToSqlStoredProcedurePtrOutputWithContext(context.Background())
}

func (o SqlStoredProcedureOutput) ToSqlStoredProcedurePtrOutputWithContext(ctx context.Context) SqlStoredProcedurePtrOutput {
	return o.ApplyT(func(v SqlStoredProcedure) *SqlStoredProcedure {
		return &v
	}).(SqlStoredProcedurePtrOutput)
}

type SqlStoredProcedurePtrOutput struct {
	*pulumi.OutputState
}

func (SqlStoredProcedurePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedure)(nil))
}

func (o SqlStoredProcedurePtrOutput) ToSqlStoredProcedurePtrOutput() SqlStoredProcedurePtrOutput {
	return o
}

func (o SqlStoredProcedurePtrOutput) ToSqlStoredProcedurePtrOutputWithContext(ctx context.Context) SqlStoredProcedurePtrOutput {
	return o
}

type SqlStoredProcedureArrayOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlStoredProcedure)(nil))
}

func (o SqlStoredProcedureArrayOutput) ToSqlStoredProcedureArrayOutput() SqlStoredProcedureArrayOutput {
	return o
}

func (o SqlStoredProcedureArrayOutput) ToSqlStoredProcedureArrayOutputWithContext(ctx context.Context) SqlStoredProcedureArrayOutput {
	return o
}

func (o SqlStoredProcedureArrayOutput) Index(i pulumi.IntInput) SqlStoredProcedureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SqlStoredProcedure {
		return vs[0].([]SqlStoredProcedure)[vs[1].(int)]
	}).(SqlStoredProcedureOutput)
}

type SqlStoredProcedureMapOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SqlStoredProcedure)(nil))
}

func (o SqlStoredProcedureMapOutput) ToSqlStoredProcedureMapOutput() SqlStoredProcedureMapOutput {
	return o
}

func (o SqlStoredProcedureMapOutput) ToSqlStoredProcedureMapOutputWithContext(ctx context.Context) SqlStoredProcedureMapOutput {
	return o
}

func (o SqlStoredProcedureMapOutput) MapIndex(k pulumi.StringInput) SqlStoredProcedureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SqlStoredProcedure {
		return vs[0].(map[string]SqlStoredProcedure)[vs[1].(string)]
	}).(SqlStoredProcedureOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlStoredProcedureOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedurePtrOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureArrayOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureMapOutput{})
}
