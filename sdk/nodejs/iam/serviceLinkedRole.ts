// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const elasticbeanstalk = new aws.iam.ServiceLinkedRole("elasticbeanstalk", {
 *     awsServiceName: "elasticbeanstalk.amazonaws.com",
 * });
 * ```
 */
export class ServiceLinkedRole extends pulumi.CustomResource {
    /**
     * Get an existing ServiceLinkedRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceLinkedRoleState, opts?: pulumi.CustomResourceOptions): ServiceLinkedRole {
        return new ServiceLinkedRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/serviceLinkedRole:ServiceLinkedRole';

    /**
     * Returns true if the given object is an instance of ServiceLinkedRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceLinkedRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceLinkedRole.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) specifying the role.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
     */
    public readonly awsServiceName!: pulumi.Output<string>;
    /**
     * The creation date of the IAM role.
     */
    public /*out*/ readonly createDate!: pulumi.Output<string>;
    /**
     * Additional string appended to the role name. Not all AWS services support custom suffixes.
     */
    public readonly customSuffix!: pulumi.Output<string | undefined>;
    /**
     * The description of the role.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the role.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The path of the role.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * The stable and unique string identifying the role.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a ServiceLinkedRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceLinkedRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceLinkedRoleArgs | ServiceLinkedRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceLinkedRoleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsServiceName"] = state ? state.awsServiceName : undefined;
            inputs["createDate"] = state ? state.createDate : undefined;
            inputs["customSuffix"] = state ? state.customSuffix : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
        } else {
            const args = argsOrState as ServiceLinkedRoleArgs | undefined;
            if (!args || args.awsServiceName === undefined) {
                throw new Error("Missing required property 'awsServiceName'");
            }
            inputs["awsServiceName"] = args ? args.awsServiceName : undefined;
            inputs["customSuffix"] = args ? args.customSuffix : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createDate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["path"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceLinkedRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceLinkedRole resources.
 */
export interface ServiceLinkedRoleState {
    /**
     * The Amazon Resource Name (ARN) specifying the role.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
     */
    readonly awsServiceName?: pulumi.Input<string>;
    /**
     * The creation date of the IAM role.
     */
    readonly createDate?: pulumi.Input<string>;
    /**
     * Additional string appended to the role name. Not all AWS services support custom suffixes.
     */
    readonly customSuffix?: pulumi.Input<string>;
    /**
     * The description of the role.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The path of the role.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The stable and unique string identifying the role.
     */
    readonly uniqueId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceLinkedRole resource.
 */
export interface ServiceLinkedRoleArgs {
    /**
     * The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
     */
    readonly awsServiceName: pulumi.Input<string>;
    /**
     * Additional string appended to the role name. Not all AWS services support custom suffixes.
     */
    readonly customSuffix?: pulumi.Input<string>;
    /**
     * The description of the role.
     */
    readonly description?: pulumi.Input<string>;
}
