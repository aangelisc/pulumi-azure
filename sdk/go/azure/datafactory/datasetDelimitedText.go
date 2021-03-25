// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Delimited Text Dataset inside an Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLinkedServiceWeb, err := datafactory.NewLinkedServiceWeb(ctx, "exampleLinkedServiceWeb", &datafactory.LinkedServiceWebArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryName:    exampleFactory.Name,
// 			AuthenticationType: pulumi.String("Anonymous"),
// 			Url:                pulumi.String("https://www.bing.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetDelimitedText(ctx, "exampleDatasetDelimitedText", &datafactory.DatasetDelimitedTextArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			LinkedServiceName: exampleLinkedServiceWeb.Name,
// 			HttpServerLocation: &datafactory.DatasetDelimitedTextHttpServerLocationArgs{
// 				RelativeUrl: pulumi.String("http://www.bing.com"),
// 				Path:        pulumi.String("foo/bar/"),
// 				Filename:    pulumi.String("fizz.txt"),
// 			},
// 			ColumnDelimiter:  pulumi.String(","),
// 			RowDelimiter:     pulumi.String("NEW"),
// 			Encoding:         pulumi.String("UTF-8"),
// 			QuoteCharacter:   pulumi.String("x"),
// 			EscapeCharacter:  pulumi.String("f"),
// 			FirstRowAsHeader: pulumi.Bool(true),
// 			NullValue:        pulumi.String("NULL"),
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
// Data Factory Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/datasetDelimitedText:DatasetDelimitedText example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type DatasetDelimitedText struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetDelimitedTextAzureBlobStorageLocationPtrOutput `pulumi:"azureBlobStorageLocation"`
	// The column delimiter.
	ColumnDelimiter pulumi.StringPtrOutput `pulumi:"columnDelimiter"`
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec pulumi.StringPtrOutput `pulumi:"compressionCodec"`
	// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
	CompressionLevel pulumi.StringPtrOutput `pulumi:"compressionLevel"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encoding format for the file.
	Encoding pulumi.StringPtrOutput `pulumi:"encoding"`
	// The escape character.
	EscapeCharacter pulumi.StringPtrOutput `pulumi:"escapeCharacter"`
	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
	FirstRowAsHeader pulumi.BoolPtrOutput `pulumi:"firstRowAsHeader"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetDelimitedTextHttpServerLocationPtrOutput `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// The null value string.
	NullValue pulumi.StringPtrOutput `pulumi:"nullValue"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The quote character.
	QuoteCharacter pulumi.StringPtrOutput `pulumi:"quoteCharacter"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The row delimiter.
	RowDelimiter pulumi.StringPtrOutput `pulumi:"rowDelimiter"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetDelimitedTextSchemaColumnArrayOutput `pulumi:"schemaColumns"`
}

// NewDatasetDelimitedText registers a new resource with the given unique name, arguments, and options.
func NewDatasetDelimitedText(ctx *pulumi.Context,
	name string, args *DatasetDelimitedTextArgs, opts ...pulumi.ResourceOption) (*DatasetDelimitedText, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.LinkedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'LinkedServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DatasetDelimitedText
	err := ctx.RegisterResource("azure:datafactory/datasetDelimitedText:DatasetDelimitedText", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetDelimitedText gets an existing DatasetDelimitedText resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetDelimitedText(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetDelimitedTextState, opts ...pulumi.ResourceOption) (*DatasetDelimitedText, error) {
	var resource DatasetDelimitedText
	err := ctx.ReadResource("azure:datafactory/datasetDelimitedText:DatasetDelimitedText", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetDelimitedText resources.
type datasetDelimitedTextState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation *DatasetDelimitedTextAzureBlobStorageLocation `pulumi:"azureBlobStorageLocation"`
	// The column delimiter.
	ColumnDelimiter *string `pulumi:"columnDelimiter"`
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec *string `pulumi:"compressionCodec"`
	// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
	CompressionLevel *string `pulumi:"compressionLevel"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The encoding format for the file.
	Encoding *string `pulumi:"encoding"`
	// The escape character.
	EscapeCharacter *string `pulumi:"escapeCharacter"`
	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
	FirstRowAsHeader *bool `pulumi:"firstRowAsHeader"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation *DatasetDelimitedTextHttpServerLocation `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The null value string.
	NullValue *string `pulumi:"nullValue"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The quote character.
	QuoteCharacter *string `pulumi:"quoteCharacter"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The row delimiter.
	RowDelimiter *string `pulumi:"rowDelimiter"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetDelimitedTextSchemaColumn `pulumi:"schemaColumns"`
}

type DatasetDelimitedTextState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetDelimitedTextAzureBlobStorageLocationPtrInput
	// The column delimiter.
	ColumnDelimiter pulumi.StringPtrInput
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec pulumi.StringPtrInput
	// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
	CompressionLevel pulumi.StringPtrInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The encoding format for the file.
	Encoding pulumi.StringPtrInput
	// The escape character.
	EscapeCharacter pulumi.StringPtrInput
	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
	FirstRowAsHeader pulumi.BoolPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetDelimitedTextHttpServerLocationPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The null value string.
	NullValue pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The quote character.
	QuoteCharacter pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The row delimiter.
	RowDelimiter pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetDelimitedTextSchemaColumnArrayInput
}

func (DatasetDelimitedTextState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDelimitedTextState)(nil)).Elem()
}

type datasetDelimitedTextArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation *DatasetDelimitedTextAzureBlobStorageLocation `pulumi:"azureBlobStorageLocation"`
	// The column delimiter.
	ColumnDelimiter *string `pulumi:"columnDelimiter"`
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec *string `pulumi:"compressionCodec"`
	// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
	CompressionLevel *string `pulumi:"compressionLevel"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The encoding format for the file.
	Encoding *string `pulumi:"encoding"`
	// The escape character.
	EscapeCharacter *string `pulumi:"escapeCharacter"`
	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
	FirstRowAsHeader *bool `pulumi:"firstRowAsHeader"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation *DatasetDelimitedTextHttpServerLocation `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The null value string.
	NullValue *string `pulumi:"nullValue"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The quote character.
	QuoteCharacter *string `pulumi:"quoteCharacter"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The row delimiter.
	RowDelimiter *string `pulumi:"rowDelimiter"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetDelimitedTextSchemaColumn `pulumi:"schemaColumns"`
}

// The set of arguments for constructing a DatasetDelimitedText resource.
type DatasetDelimitedTextArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetDelimitedTextAzureBlobStorageLocationPtrInput
	// The column delimiter.
	ColumnDelimiter pulumi.StringPtrInput
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec pulumi.StringPtrInput
	// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
	CompressionLevel pulumi.StringPtrInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The encoding format for the file.
	Encoding pulumi.StringPtrInput
	// The escape character.
	EscapeCharacter pulumi.StringPtrInput
	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
	FirstRowAsHeader pulumi.BoolPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetDelimitedTextHttpServerLocationPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The null value string.
	NullValue pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The quote character.
	QuoteCharacter pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The row delimiter.
	RowDelimiter pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetDelimitedTextSchemaColumnArrayInput
}

func (DatasetDelimitedTextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetDelimitedTextArgs)(nil)).Elem()
}

type DatasetDelimitedTextInput interface {
	pulumi.Input

	ToDatasetDelimitedTextOutput() DatasetDelimitedTextOutput
	ToDatasetDelimitedTextOutputWithContext(ctx context.Context) DatasetDelimitedTextOutput
}

func (*DatasetDelimitedText) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetDelimitedText)(nil))
}

func (i *DatasetDelimitedText) ToDatasetDelimitedTextOutput() DatasetDelimitedTextOutput {
	return i.ToDatasetDelimitedTextOutputWithContext(context.Background())
}

func (i *DatasetDelimitedText) ToDatasetDelimitedTextOutputWithContext(ctx context.Context) DatasetDelimitedTextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDelimitedTextOutput)
}

func (i *DatasetDelimitedText) ToDatasetDelimitedTextPtrOutput() DatasetDelimitedTextPtrOutput {
	return i.ToDatasetDelimitedTextPtrOutputWithContext(context.Background())
}

func (i *DatasetDelimitedText) ToDatasetDelimitedTextPtrOutputWithContext(ctx context.Context) DatasetDelimitedTextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDelimitedTextPtrOutput)
}

type DatasetDelimitedTextPtrInput interface {
	pulumi.Input

	ToDatasetDelimitedTextPtrOutput() DatasetDelimitedTextPtrOutput
	ToDatasetDelimitedTextPtrOutputWithContext(ctx context.Context) DatasetDelimitedTextPtrOutput
}

type datasetDelimitedTextPtrType DatasetDelimitedTextArgs

func (*datasetDelimitedTextPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetDelimitedText)(nil))
}

func (i *datasetDelimitedTextPtrType) ToDatasetDelimitedTextPtrOutput() DatasetDelimitedTextPtrOutput {
	return i.ToDatasetDelimitedTextPtrOutputWithContext(context.Background())
}

func (i *datasetDelimitedTextPtrType) ToDatasetDelimitedTextPtrOutputWithContext(ctx context.Context) DatasetDelimitedTextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDelimitedTextPtrOutput)
}

// DatasetDelimitedTextArrayInput is an input type that accepts DatasetDelimitedTextArray and DatasetDelimitedTextArrayOutput values.
// You can construct a concrete instance of `DatasetDelimitedTextArrayInput` via:
//
//          DatasetDelimitedTextArray{ DatasetDelimitedTextArgs{...} }
type DatasetDelimitedTextArrayInput interface {
	pulumi.Input

	ToDatasetDelimitedTextArrayOutput() DatasetDelimitedTextArrayOutput
	ToDatasetDelimitedTextArrayOutputWithContext(context.Context) DatasetDelimitedTextArrayOutput
}

type DatasetDelimitedTextArray []DatasetDelimitedTextInput

func (DatasetDelimitedTextArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatasetDelimitedText)(nil))
}

