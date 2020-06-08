// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CodeBuild Source Credentials Resource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.codebuild.SourceCredential("example", {
 *     authType: "PERSONAL_ACCESS_TOKEN",
 *     serverType: "GITHUB",
 *     token: "example",
 * });
 * ```
 *
 * ### Bitbucket Server Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.codebuild.SourceCredential("example", {
 *     authType: "BASIC_AUTH",
 *     serverType: "BITBUCKET",
 *     token: "example",
 *     userName: "test-user",
 * });
 * ```
 */
export class SourceCredential extends pulumi.CustomResource {
    /**
     * Get an existing SourceCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceCredentialState, opts?: pulumi.CustomResourceOptions): SourceCredential {
        return new SourceCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codebuild/sourceCredential:SourceCredential';

    /**
     * Returns true if the given object is an instance of SourceCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceCredential.__pulumiType;
    }

    /**
     * The ARN of Source Credential.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * The source provider used for this project.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a SourceCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceCredentialArgs | SourceCredentialState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SourceCredentialState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["authType"] = state ? state.authType : undefined;
            inputs["serverType"] = state ? state.serverType : undefined;
            inputs["token"] = state ? state.token : undefined;
            inputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as SourceCredentialArgs | undefined;
            if (!args || args.authType === undefined) {
                throw new Error("Missing required property 'authType'");
            }
            if (!args || args.serverType === undefined) {
                throw new Error("Missing required property 'serverType'");
            }
            if (!args || args.token === undefined) {
                throw new Error("Missing required property 'token'");
            }
            inputs["authType"] = args ? args.authType : undefined;
            inputs["serverType"] = args ? args.serverType : undefined;
            inputs["token"] = args ? args.token : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SourceCredential.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceCredential resources.
 */
export interface SourceCredentialState {
    /**
     * The ARN of Source Credential.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
     */
    readonly authType?: pulumi.Input<string>;
    /**
     * The source provider used for this project.
     */
    readonly serverType?: pulumi.Input<string>;
    /**
     * For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
     */
    readonly userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SourceCredential resource.
 */
export interface SourceCredentialArgs {
    /**
     * The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
     */
    readonly authType: pulumi.Input<string>;
    /**
     * The source provider used for this project.
     */
    readonly serverType: pulumi.Input<string>;
    /**
     * For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
     */
    readonly token: pulumi.Input<string>;
    /**
     * The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
     */
    readonly userName?: pulumi.Input<string>;
}
