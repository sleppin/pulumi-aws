// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage accepting a Resource Access Manager (RAM) Resource Share invitation. From a _receiver_ AWS account, accept an invitation to share resources that were shared by a _sender_ AWS account. To create a resource share in the _sender_, see the `aws.ram.ResourceShare` resource.
 *
 * > **Note:** If both AWS accounts are in the same Organization and [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), this resource is not necessary as RAM Resource Share invitations are not used.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const alternate = new aws.Provider("alternate", {
 *     profile: "profile1",
 * });
 * const senderShare = new aws.ram.ResourceShare("senderShare", {
 *     allowExternalPrincipals: true,
 *     tags: {
 *         Name: "tf-test-resource-share",
 *     },
 * }, { provider: alternate });
 * const receiver = pulumi.output(aws.getCallerIdentity({ async: true }));
 * const senderInvite = new aws.ram.PrincipalAssociation("senderInvite", {
 *     principal: receiver.accountId,
 *     resourceShareArn: senderShare.arn,
 * }, { provider: alternate });
 * const receiverAccept = new aws.ram.ResourceShareAccepter("receiverAccept", {
 *     shareArn: senderInvite.resourceShareArn,
 * });
 * ```
 */
export class ResourceShareAccepter extends pulumi.CustomResource {
    /**
     * Get an existing ResourceShareAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceShareAccepterState, opts?: pulumi.CustomResourceOptions): ResourceShareAccepter {
        return new ResourceShareAccepter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ram/resourceShareAccepter:ResourceShareAccepter';

    /**
     * Returns true if the given object is an instance of ResourceShareAccepter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceShareAccepter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceShareAccepter.__pulumiType;
    }

    /**
     * The ARN of the resource share invitation.
     */
    public /*out*/ readonly invitationArn!: pulumi.Output<string>;
    /**
     * The account ID of the receiver account which accepts the invitation.
     */
    public /*out*/ readonly receiverAccountId!: pulumi.Output<string>;
    /**
     * A list of the resource ARNs shared via the resource share.
     */
    public /*out*/ readonly resources!: pulumi.Output<string[]>;
    /**
     * The account ID of the sender account which submits the invitation.
     */
    public /*out*/ readonly senderAccountId!: pulumi.Output<string>;
    /**
     * The ARN of the resource share.
     */
    public readonly shareArn!: pulumi.Output<string>;
    /**
     * The ID of the resource share as displayed in the console.
     */
    public /*out*/ readonly shareId!: pulumi.Output<string>;
    /**
     * The name of the resource share.
     */
    public /*out*/ readonly shareName!: pulumi.Output<string>;
    /**
     * The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ResourceShareAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceShareAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceShareAccepterArgs | ResourceShareAccepterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceShareAccepterState | undefined;
            inputs["invitationArn"] = state ? state.invitationArn : undefined;
            inputs["receiverAccountId"] = state ? state.receiverAccountId : undefined;
            inputs["resources"] = state ? state.resources : undefined;
            inputs["senderAccountId"] = state ? state.senderAccountId : undefined;
            inputs["shareArn"] = state ? state.shareArn : undefined;
            inputs["shareId"] = state ? state.shareId : undefined;
            inputs["shareName"] = state ? state.shareName : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ResourceShareAccepterArgs | undefined;
            if (!args || args.shareArn === undefined) {
                throw new Error("Missing required property 'shareArn'");
            }
            inputs["shareArn"] = args ? args.shareArn : undefined;
            inputs["invitationArn"] = undefined /*out*/;
            inputs["receiverAccountId"] = undefined /*out*/;
            inputs["resources"] = undefined /*out*/;
            inputs["senderAccountId"] = undefined /*out*/;
            inputs["shareId"] = undefined /*out*/;
            inputs["shareName"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourceShareAccepter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceShareAccepter resources.
 */
export interface ResourceShareAccepterState {
    /**
     * The ARN of the resource share invitation.
     */
    readonly invitationArn?: pulumi.Input<string>;
    /**
     * The account ID of the receiver account which accepts the invitation.
     */
    readonly receiverAccountId?: pulumi.Input<string>;
    /**
     * A list of the resource ARNs shared via the resource share.
     */
    readonly resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The account ID of the sender account which submits the invitation.
     */
    readonly senderAccountId?: pulumi.Input<string>;
    /**
     * The ARN of the resource share.
     */
    readonly shareArn?: pulumi.Input<string>;
    /**
     * The ID of the resource share as displayed in the console.
     */
    readonly shareId?: pulumi.Input<string>;
    /**
     * The name of the resource share.
     */
    readonly shareName?: pulumi.Input<string>;
    /**
     * The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceShareAccepter resource.
 */
export interface ResourceShareAccepterArgs {
    /**
     * The ARN of the resource share.
     */
    readonly shareArn: pulumi.Input<string>;
}