func (i DatasetDelimitedTextArray) ToDatasetDelimitedTextArrayOutput() DatasetDelimitedTextArrayOutput {
	return i.ToDatasetDelimitedTextArrayOutputWithContext(context.Background())
}

func (i DatasetDelimitedTextArray) ToDatasetDelimitedTextArrayOutputWithContext(ctx context.Context) DatasetDelimitedTextArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDelimitedTextArrayOutput)
}

// DatasetDelimitedTextMapInput is an input type that accepts DatasetDelimitedTextMap and DatasetDelimitedTextMapOutput values.
// You can construct a concrete instance of `DatasetDelimitedTextMapInput` via:
//
//          DatasetDelimitedTextMap{ "key": DatasetDelimitedTextArgs{...} }
type DatasetDelimitedTextMapInput interface {
	pulumi.Input

	ToDatasetDelimitedTextMapOutput() DatasetDelimitedTextMapOutput
	ToDatasetDelimitedTextMapOutputWithContext(context.Context) DatasetDelimitedTextMapOutput
}

type DatasetDelimitedTextMap map[string]DatasetDelimitedTextInput

func (DatasetDelimitedTextMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatasetDelimitedText)(nil))
}

func (i DatasetDelimitedTextMap) ToDatasetDelimitedTextMapOutput() DatasetDelimitedTextMapOutput {
	return i.ToDatasetDelimitedTextMapOutputWithContext(context.Background())
}

func (i DatasetDelimitedTextMap) ToDatasetDelimitedTextMapOutputWithContext(ctx context.Context) DatasetDelimitedTextMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetDelimitedTextMapOutput)
}

type DatasetDelimitedTextOutput struct {
	*pulumi.OutputState
}

func (DatasetDelimitedTextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetDelimitedText)(nil))
}

func (o DatasetDelimitedTextOutput) ToDatasetDelimitedTextOutput() DatasetDelimitedTextOutput {
	return o
}

func (o DatasetDelimitedTextOutput) ToDatasetDelimitedTextOutputWithContext(ctx context.Context) DatasetDelimitedTextOutput {
	return o
}

func (o DatasetDelimitedTextOutput) ToDatasetDelimitedTextPtrOutput() DatasetDelimitedTextPtrOutput {
	return o.ToDatasetDelimitedTextPtrOutputWithContext(context.Background())
}

func (o DatasetDelimitedTextOutput) ToDatasetDelimitedTextPtrOutputWithContext(ctx context.Context) DatasetDelimitedTextPtrOutput {
	return o.ApplyT(func(v DatasetDelimitedText) *DatasetDelimitedText {
		return &v
	}).(DatasetDelimitedTextPtrOutput)
}

type DatasetDelimitedTextPtrOutput struct {
	*pulumi.OutputState
}

func (DatasetDelimitedTextPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetDelimitedText)(nil))
}

func (o DatasetDelimitedTextPtrOutput) ToDatasetDelimitedTextPtrOutput() DatasetDelimitedTextPtrOutput {
	return o
}

func (o DatasetDelimitedTextPtrOutput) ToDatasetDelimitedTextPtrOutputWithContext(ctx context.Context) DatasetDelimitedTextPtrOutput {
	return o
}

type DatasetDelimitedTextArrayOutput struct{ *pulumi.OutputState }

func (DatasetDelimitedTextArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetDelimitedText)(nil))
}

func (o DatasetDelimitedTextArrayOutput) ToDatasetDelimitedTextArrayOutput() DatasetDelimitedTextArrayOutput {
	return o
}

func (o DatasetDelimitedTextArrayOutput) ToDatasetDelimitedTextArrayOutputWithContext(ctx context.Context) DatasetDelimitedTextArrayOutput {
	return o
}

func (o DatasetDelimitedTextArrayOutput) Index(i pulumi.IntInput) DatasetDelimitedTextOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetDelimitedText {
		return vs[0].([]DatasetDelimitedText)[vs[1].(int)]
	}).(DatasetDelimitedTextOutput)
}

type DatasetDelimitedTextMapOutput struct{ *pulumi.OutputState }

func (DatasetDelimitedTextMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatasetDelimitedText)(nil))
}

func (o DatasetDelimitedTextMapOutput) ToDatasetDelimitedTextMapOutput() DatasetDelimitedTextMapOutput {
	return o
}

func (o DatasetDelimitedTextMapOutput) ToDatasetDelimitedTextMapOutputWithContext(ctx context.Context) DatasetDelimitedTextMapOutput {
	return o
}

func (o DatasetDelimitedTextMapOutput) MapIndex(k pulumi.StringInput) DatasetDelimitedTextOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatasetDelimitedText {
		return vs[0].(map[string]DatasetDelimitedText)[vs[1].(string)]
	}).(DatasetDelimitedTextOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetDelimitedTextOutput{})
	pulumi.RegisterOutputType(DatasetDelimitedTextPtrOutput{})
	pulumi.RegisterOutputType(DatasetDelimitedTextArrayOutput{})
	pulumi.RegisterOutputType(DatasetDelimitedTextMapOutput{})
}
