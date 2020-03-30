// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancingv2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.
//
// Provides information about a Load Balancer.
//
// This data source can prove useful when a module accepts an LB as an input
// variable and needs to, for example, determine the security groups associated
// with it, etc.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lb.html.markdown.
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws:elasticloadbalancingv2/getLoadBalancer:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerArgs struct {
	// The full ARN of the load balancer.
	Arn *string `pulumi:"arn"`
	// The unique name of the load balancer.
	Name *string                `pulumi:"name"`
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getLoadBalancer.
type LookupLoadBalancerResult struct {
	AccessLogs               GetLoadBalancerAccessLogs `pulumi:"accessLogs"`
	Arn                      string                    `pulumi:"arn"`
	ArnSuffix                string                    `pulumi:"arnSuffix"`
	DnsName                  string                    `pulumi:"dnsName"`
	DropInvalidHeaderFields  bool                      `pulumi:"dropInvalidHeaderFields"`
	EnableDeletionProtection bool                      `pulumi:"enableDeletionProtection"`
	// id is the provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	IdleTimeout      int                            `pulumi:"idleTimeout"`
	Internal         bool                           `pulumi:"internal"`
	LoadBalancerType string                         `pulumi:"loadBalancerType"`
	Name             string                         `pulumi:"name"`
	SecurityGroups   []string                       `pulumi:"securityGroups"`
	SubnetMappings   []GetLoadBalancerSubnetMapping `pulumi:"subnetMappings"`
	Subnets          []string                       `pulumi:"subnets"`
	Tags             map[string]interface{}         `pulumi:"tags"`
	VpcId            string                         `pulumi:"vpcId"`
	ZoneId           string                         `pulumi:"zoneId"`
}
