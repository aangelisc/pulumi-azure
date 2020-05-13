// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing API Management Product.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.apimanagement.getProduct({
 *     productId: "my-product",
 *     apiManagementName: "example-apim",
 *     resourceGroupName: "search-service",
 * });
 * export const productTerms = example.then(example => example.terms);
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/api_management_product.html.markdown.
 */
export function getProduct(args: GetProductArgs, opts?: pulumi.InvokeOptions): Promise<GetProductResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:apimanagement/getProduct:getProduct", {
        "apiManagementName": args.apiManagementName,
        "productId": args.productId,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getProduct.
 */
export interface GetProductArgs {
    /**
     * The Name of the API Management Service in which this Product exists.
     */
    readonly apiManagementName: string;
    /**
     * The Identifier for the API Management Product.
     */
    readonly productId: string;
    /**
     * The Name of the Resource Group in which the API Management Service exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getProduct.
 */
export interface GetProductResult {
    readonly apiManagementName: string;
    /**
     * Do subscribers need to be approved prior to being able to use the Product?
     */
    readonly approvalRequired: boolean;
    /**
     * The description of this Product, which may include HTML formatting tags.
     */
    readonly description: string;
    /**
     * The Display Name for this API Management Product.
     */
    readonly displayName: string;
    readonly productId: string;
    /**
     * Is this Product Published?
     */
    readonly published: boolean;
    readonly resourceGroupName: string;
    /**
     * Is a Subscription required to access API's included in this Product?
     */
    readonly subscriptionRequired: boolean;
    /**
     * The number of subscriptions a user can have to this Product at the same time.
     */
    readonly subscriptionsLimit: number;
    /**
     * Any Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
     */
    readonly terms: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
