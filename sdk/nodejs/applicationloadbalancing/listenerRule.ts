// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer Listener Rule resource.
 *
 * > **Note:** `aws.alb.ListenerRule` is known as `aws.lb.ListenerRule`. The functionality is identical.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEnd", {});
 * const frontEndListener = new aws.lb.Listener("frontEnd", {});
 * const static = new aws.lb.ListenerRule("static", {
 *     actions: [{
 *         targetGroupArn: aws_lb_target_group_static.arn,
 *         type: "forward",
 *     }],
 *     conditions: [
 *         {
 *             pathPattern: {
 *                 values: ["/static/*"],
 *             },
 *         },
 *         {
 *             hostHeader: {
 *                 values: ["example.com"],
 *             },
 *         },
 *     ],
 *     listenerArn: frontEndListener.arn,
 *     priority: 100,
 * });
 * const hostBasedRouting = new aws.lb.ListenerRule("hostBasedRouting", {
 *     actions: [{
 *         forward: {
 *             stickiness: {
 *                 duration: 600,
 *                 enabled: true,
 *             },
 *             targetGroups: [
 *                 {
 *                     arn: aws_lb_target_group_main.arn,
 *                     weight: 80,
 *                 },
 *                 {
 *                     arn: aws_lb_target_group_canary.arn,
 *                     weight: 20,
 *                 },
 *             ],
 *         },
 *         type: "forward",
 *     }],
 *     conditions: [{
 *         hostHeader: {
 *             values: ["my-service.*.mycompany.io"],
 *         },
 *     }],
 *     listenerArn: frontEndListener.arn,
 *     priority: 99,
 * });
 * const hostBasedWeightedRouting = new aws.lb.ListenerRule("hostBasedWeightedRouting", {
 *     actions: [{
 *         targetGroupArn: aws_lb_target_group_static.arn,
 *         type: "forward",
 *     }],
 *     conditions: [{
 *         hostHeader: {
 *             values: ["my-service.*.mydomain.io"],
 *         },
 *     }],
 *     listenerArn: frontEndListener.arn,
 *     priority: 99,
 * });
 * const redirectHttpToHttps = new aws.lb.ListenerRule("redirectHttpToHttps", {
 *     actions: [{
 *         redirect: {
 *             port: "443",
 *             protocol: "HTTPS",
 *             statusCode: "HTTP_301",
 *         },
 *         type: "redirect",
 *     }],
 *     conditions: [{
 *         httpHeader: {
 *             httpHeaderName: "X-Forwarded-For",
 *             values: ["192.168.1.*"],
 *         },
 *     }],
 *     listenerArn: frontEndListener.arn,
 * });
 * const healthCheck = new aws.lb.ListenerRule("healthCheck", {
 *     actions: [{
 *         fixedResponse: {
 *             contentType: "text/plain",
 *             messageBody: "HEALTHY",
 *             statusCode: "200",
 *         },
 *         type: "fixed-response",
 *     }],
 *     conditions: [{
 *         queryStrings: [
 *             {
 *                 key: "health",
 *                 value: "check",
 *             },
 *             {
 *                 value: "bar",
 *             },
 *         ],
 *     }],
 *     listenerArn: frontEndListener.arn,
 * });
 * const pool = new aws.cognito.UserPool("pool", {});
 * const client = new aws.cognito.UserPoolClient("client", {});
 * const domain = new aws.cognito.UserPoolDomain("domain", {});
 * const admin = new aws.lb.ListenerRule("admin", {
 *     actions: [
 *         {
 *             authenticateOidc: {
 *                 authorizationEndpoint: "https://example.com/authorization_endpoint",
 *                 clientId: "clientId",
 *                 clientSecret: "clientSecret",
 *                 issuer: "https://example.com",
 *                 tokenEndpoint: "https://example.com/token_endpoint",
 *                 userInfoEndpoint: "https://example.com/user_info_endpoint",
 *             },
 *             type: "authenticate-oidc",
 *         },
 *         {
 *             targetGroupArn: aws_lb_target_group_static.arn,
 *             type: "forward",
 *         },
 *     ],
 *     listenerArn: frontEndListener.arn,
 * });
 * ```
 *
 * @deprecated aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule
 */
export class ListenerRule extends pulumi.CustomResource {
    /**
     * Get an existing ListenerRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerRuleState, opts?: pulumi.CustomResourceOptions): ListenerRule {
        pulumi.log.warn("ListenerRule is deprecated: aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule")
        return new ListenerRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:applicationloadbalancing/listenerRule:ListenerRule';

    /**
     * Returns true if the given object is an instance of ListenerRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListenerRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListenerRule.__pulumiType;
    }

    /**
     * An Action block. Action blocks are documented below.
     */
    public readonly actions!: pulumi.Output<outputs.applicationloadbalancing.ListenerRuleAction[]>;
    /**
     * The Amazon Resource Name (ARN) of the target group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
     */
    public readonly conditions!: pulumi.Output<outputs.applicationloadbalancing.ListenerRuleCondition[]>;
    /**
     * The ARN of the listener to which to attach the rule.
     */
    public readonly listenerArn!: pulumi.Output<string>;
    /**
     * The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
     */
    public readonly priority!: pulumi.Output<number>;

    /**
     * Create a ListenerRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule */
    constructor(name: string, args: ListenerRuleArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule */
    constructor(name: string, argsOrState?: ListenerRuleArgs | ListenerRuleState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ListenerRule is deprecated: aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ListenerRuleState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["conditions"] = state ? state.conditions : undefined;
            inputs["listenerArn"] = state ? state.listenerArn : undefined;
            inputs["priority"] = state ? state.priority : undefined;
        } else {
            const args = argsOrState as ListenerRuleArgs | undefined;
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.conditions === undefined) {
                throw new Error("Missing required property 'conditions'");
            }
            if (!args || args.listenerArn === undefined) {
                throw new Error("Missing required property 'listenerArn'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["listenerArn"] = args ? args.listenerArn : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ListenerRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerRule resources.
 */
export interface ListenerRuleState {
    /**
     * An Action block. Action blocks are documented below.
     */
    readonly actions?: pulumi.Input<pulumi.Input<inputs.applicationloadbalancing.ListenerRuleAction>[]>;
    /**
     * The Amazon Resource Name (ARN) of the target group.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
     */
    readonly conditions?: pulumi.Input<pulumi.Input<inputs.applicationloadbalancing.ListenerRuleCondition>[]>;
    /**
     * The ARN of the listener to which to attach the rule.
     */
    readonly listenerArn?: pulumi.Input<string>;
    /**
     * The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
     */
    readonly priority?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ListenerRule resource.
 */
export interface ListenerRuleArgs {
    /**
     * An Action block. Action blocks are documented below.
     */
    readonly actions: pulumi.Input<pulumi.Input<inputs.applicationloadbalancing.ListenerRuleAction>[]>;
    /**
     * A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
     */
    readonly conditions: pulumi.Input<pulumi.Input<inputs.applicationloadbalancing.ListenerRuleCondition>[]>;
    /**
     * The ARN of the listener to which to attach the rule.
     */
    readonly listenerArn: pulumi.Input<string>;
    /**
     * The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
     */
    readonly priority?: pulumi.Input<number>;
}
