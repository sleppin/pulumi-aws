// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can be used to fetch information about a specific
 * IAM instance profile. By using this data source, you can reference IAM
 * instance profile properties without having to hard code ARNs as input.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.iam.getInstanceProfile({
 *     name: "anExampleInstanceProfileName",
 * }, { async: true }));
 * ```
 */
export function getInstanceProfile(args: GetInstanceProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getInstanceProfile:getInstanceProfile", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceProfile.
 */
export interface GetInstanceProfileArgs {
    /**
     * The friendly IAM instance profile name to match.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getInstanceProfile.
 */
export interface GetInstanceProfileResult {
    /**
     * The Amazon Resource Name (ARN) specifying the instance profile.
     */
    readonly arn: string;
    /**
     * The string representation of the date the instance profile
     * was created.
     */
    readonly createDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The path to the instance profile.
     */
    readonly path: string;
    /**
     * The role arn associated with this instance profile.
     */
    readonly roleArn: string;
    /**
     * The role id associated with this instance profile.
     */
    readonly roleId: string;
    /**
     * The role name associated with this instance profile.
     */
    readonly roleName: string;
}
