// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Route53 record resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_record.html.markdown.
type Record struct {
	s *pulumi.ResourceState
}

// NewRecord registers a new resource with the given unique name, arguments, and options.
func NewRecord(ctx *pulumi.Context,
	name string, args *RecordArgs, opts ...pulumi.ResourceOpt) (*Record, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["aliases"] = nil
		inputs["allowOverwrite"] = nil
		inputs["failoverRoutingPolicies"] = nil
		inputs["geolocationRoutingPolicies"] = nil
		inputs["healthCheckId"] = nil
		inputs["latencyRoutingPolicies"] = nil
		inputs["multivalueAnswerRoutingPolicy"] = nil
		inputs["name"] = nil
		inputs["records"] = nil
		inputs["setIdentifier"] = nil
		inputs["ttl"] = nil
		inputs["type"] = nil
		inputs["weightedRoutingPolicies"] = nil
		inputs["zoneId"] = nil
	} else {
		inputs["aliases"] = args.Aliases
		inputs["allowOverwrite"] = args.AllowOverwrite
		inputs["failoverRoutingPolicies"] = args.FailoverRoutingPolicies
		inputs["geolocationRoutingPolicies"] = args.GeolocationRoutingPolicies
		inputs["healthCheckId"] = args.HealthCheckId
		inputs["latencyRoutingPolicies"] = args.LatencyRoutingPolicies
		inputs["multivalueAnswerRoutingPolicy"] = args.MultivalueAnswerRoutingPolicy
		inputs["name"] = args.Name
		inputs["records"] = args.Records
		inputs["setIdentifier"] = args.SetIdentifier
		inputs["ttl"] = args.Ttl
		inputs["type"] = args.Type
		inputs["weightedRoutingPolicies"] = args.WeightedRoutingPolicies
		inputs["zoneId"] = args.ZoneId
	}
	inputs["fqdn"] = nil
	s, err := ctx.RegisterResource("aws:route53/record:Record", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Record{s: s}, nil
}

// GetRecord gets an existing Record resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RecordState, opts ...pulumi.ResourceOpt) (*Record, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["aliases"] = state.Aliases
		inputs["allowOverwrite"] = state.AllowOverwrite
		inputs["failoverRoutingPolicies"] = state.FailoverRoutingPolicies
		inputs["fqdn"] = state.Fqdn
		inputs["geolocationRoutingPolicies"] = state.GeolocationRoutingPolicies
		inputs["healthCheckId"] = state.HealthCheckId
		inputs["latencyRoutingPolicies"] = state.LatencyRoutingPolicies
		inputs["multivalueAnswerRoutingPolicy"] = state.MultivalueAnswerRoutingPolicy
		inputs["name"] = state.Name
		inputs["records"] = state.Records
		inputs["setIdentifier"] = state.SetIdentifier
		inputs["ttl"] = state.Ttl
		inputs["type"] = state.Type
		inputs["weightedRoutingPolicies"] = state.WeightedRoutingPolicies
		inputs["zoneId"] = state.ZoneId
	}
	s, err := ctx.ReadResource("aws:route53/record:Record", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Record{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Record) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Record) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An alias block. Conflicts with `ttl` & `records`.
// Alias record documented below.
func (r *Record) Aliases() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["aliases"])
}

// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
func (r *Record) AllowOverwrite() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["allowOverwrite"])
}

// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
func (r *Record) FailoverRoutingPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["failoverRoutingPolicies"])
}

// [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
func (r *Record) Fqdn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["fqdn"])
}

// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
func (r *Record) GeolocationRoutingPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["geolocationRoutingPolicies"])
}

// The health check the record should be associated with.
func (r *Record) HealthCheckId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["healthCheckId"])
}

// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
func (r *Record) LatencyRoutingPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["latencyRoutingPolicies"])
}

// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
func (r *Record) MultivalueAnswerRoutingPolicy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["multivalueAnswerRoutingPolicy"])
}

