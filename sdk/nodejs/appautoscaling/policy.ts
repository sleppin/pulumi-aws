// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Application AutoScaling Policy resource.
 *
 * ## Example Usage
 *
 * ### DynamoDB Table Autoscaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodbTableReadTarget", {
 *     maxCapacity: 100,
 *     minCapacity: 5,
 *     resourceId: "table/tableName",
 *     roleArn: aws_iam_role_DynamoDBAutoscaleRole.arn,
 *     scalableDimension: "dynamodb:table:ReadCapacityUnits",
 *     serviceNamespace: "dynamodb",
 * });
 * const dynamodbTableReadPolicy = new aws.appautoscaling.Policy("dynamodbTableReadPolicy", {
 *     policyType: "TargetTrackingScaling",
 *     resourceId: dynamodbTableReadTarget.resourceId,
 *     scalableDimension: dynamodbTableReadTarget.scalableDimension,
 *     serviceNamespace: dynamodbTableReadTarget.serviceNamespace,
 *     targetTrackingScalingPolicyConfiguration: {
 *         predefinedMetricSpecification: {
 *             predefinedMetricType: "DynamoDBReadCapacityUtilization",
 *         },
 *         targetValue: 70,
 *     },
 * });
 * ```
 *
 * ### ECS Service Autoscaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ecsTarget = new aws.appautoscaling.Target("ecsTarget", {
 *     maxCapacity: 4,
 *     minCapacity: 1,
 *     resourceId: "service/clusterName/serviceName",
 *     roleArn: var_ecs_iam_role,
 *     scalableDimension: "ecs:service:DesiredCount",
 *     serviceNamespace: "ecs",
 * });
 * const ecsPolicy = new aws.appautoscaling.Policy("ecsPolicy", {
 *     policyType: "StepScaling",
 *     resourceId: ecsTarget.resourceId,
 *     scalableDimension: ecsTarget.scalableDimension,
 *     serviceNamespace: ecsTarget.serviceNamespace,
 *     stepScalingPolicyConfiguration: {
 *         adjustmentType: "ChangeInCapacity",
 *         cooldown: 60,
 *         metricAggregationType: "Maximum",
 *         stepAdjustments: [{
 *             metricIntervalUpperBound: "0",
 *             scalingAdjustment: -1,
 *         }],
 *     },
 * });
 * ```
 *
 * ### Preserve desired count when updating an autoscaled ECS Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ecsService = new aws.ecs.Service("ecsService", {
 *     cluster: "clusterName",
 *     desiredCount: 2,
 *     taskDefinition: "taskDefinitionFamily:1",
 * }, {ignoreChanges: ["desiredCount"]});
 * ```
 *
 * ### Aurora Read Replica Autoscaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const replicasTarget = new aws.appautoscaling.Target("replicas", {
 *     maxCapacity: 15,
 *     minCapacity: 1,
 *     resourceId: pulumi.interpolate`cluster:${aws_rds_cluster_example.id}`,
 *     scalableDimension: "rds:cluster:ReadReplicaCount",
 *     serviceNamespace: "rds",
 * });
 * const replicasPolicy = new aws.appautoscaling.Policy("replicas", {
 *     policyType: "TargetTrackingScaling",
 *     resourceId: replicasTarget.resourceId,
 *     scalableDimension: replicasTarget.scalableDimension,
 *     serviceNamespace: replicasTarget.serviceNamespace,
 *     targetTrackingScalingPolicyConfiguration: {
 *         predefinedMetricSpecification: {
 *             predefinedMetricType: "RDSReaderAverageCPUUtilization",
 *         },
 *         scaleInCooldown: 300,
 *         scaleOutCooldown: 300,
 *         targetValue: 75,
 *     },
 * });
 * ```
 *
 * ## Nested fields
 *
 * ### `targetTrackingScalingPolicyConfiguration`
 *
 * * `targetValue` - (Required) The target value for the metric.
 * * `disableScaleIn` - (Optional) Indicates whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is `false`.
 * * `scaleInCooldown` - (Optional) The amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
 * * `scaleOutCooldown` - (Optional) The amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
 * * `customizedMetricSpecification` - (Optional) A custom CloudWatch metric. Documentation can be found  at: [AWS Customized Metric Specification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CustomizedMetricSpecification.html). See supported fields below.
 * * `predefinedMetricSpecification` - (Optional) A predefined metric. See supported fields below.
 *
 * ### `customizedMetricSpecification`
 *
 * * `dimensions` - (Optional) The dimensions of the metric.
 * * `metricName` - (Required) The name of the metric.
 * * `namespace` - (Required) The namespace of the metric.
 * * `statistic` - (Required) The statistic of the metric.
 * * `unit` - (Optional) The unit of the metric.
 *
 * ### `predefinedMetricSpecification`
 *
 * * `predefinedMetricType` - (Required) The metric type.
 * * `resourceLabel` - (Optional) Reserved for future use.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appautoscaling_policy.html.markdown.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appautoscaling/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly scalableDimension!: pulumi.Output<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly serviceNamespace!: pulumi.Output<string>;
    /**
     * Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
     */
    public readonly stepScalingPolicyConfiguration!: pulumi.Output<outputs.appautoscaling.PolicyStepScalingPolicyConfiguration | undefined>;
    /**
     * A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
     */
    public readonly targetTrackingScalingPolicyConfiguration!: pulumi.Output<outputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["scalableDimension"] = state ? state.scalableDimension : undefined;
            inputs["serviceNamespace"] = state ? state.serviceNamespace : undefined;
            inputs["stepScalingPolicyConfiguration"] = state ? state.stepScalingPolicyConfiguration : undefined;
            inputs["targetTrackingScalingPolicyConfiguration"] = state ? state.targetTrackingScalingPolicyConfiguration : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.scalableDimension === undefined) {
                throw new Error("Missing required property 'scalableDimension'");
            }
            if (!args || args.serviceNamespace === undefined) {
                throw new Error("Missing required property 'serviceNamespace'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["scalableDimension"] = args ? args.scalableDimension : undefined;
            inputs["serviceNamespace"] = args ? args.serviceNamespace : undefined;
            inputs["stepScalingPolicyConfiguration"] = args ? args.stepScalingPolicyConfiguration : undefined;
            inputs["targetTrackingScalingPolicyConfiguration"] = args ? args.targetTrackingScalingPolicyConfiguration : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension?: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace?: pulumi.Input<string>;
    /**
     * Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
     */
    readonly stepScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyStepScalingPolicyConfiguration>;
    /**
     * A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
     */
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace: pulumi.Input<string>;
    /**
     * Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
     */
    readonly stepScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyStepScalingPolicyConfiguration>;
    /**
     * A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
     */
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration>;
}
