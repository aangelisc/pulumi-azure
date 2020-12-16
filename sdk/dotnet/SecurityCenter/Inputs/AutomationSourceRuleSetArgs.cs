// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter.Inputs
{

    public sealed class AutomationSourceRuleSetArgs : Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.AutomationSourceRuleSetRuleArgs>? _rules;

        /// <summary>
        /// One or more `rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AutomationSourceRuleSetRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.AutomationSourceRuleSetRuleArgs>());
            set => _rules = value;
        }

        public AutomationSourceRuleSetArgs()
        {
        }
    }
}