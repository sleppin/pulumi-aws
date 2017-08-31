// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Policy extends lumi.NamedResource implements PolicyArgs {
    public readonly adjustmentType: string;
    public /*out*/ readonly arn: string;
    public readonly autoscalingGroupName: string;
    public readonly cooldown?: number;
    public readonly estimatedInstanceWarmup?: number;
    public readonly metricAggregationType: string;
    public readonly minAdjustmentMagnitude?: number;
    public readonly minAdjustmentStep?: number;
    public readonly name: string;
    public readonly policyType?: string;
    public readonly scalingAdjustment?: number;
    public readonly stepAdjustment?: { metricIntervalLowerBound?: string, metricIntervalUpperBound?: string, scalingAdjustment: number }[];

    public static get(id: lumi.ID): Policy {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Policy[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(urnName: string, args: PolicyArgs) {
        super(urnName);
        if (lumirt.defaultIfComputed(args.adjustmentType, "") === undefined) {
            throw new Error("Property argument 'adjustmentType' is required, but was missing");
        }
        this.adjustmentType = <any>args.adjustmentType;
        if (lumirt.defaultIfComputed(args.autoscalingGroupName, "") === undefined) {
            throw new Error("Property argument 'autoscalingGroupName' is required, but was missing");
        }
        this.autoscalingGroupName = <any>args.autoscalingGroupName;
        this.cooldown = <any>args.cooldown;
        this.estimatedInstanceWarmup = <any>args.estimatedInstanceWarmup;
        this.metricAggregationType = <any>args.metricAggregationType;
        this.minAdjustmentMagnitude = <any>args.minAdjustmentMagnitude;
        this.minAdjustmentStep = <any>args.minAdjustmentStep;
        this.name = <any>args.name;
        this.policyType = <any>args.policyType;
        this.scalingAdjustment = <any>args.scalingAdjustment;
        this.stepAdjustment = <any>args.stepAdjustment;
    }
}

export interface PolicyArgs {
    readonly adjustmentType: string;
    readonly autoscalingGroupName: string;
    readonly cooldown?: number;
    readonly estimatedInstanceWarmup?: number;
    readonly metricAggregationType?: string;
    readonly minAdjustmentMagnitude?: number;
    readonly minAdjustmentStep?: number;
    readonly name?: string;
    readonly policyType?: string;
    readonly scalingAdjustment?: number;
    readonly stepAdjustment?: { metricIntervalLowerBound?: string, metricIntervalUpperBound?: string, scalingAdjustment: number }[];
}