// DNS domain name for a CloudFront distribution, S3 bucket, ELB, or another resource record set in this hosted zone.
func (r *Record) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the configuration string (e.g. `"first255characters\"\"morecharacters"`).
func (r *Record) Records() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["records"])
}

// Unique identifier to differentiate records with routing policies from one another. Required if using `failover`, `geolocation`, `latency`, or `weighted` routing policies documented below.
func (r *Record) SetIdentifier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["setIdentifier"])
}

// The TTL of the record.
func (r *Record) Ttl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["ttl"])
}

// `PRIMARY` or `SECONDARY`. A `PRIMARY` record will be served if its healthcheck is passing, otherwise the `SECONDARY` will be served. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring-options.html#dns-failover-failover-rrsets
func (r *Record) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
func (r *Record) WeightedRoutingPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["weightedRoutingPolicies"])
}

// Hosted zone ID for a CloudFront distribution, S3 bucket, ELB, or Route 53 hosted zone. See [`resource_elb.zone_id`](https://www.terraform.io/docs/providers/aws/r/elb.html#zone_id) for example.
func (r *Record) ZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zoneId"])
}

// Input properties used for looking up and filtering Record resources.
type RecordState struct {
	// An alias block. Conflicts with `ttl` & `records`.
	// Alias record documented below.
	Aliases interface{}
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	AllowOverwrite interface{}
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies interface{}
	// [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
	Fqdn interface{}
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies interface{}
	// The health check the record should be associated with.
	HealthCheckId interface{}
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies interface{}
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy interface{}
	// DNS domain name for a CloudFront distribution, S3 bucket, ELB, or another resource record set in this hosted zone.
	Name interface{}
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the configuration string (e.g. `"first255characters\"\"morecharacters"`).
	Records interface{}
	// Unique identifier to differentiate records with routing policies from one another. Required if using `failover`, `geolocation`, `latency`, or `weighted` routing policies documented below.
	SetIdentifier interface{}
	// The TTL of the record.
	Ttl interface{}
	// `PRIMARY` or `SECONDARY`. A `PRIMARY` record will be served if its healthcheck is passing, otherwise the `SECONDARY` will be served. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring-options.html#dns-failover-failover-rrsets
	Type interface{}
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies interface{}
	// Hosted zone ID for a CloudFront distribution, S3 bucket, ELB, or Route 53 hosted zone. See [`resource_elb.zone_id`](https://www.terraform.io/docs/providers/aws/r/elb.html#zone_id) for example.
	ZoneId interface{}
}

// The set of arguments for constructing a Record resource.
type RecordArgs struct {
	// An alias block. Conflicts with `ttl` & `records`.
	// Alias record documented below.
	Aliases interface{}
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	AllowOverwrite interface{}
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies interface{}
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies interface{}
	// The health check the record should be associated with.
	HealthCheckId interface{}
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies interface{}
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy interface{}
	// DNS domain name for a CloudFront distribution, S3 bucket, ELB, or another resource record set in this hosted zone.
	Name interface{}
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the configuration string (e.g. `"first255characters\"\"morecharacters"`).
	Records interface{}
	// Unique identifier to differentiate records with routing policies from one another. Required if using `failover`, `geolocation`, `latency`, or `weighted` routing policies documented below.
	SetIdentifier interface{}
	// The TTL of the record.
	Ttl interface{}
	// `PRIMARY` or `SECONDARY`. A `PRIMARY` record will be served if its healthcheck is passing, otherwise the `SECONDARY` will be served. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring-options.html#dns-failover-failover-rrsets
	Type interface{}
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies interface{}
	// Hosted zone ID for a CloudFront distribution, S3 bucket, ELB, or Route 53 hosted zone. See [`resource_elb.zone_id`](https://www.terraform.io/docs/providers/aws/r/elb.html#zone_id) for example.
	ZoneId interface{}
}
