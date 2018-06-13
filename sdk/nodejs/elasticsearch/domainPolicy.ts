// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Allows setting policy to an ElasticSearch domain while referencing domain attributes (e.g. ARN)
 */
export class DomainPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DomainPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainPolicyState): DomainPolicy {
        return new DomainPolicy(name, <any>state, { id });
    }

    /**
     * IAM policy document specifying the access policies for the domain
     */
    public readonly accessPolicies: pulumi.Output<string>;
    /**
     * Name of the domain.
     */
    public readonly domainName: pulumi.Output<string>;

    /**
     * Create a DomainPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainPolicyArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: DomainPolicyArgs | DomainPolicyState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DomainPolicyState = argsOrState as DomainPolicyState | undefined;
            inputs["accessPolicies"] = state ? state.accessPolicies : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
        } else {
            const args = argsOrState as DomainPolicyArgs | undefined;
            if (!args || args.accessPolicies === undefined) {
                throw new Error("Missing required property 'accessPolicies'");
            }
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["accessPolicies"] = args ? args.accessPolicies : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
        }
        super("aws:elasticsearch/domainPolicy:DomainPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainPolicy resources.
 */
export interface DomainPolicyState {
    /**
     * IAM policy document specifying the access policies for the domain
     */
    readonly accessPolicies?: pulumi.Input<string>;
    /**
     * Name of the domain.
     */
    readonly domainName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainPolicy resource.
 */
export interface DomainPolicyArgs {
    /**
     * IAM policy document specifying the access policies for the domain
     */
    readonly accessPolicies: pulumi.Input<string>;
    /**
     * Name of the domain.
     */
    readonly domainName: pulumi.Input<string>;
}