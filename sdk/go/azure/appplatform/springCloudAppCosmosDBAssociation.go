// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates a Spring Cloud Application with a CosmosDB Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appplatform"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudService, err := appplatform.NewSpringCloudService(ctx, "exampleSpringCloudService", &appplatform.SpringCloudServiceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudApp, err := appplatform.NewSpringCloudApp(ctx, "exampleSpringCloudApp", &appplatform.SpringCloudAppArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServiceName:       exampleSpringCloudService.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := cosmosdb.NewAccount(ctx, "exampleAccount", &cosmosdb.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			OfferType:         pulumi.String("Standard"),
// 			Kind:              pulumi.String("GlobalDocumentDB"),
// 			ConsistencyPolicy: &cosmosdb.AccountConsistencyPolicyArgs{
// 				ConsistencyLevel: pulumi.String("Strong"),
// 			},
// 			GeoLocations: cosmosdb.AccountGeoLocationArray{
// 				&cosmosdb.AccountGeoLocationArgs{
// 					Location:         exampleResourceGroup.Location,
// 					FailoverPriority: pulumi.Int(0),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appplatform.NewSpringCloudAppCosmosDBAssociation(ctx, "exampleSpringCloudAppCosmosDBAssociation", &appplatform.SpringCloudAppCosmosDBAssociationArgs{
// 			SpringCloudAppId:  exampleSpringCloudApp.ID(),
// 			CosmosdbAccountId: exampleAccount.ID(),
// 			ApiType:           pulumi.String("table"),
// 			CosmosdbAccessKey: exampleAccount.PrimaryKey,
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
// Spring Cloud Application CosmosDB Association can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appplatform/springCloudAppCosmosDBAssociation:SpringCloudAppCosmosDBAssociation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/service1/apps/app1/bindings/bind1
// ```
type SpringCloudAppCosmosDBAssociation struct {
	pulumi.CustomResourceState

	// Specifies the api type which should be used when connecting to the CosmosDB Account. Possible values are `cassandra`, `gremlin`, `mongo`, `sql` or `table`. Changing this forces a new resource to be created.
	ApiType pulumi.StringOutput `pulumi:"apiType"`
	// Specifies the CosmosDB Account access key.
	CosmosdbAccessKey pulumi.StringOutput `pulumi:"cosmosdbAccessKey"`
	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosdbAccountId pulumi.StringOutput `pulumi:"cosmosdbAccountId"`
	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when `apiType` is `cassandra`.
	CosmosdbCassandraKeyspaceName pulumi.StringPtrOutput `pulumi:"cosmosdbCassandraKeyspaceName"`
	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinDatabaseName pulumi.StringPtrOutput `pulumi:"cosmosdbGremlinDatabaseName"`
	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinGraphName pulumi.StringPtrOutput `pulumi:"cosmosdbGremlinGraphName"`
	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `mongo`.
	CosmosdbMongoDatabaseName pulumi.StringPtrOutput `pulumi:"cosmosdbMongoDatabaseName"`
	// Specifies the name of the Sql Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `sql`.
	CosmosdbSqlDatabaseName pulumi.StringPtrOutput `pulumi:"cosmosdbSqlDatabaseName"`
	// Specifies the name of the Spring Cloud Application Association. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringOutput `pulumi:"springCloudAppId"`
}

