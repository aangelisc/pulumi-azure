// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getClusterNodePool";
export * from "./getKubernetesCluster";
export * from "./getKubernetesServiceVersions";
export * from "./getRegistry";
export * from "./getRegistryScopeMap";
export * from "./getRegistryToken";
export * from "./group";
export * from "./kubernetesCluster";
export * from "./kubernetesClusterNodePool";
export * from "./registry";
export * from "./registryScopeMap";
export * from "./registryToken";
export * from "./registryWebhook";
export * from "./registryWebook";

// Import resources to register:
import { Group } from "./group";
import { KubernetesCluster } from "./kubernetesCluster";
import { KubernetesClusterNodePool } from "./kubernetesClusterNodePool";
import { Registry } from "./registry";
import { RegistryScopeMap } from "./registryScopeMap";
import { RegistryToken } from "./registryToken";
import { RegistryWebhook } from "./registryWebhook";
import { RegistryWebook } from "./registryWebook";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:containerservice/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "azure:containerservice/kubernetesCluster:KubernetesCluster":
                return new KubernetesCluster(name, <any>undefined, { urn })
            case "azure:containerservice/kubernetesClusterNodePool:KubernetesClusterNodePool":
                return new KubernetesClusterNodePool(name, <any>undefined, { urn })
            case "azure:containerservice/registry:Registry":
                return new Registry(name, <any>undefined, { urn })
            case "azure:containerservice/registryScopeMap:RegistryScopeMap":
                return new RegistryScopeMap(name, <any>undefined, { urn })
            case "azure:containerservice/registryToken:RegistryToken":
                return new RegistryToken(name, <any>undefined, { urn })
            case "azure:containerservice/registryWebhook:RegistryWebhook":
                return new RegistryWebhook(name, <any>undefined, { urn })
            case "azure:containerservice/registryWebook:RegistryWebook":
                return new RegistryWebook(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "containerservice/group", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/kubernetesCluster", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/kubernetesClusterNodePool", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/registry", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/registryScopeMap", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/registryToken", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/registryWebhook", _module)
pulumi.runtime.registerResourceModule("azure", "containerservice/registryWebook", _module)
