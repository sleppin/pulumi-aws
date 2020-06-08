// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a single EBS volume.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ebs.Volume("example", {
 *     availabilityZone: "us-west-2a",
 *     size: 40,
 *     tags: {
 *         Name: "HelloWorld",
 *     },
 * });
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ebs/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AZ where the EBS volume will exist.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * If true, the disk will be encrypted.
     */
    public readonly encrypted!: pulumi.Output<boolean>;
    /**
     * The amount of IOPS to provision for the disk.
     */
    public readonly iops!: pulumi.Output<number>;
    /**
     * The ARN for the KMS encryption key. When specifying `kmsKeyId`, `encrypted` needs to be set to true.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported exclusively on `io1` volumes.
     */
    public readonly multiAttachEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    /**
     * The size of the drive in GiBs.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * A snapshot to base the EBS volume off of.
     */
    public readonly snapshotId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The type of EBS volume. Can be "standard", "gp2", "io1", "sc1" or "st1" (Default: "gp2").
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["iops"] = state ? state.iops : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["multiAttachEnabled"] = state ? state.multiAttachEnabled : undefined;
            inputs["outpostArn"] = state ? state.outpostArn : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if (!args || args.availabilityZone === undefined) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["iops"] = args ? args.iops : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["multiAttachEnabled"] = args ? args.multiAttachEnabled : undefined;
            inputs["outpostArn"] = args ? args.outpostArn : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The AZ where the EBS volume will exist.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The amount of IOPS to provision for the disk.
     */
    readonly iops?: pulumi.Input<number>;
    /**
     * The ARN for the KMS encryption key. When specifying `kmsKeyId`, `encrypted` needs to be set to true.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported exclusively on `io1` volumes.
     */
    readonly multiAttachEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    readonly outpostArn?: pulumi.Input<string>;
    /**
     * The size of the drive in GiBs.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * A snapshot to base the EBS volume off of.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of EBS volume. Can be "standard", "gp2", "io1", "sc1" or "st1" (Default: "gp2").
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The AZ where the EBS volume will exist.
     */
    readonly availabilityZone: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The amount of IOPS to provision for the disk.
     */
    readonly iops?: pulumi.Input<number>;
    /**
     * The ARN for the KMS encryption key. When specifying `kmsKeyId`, `encrypted` needs to be set to true.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported exclusively on `io1` volumes.
     */
    readonly multiAttachEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    readonly outpostArn?: pulumi.Input<string>;
    /**
     * The size of the drive in GiBs.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * A snapshot to base the EBS volume off of.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of EBS volume. Can be "standard", "gp2", "io1", "sc1" or "st1" (Default: "gp2").
     */
    readonly type?: pulumi.Input<string>;
}
