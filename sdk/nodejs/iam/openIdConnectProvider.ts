// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM OpenID Connect provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultOpenIdConnectProvider = new aws.iam.OpenIdConnectProvider("default", {
 *     clientIdLists: ["266362248691-342342xasdasdasda-apps.googleusercontent.com"],
 *     thumbprintLists: [],
 *     url: "https://accounts.google.com",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_openid_connect_provider.html.markdown.
 */
export class OpenIdConnectProvider extends pulumi.CustomResource {
    /**
     * Get an existing OpenIdConnectProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OpenIdConnectProviderState, opts?: pulumi.CustomResourceOptions): OpenIdConnectProvider {
        return new OpenIdConnectProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/openIdConnectProvider:OpenIdConnectProvider';

    /**
     * Returns true if the given object is an instance of OpenIdConnectProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenIdConnectProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenIdConnectProvider.__pulumiType;
    }

    /**
     * The ARN assigned by AWS for this provider.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
     */
    public readonly clientIdLists!: pulumi.Output<string[]>;
    /**
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s). 
     */
    public readonly thumbprintLists!: pulumi.Output<string[]>;
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a OpenIdConnectProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenIdConnectProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpenIdConnectProviderArgs | OpenIdConnectProviderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OpenIdConnectProviderState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["clientIdLists"] = state ? state.clientIdLists : undefined;
            inputs["thumbprintLists"] = state ? state.thumbprintLists : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as OpenIdConnectProviderArgs | undefined;
            if (!args || args.clientIdLists === undefined) {
                throw new Error("Missing required property 'clientIdLists'");
            }
            if (!args || args.thumbprintLists === undefined) {
                throw new Error("Missing required property 'thumbprintLists'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["clientIdLists"] = args ? args.clientIdLists : undefined;
            inputs["thumbprintLists"] = args ? args.thumbprintLists : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OpenIdConnectProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OpenIdConnectProvider resources.
 */
export interface OpenIdConnectProviderState {
    /**
     * The ARN assigned by AWS for this provider.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
     */
    readonly clientIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s). 
     */
    readonly thumbprintLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OpenIdConnectProvider resource.
 */
export interface OpenIdConnectProviderArgs {
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
     */
    readonly clientIdLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s). 
     */
    readonly thumbprintLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     */
    readonly url: pulumi.Input<string>;
}
