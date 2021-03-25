// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package advisor

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Advisor Recommendations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/advisor"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := advisor.GetRecommendations(ctx, &advisor.GetRecommendationsArgs{
// 			FilterByCategories: []string{
// 				"security",
// 				"cost",
// 			},
// 			FilterByResourceGroups: []string{
// 				"example-resgroups",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("recommendations", example.Recommendations)
// 		return nil
// 	})
// }
// ```
func GetRecommendations(ctx *pulumi.Context, args *GetRecommendationsArgs, opts ...pulumi.InvokeOption) (*GetRecommendationsResult, error) {
	var rv GetRecommendationsResult
	err := ctx.Invoke("azure:advisor/getRecommendations:getRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecommendations.
type GetRecommendationsArgs struct {
	// Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
	FilterByCategories []string `pulumi:"filterByCategories"`
	// Specifies a list of resource groups about which the Advisor Recommendations will be listed.
	FilterByResourceGroups []string `pulumi:"filterByResourceGroups"`
}

// A collection of values returned by getRecommendations.
type GetRecommendationsResult struct {
	FilterByCategories     []string `pulumi:"filterByCategories"`
	FilterByResourceGroups []string `pulumi:"filterByResourceGroups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// One or more `recommendations` blocks as defined below.
	Recommendations []GetRecommendationsRecommendation `pulumi:"recommendations"`
}
