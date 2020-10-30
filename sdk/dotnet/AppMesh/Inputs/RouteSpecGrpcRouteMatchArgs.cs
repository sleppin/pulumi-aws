// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class RouteSpecGrpcRouteMatchArgs : Pulumi.ResourceArgs
    {
        [Input("metadatas")]
        private InputList<Inputs.RouteSpecGrpcRouteMatchMetadataArgs>? _metadatas;

        /// <summary>
        /// The data to match from the gRPC request.
        /// </summary>
        public InputList<Inputs.RouteSpecGrpcRouteMatchMetadataArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.RouteSpecGrpcRouteMatchMetadataArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// The method name to match from the request. If you specify a name, you must also specify a `service_name`.
        /// </summary>
        [Input("methodName")]
        public Input<string>? MethodName { get; set; }

        /// <summary>
        /// The value sent by the client must begin with the specified characters.
        /// This parameter must always start with /, which by itself matches all requests to the virtual router service name.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The fully qualified domain name for the service to match from the request.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public RouteSpecGrpcRouteMatchArgs()
        {
        }
    }
}