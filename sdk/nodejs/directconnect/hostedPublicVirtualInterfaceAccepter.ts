// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage the accepter's side of a Direct Connect hosted public virtual interface.
 * This resource accepts ownership of a public virtual interface created by another AWS account.
 */
export class HostedPublicVirtualInterfaceAccepter extends pulumi.CustomResource {
    /**
     * Get an existing HostedPublicVirtualInterfaceAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedPublicVirtualInterfaceAccepterState, opts?: pulumi.CustomResourceOptions): HostedPublicVirtualInterfaceAccepter {
        return new HostedPublicVirtualInterfaceAccepter(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    public readonly virtualInterfaceId: pulumi.Output<string>;

    /**
     * Create a HostedPublicVirtualInterfaceAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedPublicVirtualInterfaceAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedPublicVirtualInterfaceAccepterArgs | HostedPublicVirtualInterfaceAccepterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: HostedPublicVirtualInterfaceAccepterState = argsOrState as HostedPublicVirtualInterfaceAccepterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualInterfaceId"] = state ? state.virtualInterfaceId : undefined;
        } else {
            const args = argsOrState as HostedPublicVirtualInterfaceAccepterArgs | undefined;
            if (!args || args.virtualInterfaceId === undefined) {
                throw new Error("Missing required property 'virtualInterfaceId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualInterfaceId"] = args ? args.virtualInterfaceId : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedPublicVirtualInterfaceAccepter resources.
 */
export interface HostedPublicVirtualInterfaceAccepterState {
    /**
     * The ARN of the virtual interface.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    readonly virtualInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostedPublicVirtualInterfaceAccepter resource.
 */
export interface HostedPublicVirtualInterfaceAccepterArgs {
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    readonly virtualInterfaceId: pulumi.Input<string>;
}
