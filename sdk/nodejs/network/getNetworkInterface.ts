// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to access the properties of an Azure Network Interface.
 */
export function getNetworkInterface(args: GetNetworkInterfaceArgs): Promise<GetNetworkInterfaceResult> {
    return pulumi.runtime.invoke("azure:network/getNetworkInterface:getNetworkInterface", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    });
}

/**
 * A collection of arguments for invoking getNetworkInterface.
 */
export interface GetNetworkInterfaceArgs {
    /**
     * Specifies the name of the Network Interface.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group the Network Interface is located in.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}

/**
 * A collection of values returned by getNetworkInterface.
 */
export interface GetNetworkInterfaceResult {
    /**
     * List of DNS servers applied to the specified network interface.
     */
    readonly appliedDnsServers: string[];
    /**
     * The list of DNS servers used by the specified network interface.
     */
    readonly dnsServers: string[];
    /**
     * Indicates if accelerated networking is set on the specified network interface.
     */
    readonly enableAcceleratedNetworking: boolean;
    /**
     * Indicate if IP forwarding is set on the specified network interface.
     */
    readonly enableIpForwarding: boolean;
    /**
     * The internal dns name label of the specified network interface.
     */
    readonly internalDnsNameLabel: string;
    /**
     * The internal FQDN associated to the specified network interface.
     */
    readonly internalFqdn: string;
    /**
     * The list of IP configurations associated to the specified network interface.
     */
    readonly ipConfigurations: { applicationGatewayBackendAddressPoolsIds: string[], loadBalancerBackendAddressPoolsIds: string[], loadBalancerInboundNatRulesIds: string[], name: string, primary: boolean, privateIpAddress: string, privateIpAddressAllocation: string, publicIpAddressId: string, subnetId: string }[];
    /**
     * The location of the specified network interface.
     */
    readonly location: string;
    /**
     * The MAC address used by the specified network interface.
     */
    readonly macAddress: string;
    /**
     * The ID of the network security group associated to the specified network interface.
     */
    readonly networkSecurityGroupId: string;
    /**
     * The primary private ip address associated to the specified network interface.
     */
    readonly privateIpAddress: string;
    /**
     * The list of private ip addresses associates to the specified network interface.
     */
    readonly privateIpAddresses: string[];
    /**
     * List the tags assocatied to the specified network interface.
     */
    readonly tags: {[key: string]: any};
    /**
     * The ID of the virtual machine that the specified network interface is attached to.
     */
    readonly virtualMachineId: string;
}
