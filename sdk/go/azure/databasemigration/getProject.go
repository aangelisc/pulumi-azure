// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Database Migration Project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/databasemigration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := databasemigration.LookupProject(ctx, &databasemigration.LookupProjectArgs{
// 			Name:              "example-dbms-project",
// 			ResourceGroupName: "example-rg",
// 			ServiceName:       "example-dbms",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("name", example.Name)
// 		return nil
// 	})
// }
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure:databasemigration/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// Name of the database migration project.
	Name string `pulumi:"name"`
	// Name of the resource group where resource belongs to.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the database migration service where resource belongs to.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Azure location where the resource exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	// The platform type of the migration source.
	SourcePlatform string `pulumi:"sourcePlatform"`
	// A mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The platform type of the migration target.
	TargetPlatform string `pulumi:"targetPlatform"`
}
