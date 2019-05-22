// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Direct Connect Gateway Association Proposal, typically for enabling cross-account associations. For single account associations, see the [`aws_dx_gateway_association` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleGateway = new aws.directconnect.Gateway("example", {
 *     amazonSideAsn: "64512",
 * });
 * const exampleVpc = new aws.ec2.Vpc("example", {
 *     cidrBlock: "10.255.255.0/28",
 * });
 * const exampleVpnGateway = new aws.ec2.VpnGateway("example", {
 *     vpcId: exampleVpc.id,
 * });
 * const exampleGatewayAssociationProposal = new aws.directconnect.GatewayAssociationProposal("example", {
 *     dxGatewayId: exampleGateway.id,
 *     dxGatewayOwnerAccountId: exampleGateway.ownerAccountId,
 *     vpnGatewayId: exampleVpnGateway.id,
 * });
 * ```
 */
export class GatewayAssociationProposal extends pulumi.CustomResource {
    /**
     * Get an existing GatewayAssociationProposal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayAssociationProposalState, opts?: pulumi.CustomResourceOptions): GatewayAssociationProposal {
        return new GatewayAssociationProposal(name, <any>state, { ...opts, id: id });
    }

    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     */
    public readonly allowedPrefixes!: pulumi.Output<string[]>;
    /**
     * Direct Connect Gateway identifier.
     */
    public readonly dxGatewayId!: pulumi.Output<string>;
    /**
     * AWS Account identifier of the Direct Connect Gateway.
     */
    public readonly dxGatewayOwnerAccountId!: pulumi.Output<string>;
    /**
     * Virtual Gateway identifier to associate with the Direct Connect Gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string>;

    /**
     * Create a GatewayAssociationProposal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayAssociationProposalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayAssociationProposalArgs | GatewayAssociationProposalState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GatewayAssociationProposalState | undefined;
            inputs["allowedPrefixes"] = state ? state.allowedPrefixes : undefined;
            inputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            inputs["dxGatewayOwnerAccountId"] = state ? state.dxGatewayOwnerAccountId : undefined;
            inputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as GatewayAssociationProposalArgs | undefined;
            if (!args || args.dxGatewayId === undefined) {
                throw new Error("Missing required property 'dxGatewayId'");
            }
            if (!args || args.dxGatewayOwnerAccountId === undefined) {
                throw new Error("Missing required property 'dxGatewayOwnerAccountId'");
            }
            if (!args || args.vpnGatewayId === undefined) {
                throw new Error("Missing required property 'vpnGatewayId'");
            }
            inputs["allowedPrefixes"] = args ? args.allowedPrefixes : undefined;
            inputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            inputs["dxGatewayOwnerAccountId"] = args ? args.dxGatewayOwnerAccountId : undefined;
            inputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
        }
        super("aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayAssociationProposal resources.
 */
export interface GatewayAssociationProposalState {
    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     */
    readonly allowedPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direct Connect Gateway identifier.
     */
    readonly dxGatewayId?: pulumi.Input<string>;
    /**
     * AWS Account identifier of the Direct Connect Gateway.
     */
    readonly dxGatewayOwnerAccountId?: pulumi.Input<string>;
    /**
     * Virtual Gateway identifier to associate with the Direct Connect Gateway.
     */
    readonly vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewayAssociationProposal resource.
 */
export interface GatewayAssociationProposalArgs {
    /**
     * VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
     */
    readonly allowedPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Direct Connect Gateway identifier.
     */
    readonly dxGatewayId: pulumi.Input<string>;
    /**
     * AWS Account identifier of the Direct Connect Gateway.
     */
    readonly dxGatewayOwnerAccountId: pulumi.Input<string>;
    /**
     * Virtual Gateway identifier to associate with the Direct Connect Gateway.
     */
    readonly vpnGatewayId: pulumi.Input<string>;
}