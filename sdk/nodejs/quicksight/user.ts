// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing QuickSight User
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.quicksight.User("example", {
 *     email: "author@example.com",
 *     identityType: "IAM",
 *     userName: "an-author",
 *     userRole: "AUTHOR",
 * });
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:quicksight/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the user
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * The email address of the user that you want to register.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The ARN of the IAM user or role that you are registering with Amazon QuickSight.
     */
    public readonly iamArn!: pulumi.Output<string | undefined>;
    /**
     * Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
     */
    public readonly identityType!: pulumi.Output<string>;
    /**
     * The namespace. Currently, you should set this to `default`.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
     */
    public readonly sessionName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon QuickSight user name that you want to create for the user you are registering.
     */
    public readonly userName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
     */
    public readonly userRole!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsAccountId"] = state ? state.awsAccountId : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["iamArn"] = state ? state.iamArn : undefined;
            inputs["identityType"] = state ? state.identityType : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
            inputs["sessionName"] = state ? state.sessionName : undefined;
            inputs["userName"] = state ? state.userName : undefined;
            inputs["userRole"] = state ? state.userRole : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.email === undefined) {
                throw new Error("Missing required property 'email'");
            }
            if (!args || args.identityType === undefined) {
                throw new Error("Missing required property 'identityType'");
            }
            if (!args || args.userRole === undefined) {
                throw new Error("Missing required property 'userRole'");
            }
            inputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["iamArn"] = args ? args.iamArn : undefined;
            inputs["identityType"] = args ? args.identityType : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["sessionName"] = args ? args.sessionName : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["userRole"] = args ? args.userRole : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Amazon Resource Name (ARN) of the user
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    readonly awsAccountId?: pulumi.Input<string>;
    /**
     * The email address of the user that you want to register.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The ARN of the IAM user or role that you are registering with Amazon QuickSight.
     */
    readonly iamArn?: pulumi.Input<string>;
    /**
     * Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
     */
    readonly identityType?: pulumi.Input<string>;
    /**
     * The namespace. Currently, you should set this to `default`.
     */
    readonly namespace?: pulumi.Input<string>;
    /**
     * The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
     */
    readonly sessionName?: pulumi.Input<string>;
    /**
     * The Amazon QuickSight user name that you want to create for the user you are registering.
     */
    readonly userName?: pulumi.Input<string>;
    /**
     * The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
     */
    readonly userRole?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    readonly awsAccountId?: pulumi.Input<string>;
    /**
     * The email address of the user that you want to register.
     */
    readonly email: pulumi.Input<string>;
    /**
     * The ARN of the IAM user or role that you are registering with Amazon QuickSight.
     */
    readonly iamArn?: pulumi.Input<string>;
    /**
     * Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
     */
    readonly identityType: pulumi.Input<string>;
    /**
     * The namespace. Currently, you should set this to `default`.
     */
    readonly namespace?: pulumi.Input<string>;
    /**
     * The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
     */
    readonly sessionName?: pulumi.Input<string>;
    /**
     * The Amazon QuickSight user name that you want to create for the user you are registering.
     */
    readonly userName?: pulumi.Input<string>;
    /**
     * The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
     */
    readonly userRole: pulumi.Input<string>;
}
