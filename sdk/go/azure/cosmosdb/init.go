// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:cosmosdb/account:Account":
		r = &Account{}
	case "azure:cosmosdb/cassandraKeyspace:CassandraKeyspace":
		r = &CassandraKeyspace{}
	case "azure:cosmosdb/cassandraTable:CassandraTable":
		r = &CassandraTable{}
	case "azure:cosmosdb/gremlinDatabase:GremlinDatabase":
		r = &GremlinDatabase{}
	case "azure:cosmosdb/gremlinGraph:GremlinGraph":
		r = &GremlinGraph{}
	case "azure:cosmosdb/mongoCollection:MongoCollection":
		r = &MongoCollection{}
	case "azure:cosmosdb/mongoDatabase:MongoDatabase":
		r = &MongoDatabase{}
	case "azure:cosmosdb/sqlContainer:SqlContainer":
		r = &SqlContainer{}
	case "azure:cosmosdb/sqlDatabase:SqlDatabase":
		r = &SqlDatabase{}
	case "azure:cosmosdb/sqlStoredProcedure:SqlStoredProcedure":
		r = &SqlStoredProcedure{}
	case "azure:cosmosdb/table:Table":
		r = &Table{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/cassandraKeyspace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/cassandraTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/gremlinDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/gremlinGraph",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/mongoCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/mongoDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/sqlContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/sqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/sqlStoredProcedure",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"cosmosdb/table",
		&module{version},
	)
}
