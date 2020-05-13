// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages Azure Batch Application instance.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleBatch/accountAccount = new azure.batch.Account("exampleBatch/accountAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     poolAllocationMode: "BatchService",
 *     storageAccountId: exampleAccount.id,
 * });
 * const exampleApplication = new azure.batch.Application("exampleApplication", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     accountName: exampleBatch / accountAccount.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/batch_application.html.markdown.
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:batch/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * The name of the Batch account. Changing this forces a new resource to be created.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * A value indicating whether packages within the application may be overwritten using the same version string. Defaults to `true`.
     */
    public readonly allowUpdates!: pulumi.Output<boolean | undefined>;
    /**
     * The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
     */
    public readonly defaultVersion!: pulumi.Output<string | undefined>;
    /**
     * The display name for the application.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The name of the application. This must be unique within the account. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group that contains the Batch account. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["allowUpdates"] = state ? state.allowUpdates : undefined;
            inputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["allowUpdates"] = args ? args.allowUpdates : undefined;
            inputs["defaultVersion"] = args ? args.defaultVersion : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * The name of the Batch account. Changing this forces a new resource to be created.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * A value indicating whether packages within the application may be overwritten using the same version string. Defaults to `true`.
     */
    readonly allowUpdates?: pulumi.Input<boolean>;
    /**
     * The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
     */
    readonly defaultVersion?: pulumi.Input<string>;
    /**
     * The display name for the application.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name of the application. This must be unique within the account. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The name of the Batch account. Changing this forces a new resource to be created.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * A value indicating whether packages within the application may be overwritten using the same version string. Defaults to `true`.
     */
    readonly allowUpdates?: pulumi.Input<boolean>;
    /**
     * The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
     */
    readonly defaultVersion?: pulumi.Input<string>;
    /**
     * The display name for the application.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name of the application. This must be unique within the account. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
