// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/s3_bucket_objects.html.markdown.
 */
export function getBucketObjects(args: GetBucketObjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketObjectsResult> & GetBucketObjectsResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetBucketObjectsResult> = pulumi.runtime.invoke("aws:s3/getBucketObjects:getBucketObjects", {
        "bucket": args.bucket,
        "delimiter": args.delimiter,
        "encodingType": args.encodingType,
        "fetchOwner": args.fetchOwner,
        "maxKeys": args.maxKeys,
        "prefix": args.prefix,
        "startAfter": args.startAfter,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getBucketObjects.
 */
export interface GetBucketObjectsArgs {
    /**
     * Lists object keys in this S3 bucket
     */
    readonly bucket: string;
    /**
     * A character used to group keys (Default: none)
     */
    readonly delimiter?: string;
    /**
     * Encodes keys using this method (Default: none; besides none, only "url" can be used)
     */
    readonly encodingType?: string;
    /**
     * Boolean specifying whether to populate the owner list (Default: false)
     */
    readonly fetchOwner?: boolean;
    /**
     * Maximum object keys to return (Default: 1000)
     */
    readonly maxKeys?: number;
    /**
     * Limits results to object keys with this prefix (Default: none)
     */
    readonly prefix?: string;
    /**
     * Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
     */
    readonly startAfter?: string;
}

/**
 * A collection of values returned by getBucketObjects.
 */
export interface GetBucketObjectsResult {
    readonly bucket: string;
    /**
     * List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
     */
    readonly commonPrefixes: string[];
    readonly delimiter?: string;
    readonly encodingType?: string;
    readonly fetchOwner?: boolean;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of strings representing object keys
     */
    readonly keys: string[];
    readonly maxKeys?: number;
    /**
     * List of strings representing object owner IDs (see `fetchOwner` above)
     */
    readonly owners: string[];
    readonly prefix?: string;
    readonly startAfter?: string;
}
