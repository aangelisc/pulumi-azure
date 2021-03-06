// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Enrollment Account Billing Scope.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/billing"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := billing.GetEnrollmentAccountScope(ctx, &billing.GetEnrollmentAccountScopeArgs{
// 			BillingAccountName:    "existing",
// 			EnrollmentAccountName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func GetEnrollmentAccountScope(ctx *pulumi.Context, args *GetEnrollmentAccountScopeArgs, opts ...pulumi.InvokeOption) (*GetEnrollmentAccountScopeResult, error) {
	var rv GetEnrollmentAccountScopeResult
	err := ctx.Invoke("azure:billing/getEnrollmentAccountScope:getEnrollmentAccountScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnrollmentAccountScope.
type GetEnrollmentAccountScopeArgs struct {
	// The Billing Account Name of the Enterprise Account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The Enrollment Account Name in the above Enterprise Account.
	EnrollmentAccountName string `pulumi:"enrollmentAccountName"`
}

// A collection of values returned by getEnrollmentAccountScope.
type GetEnrollmentAccountScopeResult struct {
	BillingAccountName    string `pulumi:"billingAccountName"`
	EnrollmentAccountName string `pulumi:"enrollmentAccountName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
