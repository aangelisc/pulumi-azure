// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Create a LoadBalancer Rule.
 * 
 * ~> **NOTE** When using this resource, the LoadBalancer needs to have a FrontEnd IP Configuration Attached
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState): Rule {
        return new Rule(name, <any>state, { id });
    }

    /**
     * A reference to a Backend Address Pool over which this Load Balancing Rule operates.
     */
    public readonly backendAddressPoolId: pulumi.Output<string>;
    /**
     * The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
     */
    public readonly backendPort: pulumi.Output<number>;
    /**
     * Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.
     */
    public readonly enableFloatingIp: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly frontendIpConfigurationId: pulumi.Output<string>;
    /**
     * The name of the frontend IP configuration to which the rule is associated.
     */
    public readonly frontendIpConfigurationName: pulumi.Output<string>;
    /**
     * The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
     */
    public readonly frontendPort: pulumi.Output<number>;
    /**
     * Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.
     */
    public readonly idleTimeoutInMinutes: pulumi.Output<number>;
    /**
     * Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
     */
    public readonly loadDistribution: pulumi.Output<string>;
    /**
     * The ID of the LoadBalancer in which to create the Rule.
     */
    public readonly loadbalancerId: pulumi.Output<string>;
    public readonly location: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the LB Rule.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A reference to a Probe used by this Load Balancing Rule.
     */
    public readonly probeId: pulumi.Output<string>;
    /**
     * The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
     */
    public readonly protocol: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the resource.
     */
    public readonly resourceGroupName: pulumi.Output<string>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RuleState = argsOrState as RuleState | undefined;
            inputs["backendAddressPoolId"] = state ? state.backendAddressPoolId : undefined;
            inputs["backendPort"] = state ? state.backendPort : undefined;
            inputs["enableFloatingIp"] = state ? state.enableFloatingIp : undefined;
            inputs["frontendIpConfigurationId"] = state ? state.frontendIpConfigurationId : undefined;
            inputs["frontendIpConfigurationName"] = state ? state.frontendIpConfigurationName : undefined;
            inputs["frontendPort"] = state ? state.frontendPort : undefined;
            inputs["idleTimeoutInMinutes"] = state ? state.idleTimeoutInMinutes : undefined;
            inputs["loadDistribution"] = state ? state.loadDistribution : undefined;
            inputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["probeId"] = state ? state.probeId : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if (!args || args.backendPort === undefined) {
                throw new Error("Missing required property 'backendPort'");
            }
            if (!args || args.frontendIpConfigurationName === undefined) {
                throw new Error("Missing required property 'frontendIpConfigurationName'");
            }
            if (!args || args.frontendPort === undefined) {
                throw new Error("Missing required property 'frontendPort'");
            }
            if (!args || args.loadbalancerId === undefined) {
                throw new Error("Missing required property 'loadbalancerId'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backendAddressPoolId"] = args ? args.backendAddressPoolId : undefined;
            inputs["backendPort"] = args ? args.backendPort : undefined;
            inputs["enableFloatingIp"] = args ? args.enableFloatingIp : undefined;
            inputs["frontendIpConfigurationName"] = args ? args.frontendIpConfigurationName : undefined;
            inputs["frontendPort"] = args ? args.frontendPort : undefined;
            inputs["idleTimeoutInMinutes"] = args ? args.idleTimeoutInMinutes : undefined;
            inputs["loadDistribution"] = args ? args.loadDistribution : undefined;
            inputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["probeId"] = args ? args.probeId : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["frontendIpConfigurationId"] = undefined /*out*/;
        }
        super("azure:lb/rule:Rule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * A reference to a Backend Address Pool over which this Load Balancing Rule operates.
     */
    readonly backendAddressPoolId?: pulumi.Input<string>;
    /**
     * The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
     */
    readonly backendPort?: pulumi.Input<number>;
    /**
     * Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.
     */
    readonly enableFloatingIp?: pulumi.Input<boolean>;
    readonly frontendIpConfigurationId?: pulumi.Input<string>;
    /**
     * The name of the frontend IP configuration to which the rule is associated.
     */
    readonly frontendIpConfigurationName?: pulumi.Input<string>;
    /**
     * The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
     */
    readonly frontendPort?: pulumi.Input<number>;
    /**
     * Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.
     */
    readonly idleTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
     */
    readonly loadDistribution?: pulumi.Input<string>;
    /**
     * The ID of the LoadBalancer in which to create the Rule.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the LB Rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A reference to a Probe used by this Load Balancing Rule.
     */
    readonly probeId?: pulumi.Input<string>;
    /**
     * The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * A reference to a Backend Address Pool over which this Load Balancing Rule operates.
     */
    readonly backendAddressPoolId?: pulumi.Input<string>;
    /**
     * The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
     */
    readonly backendPort: pulumi.Input<number>;
    /**
     * Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.
     */
    readonly enableFloatingIp?: pulumi.Input<boolean>;
    /**
     * The name of the frontend IP configuration to which the rule is associated.
     */
    readonly frontendIpConfigurationName: pulumi.Input<string>;
    /**
     * The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
     */
    readonly frontendPort: pulumi.Input<number>;
    /**
     * Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.
     */
    readonly idleTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: `Default` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. `SourceIP` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. `SourceIPProtocol` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where  the options are called `None`, `Client IP` and `Client IP and Protocol` respectively.
     */
    readonly loadDistribution?: pulumi.Input<string>;
    /**
     * The ID of the LoadBalancer in which to create the Rule.
     */
    readonly loadbalancerId: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the LB Rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A reference to a Probe used by this Load Balancing Rule.
     */
    readonly probeId?: pulumi.Input<string>;
    /**
     * The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
