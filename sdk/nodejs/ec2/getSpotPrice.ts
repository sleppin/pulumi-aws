// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Information about most recent Spot Price for a given EC2 instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ec2.getSpotPrice({
 *     availabilityZone: "us-west-2a",
 *     filters: [{
 *         name: "product-description",
 *         values: ["Linux/UNIX"],
 *     }],
 *     instanceType: "t3.medium",
 * }, { async: true }));
 * ```
 */
export function getSpotPrice(args?: GetSpotPriceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpotPriceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getSpotPrice:getSpotPrice", {
        "availabilityZone": args.availabilityZone,
        "filters": args.filters,
        "instanceType": args.instanceType,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpotPrice.
 */
export interface GetSpotPriceArgs {
    /**
     * The availability zone in which to query Spot price information.
     */
    readonly availabilityZone?: string;
    /**
     * One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSpotPriceHistory.html) for supported filters. Detailed below.
     */
    readonly filters?: inputs.ec2.GetSpotPriceFilter[];
    /**
     * The type of instance for which to query Spot Price information.
     */
    readonly instanceType?: string;
}

/**
 * A collection of values returned by getSpotPrice.
 */
export interface GetSpotPriceResult {
    readonly availabilityZone?: string;
    readonly filters?: outputs.ec2.GetSpotPriceFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceType?: string;
    /**
     * The most recent Spot Price value for the given instance type and AZ.
     */
    readonly spotPrice: string;
    /**
     * The timestamp at which the Spot Price value was published.
     */
    readonly spotPriceTimestamp: string;
}
