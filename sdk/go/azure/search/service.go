// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package search

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Search Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/search_service.html.markdown.
type Service struct {
	pulumi.CustomResourceState

	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of partitions which should be created.
	PartitionCount pulumi.IntOutput `pulumi:"partitionCount"`
	// The Primary Key used for Search Service Administration.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// A `queryKeys` block as defined below.
	QueryKeys ServiceQueryKeyArrayOutput `pulumi:"queryKeys"`
	// The number of replica's which should be created.
	ReplicaCount pulumi.IntOutput `pulumi:"replicaCount"`
	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Key used for Search Service Administration.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags which should be assigned to the Search Service.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("azure:search/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:search/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name *string `pulumi:"name"`
	// The number of partitions which should be created.
	PartitionCount *int `pulumi:"partitionCount"`
	// The Primary Key used for Search Service Administration.
	PrimaryKey *string `pulumi:"primaryKey"`
	// A `queryKeys` block as defined below.
	QueryKeys []ServiceQueryKey `pulumi:"queryKeys"`
	// The number of replica's which should be created.
	ReplicaCount *int `pulumi:"replicaCount"`
	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Key used for Search Service Administration.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags which should be assigned to the Search Service.
	Tags map[string]string `pulumi:"tags"`
}

type ServiceState struct {
	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name pulumi.StringPtrInput
	// The number of partitions which should be created.
	PartitionCount pulumi.IntPtrInput
	// The Primary Key used for Search Service Administration.
	PrimaryKey pulumi.StringPtrInput
	// A `queryKeys` block as defined below.
	QueryKeys ServiceQueryKeyArrayInput
	// The number of replica's which should be created.
	ReplicaCount pulumi.IntPtrInput
	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Key used for Search Service Administration.
	SecondaryKey pulumi.StringPtrInput
	// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Search Service.
	Tags pulumi.StringMapInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name *string `pulumi:"name"`
	// The number of partitions which should be created.
	PartitionCount *int `pulumi:"partitionCount"`
	// The number of replica's which should be created.
	ReplicaCount *int `pulumi:"replicaCount"`
	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
	Sku string `pulumi:"sku"`
	// A mapping of tags which should be assigned to the Search Service.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name pulumi.StringPtrInput
	// The number of partitions which should be created.
	PartitionCount pulumi.IntPtrInput
	// The number of replica's which should be created.
	ReplicaCount pulumi.IntPtrInput
	// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2` and `standard3` Changing this forces a new Search Service to be created.
	Sku pulumi.StringInput
	// A mapping of tags which should be assigned to the Search Service.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

