// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Linked Service (connection) between Azure Databricks and Azure Data Factory.
 *
 * ## Example Usage
 * ### With Managed Identity & New Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "East US"});
 * //Create a Linked Service using managed identity and new cluster config
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * //Create a databricks instance
 * const exampleWorkspace = new azure.databricks.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: "standard",
 * });
 * const msiLinked = new azure.datafactory.LinkedServiceAzureDatabricks("msiLinked", {
 *     dataFactoryName: exampleFactory.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     description: "ADB Linked Service via MSI",
 *     adbDomain: pulumi.interpolate`https://${exampleWorkspace.workspaceUrl}`,
 *     msiWorkSpaceResourceId: exampleWorkspace.id,
 *     newClusterConfig: {
 *         nodeType: "Standard_NC12",
 *         clusterVersion: "5.5.x-gpu-scala2.11",
 *         minNumberOfWorkers: 1,
 *         maxNumberOfWorkers: 5,
 *         driverNodeType: "Standard_NC12",
 *         logDestination: "dbfs:/logs",
 *         customTags: {
 *             custom_tag1: "sct_value_1",
 *             custom_tag2: "sct_value_2",
 *         },
 *         sparkConfig: {
 *             config1: "value1",
 *             config2: "value2",
 *         },
 *         sparkEnvironmentVariables: {
 *             envVar1: "value1",
 *             envVar2: "value2",
 *         },
 *         initScripts: [
 *             "init.sh",
 *             "init2.sh",
 *         ],
 *     },
 * });
 * ```
 * ### With Access Token & Existing Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "East US"});
 * //Link to an existing cluster via access token
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * //Create a databricks instance
 * const exampleWorkspace = new azure.databricks.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: "standard",
 * });
 * const atLinked = new azure.datafactory.LinkedServiceAzureDatabricks("atLinked", {
 *     dataFactoryName: exampleFactory.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     description: "ADB Linked Service via Access Token",
 *     existingClusterId: "0308-201146-sly615",
 *     accessToken: "SomeDatabricksAccessToken",
 *     adbDomain: pulumi.interpolate`https://${exampleWorkspace.workspaceUrl}`,
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Linked Services can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/linkedServiceAzureDatabricks:LinkedServiceAzureDatabricks example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
 * ```
 */
export class LinkedServiceAzureDatabricks extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServiceAzureDatabricks resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServiceAzureDatabricksState, opts?: pulumi.CustomResourceOptions): LinkedServiceAzureDatabricks {
        return new LinkedServiceAzureDatabricks(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/linkedServiceAzureDatabricks:LinkedServiceAzureDatabricks';

    /**
     * Returns true if the given object is an instance of LinkedServiceAzureDatabricks.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServiceAzureDatabricks {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServiceAzureDatabricks.__pulumiType;
    }

    /**
     * Authenticate to ADB via an access token.
     */
    public readonly accessToken!: pulumi.Output<string | undefined>;
    /**
     * The domain URL of the databricks instance.
     */
    public readonly adbDomain!: pulumi.Output<string>;
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The clusterId of an existing cluster within the linked ADB instance.
     */
    public readonly existingClusterId!: pulumi.Output<string | undefined>;
    /**
     * Leverages an instance pool within the linked ADB instance as defined by  `instancePool` block below.
     */
    public readonly instancePool!: pulumi.Output<outputs.datafactory.LinkedServiceAzureDatabricksInstancePool | undefined>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    public readonly integrationRuntimeName!: pulumi.Output<string | undefined>;
    /**
     * Authenticate to ADB via Azure Key Vault Linked Service as defined in the `keyVaultPassword` block below.
     */
    public readonly keyVaultPassword!: pulumi.Output<outputs.datafactory.LinkedServiceAzureDatabricksKeyVaultPassword | undefined>;
    /**
     * Authenticate to ADB via managed service identity.
     */
    public readonly msiWorkSpaceResourceId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates new clusters within the linked ADB instance as defined in the  `newClusterConfig` block below.
     */
    public readonly newClusterConfig!: pulumi.Output<outputs.datafactory.LinkedServiceAzureDatabricksNewClusterConfig | undefined>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a LinkedServiceAzureDatabricks resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceAzureDatabricksArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceAzureDatabricksArgs | LinkedServiceAzureDatabricksState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LinkedServiceAzureDatabricksState | undefined;
            inputs["accessToken"] = state ? state.accessToken : undefined;
            inputs["adbDomain"] = state ? state.adbDomain : undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["existingClusterId"] = state ? state.existingClusterId : undefined;
            inputs["instancePool"] = state ? state.instancePool : undefined;
            inputs["integrationRuntimeName"] = state ? state.integrationRuntimeName : undefined;
            inputs["keyVaultPassword"] = state ? state.keyVaultPassword : undefined;
            inputs["msiWorkSpaceResourceId"] = state ? state.msiWorkSpaceResourceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["newClusterConfig"] = state ? state.newClusterConfig : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as LinkedServiceAzureDatabricksArgs | undefined;
            if ((!args || args.adbDomain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adbDomain'");
            }
            if ((!args || args.dataFactoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accessToken"] = args ? args.accessToken : undefined;
            inputs["adbDomain"] = args ? args.adbDomain : undefined;
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["existingClusterId"] = args ? args.existingClusterId : undefined;
            inputs["instancePool"] = args ? args.instancePool : undefined;
            inputs["integrationRuntimeName"] = args ? args.integrationRuntimeName : undefined;
            inputs["keyVaultPassword"] = args ? args.keyVaultPassword : undefined;
            inputs["msiWorkSpaceResourceId"] = args ? args.msiWorkSpaceResourceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["newClusterConfig"] = args ? args.newClusterConfig : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LinkedServiceAzureDatabricks.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServiceAzureDatabricks resources.
 */
export interface LinkedServiceAzureDatabricksState {
    /**
     * Authenticate to ADB via an access token.
     */
    accessToken?: pulumi.Input<string>;
    /**
     * The domain URL of the databricks instance.
     */
    adbDomain?: pulumi.Input<string>;
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    description?: pulumi.Input<string>;
    /**
     * The clusterId of an existing cluster within the linked ADB instance.
     */
    existingClusterId?: pulumi.Input<string>;
    /**
     * Leverages an instance pool within the linked ADB instance as defined by  `instancePool` block below.
     */
    instancePool?: pulumi.Input<inputs.datafactory.LinkedServiceAzureDatabricksInstancePool>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Authenticate to ADB via Azure Key Vault Linked Service as defined in the `keyVaultPassword` block below.
     */
    keyVaultPassword?: pulumi.Input<inputs.datafactory.LinkedServiceAzureDatabricksKeyVaultPassword>;
    /**
     * Authenticate to ADB via managed service identity.
     */
    msiWorkSpaceResourceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates new clusters within the linked ADB instance as defined in the  `newClusterConfig` block below.
     */
    newClusterConfig?: pulumi.Input<inputs.datafactory.LinkedServiceAzureDatabricksNewClusterConfig>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource.
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinkedServiceAzureDatabricks resource.
 */
export interface LinkedServiceAzureDatabricksArgs {
    /**
     * Authenticate to ADB via an access token.
     */
    accessToken?: pulumi.Input<string>;
    /**
     * The domain URL of the databricks instance.
     */
    adbDomain: pulumi.Input<string>;
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    description?: pulumi.Input<string>;
    /**
     * The clusterId of an existing cluster within the linked ADB instance.
     */
    existingClusterId?: pulumi.Input<string>;
    /**
     * Leverages an instance pool within the linked ADB instance as defined by  `instancePool` block below.
     */
    instancePool?: pulumi.Input<inputs.datafactory.LinkedServiceAzureDatabricksInstancePool>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Authenticate to ADB via Azure Key Vault Linked Service as defined in the `keyVaultPassword` block below.
     */
    keyVaultPassword?: pulumi.Input<inputs.datafactory.LinkedServiceAzureDatabricksKeyVaultPassword>;
    /**
     * Authenticate to ADB via managed service identity.
     */
    msiWorkSpaceResourceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates new clusters within the linked ADB instance as defined in the  `newClusterConfig` block below.
     */
    newClusterConfig?: pulumi.Input<inputs.datafactory.LinkedServiceAzureDatabricksNewClusterConfig>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource.
     */
    resourceGroupName: pulumi.Input<string>;
}
