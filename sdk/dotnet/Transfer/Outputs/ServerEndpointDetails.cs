// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer.Outputs
{

    [OutputType]
    public sealed class ServerEndpointDetails
    {
        /// <summary>
        /// A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when `endpoint_type` is set to `VPC`.
        /// </summary>
        public readonly ImmutableArray<string> AddressAllocationIds;
        /// <summary>
        /// A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when `endpoint_type` is set to `VPC`.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The ID of the VPC endpoint. This property can only be used when `endpoint_type` is set to `VPC_ENDPOINT`
        /// </summary>
        public readonly string? VpcEndpointId;
        /// <summary>
        /// The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when `endpoint_type` is set to `VPC`.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private ServerEndpointDetails(
            ImmutableArray<string> addressAllocationIds,

            ImmutableArray<string> subnetIds,

            string? vpcEndpointId,

            string? vpcId)
        {
            AddressAllocationIds = addressAllocationIds;
            SubnetIds = subnetIds;
            VpcEndpointId = vpcEndpointId;
            VpcId = vpcId;
        }
    }
}
