// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Sets a PostgreSQL Configuration value on a PostgreSQL Server.
 *
 * ## Disclaimers
 *
 * > **Note:** Since this resource is provisioned by default, the Azure Provider will not check for the presence of an existing resource prior to attempting to create it.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleServer = new azure.postgresql.Server("exampleServer", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "B_Gen5_2",
 *     storageMb: 5120,
 *     backupRetentionDays: 7,
 *     geoRedundantBackupEnabled: false,
 *     autoGrowEnabled: true,
 *     administratorLogin: "psqladminun",
 *     administratorLoginPassword: "H@Sh1CoR3!",
 *     version: "9.5",
 *     sslEnforcementEnabled: true,
 * });
 * const exampleConfiguration = new azure.postgresql.Configuration("exampleConfiguration", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serverName: exampleServer.name,
 *     value: "on",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/postgresql_configuration.html.markdown.
 */
export class Configuration extends pulumi.CustomResource {
    /**
     * Get an existing Configuration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationState, opts?: pulumi.CustomResourceOptions): Configuration {
        return new Configuration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:postgresql/configuration:Configuration';

    /**
     * Returns true if the given object is an instance of Configuration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Configuration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Configuration.__pulumiType;
    }

    /**
     * Specifies the name of the PostgreSQL Configuration, which needs [to be a valid PostgreSQL configuration name](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIER). Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Configuration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigurationArgs | ConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConfigurationState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ConfigurationArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Configuration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Configuration resources.
 */
export interface ConfigurationState {
    /**
     * Specifies the name of the PostgreSQL Configuration, which needs [to be a valid PostgreSQL configuration name](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIER). Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
     */
    readonly serverName?: pulumi.Input<string>;
    /**
     * Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Configuration resource.
 */
export interface ConfigurationArgs {
    /**
     * Specifies the name of the PostgreSQL Configuration, which needs [to be a valid PostgreSQL configuration name](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIER). Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
     */
    readonly value: pulumi.Input<string>;
}
