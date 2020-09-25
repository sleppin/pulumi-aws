// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer.Inputs
{

    public sealed class ServerEndpointDetailsGetArgs : Pulumi.ResourceArgs
    {
        [Input("addressAllocationIds")]
        private InputList<string>? _addressAllocationIds;

        /// <summary>
        /// A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when `endpoint_type` is set to `VPC`.
        /// </summary>
        public InputList<string> AddressAllocationIds
        {
            get => _addressAllocationIds ?? (_addressAllocationIds = new InputList<string>());
            set => _addressAllocationIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when `endpoint_type` is set to `VPC`.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The ID of the VPC endpoint. This property can only be used when `endpoint_type` is set to `VPC_ENDPOINT`
        /// </summary>
        [Input("vpcEndpointId")]
        public Input<string>? VpcEndpointId { get; set; }

        /// <summary>
        /// The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when `endpoint_type` is set to `VPC`.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ServerEndpointDetailsGetArgs()
        {
        }
    }
}
