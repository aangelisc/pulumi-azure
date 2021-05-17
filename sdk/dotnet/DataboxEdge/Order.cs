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
    /// Manages a Databox Edge Order.
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
    ///         var exampleOrder = new Azure.DataboxEdge.Order("exampleOrder", new Azure.DataboxEdge.OrderArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             DeviceName = exampleDevice.Name,
    ///             Contact = new Azure.DataboxEdge.Inputs.OrderContactArgs
    ///             {
    ///                 CompanyName = "Contoso Corporation",
    ///                 Name = "Bart",
    ///                 EmailLists = 
    ///                 {
    ///                     "bart@example.com",
    ///                 },
    ///                 Phone = "(800) 867-5309",
    ///             },
    ///             ShipmentAddress = new Azure.DataboxEdge.Inputs.OrderShipmentAddressArgs
    ///             {
    ///                 Addresses = 
    ///                 {
    ///                     "740 Evergreen Terrace",
    ///                 },
    ///                 City = "Springfield",
    ///                 Country = "United States",
    ///                 PostalCode = "97403",
    ///                 State = "OR",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Databox Edge Orders can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:databoxedge/order:Order example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1/orders/default
    /// ```
    /// </summary>
    [AzureResourceType("azure:databoxedge/order:Order")]
    public partial class Order : Pulumi.CustomResource
    {
        /// <summary>
        /// A `contact` block as defined below.
        /// </summary>
        [Output("contact")]
        public Output<Outputs.OrderContact> Contact { get; private set; } = null!;

        /// <summary>
        /// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Output("deviceName")]
        public Output<string> DeviceName { get; private set; } = null!;

        /// <summary>
        /// The contact person name. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Tracking information for the package returned from the customer whether it has an original or a replacement device. A `return_tracking` block as defined below.
        /// </summary>
        [Output("returnTrackings")]
        public Output<ImmutableArray<Outputs.OrderReturnTracking>> ReturnTrackings { get; private set; } = null!;

        /// <summary>
        /// Serial number of the device being tracked.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// A `shipment_address block as defined below.
        /// </summary>
        [Output("shipmentAddress")]
        public Output<Outputs.OrderShipmentAddress> ShipmentAddress { get; private set; } = null!;

        /// <summary>
        /// List of status changes in the order. A `shipment_history` block as defined below.
        /// </summary>
        [Output("shipmentHistories")]
        public Output<ImmutableArray<Outputs.OrderShipmentHistory>> ShipmentHistories { get; private set; } = null!;

        /// <summary>
        /// Tracking information for the package delivered to the customer whether it has an original or a replacement device. A `shipment_tracking` block as defined below.
        /// </summary>
        [Output("shipmentTrackings")]
        public Output<ImmutableArray<Outputs.OrderShipmentTracking>> ShipmentTrackings { get; private set; } = null!;

        /// <summary>
        /// The current status of the order. A `status` block as defined below.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.OrderStatus>> Statuses { get; private set; } = null!;


        /// <summary>
        /// Create a Order resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Order(string name, OrderArgs args, CustomResourceOptions? options = null)
            : base("azure:databoxedge/order:Order", name, args ?? new OrderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Order(string name, Input<string> id, OrderState? state = null, CustomResourceOptions? options = null)
            : base("azure:databoxedge/order:Order", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Order resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Order Get(string name, Input<string> id, OrderState? state = null, CustomResourceOptions? options = null)
        {
            return new Order(name, id, state, options);
        }
    }

    public sealed class OrderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `contact` block as defined below.
        /// </summary>
        [Input("contact", required: true)]
        public Input<Inputs.OrderContactArgs> Contact { get; set; } = null!;

        /// <summary>
        /// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `shipment_address block as defined below.
        /// </summary>
        [Input("shipmentAddress", required: true)]
        public Input<Inputs.OrderShipmentAddressArgs> ShipmentAddress { get; set; } = null!;

        public OrderArgs()
        {
        }
    }

    public sealed class OrderState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `contact` block as defined below.
        /// </summary>
        [Input("contact")]
        public Input<Inputs.OrderContactGetArgs>? Contact { get; set; }

        /// <summary>
        /// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// The contact person name. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("returnTrackings")]
        private InputList<Inputs.OrderReturnTrackingGetArgs>? _returnTrackings;

        /// <summary>
        /// Tracking information for the package returned from the customer whether it has an original or a replacement device. A `return_tracking` block as defined below.
        /// </summary>
        public InputList<Inputs.OrderReturnTrackingGetArgs> ReturnTrackings
        {
            get => _returnTrackings ?? (_returnTrackings = new InputList<Inputs.OrderReturnTrackingGetArgs>());
            set => _returnTrackings = value;
        }

        /// <summary>
        /// Serial number of the device being tracked.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// A `shipment_address block as defined below.
        /// </summary>
        [Input("shipmentAddress")]
        public Input<Inputs.OrderShipmentAddressGetArgs>? ShipmentAddress { get; set; }

        [Input("shipmentHistories")]
        private InputList<Inputs.OrderShipmentHistoryGetArgs>? _shipmentHistories;

        /// <summary>
        /// List of status changes in the order. A `shipment_history` block as defined below.
        /// </summary>
        public InputList<Inputs.OrderShipmentHistoryGetArgs> ShipmentHistories
        {
            get => _shipmentHistories ?? (_shipmentHistories = new InputList<Inputs.OrderShipmentHistoryGetArgs>());
            set => _shipmentHistories = value;
        }

        [Input("shipmentTrackings")]
        private InputList<Inputs.OrderShipmentTrackingGetArgs>? _shipmentTrackings;

        /// <summary>
        /// Tracking information for the package delivered to the customer whether it has an original or a replacement device. A `shipment_tracking` block as defined below.
        /// </summary>
        public InputList<Inputs.OrderShipmentTrackingGetArgs> ShipmentTrackings
        {
            get => _shipmentTrackings ?? (_shipmentTrackings = new InputList<Inputs.OrderShipmentTrackingGetArgs>());
            set => _shipmentTrackings = value;
        }

        [Input("statuses")]
        private InputList<Inputs.OrderStatusGetArgs>? _statuses;

        /// <summary>
        /// The current status of the order. A `status` block as defined below.
        /// </summary>
        public InputList<Inputs.OrderStatusGetArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.OrderStatusGetArgs>());
            set => _statuses = value;
        }

        public OrderState()
        {
        }
    }
}
