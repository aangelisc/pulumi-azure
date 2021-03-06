// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./analyticsItem";
export * from "./apiKey";
export * from "./getInsights";
export * from "./insights";
export * from "./smartDetectionRule";
export * from "./webTest";

// Import resources to register:
import { AnalyticsItem } from "./analyticsItem";
import { ApiKey } from "./apiKey";
import { Insights } from "./insights";
import { SmartDetectionRule } from "./smartDetectionRule";
import { WebTest } from "./webTest";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:appinsights/analyticsItem:AnalyticsItem":
                return new AnalyticsItem(name, <any>undefined, { urn })
            case "azure:appinsights/apiKey:ApiKey":
                return new ApiKey(name, <any>undefined, { urn })
            case "azure:appinsights/insights:Insights":
                return new Insights(name, <any>undefined, { urn })
            case "azure:appinsights/smartDetectionRule:SmartDetectionRule":
                return new SmartDetectionRule(name, <any>undefined, { urn })
            case "azure:appinsights/webTest:WebTest":
                return new WebTest(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "appinsights/analyticsItem", _module)
pulumi.runtime.registerResourceModule("azure", "appinsights/apiKey", _module)
pulumi.runtime.registerResourceModule("azure", "appinsights/insights", _module)
pulumi.runtime.registerResourceModule("azure", "appinsights/smartDetectionRule", _module)
pulumi.runtime.registerResourceModule("azure", "appinsights/webTest", _module)
