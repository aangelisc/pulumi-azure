// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Management Group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const current = pulumi.output(azure.core.getSubscription({}));
 * const exampleParent = new azure.management.Group("exampleParent", {
 *     displayName: "ParentGroup",
 *     subscriptionIds: [current.subscriptionId],
 * });
 * const exampleChild = new azure.management.Group("exampleChild", {
 *     displayName: "ChildGroup",
 *     parentManagementGroupId: exampleParent.id,
 *     subscriptionIds: [current.subscriptionId],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/management_group.html.markdown.
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:management/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * A friendly name for this Management Group. If not specified, this'll be the same as the `groupId`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The ID of the Parent Management Group. Changing this forces a new resource to be created.
     */
    public readonly parentManagementGroupId!: pulumi.Output<string>;
    /**
     * A list of Subscription GUIDs which should be assigned to the Management Group.
     */
    public readonly subscriptionIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["parentManagementGroupId"] = state ? state.parentManagementGroupId : undefined;
            inputs["subscriptionIds"] = state ? state.subscriptionIds : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["parentManagementGroupId"] = args ? args.parentManagementGroupId : undefined;
            inputs["subscriptionIds"] = args ? args.subscriptionIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure:managementgroups/managementGroup:ManagementGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * A friendly name for this Management Group. If not specified, this'll be the same as the `groupId`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The ID of the Parent Management Group. Changing this forces a new resource to be created.
     */
    readonly parentManagementGroupId?: pulumi.Input<string>;
    /**
     * A list of Subscription GUIDs which should be assigned to the Management Group.
     */
    readonly subscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * A friendly name for this Management Group. If not specified, this'll be the same as the `groupId`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The ID of the Parent Management Group. Changing this forces a new resource to be created.
     */
    readonly parentManagementGroupId?: pulumi.Input<string>;
    /**
     * A list of Subscription GUIDs which should be assigned to the Management Group.
     */
    readonly subscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
}
