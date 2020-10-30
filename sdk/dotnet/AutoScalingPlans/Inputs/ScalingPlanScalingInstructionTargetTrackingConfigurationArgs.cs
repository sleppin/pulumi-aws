// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScalingPlans.Inputs
{

    public sealed class ScalingPlanScalingInstructionTargetTrackingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A customized metric. You can specify either `customized_scaling_metric_specification` or `predefined_scaling_metric_specification`.
        /// More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_CustomizedScalingMetricSpecification.html).
        /// </summary>
        [Input("customizedScalingMetricSpecification")]
        public Input<Inputs.ScalingPlanScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationArgs>? CustomizedScalingMetricSpecification { get; set; }

        /// <summary>
        /// Boolean indicating whether scale in by the target tracking scaling policy is disabled. Defaults to `false`.
        /// </summary>
        [Input("disableScaleIn")]
        public Input<bool>? DisableScaleIn { get; set; }

        /// <summary>
        /// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
        /// This value is used only if the resource is an Auto Scaling group.
        /// </summary>
        [Input("estimatedInstanceWarmup")]
        public Input<int>? EstimatedInstanceWarmup { get; set; }

        /// <summary>
        /// A predefined metric. You can specify either `predefined_scaling_metric_specification` or `customized_scaling_metric_specification`.
        /// More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_PredefinedScalingMetricSpecification.html).
        /// </summary>
        [Input("predefinedScalingMetricSpecification")]
        public Input<Inputs.ScalingPlanScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecificationArgs>? PredefinedScalingMetricSpecification { get; set; }

        /// <summary>
        /// The amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
        /// This value is not used if the scalable resource is an Auto Scaling group.
        /// </summary>
        [Input("scaleInCooldown")]
        public Input<int>? ScaleInCooldown { get; set; }

        /// <summary>
        /// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
        /// This value is not used if the scalable resource is an Auto Scaling group.
        /// </summary>
        [Input("scaleOutCooldown")]
        public Input<int>? ScaleOutCooldown { get; set; }

        /// <summary>
        /// The target value for the metric.
        /// </summary>
        [Input("targetValue", required: true)]
        public Input<double> TargetValue { get; set; } = null!;

        public ScalingPlanScalingInstructionTargetTrackingConfigurationArgs()
        {
        }
    }
}