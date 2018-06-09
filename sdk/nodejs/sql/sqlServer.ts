// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a SQL Azure Database Server.
 * 
 * ~> **Note:** All arguments including the administrator login and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](/docs/state/sensitive-data.html).
 */
export class SqlServer extends pulumi.CustomResource {
    /**
     * Get an existing SqlServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SqlServerState): SqlServer {
        return new SqlServer(name, <any>state, { id });
    }

    /**
     * The administrator login name for the new server. Changing this forces a new resource to be created.
     */
    public readonly administratorLogin: pulumi.Output<string>;
    /**
     * The password associated with the `administrator_login` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
     */
    public readonly administratorLoginPassword: pulumi.Output<string>;
    /**
     * The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
     */
    public /*out*/ readonly fullyQualifiedDomainName: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * The name of the SQL Server. This needs to be globally unique within Azure.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the SQL Server.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;
    /**
     * The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
     */
    public readonly version: pulumi.Output<string>;

    /**
     * Create a SqlServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlServerArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: SqlServerArgs | SqlServerState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SqlServerState = argsOrState as SqlServerState | undefined;
            inputs["administratorLogin"] = state ? state.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = state ? state.administratorLoginPassword : undefined;
            inputs["fullyQualifiedDomainName"] = state ? state.fullyQualifiedDomainName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as SqlServerArgs | undefined;
            if (!args || args.administratorLogin === undefined) {
                throw new Error("Missing required property 'administratorLogin'");
            }
            if (!args || args.administratorLoginPassword === undefined) {
                throw new Error("Missing required property 'administratorLoginPassword'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["administratorLogin"] = args ? args.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = args ? args.administratorLoginPassword : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["fullyQualifiedDomainName"] = undefined /*out*/;
        }
        super("azure:sql/sqlServer:SqlServer", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SqlServer resources.
 */
export interface SqlServerState {
    /**
     * The administrator login name for the new server. Changing this forces a new resource to be created.
     */
    readonly administratorLogin?: pulumi.Input<string>;
    /**
     * The password associated with the `administrator_login` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
     */
    readonly administratorLoginPassword?: pulumi.Input<string>;
    /**
     * The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
     */
    readonly fullyQualifiedDomainName?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the SQL Server. This needs to be globally unique within Azure.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the SQL Server.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SqlServer resource.
 */
export interface SqlServerArgs {
    /**
     * The administrator login name for the new server. Changing this forces a new resource to be created.
     */
    readonly administratorLogin: pulumi.Input<string>;
    /**
     * The password associated with the `administrator_login` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
     */
    readonly administratorLoginPassword: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the SQL Server. This needs to be globally unique within Azure.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the SQL Server.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
     */
    readonly version: pulumi.Input<string>;
}
