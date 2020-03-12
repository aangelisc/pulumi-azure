// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages an API Schema within an API Management Service.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_api_schema.html.markdown.
    /// </summary>
    public partial class ApiSchema : Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiName")]
        public Output<string> ApiName { get; private set; } = null!;

        /// <summary>
        /// The content type of the API Schema.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for this API Schema. Changing this forces a new resource to be created.
        /// </summary>
        [Output("schemaId")]
        public Output<string> SchemaId { get; private set; } = null!;

        /// <summary>
        /// The JSON escaped string defining the document representing the Schema.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ApiSchema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiSchema(string name, ApiSchemaArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/apiSchema:ApiSchema", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ApiSchema(string name, Input<string> id, ApiSchemaState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/apiSchema:ApiSchema", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiSchema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiSchema Get(string name, Input<string> id, ApiSchemaState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiSchema(name, id, state, options);
        }
    }

    public sealed class ApiSchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiName", required: true)]
        public Input<string> ApiName { get; set; } = null!;

        /// <summary>
        /// The content type of the API Schema.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A unique identifier for this API Schema. Changing this forces a new resource to be created.
        /// </summary>
        [Input("schemaId", required: true)]
        public Input<string> SchemaId { get; set; } = null!;

        /// <summary>
        /// The JSON escaped string defining the document representing the Schema.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ApiSchemaArgs()
        {
        }
    }

    public sealed class ApiSchemaState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiName")]
        public Input<string>? ApiName { get; set; }

        /// <summary>
        /// The content type of the API Schema.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A unique identifier for this API Schema. Changing this forces a new resource to be created.
        /// </summary>
        [Input("schemaId")]
        public Input<string>? SchemaId { get; set; }

        /// <summary>
        /// The JSON escaped string defining the document representing the Schema.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApiSchemaState()
        {
        }
    }
}
