// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an API Gateway VPC Link.
 *
 * > **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
 * To enable private integration for HTTP APIs, use the `Amazon API Gateway Version 2 VPC Link` resource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLoadBalancer = new aws.lb.LoadBalancer("example", {
 *     internal: true,
 *     loadBalancerType: "network",
 *     subnetMappings: [{
 *         subnetId: "12345",
 *     }],
 * });
 * const exampleVpcLink = new aws.apigateway.VpcLink("example", {
 *     description: "example description",
 *     targetArn: exampleLoadBalancer.arn,
 * });
 * ```
 */
export class VpcLink extends pulumi.CustomResource {
    /**
     * Get an existing VpcLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcLinkState, opts?: pulumi.CustomResourceOptions): VpcLink {
        return new VpcLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/vpcLink:VpcLink';

    /**
     * Returns true if the given object is an instance of VpcLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcLink.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the VPC link.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name used to label and identify the VPC link.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
     */
    public readonly targetArn!: pulumi.Output<string>;

    /**
     * Create a VpcLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcLinkArgs | VpcLinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcLinkState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetArn"] = state ? state.targetArn : undefined;
        } else {
            const args = argsOrState as VpcLinkArgs | undefined;
            if (!args || args.targetArn === undefined) {
                throw new Error("Missing required property 'targetArn'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetArn"] = args ? args.targetArn : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcLink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcLink resources.
 */
export interface VpcLinkState {
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the VPC link.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name used to label and identify the VPC link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
     */
    readonly targetArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcLink resource.
 */
export interface VpcLinkArgs {
    /**
     * The description of the VPC link.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name used to label and identify the VPC link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
     */
    readonly targetArn: pulumi.Input<string>;
}
