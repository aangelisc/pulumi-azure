// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState): Group {
        return new Group(name, <any>state, { id });
    }

    public readonly containers: pulumi.Output<{ command?: string, cpu: number, environmentVariables?: {[key: string]: any}, image: string, memory: number, name: string, port?: number, protocol?: string, volumes?: { mountPath: string, name: string, readOnly?: boolean, shareName: string, storageAccountKey: string, storageAccountName: string }[] }[]>;
    public readonly dnsNameLabel: pulumi.Output<string | undefined>;
    public /*out*/ readonly fqdn: pulumi.Output<string>;
    public readonly imageRegistryCredentials: pulumi.Output<{ password: string, server: string, username: string }[] | undefined>;
    public /*out*/ readonly ipAddress: pulumi.Output<string>;
    public readonly ipAddressType: pulumi.Output<string | undefined>;
    public readonly location: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public readonly osType: pulumi.Output<string>;
    public readonly resourceGroupName: pulumi.Output<string>;
    public readonly restartPolicy: pulumi.Output<string | undefined>;
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: GroupState = argsOrState as GroupState | undefined;
            inputs["containers"] = state ? state.containers : undefined;
            inputs["dnsNameLabel"] = state ? state.dnsNameLabel : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["imageRegistryCredentials"] = state ? state.imageRegistryCredentials : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["restartPolicy"] = state ? state.restartPolicy : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if (!args || args.containers === undefined) {
                throw new Error("Missing required property 'containers'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.osType === undefined) {
                throw new Error("Missing required property 'osType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["containers"] = args ? args.containers : undefined;
            inputs["dnsNameLabel"] = args ? args.dnsNameLabel : undefined;
            inputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            inputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restartPolicy"] = args ? args.restartPolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["fqdn"] = undefined /*out*/;
            inputs["ipAddress"] = undefined /*out*/;
        }
        super("azure:containerservice/group:Group", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    readonly containers?: pulumi.Input<pulumi.Input<{ command?: pulumi.Input<string>, cpu: pulumi.Input<number>, environmentVariables?: pulumi.Input<{[key: string]: any}>, image: pulumi.Input<string>, memory: pulumi.Input<number>, name: pulumi.Input<string>, port?: pulumi.Input<number>, protocol?: pulumi.Input<string>, volumes?: pulumi.Input<pulumi.Input<{ mountPath: pulumi.Input<string>, name: pulumi.Input<string>, readOnly?: pulumi.Input<boolean>, shareName: pulumi.Input<string>, storageAccountKey: pulumi.Input<string>, storageAccountName: pulumi.Input<string> }>[]> }>[]>;
    readonly dnsNameLabel?: pulumi.Input<string>;
    readonly fqdn?: pulumi.Input<string>;
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<{ password: pulumi.Input<string>, server: pulumi.Input<string>, username: pulumi.Input<string> }>[]>;
    readonly ipAddress?: pulumi.Input<string>;
    readonly ipAddressType?: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly osType?: pulumi.Input<string>;
    readonly resourceGroupName?: pulumi.Input<string>;
    readonly restartPolicy?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    readonly containers: pulumi.Input<pulumi.Input<{ command?: pulumi.Input<string>, cpu: pulumi.Input<number>, environmentVariables?: pulumi.Input<{[key: string]: any}>, image: pulumi.Input<string>, memory: pulumi.Input<number>, name: pulumi.Input<string>, port?: pulumi.Input<number>, protocol?: pulumi.Input<string>, volumes?: pulumi.Input<pulumi.Input<{ mountPath: pulumi.Input<string>, name: pulumi.Input<string>, readOnly?: pulumi.Input<boolean>, shareName: pulumi.Input<string>, storageAccountKey: pulumi.Input<string>, storageAccountName: pulumi.Input<string> }>[]> }>[]>;
    readonly dnsNameLabel?: pulumi.Input<string>;
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<{ password: pulumi.Input<string>, server: pulumi.Input<string>, username: pulumi.Input<string> }>[]>;
    readonly ipAddressType?: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly osType: pulumi.Input<string>;
    readonly resourceGroupName: pulumi.Input<string>;
    readonly restartPolicy?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
