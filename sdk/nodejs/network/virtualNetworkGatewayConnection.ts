// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a connection in an existing Virtual Network Gateway.
 *
 * ## Example Usage
 *
 * ### Site-to-Site connection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefix: "10.0.1.0/24",
 * });
 * const onpremiseLocalNetworkGateway = new azure.network.LocalNetworkGateway("onpremiseLocalNetworkGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     gatewayAddress: "168.62.225.23",
 *     addressSpaces: ["10.1.1.0/24"],
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const exampleVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("exampleVirtualNetworkGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     activeActive: false,
 *     enableBgp: false,
 *     sku: "Basic",
 *     ip_configuration: [{
 *         publicIpAddressId: examplePublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: exampleSubnet.id,
 *     }],
 * });
 * const onpremiseVirtualNetworkGatewayConnection = new azure.network.VirtualNetworkGatewayConnection("onpremiseVirtualNetworkGatewayConnection", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     type: "IPsec",
 *     virtualNetworkGatewayId: exampleVirtualNetworkGateway.id,
 *     localNetworkGatewayId: onpremiseLocalNetworkGateway.id,
 *     sharedKey: "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
 * });
 * ```
 *
 * ### VNet-to-VNet connection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const usResourceGroup = new azure.core.ResourceGroup("usResourceGroup", {location: "East US"});
 * const usVirtualNetwork = new azure.network.VirtualNetwork("usVirtualNetwork", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const usGateway = new azure.network.Subnet("usGateway", {
 *     resourceGroupName: usResourceGroup.name,
 *     virtualNetworkName: usVirtualNetwork.name,
 *     addressPrefix: "10.0.1.0/24",
 * });
 * const usPublicIp = new azure.network.PublicIp("usPublicIp", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const usVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("usVirtualNetworkGateway", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     sku: "Basic",
 *     ip_configuration: [{
 *         publicIpAddressId: usPublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: usGateway.id,
 *     }],
 * });
 * const europeResourceGroup = new azure.core.ResourceGroup("europeResourceGroup", {location: "West Europe"});
 * const europeVirtualNetwork = new azure.network.VirtualNetwork("europeVirtualNetwork", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     addressSpaces: ["10.1.0.0/16"],
 * });
 * const europeGateway = new azure.network.Subnet("europeGateway", {
 *     resourceGroupName: europeResourceGroup.name,
 *     virtualNetworkName: europeVirtualNetwork.name,
 *     addressPrefix: "10.1.1.0/24",
 * });
 * const europePublicIp = new azure.network.PublicIp("europePublicIp", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const europeVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("europeVirtualNetworkGateway", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     sku: "Basic",
 *     ip_configuration: [{
 *         publicIpAddressId: europePublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: europeGateway.id,
 *     }],
 * });
 * const usToEurope = new azure.network.VirtualNetworkGatewayConnection("usToEurope", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     type: "Vnet2Vnet",
 *     virtualNetworkGatewayId: usVirtualNetworkGateway.id,
 *     peerVirtualNetworkGatewayId: europeVirtualNetworkGateway.id,
 *     sharedKey: "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
 * });
 * const europeToUs = new azure.network.VirtualNetworkGatewayConnection("europeToUs", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     type: "Vnet2Vnet",
 *     virtualNetworkGatewayId: europeVirtualNetworkGateway.id,
 *     peerVirtualNetworkGatewayId: usVirtualNetworkGateway.id,
 *     sharedKey: "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_network_gateway_connection.html.markdown.
 */
export class VirtualNetworkGatewayConnection extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualNetworkGatewayConnectionState, opts?: pulumi.CustomResourceOptions): VirtualNetworkGatewayConnection {
        return new VirtualNetworkGatewayConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection';

    /**
     * Returns true if the given object is an instance of VirtualNetworkGatewayConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkGatewayConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkGatewayConnection.__pulumiType;
    }

    /**
     * The authorization key associated with the
     * Express Route Circuit. This field is required only if the type is an
     * ExpressRoute connection.
     */
    public readonly authorizationKey!: pulumi.Output<string | undefined>;
    /**
     * The IKE protocol version to use. Possible
     * values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
     * Changing this value will force a resource to be created.
     * > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
     */
    public readonly connectionProtocol!: pulumi.Output<string>;
    /**
     * If `true`, BGP (Border Gateway Protocol) is enabled
     * for this connection. Defaults to `false`.
     */
    public readonly enableBgp!: pulumi.Output<boolean>;
    /**
     * The ID of the Express Route Circuit
     * when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
     * The Express Route Circuit can be in the same or in a different subscription.
     */
    public readonly expressRouteCircuitId!: pulumi.Output<string | undefined>;
    /**
     * If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
     */
    public readonly expressRouteGatewayBypass!: pulumi.Output<boolean>;
    /**
     * A `ipsecPolicy` block which is documented below.
     * Only a single policy can be defined for a connection. For details on
     * custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
     */
    public readonly ipsecPolicy!: pulumi.Output<outputs.network.VirtualNetworkGatewayConnectionIpsecPolicy | undefined>;
    /**
     * The ID of the local network gateway
     * when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
     */
    public readonly localNetworkGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The location/region where the connection is
     * located. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the connection. Changing the name forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the peer virtual
     * network gateway when creating a VNet-to-VNet connection (i.e. when `type`
     * is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
     * in a different subscription.
     */
    public readonly peerVirtualNetworkGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to
     * create the connection Changing the name forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The routing weight. Defaults to `10`.
     */
    public readonly routingWeight!: pulumi.Output<number>;
    /**
     * The shared IPSec key. A key could be provided if a
     * Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
     */
    public readonly sharedKey!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of connection. Valid options are `IPsec`
     * (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
     * Each connection type requires different mandatory arguments (refer to the
     * examples above). Changing the connection type will force a new connection
     * to be created.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * If `true`, policy-based traffic
     * selectors are enabled for this connection. Enabling policy-based traffic
     * selectors requires an `ipsecPolicy` block. Defaults to `false`.
     */
    public readonly usePolicyBasedTrafficSelectors!: pulumi.Output<boolean>;
    /**
     * The ID of the Virtual Network Gateway
     * in which the connection will be created. Changing the gateway forces a new
     * resource to be created.
     */
    public readonly virtualNetworkGatewayId!: pulumi.Output<string>;

    /**
     * Create a VirtualNetworkGatewayConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkGatewayConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkGatewayConnectionArgs | VirtualNetworkGatewayConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualNetworkGatewayConnectionState | undefined;
            inputs["authorizationKey"] = state ? state.authorizationKey : undefined;
            inputs["connectionProtocol"] = state ? state.connectionProtocol : undefined;
            inputs["enableBgp"] = state ? state.enableBgp : undefined;
            inputs["expressRouteCircuitId"] = state ? state.expressRouteCircuitId : undefined;
            inputs["expressRouteGatewayBypass"] = state ? state.expressRouteGatewayBypass : undefined;
            inputs["ipsecPolicy"] = state ? state.ipsecPolicy : undefined;
            inputs["localNetworkGatewayId"] = state ? state.localNetworkGatewayId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["peerVirtualNetworkGatewayId"] = state ? state.peerVirtualNetworkGatewayId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["routingWeight"] = state ? state.routingWeight : undefined;
            inputs["sharedKey"] = state ? state.sharedKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["usePolicyBasedTrafficSelectors"] = state ? state.usePolicyBasedTrafficSelectors : undefined;
            inputs["virtualNetworkGatewayId"] = state ? state.virtualNetworkGatewayId : undefined;
        } else {
            const args = argsOrState as VirtualNetworkGatewayConnectionArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.virtualNetworkGatewayId === undefined) {
                throw new Error("Missing required property 'virtualNetworkGatewayId'");
            }
            inputs["authorizationKey"] = args ? args.authorizationKey : undefined;
            inputs["connectionProtocol"] = args ? args.connectionProtocol : undefined;
            inputs["enableBgp"] = args ? args.enableBgp : undefined;
            inputs["expressRouteCircuitId"] = args ? args.expressRouteCircuitId : undefined;
            inputs["expressRouteGatewayBypass"] = args ? args.expressRouteGatewayBypass : undefined;
            inputs["ipsecPolicy"] = args ? args.ipsecPolicy : undefined;
            inputs["localNetworkGatewayId"] = args ? args.localNetworkGatewayId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peerVirtualNetworkGatewayId"] = args ? args.peerVirtualNetworkGatewayId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routingWeight"] = args ? args.routingWeight : undefined;
            inputs["sharedKey"] = args ? args.sharedKey : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["usePolicyBasedTrafficSelectors"] = args ? args.usePolicyBasedTrafficSelectors : undefined;
            inputs["virtualNetworkGatewayId"] = args ? args.virtualNetworkGatewayId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualNetworkGatewayConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualNetworkGatewayConnection resources.
 */
export interface VirtualNetworkGatewayConnectionState {
    /**
     * The authorization key associated with the
     * Express Route Circuit. This field is required only if the type is an
     * ExpressRoute connection.
     */
    readonly authorizationKey?: pulumi.Input<string>;
    /**
     * The IKE protocol version to use. Possible
     * values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
     * Changing this value will force a resource to be created.
     * > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
     */
    readonly connectionProtocol?: pulumi.Input<string>;
    /**
     * If `true`, BGP (Border Gateway Protocol) is enabled
     * for this connection. Defaults to `false`.
     */
    readonly enableBgp?: pulumi.Input<boolean>;
    /**
     * The ID of the Express Route Circuit
     * when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
     * The Express Route Circuit can be in the same or in a different subscription.
     */
    readonly expressRouteCircuitId?: pulumi.Input<string>;
    /**
     * If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
     */
    readonly expressRouteGatewayBypass?: pulumi.Input<boolean>;
    /**
     * A `ipsecPolicy` block which is documented below.
     * Only a single policy can be defined for a connection. For details on
     * custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
     */
    readonly ipsecPolicy?: pulumi.Input<inputs.network.VirtualNetworkGatewayConnectionIpsecPolicy>;
    /**
     * The ID of the local network gateway
     * when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
     */
    readonly localNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The location/region where the connection is
     * located. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the connection. Changing the name forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the peer virtual
     * network gateway when creating a VNet-to-VNet connection (i.e. when `type`
     * is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
     * in a different subscription.
     */
    readonly peerVirtualNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the connection Changing the name forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The routing weight. Defaults to `10`.
     */
    readonly routingWeight?: pulumi.Input<number>;
    /**
     * The shared IPSec key. A key could be provided if a
     * Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
     */
    readonly sharedKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of connection. Valid options are `IPsec`
     * (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
     * Each connection type requires different mandatory arguments (refer to the
     * examples above). Changing the connection type will force a new connection
     * to be created.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * If `true`, policy-based traffic
     * selectors are enabled for this connection. Enabling policy-based traffic
     * selectors requires an `ipsecPolicy` block. Defaults to `false`.
     */
    readonly usePolicyBasedTrafficSelectors?: pulumi.Input<boolean>;
    /**
     * The ID of the Virtual Network Gateway
     * in which the connection will be created. Changing the gateway forces a new
     * resource to be created.
     */
    readonly virtualNetworkGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualNetworkGatewayConnection resource.
 */
export interface VirtualNetworkGatewayConnectionArgs {
    /**
     * The authorization key associated with the
     * Express Route Circuit. This field is required only if the type is an
     * ExpressRoute connection.
     */
    readonly authorizationKey?: pulumi.Input<string>;
    /**
     * The IKE protocol version to use. Possible
     * values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
     * Changing this value will force a resource to be created.
     * > **Note**: Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
     */
    readonly connectionProtocol?: pulumi.Input<string>;
    /**
     * If `true`, BGP (Border Gateway Protocol) is enabled
     * for this connection. Defaults to `false`.
     */
    readonly enableBgp?: pulumi.Input<boolean>;
    /**
     * The ID of the Express Route Circuit
     * when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
     * The Express Route Circuit can be in the same or in a different subscription.
     */
    readonly expressRouteCircuitId?: pulumi.Input<string>;
    /**
     * If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
     */
    readonly expressRouteGatewayBypass?: pulumi.Input<boolean>;
    /**
     * A `ipsecPolicy` block which is documented below.
     * Only a single policy can be defined for a connection. For details on
     * custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
     */
    readonly ipsecPolicy?: pulumi.Input<inputs.network.VirtualNetworkGatewayConnectionIpsecPolicy>;
    /**
     * The ID of the local network gateway
     * when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
     */
    readonly localNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The location/region where the connection is
     * located. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the connection. Changing the name forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the peer virtual
     * network gateway when creating a VNet-to-VNet connection (i.e. when `type`
     * is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
     * in a different subscription.
     */
    readonly peerVirtualNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the connection Changing the name forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The routing weight. Defaults to `10`.
     */
    readonly routingWeight?: pulumi.Input<number>;
    /**
     * The shared IPSec key. A key could be provided if a
     * Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
     */
    readonly sharedKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of connection. Valid options are `IPsec`
     * (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
     * Each connection type requires different mandatory arguments (refer to the
     * examples above). Changing the connection type will force a new connection
     * to be created.
     */
    readonly type: pulumi.Input<string>;
    /**
     * If `true`, policy-based traffic
     * selectors are enabled for this connection. Enabling policy-based traffic
     * selectors requires an `ipsecPolicy` block. Defaults to `false`.
     */
    readonly usePolicyBasedTrafficSelectors?: pulumi.Input<boolean>;
    /**
     * The ID of the Virtual Network Gateway
     * in which the connection will be created. Changing the gateway forces a new
     * resource to be created.
     */
    readonly virtualNetworkGatewayId: pulumi.Input<string>;
}
