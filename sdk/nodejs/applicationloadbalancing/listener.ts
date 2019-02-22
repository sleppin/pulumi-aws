// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer Listener resource.
 * 
 * > **Note:** `aws_alb_listener` is known as `aws_lb_listener`. The functionality is identical.
 * 
 * ## Example Usage
 * 
 * ### Forward Action
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const frontEndLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("front_end", {});
 * const frontEndTargetGroup = new aws.elasticloadbalancingv2.TargetGroup("front_end", {});
 * const frontEndListener = new aws.elasticloadbalancingv2.Listener("front_end", {
 *     certificateArn: "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
 *     defaultAction: {
 *         targetGroupArn: frontEndTargetGroup.arn,
 *         type: "forward",
 *     },
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: 443,
 *     protocol: "HTTPS",
 *     sslPolicy: "ELBSecurityPolicy-2016-08",
 * });
 * ```
 * 
 * ### Redirect Action
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const frontEndLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("front_end", {});
 * const frontEndListener = new aws.elasticloadbalancingv2.Listener("front_end", {
 *     defaultAction: {
 *         redirect: {
 *             port: "443",
 *             protocol: "HTTPS",
 *             statusCode: "HTTP_301",
 *         },
 *         type: "redirect",
 *     },
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: 80,
 *     protocol: "HTTP",
 * });
 * ```
 * 
 * ### Fixed-response Action
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const frontEndLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("front_end", {});
 * const frontEndListener = new aws.elasticloadbalancingv2.Listener("front_end", {
 *     defaultAction: {
 *         fixedResponse: {
 *             contentType: "text/plain",
 *             messageBody: "Fixed response content",
 *             statusCode: "200",
 *         },
 *         type: "fixed-response",
 *     },
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: 80,
 *     protocol: "HTTP",
 * });
 * ```
 * 
 * ### Authenticate-cognito Action
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const pool = new aws.cognito.UserPool("pool", {});
 * const client = new aws.cognito.UserPoolClient("client", {});
 * const domain = new aws.cognito.UserPoolDomain("domain", {});
 * const frontEndLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("front_end", {});
 * const frontEndTargetGroup = new aws.elasticloadbalancingv2.TargetGroup("front_end", {});
 * const frontEndListener = new aws.elasticloadbalancingv2.Listener("front_end", {
 *     defaultAction: pulumi.all([pool.arn, client.id, domain.domain, frontEndTargetGroup.arn]).apply(([poolArn, id, domain, frontEndTargetGroupArn]) => (() => {
 *         throw "tf2pulumi error: aws_lb_listener.front_end.default_action: expected at most one item in list, got 2";
 *         return [
 *             {
 *                 authenticateCognito: {
 *                     userPoolArn: poolArn,
 *                     userPoolClientId: id,
 *                     userPoolDomain: domain,
 *                 },
 *                 type: "authenticate-cognito",
 *             },
 *             {
 *                 targetGroupArn: frontEndTargetGroupArn,
 *                 type: "forward",
 *             },
 *         ];
 *     })()),
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: 80,
 *     protocol: "HTTP",
 * });
 * ```
 * 
 * ### Authenticate-oidc Action
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const frontEndLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("front_end", {});
 * const frontEndTargetGroup = new aws.elasticloadbalancingv2.TargetGroup("front_end", {});
 * const frontEndListener = new aws.elasticloadbalancingv2.Listener("front_end", {
 *     defaultAction: frontEndTargetGroup.arn.apply(arn => (() => {
 *         throw "tf2pulumi error: aws_lb_listener.front_end.default_action: expected at most one item in list, got 2";
 *         return [
 *             {
 *                 authenticateOidc: {
 *                     authorizationEndpoint: "https://example.com/authorization_endpoint",
 *                     clientId: "client_id",
 *                     clientSecret: "client_secret",
 *                     issuer: "https://example.com",
 *                     tokenEndpoint: "https://example.com/token_endpoint",
 *                     userInfoEndpoint: "https://example.com/user_info_endpoint",
 *                 },
 *                 type: "authenticate-oidc",
 *             },
 *             {
 *                 targetGroupArn: arn,
 *                 type: "forward",
 *             },
 *         ];
 *     })()),
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: 80,
 *     protocol: "HTTP",
 * });
 * ```
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the listener (matches `id`)
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
     */
    public readonly certificateArn: pulumi.Output<string | undefined>;
    /**
     * An Action block. Action blocks are documented below.
     */
    public readonly defaultAction: pulumi.Output<{ authenticateCognito?: { authenticationRequestExtraParams?: {[key: string]: any}, onUnauthenticatedRequest: string, scope: string, sessionCookieName: string, sessionTimeout: number, userPoolArn: string, userPoolClientId: string, userPoolDomain: string }, authenticateOidc?: { authenticationRequestExtraParams?: {[key: string]: any}, authorizationEndpoint: string, clientId: string, clientSecret: string, issuer: string, onUnauthenticatedRequest: string, scope: string, sessionCookieName: string, sessionTimeout: number, tokenEndpoint: string, userInfoEndpoint: string }, fixedResponse?: { contentType: string, messageBody?: string, statusCode: string }, order: number, redirect?: { host?: string, path?: string, port?: string, protocol?: string, query?: string, statusCode: string }, targetGroupArn?: string, type: string }>;
    /**
     * The ARN of the load balancer.
     */
    public readonly loadBalancerArn: pulumi.Output<string>;
    /**
     * The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    public readonly port: pulumi.Output<number>;
    /**
     * The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    public readonly protocol: pulumi.Output<string | undefined>;
    /**
     * The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     */
    public readonly sslPolicy: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ListenerState = argsOrState as ListenerState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["defaultAction"] = state ? state.defaultAction : undefined;
            inputs["loadBalancerArn"] = state ? state.loadBalancerArn : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["sslPolicy"] = state ? state.sslPolicy : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if (!args || args.defaultAction === undefined) {
                throw new Error("Missing required property 'defaultAction'");
            }
            if (!args || args.loadBalancerArn === undefined) {
                throw new Error("Missing required property 'loadBalancerArn'");
            }
            if (!args || args.port === undefined) {
                throw new Error("Missing required property 'port'");
            }
            inputs["certificateArn"] = args ? args.certificateArn : undefined;
            inputs["defaultAction"] = args ? args.defaultAction : undefined;
            inputs["loadBalancerArn"] = args ? args.loadBalancerArn : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:applicationloadbalancing/listener:Listener", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * The ARN of the listener (matches `id`)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
     */
    readonly certificateArn?: pulumi.Input<string>;
    /**
     * An Action block. Action blocks are documented below.
     */
    readonly defaultAction?: pulumi.Input<{ authenticateCognito?: pulumi.Input<{ authenticationRequestExtraParams?: pulumi.Input<{[key: string]: any}>, onUnauthenticatedRequest?: pulumi.Input<string>, scope?: pulumi.Input<string>, sessionCookieName?: pulumi.Input<string>, sessionTimeout?: pulumi.Input<number>, userPoolArn: pulumi.Input<string>, userPoolClientId: pulumi.Input<string>, userPoolDomain: pulumi.Input<string> }>, authenticateOidc?: pulumi.Input<{ authenticationRequestExtraParams?: pulumi.Input<{[key: string]: any}>, authorizationEndpoint: pulumi.Input<string>, clientId: pulumi.Input<string>, clientSecret: pulumi.Input<string>, issuer: pulumi.Input<string>, onUnauthenticatedRequest?: pulumi.Input<string>, scope?: pulumi.Input<string>, sessionCookieName?: pulumi.Input<string>, sessionTimeout?: pulumi.Input<number>, tokenEndpoint: pulumi.Input<string>, userInfoEndpoint: pulumi.Input<string> }>, fixedResponse?: pulumi.Input<{ contentType: pulumi.Input<string>, messageBody?: pulumi.Input<string>, statusCode?: pulumi.Input<string> }>, order?: pulumi.Input<number>, redirect?: pulumi.Input<{ host?: pulumi.Input<string>, path?: pulumi.Input<string>, port?: pulumi.Input<string>, protocol?: pulumi.Input<string>, query?: pulumi.Input<string>, statusCode: pulumi.Input<string> }>, targetGroupArn?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The ARN of the load balancer.
     */
    readonly loadBalancerArn?: pulumi.Input<string>;
    /**
     * The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     */
    readonly sslPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
     */
    readonly certificateArn?: pulumi.Input<string>;
    /**
     * An Action block. Action blocks are documented below.
     */
    readonly defaultAction: pulumi.Input<{ authenticateCognito?: pulumi.Input<{ authenticationRequestExtraParams?: pulumi.Input<{[key: string]: any}>, onUnauthenticatedRequest?: pulumi.Input<string>, scope?: pulumi.Input<string>, sessionCookieName?: pulumi.Input<string>, sessionTimeout?: pulumi.Input<number>, userPoolArn: pulumi.Input<string>, userPoolClientId: pulumi.Input<string>, userPoolDomain: pulumi.Input<string> }>, authenticateOidc?: pulumi.Input<{ authenticationRequestExtraParams?: pulumi.Input<{[key: string]: any}>, authorizationEndpoint: pulumi.Input<string>, clientId: pulumi.Input<string>, clientSecret: pulumi.Input<string>, issuer: pulumi.Input<string>, onUnauthenticatedRequest?: pulumi.Input<string>, scope?: pulumi.Input<string>, sessionCookieName?: pulumi.Input<string>, sessionTimeout?: pulumi.Input<number>, tokenEndpoint: pulumi.Input<string>, userInfoEndpoint: pulumi.Input<string> }>, fixedResponse?: pulumi.Input<{ contentType: pulumi.Input<string>, messageBody?: pulumi.Input<string>, statusCode?: pulumi.Input<string> }>, order?: pulumi.Input<number>, redirect?: pulumi.Input<{ host?: pulumi.Input<string>, path?: pulumi.Input<string>, port?: pulumi.Input<string>, protocol?: pulumi.Input<string>, query?: pulumi.Input<string>, statusCode: pulumi.Input<string> }>, targetGroupArn?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * The ARN of the load balancer.
     */
    readonly loadBalancerArn: pulumi.Input<string>;
    /**
     * The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     */
    readonly sslPolicy?: pulumi.Input<string>;
}
