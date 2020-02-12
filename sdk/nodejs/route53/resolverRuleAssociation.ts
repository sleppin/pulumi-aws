// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Route53 Resolver rule association.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.route53.ResolverRuleAssociation("example", {
 *     resolverRuleId: aws_route53_resolver_rule_sys.id,
 *     vpcId: aws_vpc_foo.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_rule_association.html.markdown.
 */
export class ResolverRuleAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ResolverRuleAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResolverRuleAssociationState, opts?: pulumi.CustomResourceOptions): ResolverRuleAssociation {
        return new ResolverRuleAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/resolverRuleAssociation:ResolverRuleAssociation';

    /**
     * Returns true if the given object is an instance of ResolverRuleAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverRuleAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverRuleAssociation.__pulumiType;
    }

    /**
     * A name for the association that you're creating between a resolver rule and a VPC.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the resolver rule that you want to associate with the VPC.
     */
    public readonly resolverRuleId!: pulumi.Output<string>;
    /**
     * The ID of the VPC that you want to associate the resolver rule with.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a ResolverRuleAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResolverRuleAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResolverRuleAssociationArgs | ResolverRuleAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResolverRuleAssociationState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resolverRuleId"] = state ? state.resolverRuleId : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ResolverRuleAssociationArgs | undefined;
            if (!args || args.resolverRuleId === undefined) {
                throw new Error("Missing required property 'resolverRuleId'");
            }
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["resolverRuleId"] = args ? args.resolverRuleId : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResolverRuleAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResolverRuleAssociation resources.
 */
export interface ResolverRuleAssociationState {
    /**
     * A name for the association that you're creating between a resolver rule and a VPC.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the resolver rule that you want to associate with the VPC.
     */
    readonly resolverRuleId?: pulumi.Input<string>;
    /**
     * The ID of the VPC that you want to associate the resolver rule with.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResolverRuleAssociation resource.
 */
export interface ResolverRuleAssociationArgs {
    /**
     * A name for the association that you're creating between a resolver rule and a VPC.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the resolver rule that you want to associate with the VPC.
     */
    readonly resolverRuleId: pulumi.Input<string>;
    /**
     * The ID of the VPC that you want to associate the resolver rule with.
     */
    readonly vpcId: pulumi.Input<string>;
}
