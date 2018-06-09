// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Gets information about a Route Table
 */
export function getRouteTable(args: GetRouteTableArgs): Promise<GetRouteTableResult> {
    return pulumi.runtime.invoke("azure:network/getRouteTable:getRouteTable", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    });
}

/**
 * A collection of arguments for invoking getRouteTable.
 */
export interface GetRouteTableArgs {
    /**
     * The name of the Route Table.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Route Table exists.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}

/**
 * A collection of values returned by getRouteTable.
 */
export interface GetRouteTableResult {
    /**
     * The Azure Region in which the Route Table exists.
     */
    readonly location: string;
    /**
     * One or more `route` blocks as documented below.
     */
    readonly routes: { addressPrefix: string, name: string, nextHopInIpAddress: string, nextHopType: string }[];
    /**
     * The collection of Subnets associated with this route table.
     */
    readonly subnets: string[];
    /**
     * A mapping of tags assigned to the Route Table.
     */
    readonly tags: {[key: string]: any};
}
