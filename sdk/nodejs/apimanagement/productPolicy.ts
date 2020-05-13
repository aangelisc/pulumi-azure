// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an API Management Product Policy
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleProduct = azure.apimanagement.getProduct({
 *     productId: "my-product",
 *     apiManagementName: "example-apim",
 *     resourceGroupName: "search-service",
 * });
 * const exampleProductPolicy = new azure.apimanagement.ProductPolicy("exampleProductPolicy", {
 *     productId: exampleProduct.then(exampleProduct => exampleProduct.productId),
 *     apiManagementName: exampleProduct.then(exampleProduct => exampleProduct.apiManagementName),
 *     resourceGroupName: exampleProduct.then(exampleProduct => exampleProduct.resourceGroupName),
 *     xmlContent: `<policies>
 *   <inbound>
 *     <find-and-replace from="xyz" to="abc" />
 *   </inbound>
 * </policies>
 * `,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_product_policy.html.markdown.
 */
export class ProductPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ProductPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductPolicyState, opts?: pulumi.CustomResourceOptions): ProductPolicy {
        return new ProductPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/productPolicy:ProductPolicy';

    /**
     * Returns true if the given object is an instance of ProductPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductPolicy.__pulumiType;
    }

    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The XML Content for this Policy.
     */
    public readonly xmlContent!: pulumi.Output<string>;
    /**
     * A link to a Policy XML Document, which must be publicly available.
     */
    public readonly xmlLink!: pulumi.Output<string | undefined>;

    /**
     * Create a ProductPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductPolicyArgs | ProductPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProductPolicyState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["productId"] = state ? state.productId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["xmlContent"] = state ? state.xmlContent : undefined;
            inputs["xmlLink"] = state ? state.xmlLink : undefined;
        } else {
            const args = argsOrState as ProductPolicyArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.productId === undefined) {
                throw new Error("Missing required property 'productId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["xmlContent"] = args ? args.xmlContent : undefined;
            inputs["xmlLink"] = args ? args.xmlLink : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProductPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductPolicy resources.
 */
export interface ProductPolicyState {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly productId?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The XML Content for this Policy.
     */
    readonly xmlContent?: pulumi.Input<string>;
    /**
     * A link to a Policy XML Document, which must be publicly available.
     */
    readonly xmlLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProductPolicy resource.
 */
export interface ProductPolicyArgs {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly productId: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The XML Content for this Policy.
     */
    readonly xmlContent?: pulumi.Input<string>;
    /**
     * A link to a Policy XML Document, which must be publicly available.
     */
    readonly xmlLink?: pulumi.Input<string>;
}