// NewSpringCloudAppCosmosDBAssociation registers a new resource with the given unique name, arguments, and options.
func NewSpringCloudAppCosmosDBAssociation(ctx *pulumi.Context,
	name string, args *SpringCloudAppCosmosDBAssociationArgs, opts ...pulumi.ResourceOption) (*SpringCloudAppCosmosDBAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiType == nil {
		return nil, errors.New("invalid value for required argument 'ApiType'")
	}
	if args.CosmosdbAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'CosmosdbAccessKey'")
	}
	if args.CosmosdbAccountId == nil {
		return nil, errors.New("invalid value for required argument 'CosmosdbAccountId'")
	}
	if args.SpringCloudAppId == nil {
		return nil, errors.New("invalid value for required argument 'SpringCloudAppId'")
	}
	var resource SpringCloudAppCosmosDBAssociation
	err := ctx.RegisterResource("azure:appplatform/springCloudAppCosmosDBAssociation:SpringCloudAppCosmosDBAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringCloudAppCosmosDBAssociation gets an existing SpringCloudAppCosmosDBAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringCloudAppCosmosDBAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringCloudAppCosmosDBAssociationState, opts ...pulumi.ResourceOption) (*SpringCloudAppCosmosDBAssociation, error) {
	var resource SpringCloudAppCosmosDBAssociation
	err := ctx.ReadResource("azure:appplatform/springCloudAppCosmosDBAssociation:SpringCloudAppCosmosDBAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpringCloudAppCosmosDBAssociation resources.
type springCloudAppCosmosDBAssociationState struct {
	// Specifies the api type which should be used when connecting to the CosmosDB Account. Possible values are `cassandra`, `gremlin`, `mongo`, `sql` or `table`. Changing this forces a new resource to be created.
	ApiType *string `pulumi:"apiType"`
	// Specifies the CosmosDB Account access key.
	CosmosdbAccessKey *string `pulumi:"cosmosdbAccessKey"`
	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosdbAccountId *string `pulumi:"cosmosdbAccountId"`
	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when `apiType` is `cassandra`.
	CosmosdbCassandraKeyspaceName *string `pulumi:"cosmosdbCassandraKeyspaceName"`
	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinDatabaseName *string `pulumi:"cosmosdbGremlinDatabaseName"`
	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinGraphName *string `pulumi:"cosmosdbGremlinGraphName"`
	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `mongo`.
	CosmosdbMongoDatabaseName *string `pulumi:"cosmosdbMongoDatabaseName"`
	// Specifies the name of the Sql Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `sql`.
	CosmosdbSqlDatabaseName *string `pulumi:"cosmosdbSqlDatabaseName"`
	// Specifies the name of the Spring Cloud Application Association. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	SpringCloudAppId *string `pulumi:"springCloudAppId"`
}

type SpringCloudAppCosmosDBAssociationState struct {
	// Specifies the api type which should be used when connecting to the CosmosDB Account. Possible values are `cassandra`, `gremlin`, `mongo`, `sql` or `table`. Changing this forces a new resource to be created.
	ApiType pulumi.StringPtrInput
	// Specifies the CosmosDB Account access key.
	CosmosdbAccessKey pulumi.StringPtrInput
	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosdbAccountId pulumi.StringPtrInput
	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when `apiType` is `cassandra`.
	CosmosdbCassandraKeyspaceName pulumi.StringPtrInput
	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinDatabaseName pulumi.StringPtrInput
	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinGraphName pulumi.StringPtrInput
	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `mongo`.
	CosmosdbMongoDatabaseName pulumi.StringPtrInput
	// Specifies the name of the Sql Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `sql`.
	CosmosdbSqlDatabaseName pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Application Association. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringPtrInput
}

func (SpringCloudAppCosmosDBAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudAppCosmosDBAssociationState)(nil)).Elem()
}

type springCloudAppCosmosDBAssociationArgs struct {
	// Specifies the api type which should be used when connecting to the CosmosDB Account. Possible values are `cassandra`, `gremlin`, `mongo`, `sql` or `table`. Changing this forces a new resource to be created.
	ApiType string `pulumi:"apiType"`
	// Specifies the CosmosDB Account access key.
	CosmosdbAccessKey string `pulumi:"cosmosdbAccessKey"`
	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosdbAccountId string `pulumi:"cosmosdbAccountId"`
	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when `apiType` is `cassandra`.
	CosmosdbCassandraKeyspaceName *string `pulumi:"cosmosdbCassandraKeyspaceName"`
	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinDatabaseName *string `pulumi:"cosmosdbGremlinDatabaseName"`
	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinGraphName *string `pulumi:"cosmosdbGremlinGraphName"`
	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `mongo`.
	CosmosdbMongoDatabaseName *string `pulumi:"cosmosdbMongoDatabaseName"`
	// Specifies the name of the Sql Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `sql`.
	CosmosdbSqlDatabaseName *string `pulumi:"cosmosdbSqlDatabaseName"`
	// Specifies the name of the Spring Cloud Application Association. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	SpringCloudAppId string `pulumi:"springCloudAppId"`
}

// The set of arguments for constructing a SpringCloudAppCosmosDBAssociation resource.
type SpringCloudAppCosmosDBAssociationArgs struct {
	// Specifies the api type which should be used when connecting to the CosmosDB Account. Possible values are `cassandra`, `gremlin`, `mongo`, `sql` or `table`. Changing this forces a new resource to be created.
	ApiType pulumi.StringInput
	// Specifies the CosmosDB Account access key.
	CosmosdbAccessKey pulumi.StringInput
	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosdbAccountId pulumi.StringInput
	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when `apiType` is `cassandra`.
	CosmosdbCassandraKeyspaceName pulumi.StringPtrInput
	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinDatabaseName pulumi.StringPtrInput
	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when `apiType` is `gremlin`.
	CosmosdbGremlinGraphName pulumi.StringPtrInput
	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `mongo`.
	CosmosdbMongoDatabaseName pulumi.StringPtrInput
	// Specifies the name of the Sql Database which the Spring Cloud App should be associated with. Should only be set when `apiType` is `sql`.
	CosmosdbSqlDatabaseName pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Application Association. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringInput
}

func (SpringCloudAppCosmosDBAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudAppCosmosDBAssociationArgs)(nil)).Elem()
}

type SpringCloudAppCosmosDBAssociationInput interface {
	pulumi.Input

	ToSpringCloudAppCosmosDBAssociationOutput() SpringCloudAppCosmosDBAssociationOutput
	ToSpringCloudAppCosmosDBAssociationOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationOutput
}

func (*SpringCloudAppCosmosDBAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudAppCosmosDBAssociation)(nil))
}

func (i *SpringCloudAppCosmosDBAssociation) ToSpringCloudAppCosmosDBAssociationOutput() SpringCloudAppCosmosDBAssociationOutput {
	return i.ToSpringCloudAppCosmosDBAssociationOutputWithContext(context.Background())
}

func (i *SpringCloudAppCosmosDBAssociation) ToSpringCloudAppCosmosDBAssociationOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudAppCosmosDBAssociationOutput)
}

func (i *SpringCloudAppCosmosDBAssociation) ToSpringCloudAppCosmosDBAssociationPtrOutput() SpringCloudAppCosmosDBAssociationPtrOutput {
	return i.ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(context.Background())
}

func (i *SpringCloudAppCosmosDBAssociation) ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudAppCosmosDBAssociationPtrOutput)
}

type SpringCloudAppCosmosDBAssociationPtrInput interface {
	pulumi.Input

	ToSpringCloudAppCosmosDBAssociationPtrOutput() SpringCloudAppCosmosDBAssociationPtrOutput
	ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationPtrOutput
}

type springCloudAppCosmosDBAssociationPtrType SpringCloudAppCosmosDBAssociationArgs

func (*springCloudAppCosmosDBAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudAppCosmosDBAssociation)(nil))
}

func (i *springCloudAppCosmosDBAssociationPtrType) ToSpringCloudAppCosmosDBAssociationPtrOutput() SpringCloudAppCosmosDBAssociationPtrOutput {
	return i.ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(context.Background())
}

func (i *springCloudAppCosmosDBAssociationPtrType) ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudAppCosmosDBAssociationPtrOutput)
}

// SpringCloudAppCosmosDBAssociationArrayInput is an input type that accepts SpringCloudAppCosmosDBAssociationArray and SpringCloudAppCosmosDBAssociationArrayOutput values.
// You can construct a concrete instance of `SpringCloudAppCosmosDBAssociationArrayInput` via:
//
//          SpringCloudAppCosmosDBAssociationArray{ SpringCloudAppCosmosDBAssociationArgs{...} }
type SpringCloudAppCosmosDBAssociationArrayInput interface {
	pulumi.Input

	ToSpringCloudAppCosmosDBAssociationArrayOutput() SpringCloudAppCosmosDBAssociationArrayOutput
	ToSpringCloudAppCosmosDBAssociationArrayOutputWithContext(context.Context) SpringCloudAppCosmosDBAssociationArrayOutput
}

type SpringCloudAppCosmosDBAssociationArray []SpringCloudAppCosmosDBAssociationInput

func (SpringCloudAppCosmosDBAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SpringCloudAppCosmosDBAssociation)(nil))
}

func (i SpringCloudAppCosmosDBAssociationArray) ToSpringCloudAppCosmosDBAssociationArrayOutput() SpringCloudAppCosmosDBAssociationArrayOutput {
	return i.ToSpringCloudAppCosmosDBAssociationArrayOutputWithContext(context.Background())
}

func (i SpringCloudAppCosmosDBAssociationArray) ToSpringCloudAppCosmosDBAssociationArrayOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudAppCosmosDBAssociationArrayOutput)
}

// SpringCloudAppCosmosDBAssociationMapInput is an input type that accepts SpringCloudAppCosmosDBAssociationMap and SpringCloudAppCosmosDBAssociationMapOutput values.
// You can construct a concrete instance of `SpringCloudAppCosmosDBAssociationMapInput` via:
//
//          SpringCloudAppCosmosDBAssociationMap{ "key": SpringCloudAppCosmosDBAssociationArgs{...} }
type SpringCloudAppCosmosDBAssociationMapInput interface {
	pulumi.Input

	ToSpringCloudAppCosmosDBAssociationMapOutput() SpringCloudAppCosmosDBAssociationMapOutput
	ToSpringCloudAppCosmosDBAssociationMapOutputWithContext(context.Context) SpringCloudAppCosmosDBAssociationMapOutput
}

