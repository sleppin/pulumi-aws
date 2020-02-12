// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gamelift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Gamelift Fleet resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_fleet.html.markdown.
type Fleet struct {
	pulumi.CustomResourceState

	// Fleet ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId pulumi.StringOutput `pulumi:"buildId"`
	// Human-readable description of the fleet.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions FleetEc2InboundPermissionArrayOutput `pulumi:"ec2InboundPermissions"`
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType pulumi.StringOutput `pulumi:"ec2InstanceType"`
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType pulumi.StringPtrOutput `pulumi:"fleetType"`
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn pulumi.StringPtrOutput `pulumi:"instanceRoleArn"`
	LogPaths pulumi.StringArrayOutput `pulumi:"logPaths"`
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups pulumi.StringArrayOutput `pulumi:"metricGroups"`
	// The name of the fleet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy pulumi.StringPtrOutput `pulumi:"newGameSessionProtectionPolicy"`
	// Operating system of the fleet's computing resources.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy FleetResourceCreationLimitPolicyPtrOutput `pulumi:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration FleetRuntimeConfigurationPtrOutput `pulumi:"runtimeConfiguration"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil || args.BuildId == nil {
		return nil, errors.New("missing required argument 'BuildId'")
	}
	if args == nil || args.Ec2InstanceType == nil {
		return nil, errors.New("missing required argument 'Ec2InstanceType'")
	}
	if args == nil {
		args = &FleetArgs{}
	}
	var resource Fleet
	err := ctx.RegisterResource("aws:gamelift/fleet:Fleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetState, opts ...pulumi.ResourceOption) (*Fleet, error) {
	var resource Fleet
	err := ctx.ReadResource("aws:gamelift/fleet:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fleet resources.
type fleetState struct {
	// Fleet ARN.
	Arn *string `pulumi:"arn"`
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId *string `pulumi:"buildId"`
	// Human-readable description of the fleet.
	Description *string `pulumi:"description"`
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions []FleetEc2InboundPermission `pulumi:"ec2InboundPermissions"`
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType *string `pulumi:"ec2InstanceType"`
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType *string `pulumi:"fleetType"`
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn *string `pulumi:"instanceRoleArn"`
	LogPaths []string `pulumi:"logPaths"`
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups []string `pulumi:"metricGroups"`
	// The name of the fleet.
	Name *string `pulumi:"name"`
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy *string `pulumi:"newGameSessionProtectionPolicy"`
	// Operating system of the fleet's computing resources.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy *FleetResourceCreationLimitPolicy `pulumi:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration *FleetRuntimeConfiguration `pulumi:"runtimeConfiguration"`
}

type FleetState struct {
	// Fleet ARN.
	Arn pulumi.StringPtrInput
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId pulumi.StringPtrInput
	// Human-readable description of the fleet.
	Description pulumi.StringPtrInput
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions FleetEc2InboundPermissionArrayInput
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType pulumi.StringPtrInput
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType pulumi.StringPtrInput
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn pulumi.StringPtrInput
	LogPaths pulumi.StringArrayInput
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups pulumi.StringArrayInput
	// The name of the fleet.
	Name pulumi.StringPtrInput
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy pulumi.StringPtrInput
	// Operating system of the fleet's computing resources.
	OperatingSystem pulumi.StringPtrInput
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy FleetResourceCreationLimitPolicyPtrInput
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration FleetRuntimeConfigurationPtrInput
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId string `pulumi:"buildId"`
	// Human-readable description of the fleet.
	Description *string `pulumi:"description"`
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions []FleetEc2InboundPermission `pulumi:"ec2InboundPermissions"`
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType string `pulumi:"ec2InstanceType"`
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType *string `pulumi:"fleetType"`
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn *string `pulumi:"instanceRoleArn"`
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups []string `pulumi:"metricGroups"`
	// The name of the fleet.
	Name *string `pulumi:"name"`
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy *string `pulumi:"newGameSessionProtectionPolicy"`
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy *FleetResourceCreationLimitPolicy `pulumi:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration *FleetRuntimeConfiguration `pulumi:"runtimeConfiguration"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId pulumi.StringInput
	// Human-readable description of the fleet.
	Description pulumi.StringPtrInput
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions FleetEc2InboundPermissionArrayInput
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType pulumi.StringInput
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType pulumi.StringPtrInput
	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn pulumi.StringPtrInput
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups pulumi.StringArrayInput
	// The name of the fleet.
	Name pulumi.StringPtrInput
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy pulumi.StringPtrInput
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy FleetResourceCreationLimitPolicyPtrInput
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration FleetRuntimeConfigurationPtrInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

