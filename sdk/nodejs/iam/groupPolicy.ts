// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {PolicyDocument} from "./documents";

/**
 * Provides an IAM policy attached to a group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const myDevelopers = new aws.iam.Group("my_developers", {
 *     path: "/users/",
 * });
 * const myDeveloperPolicy = new aws.iam.GroupPolicy("my_developer_policy", {
 *     group: myDevelopers.id,
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "ec2:Describe*"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "*"
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 */
export class GroupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing GroupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupPolicyState, opts?: pulumi.CustomResourceOptions): GroupPolicy {
        return new GroupPolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * The IAM group to attach to the policy.
     */
    public readonly group: pulumi.Output<string>;
    /**
     * The name of the policy. If omitted, Terraform will
     * assign a random, unique name.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    public readonly namePrefix: pulumi.Output<string | undefined>;
    /**
     * The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
     */
    public readonly policy: pulumi.Output<string>;

    /**
     * Create a GroupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupPolicyArgs | GroupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: GroupPolicyState = argsOrState as GroupPolicyState | undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as GroupPolicyArgs | undefined;
            if (!args || args.group === undefined) {
                throw new Error("Missing required property 'group'");
            }
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            inputs["group"] = args ? args.group : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["policy"] = args ? args.policy : undefined;
        }
        super("aws:iam/groupPolicy:GroupPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupPolicy resources.
 */
export interface GroupPolicyState {
    /**
     * The IAM group to attach to the policy.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * The name of the policy. If omitted, Terraform will
     * assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
     */
    readonly policy?: pulumi.Input<string | PolicyDocument>;
}

/**
 * The set of arguments for constructing a GroupPolicy resource.
 */
export interface GroupPolicyArgs {
    /**
     * The IAM group to attach to the policy.
     */
    readonly group: pulumi.Input<string>;
    /**
     * The name of the policy. If omitted, Terraform will
     * assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
     */
    readonly policy: pulumi.Input<string | PolicyDocument>;
}
