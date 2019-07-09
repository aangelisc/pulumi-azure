// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing App Service.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.output(azure.appservice.getAppService({
 *     name: "search-app-service",
 *     resourceGroupName: "search-service",
 * }));
 * 
 * export const appServiceId = test.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/app_service.html.markdown.
 */
export function getAppService(args: GetAppServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceResult> {
    return pulumi.runtime.invoke("azure:appservice/getAppService:getAppService", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppService.
 */
export interface GetAppServiceArgs {
    /**
     * The name of the App Service.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group where the App Service exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getAppService.
 */
export interface GetAppServiceResult {
    /**
     * The ID of the App Service Plan within which the App Service exists.
     */
    readonly appServicePlanId: string;
    /**
     * A key-value pair of App Settings for the App Service.
     */
    readonly appSettings: {[key: string]: any};
    /**
     * Does the App Service send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled: boolean;
    /**
     * Does the App Service require client certificates for incoming requests?
     */
    readonly clientCertEnabled: boolean;
    /**
     * An `connection_string` block as defined below.
     */
    readonly connectionStrings: { name: string, type: string, value: string }[];
    readonly defaultSiteHostname: string;
    /**
     * Is the App Service Enabled?
     */
    readonly enabled: boolean;
    /**
     * Can the App Service only be accessed via HTTPS?
     */
    readonly httpsOnly: boolean;
    /**
     * The Azure location where the App Service exists.
     */
    readonly location: string;
    /**
     * The name of the Connection String.
     */
    readonly name: string;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    readonly outboundIpAddresses: string;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
     */
    readonly possibleOutboundIpAddresses: string;
    readonly resourceGroupName: string;
    /**
     * A `site_config` block as defined below.
     */
    readonly siteConfigs: { alwaysOn: boolean, appCommandLine: string, cors: { allowedOrigins: string[], supportCredentials: boolean }, defaultDocuments: string[], dotnetFrameworkVersion: string, ftpsState: string, http2Enabled: boolean, ipRestrictions: { ipAddress: string, subnetMask: string }[], javaContainer: string, javaContainerVersion: string, javaVersion: string, linuxFxVersion: string, localMysqlEnabled: boolean, managedPipelineMode: string, minTlsVersion: string, phpVersion: string, pythonVersion: string, remoteDebuggingEnabled: boolean, remoteDebuggingVersion: string, scmType: string, use32BitWorkerProcess: boolean, virtualNetworkName: string, websocketsEnabled: boolean, windowsFxVersion: string }[];
    readonly siteCredentials: { password: string, username: string }[];
    readonly sourceControls: { branch: string, repoUrl: string }[];
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
