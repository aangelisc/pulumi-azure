// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to obtain information about a DNS Zone.
 */
export function getZone(args: GetZoneArgs): Promise<GetZoneResult> {
    return pulumi.runtime.invoke("azure:dns/getZone:getZone", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    });
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    /**
     * The name of the DNS Zone.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the DNS Zone exists.
     * If the Name of the Resource Group is not provided, the first DNS Zone from the list of DNS Zones
     * in your subscription that matches `name` will be returned.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    /**
     * Maximum number of Records in the zone.
     */
    readonly maxNumberOfRecordSets: string;
    /**
     * A list of values that make up the NS record for the zone.
     */
    readonly nameServers: string[];
    /**
     * The number of records already in the zone.
     */
    readonly numberOfRecordSets: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags to assign to the EventHub Namespace.
     */
    readonly tags: {[key: string]: any};
}