type SpringCloudAppCosmosDBAssociationMap map[string]SpringCloudAppCosmosDBAssociationInput

func (SpringCloudAppCosmosDBAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SpringCloudAppCosmosDBAssociation)(nil))
}

func (i SpringCloudAppCosmosDBAssociationMap) ToSpringCloudAppCosmosDBAssociationMapOutput() SpringCloudAppCosmosDBAssociationMapOutput {
	return i.ToSpringCloudAppCosmosDBAssociationMapOutputWithContext(context.Background())
}

func (i SpringCloudAppCosmosDBAssociationMap) ToSpringCloudAppCosmosDBAssociationMapOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudAppCosmosDBAssociationMapOutput)
}

type SpringCloudAppCosmosDBAssociationOutput struct {
	*pulumi.OutputState
}

func (SpringCloudAppCosmosDBAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudAppCosmosDBAssociation)(nil))
}

func (o SpringCloudAppCosmosDBAssociationOutput) ToSpringCloudAppCosmosDBAssociationOutput() SpringCloudAppCosmosDBAssociationOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationOutput) ToSpringCloudAppCosmosDBAssociationOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationOutput) ToSpringCloudAppCosmosDBAssociationPtrOutput() SpringCloudAppCosmosDBAssociationPtrOutput {
	return o.ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(context.Background())
}

func (o SpringCloudAppCosmosDBAssociationOutput) ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationPtrOutput {
	return o.ApplyT(func(v SpringCloudAppCosmosDBAssociation) *SpringCloudAppCosmosDBAssociation {
		return &v
	}).(SpringCloudAppCosmosDBAssociationPtrOutput)
}

type SpringCloudAppCosmosDBAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (SpringCloudAppCosmosDBAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudAppCosmosDBAssociation)(nil))
}

func (o SpringCloudAppCosmosDBAssociationPtrOutput) ToSpringCloudAppCosmosDBAssociationPtrOutput() SpringCloudAppCosmosDBAssociationPtrOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationPtrOutput) ToSpringCloudAppCosmosDBAssociationPtrOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationPtrOutput {
	return o
}

type SpringCloudAppCosmosDBAssociationArrayOutput struct{ *pulumi.OutputState }

func (SpringCloudAppCosmosDBAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpringCloudAppCosmosDBAssociation)(nil))
}

func (o SpringCloudAppCosmosDBAssociationArrayOutput) ToSpringCloudAppCosmosDBAssociationArrayOutput() SpringCloudAppCosmosDBAssociationArrayOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationArrayOutput) ToSpringCloudAppCosmosDBAssociationArrayOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationArrayOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationArrayOutput) Index(i pulumi.IntInput) SpringCloudAppCosmosDBAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpringCloudAppCosmosDBAssociation {
		return vs[0].([]SpringCloudAppCosmosDBAssociation)[vs[1].(int)]
	}).(SpringCloudAppCosmosDBAssociationOutput)
}

type SpringCloudAppCosmosDBAssociationMapOutput struct{ *pulumi.OutputState }

func (SpringCloudAppCosmosDBAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SpringCloudAppCosmosDBAssociation)(nil))
}

func (o SpringCloudAppCosmosDBAssociationMapOutput) ToSpringCloudAppCosmosDBAssociationMapOutput() SpringCloudAppCosmosDBAssociationMapOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationMapOutput) ToSpringCloudAppCosmosDBAssociationMapOutputWithContext(ctx context.Context) SpringCloudAppCosmosDBAssociationMapOutput {
	return o
}

func (o SpringCloudAppCosmosDBAssociationMapOutput) MapIndex(k pulumi.StringInput) SpringCloudAppCosmosDBAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SpringCloudAppCosmosDBAssociation {
		return vs[0].(map[string]SpringCloudAppCosmosDBAssociation)[vs[1].(string)]
	}).(SpringCloudAppCosmosDBAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(SpringCloudAppCosmosDBAssociationOutput{})
	pulumi.RegisterOutputType(SpringCloudAppCosmosDBAssociationPtrOutput{})
	pulumi.RegisterOutputType(SpringCloudAppCosmosDBAssociationArrayOutput{})
	pulumi.RegisterOutputType(SpringCloudAppCosmosDBAssociationMapOutput{})
}