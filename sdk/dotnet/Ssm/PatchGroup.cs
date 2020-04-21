// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    /// <summary>
    /// Provides an SSM Patch Group resource
    /// </summary>
    public partial class PatchGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the patch baseline to register the patch group with.
        /// </summary>
        [Output("baselineId")]
        public Output<string> BaselineId { get; private set; } = null!;

        /// <summary>
        /// The name of the patch group that should be registered with the patch baseline.
        /// </summary>
        [Output("patchGroup")]
        public Output<string> PatchGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a PatchGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PatchGroup(string name, PatchGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:ssm/patchGroup:PatchGroup", name, args ?? new PatchGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PatchGroup(string name, Input<string> id, PatchGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssm/patchGroup:PatchGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PatchGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PatchGroup Get(string name, Input<string> id, PatchGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new PatchGroup(name, id, state, options);
        }
    }

    public sealed class PatchGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the patch baseline to register the patch group with.
        /// </summary>
        [Input("baselineId", required: true)]
        public Input<string> BaselineId { get; set; } = null!;

        /// <summary>
        /// The name of the patch group that should be registered with the patch baseline.
        /// </summary>
        [Input("patchGroup", required: true)]
        public Input<string> PatchGroupName { get; set; } = null!;

        public PatchGroupArgs()
        {
        }
    }

    public sealed class PatchGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the patch baseline to register the patch group with.
        /// </summary>
        [Input("baselineId")]
        public Input<string>? BaselineId { get; set; }

        /// <summary>
        /// The name of the patch group that should be registered with the patch baseline.
        /// </summary>
        [Input("patchGroup")]
        public Input<string>? PatchGroup { get; set; }

        public PatchGroupState()
        {
        }
    }
}
