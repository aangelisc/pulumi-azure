// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages as an Azure Container Group instance.
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
 * const exampleGroup = new azure.containerservice.Group("exampleGroup", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ipAddressType: "public",
 *     dnsNameLabel: "aci-label",
 *     osType: "Linux",
 *     container: [
 *         {
 *             name: "hello-world",
 *             image: "microsoft/aci-helloworld:latest",
 *             cpu: "0.5",
 *             memory: "1.5",
 *             ports: [{
 *                 port: 443,
 *                 protocol: "TCP",
 *             }],
 *         },
 *         {
 *             name: "sidecar",
 *             image: "microsoft/aci-tutorial-sidecar",
 *             cpu: "0.5",
 *             memory: "1.5",
 *         },
 *     ],
 *     tags: {
 *         environment: "testing",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_group.html.markdown.
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
    public static readonly __pulumiType = 'azure:containerservice/group:Group';

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
     * The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
     */
    public readonly containers!: pulumi.Output<outputs.containerservice.GroupContainer[]>;
    /**
     * A `diagnostics` block as documented below.
     */
    public readonly diagnostics!: pulumi.Output<outputs.containerservice.GroupDiagnostics | undefined>;
    /**
     * The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
     */
    public readonly dnsNameLabel!: pulumi.Output<string | undefined>;
    /**
     * The FQDN of the container group derived from `dnsNameLabel`.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.containerservice.GroupIdentity>;
    /**
     * A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
     */
    public readonly imageRegistryCredentials!: pulumi.Output<outputs.containerservice.GroupImageRegistryCredential[] | undefined>;
    /**
     * The IP address allocated to the container group.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
     */
    public readonly ipAddressType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Container Group. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network profile ID for deploying to virtual network.
     */
    public readonly networkProfileId!: pulumi.Output<string | undefined>;
    /**
     * The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
     */
    public readonly restartPolicy!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

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
            const state = argsOrState as GroupState | undefined;
            inputs["containers"] = state ? state.containers : undefined;
            inputs["diagnostics"] = state ? state.diagnostics : undefined;
            inputs["dnsNameLabel"] = state ? state.dnsNameLabel : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["imageRegistryCredentials"] = state ? state.imageRegistryCredentials : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkProfileId"] = state ? state.networkProfileId : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["restartPolicy"] = state ? state.restartPolicy : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if (!args || args.containers === undefined) {
                throw new Error("Missing required property 'containers'");
            }
            if (!args || args.osType === undefined) {
                throw new Error("Missing required property 'osType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["containers"] = args ? args.containers : undefined;
            inputs["diagnostics"] = args ? args.diagnostics : undefined;
            inputs["dnsNameLabel"] = args ? args.dnsNameLabel : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            inputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkProfileId"] = args ? args.networkProfileId : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restartPolicy"] = args ? args.restartPolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["fqdn"] = undefined /*out*/;
            inputs["ipAddress"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
     */
    readonly containers?: pulumi.Input<pulumi.Input<inputs.containerservice.GroupContainer>[]>;
    /**
     * A `diagnostics` block as documented below.
     */
    readonly diagnostics?: pulumi.Input<inputs.containerservice.GroupDiagnostics>;
    /**
     * The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
     */
    readonly dnsNameLabel?: pulumi.Input<string>;
    /**
     * The FQDN of the container group derived from `dnsNameLabel`.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.containerservice.GroupIdentity>;
    /**
     * A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
     */
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.containerservice.GroupImageRegistryCredential>[]>;
    /**
     * The IP address allocated to the container group.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
     */
    readonly ipAddressType?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network profile ID for deploying to virtual network.
     */
    readonly networkProfileId?: pulumi.Input<string>;
    /**
     * The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
     */
    readonly restartPolicy?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
     */
    readonly containers: pulumi.Input<pulumi.Input<inputs.containerservice.GroupContainer>[]>;
    /**
     * A `diagnostics` block as documented below.
     */
    readonly diagnostics?: pulumi.Input<inputs.containerservice.GroupDiagnostics>;
    /**
     * The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
     */
    readonly dnsNameLabel?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.containerservice.GroupIdentity>;
    /**
     * A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
     */
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.containerservice.GroupImageRegistryCredential>[]>;
    /**
     * Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
     */
    readonly ipAddressType?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network profile ID for deploying to virtual network.
     */
    readonly networkProfileId?: pulumi.Input<string>;
    /**
     * The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
     */
    readonly osType: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
     */
    readonly restartPolicy?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
