// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The VPC Peering Connection data source provides details about
// a specific VPC peering connection.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc_peering_connection.html.markdown.
func LookupVpcPeeringConnection(ctx *pulumi.Context, args *LookupVpcPeeringConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpcPeeringConnectionResult, error) {
	var rv LookupVpcPeeringConnectionResult
	err := ctx.Invoke("aws:ec2/getVpcPeeringConnection:getVpcPeeringConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPeeringConnection.
type LookupVpcPeeringConnectionArgs struct {
	// The CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Custom filter block as described below.
	Filters []GetVpcPeeringConnectionFilter `pulumi:"filters"`
	// The ID of the specific VPC Peering Connection to retrieve.
	Id *string `pulumi:"id"`
	// The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.
	OwnerId *string `pulumi:"ownerId"`
	// The CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerCidrBlock *string `pulumi:"peerCidrBlock"`
	// The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerOwnerId *string `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerRegion *string `pulumi:"peerRegion"`
	// The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId *string `pulumi:"peerVpcId"`
	// The region of the requester VPC of the specific VPC Peering Connection to retrieve.
	Region *string `pulumi:"region"`
	// The status of the specific VPC Peering Connection to retrieve.
	Status *string `pulumi:"status"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired VPC Peering Connection.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId *string `pulumi:"vpcId"`
}


// A collection of values returned by getVpcPeeringConnection.
type LookupVpcPeeringConnectionResult struct {
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter map[string]bool `pulumi:"accepter"`
	CidrBlock string `pulumi:"cidrBlock"`
	Filters []GetVpcPeeringConnectionFilter `pulumi:"filters"`
	Id string `pulumi:"id"`
	OwnerId string `pulumi:"ownerId"`
	PeerCidrBlock string `pulumi:"peerCidrBlock"`
	PeerOwnerId string `pulumi:"peerOwnerId"`
	PeerRegion string `pulumi:"peerRegion"`
	PeerVpcId string `pulumi:"peerVpcId"`
	Region string `pulumi:"region"`
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester map[string]bool `pulumi:"requester"`
	Status string `pulumi:"status"`
	Tags map[string]interface{} `pulumi:"tags"`
	VpcId string `pulumi:"vpcId"`
}

