// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an individual Service Quota.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.servicequotas.ServiceQuota("example", {
 *     quotaCode: "L-F678F1CE",
 *     serviceCode: "vpc",
 *     value: 75,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/servicequotas_service_quota.html.markdown.
 */
export class ServiceQuota extends pulumi.CustomResource {
    /**
     * Get an existing ServiceQuota resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceQuotaState, opts?: pulumi.CustomResourceOptions): ServiceQuota {
        return new ServiceQuota(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicequotas/serviceQuota:ServiceQuota';

    /**
     * Returns true if the given object is an instance of ServiceQuota.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceQuota {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceQuota.__pulumiType;
    }

    /**
     * Whether the service quota can be increased.
     */
    public /*out*/ readonly adjustable!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of the service quota.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Default value of the service quota.
     */
    public /*out*/ readonly defaultValue!: pulumi.Output<number>;
    /**
     * Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     */
    public readonly quotaCode!: pulumi.Output<string>;
    /**
     * Name of the quota.
     */
    public /*out*/ readonly quotaName!: pulumi.Output<string>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly requestStatus!: pulumi.Output<string>;
    /**
     * Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     */
    public readonly serviceCode!: pulumi.Output<string>;
    /**
     * Name of the service.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
     */
    public readonly value!: pulumi.Output<number>;

    /**
     * Create a ServiceQuota resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceQuotaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceQuotaArgs | ServiceQuotaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceQuotaState | undefined;
            inputs["adjustable"] = state ? state.adjustable : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["defaultValue"] = state ? state.defaultValue : undefined;
            inputs["quotaCode"] = state ? state.quotaCode : undefined;
            inputs["quotaName"] = state ? state.quotaName : undefined;
            inputs["requestId"] = state ? state.requestId : undefined;
            inputs["requestStatus"] = state ? state.requestStatus : undefined;
            inputs["serviceCode"] = state ? state.serviceCode : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ServiceQuotaArgs | undefined;
            if (!args || args.quotaCode === undefined) {
                throw new Error("Missing required property 'quotaCode'");
            }
            if (!args || args.serviceCode === undefined) {
                throw new Error("Missing required property 'serviceCode'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["quotaCode"] = args ? args.quotaCode : undefined;
            inputs["serviceCode"] = args ? args.serviceCode : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["adjustable"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["defaultValue"] = undefined /*out*/;
            inputs["quotaName"] = undefined /*out*/;
            inputs["requestId"] = undefined /*out*/;
            inputs["requestStatus"] = undefined /*out*/;
            inputs["serviceName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceQuota.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceQuota resources.
 */
export interface ServiceQuotaState {
    /**
     * Whether the service quota can be increased.
     */
    readonly adjustable?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of the service quota.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Default value of the service quota.
     */
    readonly defaultValue?: pulumi.Input<number>;
    /**
     * Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     */
    readonly quotaCode?: pulumi.Input<string>;
    /**
     * Name of the quota.
     */
    readonly quotaName?: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    readonly requestStatus?: pulumi.Input<string>;
    /**
     * Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     */
    readonly serviceCode?: pulumi.Input<string>;
    /**
     * Name of the service.
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
     */
    readonly value?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ServiceQuota resource.
 */
export interface ServiceQuotaArgs {
    /**
     * Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     */
    readonly quotaCode: pulumi.Input<string>;
    /**
     * Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     */
    readonly serviceCode: pulumi.Input<string>;
    /**
     * Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
     */
    readonly value: pulumi.Input<number>;
}
