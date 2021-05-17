// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataboxEdge
{
    /// <summary>
    /// Manages a Databox Edge Device.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleDevice = new Azure.DataboxEdge.Device("exampleDevice", new Azure.DataboxEdge.DeviceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             SkuName = "EdgeP_Base-Standard",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Databox Edge Devices can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:databoxedge/device:Device example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1
    /// ```
    /// </summary>
    [AzureResourceType("azure:databoxedge/device:Device")]
    public partial class Device : Pulumi.CustomResource
    {
        /// <summary>
        /// A `device_properties` block as defined below.
        /// </summary>
        [Output("deviceProperties")]
        public Output<ImmutableArray<Outputs.DeviceDeviceProperty>> DeviceProperties { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The `sku_name` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `sku_name` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `sku_name`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Databox Edge Device.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs args, CustomResourceOptions? options = null)
            : base("azure:databoxedge/device:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
            : base("azure:databoxedge/device:Device", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Device(name, id, state, options);
        }
    }

    public sealed class DeviceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The `sku_name` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `sku_name` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `sku_name`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Databox Edge Device.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DeviceArgs()
        {
        }
    }

    public sealed class DeviceState : Pulumi.ResourceArgs
    {
        [Input("deviceProperties")]
        private InputList<Inputs.DeviceDevicePropertyGetArgs>? _deviceProperties;

        /// <summary>
        /// A `device_properties` block as defined below.
        /// </summary>
        public InputList<Inputs.DeviceDevicePropertyGetArgs> DeviceProperties
        {
            get => _deviceProperties ?? (_deviceProperties = new InputList<Inputs.DeviceDevicePropertyGetArgs>());
            set => _deviceProperties = value;
        }

        /// <summary>
        /// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The `sku_name` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `sku_name` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `sku_name`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Databox Edge Device.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DeviceState()
        {
        }
    }
}
