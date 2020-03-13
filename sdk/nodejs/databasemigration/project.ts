// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:databasemigration/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specify the name of the database migration project. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
     */
    public readonly sourcePlatform!: pulumi.Output<string>;
    /**
     * A mapping of tags to assigned to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
     */
    public readonly targetPlatform!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["sourcePlatform"] = state ? state.sourcePlatform : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetPlatform"] = state ? state.targetPlatform : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.sourcePlatform === undefined) {
                throw new Error("Missing required property 'sourcePlatform'");
            }
            if (!args || args.targetPlatform === undefined) {
                throw new Error("Missing required property 'targetPlatform'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["sourcePlatform"] = args ? args.sourcePlatform : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetPlatform"] = args ? args.targetPlatform : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specify the name of the database migration project. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
     */
    readonly sourcePlatform?: pulumi.Input<string>;
    /**
     * A mapping of tags to assigned to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
     */
    readonly targetPlatform?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specify the name of the database migration project. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
     */
    readonly sourcePlatform: pulumi.Input<string>;
    /**
     * A mapping of tags to assigned to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
     */
    readonly targetPlatform: pulumi.Input<string>;
}
