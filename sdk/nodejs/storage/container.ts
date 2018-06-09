// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Create an Azure Storage Container.
 */
export class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerState): Container {
        return new Container(name, <any>state, { id });
    }

    /**
     * The 'interface' for access the container provides. Can be either `blob`, `container` or `private`. Defaults to `private`. Changing this forces a new resource to be created.
     */
    public readonly containerAccessType: pulumi.Output<string | undefined>;
    /**
     * The name of the storage container. Must be unique within the storage service the container is located.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Key-value definition of additional properties associated to the storage container
     */
    public /*out*/ readonly properties: pulumi.Output<{[key: string]: any}>;
    /**
     * The name of the resource group in which to
     * create the storage container. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * Specifies the storage account in which to create the storage container.
     * Changing this forces a new resource to be created.
     */
    public readonly storageAccountName: pulumi.Output<string>;

    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: ContainerArgs | ContainerState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ContainerState = argsOrState as ContainerState | undefined;
            inputs["containerAccessType"] = state ? state.containerAccessType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
        } else {
            const args = argsOrState as ContainerArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageAccountName === undefined) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            inputs["containerAccessType"] = args ? args.containerAccessType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            inputs["properties"] = undefined /*out*/;
        }
        super("azure:storage/container:Container", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Container resources.
 */
export interface ContainerState {
    /**
     * The 'interface' for access the container provides. Can be either `blob`, `container` or `private`. Defaults to `private`. Changing this forces a new resource to be created.
     */
    readonly containerAccessType?: pulumi.Input<string>;
    /**
     * The name of the storage container. Must be unique within the storage service the container is located.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value definition of additional properties associated to the storage container
     */
    readonly properties?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the resource group in which to
     * create the storage container. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the storage account in which to create the storage container.
     * Changing this forces a new resource to be created.
     */
    readonly storageAccountName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Container resource.
 */
export interface ContainerArgs {
    /**
     * The 'interface' for access the container provides. Can be either `blob`, `container` or `private`. Defaults to `private`. Changing this forces a new resource to be created.
     */
    readonly containerAccessType?: pulumi.Input<string>;
    /**
     * The name of the storage container. Must be unique within the storage service the container is located.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the storage container. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the storage account in which to create the storage container.
     * Changing this forces a new resource to be created.
     */
    readonly storageAccountName: pulumi.Input<string>;
}
