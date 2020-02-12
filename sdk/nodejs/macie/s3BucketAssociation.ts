// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Associates an S3 resource with Amazon Macie for monitoring and data classification.
 *
 * > **NOTE:** Before using Amazon Macie for the first time it must be enabled manually. Instructions are [here](https://docs.aws.amazon.com/macie/latest/userguide/macie-setting-up.html#macie-setting-up-enable).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.macie.S3BucketAssociation("example", {
 *     bucketName: "tf-macie-example",
 *     classificationType: {
 *         oneTime: "FULL",
 *     },
 *     prefix: "data",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/macie_s3_bucket_association.html.markdown.
 */
export class S3BucketAssociation extends pulumi.CustomResource {
    /**
     * Get an existing S3BucketAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3BucketAssociationState, opts?: pulumi.CustomResourceOptions): S3BucketAssociation {
        return new S3BucketAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:macie/s3BucketAssociation:S3BucketAssociation';

    /**
     * Returns true if the given object is an instance of S3BucketAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3BucketAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3BucketAssociation.__pulumiType;
    }

    /**
     * The name of the S3 bucket that you want to associate with Amazon Macie.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * The configuration of how Amazon Macie classifies the S3 objects.
     */
    public readonly classificationType!: pulumi.Output<outputs.macie.S3BucketAssociationClassificationType>;
    /**
     * The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
     */
    public readonly memberAccountId!: pulumi.Output<string | undefined>;
    /**
     * Object key prefix identifying one or more S3 objects to which the association applies.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;

    /**
     * Create a S3BucketAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3BucketAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3BucketAssociationArgs | S3BucketAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as S3BucketAssociationState | undefined;
            inputs["bucketName"] = state ? state.bucketName : undefined;
            inputs["classificationType"] = state ? state.classificationType : undefined;
            inputs["memberAccountId"] = state ? state.memberAccountId : undefined;
            inputs["prefix"] = state ? state.prefix : undefined;
        } else {
            const args = argsOrState as S3BucketAssociationArgs | undefined;
            if (!args || args.bucketName === undefined) {
                throw new Error("Missing required property 'bucketName'");
            }
            inputs["bucketName"] = args ? args.bucketName : undefined;
            inputs["classificationType"] = args ? args.classificationType : undefined;
            inputs["memberAccountId"] = args ? args.memberAccountId : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(S3BucketAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3BucketAssociation resources.
 */
export interface S3BucketAssociationState {
    /**
     * The name of the S3 bucket that you want to associate with Amazon Macie.
     */
    readonly bucketName?: pulumi.Input<string>;
    /**
     * The configuration of how Amazon Macie classifies the S3 objects.
     */
    readonly classificationType?: pulumi.Input<inputs.macie.S3BucketAssociationClassificationType>;
    /**
     * The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
     */
    readonly memberAccountId?: pulumi.Input<string>;
    /**
     * Object key prefix identifying one or more S3 objects to which the association applies.
     */
    readonly prefix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3BucketAssociation resource.
 */
export interface S3BucketAssociationArgs {
    /**
     * The name of the S3 bucket that you want to associate with Amazon Macie.
     */
    readonly bucketName: pulumi.Input<string>;
    /**
     * The configuration of how Amazon Macie classifies the S3 objects.
     */
    readonly classificationType?: pulumi.Input<inputs.macie.S3BucketAssociationClassificationType>;
    /**
     * The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
     */
    readonly memberAccountId?: pulumi.Input<string>;
    /**
     * Object key prefix identifying one or more S3 objects to which the association applies.
     */
    readonly prefix?: pulumi.Input<string>;
}
