// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks Java application layer resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_java_app_layer.html.markdown.
type JavaAppLayer struct {
	s *pulumi.ResourceState
}

// NewJavaAppLayer registers a new resource with the given unique name, arguments, and options.
func NewJavaAppLayer(ctx *pulumi.Context,
	name string, args *JavaAppLayerArgs, opts ...pulumi.ResourceOpt) (*JavaAppLayer, error) {
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appServer"] = nil
		inputs["appServerVersion"] = nil
		inputs["autoAssignElasticIps"] = nil
		inputs["autoAssignPublicIps"] = nil
		inputs["autoHealing"] = nil
		inputs["customConfigureRecipes"] = nil
		inputs["customDeployRecipes"] = nil
		inputs["customInstanceProfileArn"] = nil
		inputs["customJson"] = nil
		inputs["customSecurityGroupIds"] = nil
		inputs["customSetupRecipes"] = nil
		inputs["customShutdownRecipes"] = nil
		inputs["customUndeployRecipes"] = nil
		inputs["drainElbOnShutdown"] = nil
		inputs["ebsVolumes"] = nil
		inputs["elasticLoadBalancer"] = nil
		inputs["installUpdatesOnBoot"] = nil
		inputs["instanceShutdownTimeout"] = nil
		inputs["jvmOptions"] = nil
		inputs["jvmType"] = nil
		inputs["jvmVersion"] = nil
		inputs["name"] = nil
		inputs["stackId"] = nil
		inputs["systemPackages"] = nil
		inputs["useEbsOptimizedInstances"] = nil
	} else {
		inputs["appServer"] = args.AppServer
		inputs["appServerVersion"] = args.AppServerVersion
		inputs["autoAssignElasticIps"] = args.AutoAssignElasticIps
		inputs["autoAssignPublicIps"] = args.AutoAssignPublicIps
		inputs["autoHealing"] = args.AutoHealing
		inputs["customConfigureRecipes"] = args.CustomConfigureRecipes
		inputs["customDeployRecipes"] = args.CustomDeployRecipes
		inputs["customInstanceProfileArn"] = args.CustomInstanceProfileArn
		inputs["customJson"] = args.CustomJson
		inputs["customSecurityGroupIds"] = args.CustomSecurityGroupIds
		inputs["customSetupRecipes"] = args.CustomSetupRecipes
		inputs["customShutdownRecipes"] = args.CustomShutdownRecipes
		inputs["customUndeployRecipes"] = args.CustomUndeployRecipes
		inputs["drainElbOnShutdown"] = args.DrainElbOnShutdown
		inputs["ebsVolumes"] = args.EbsVolumes
		inputs["elasticLoadBalancer"] = args.ElasticLoadBalancer
		inputs["installUpdatesOnBoot"] = args.InstallUpdatesOnBoot
		inputs["instanceShutdownTimeout"] = args.InstanceShutdownTimeout
		inputs["jvmOptions"] = args.JvmOptions
		inputs["jvmType"] = args.JvmType
		inputs["jvmVersion"] = args.JvmVersion
		inputs["name"] = args.Name
		inputs["stackId"] = args.StackId
		inputs["systemPackages"] = args.SystemPackages
		inputs["useEbsOptimizedInstances"] = args.UseEbsOptimizedInstances
	}
	s, err := ctx.RegisterResource("aws:opsworks/javaAppLayer:JavaAppLayer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JavaAppLayer{s: s}, nil
}

// GetJavaAppLayer gets an existing JavaAppLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJavaAppLayer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JavaAppLayerState, opts ...pulumi.ResourceOpt) (*JavaAppLayer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appServer"] = state.AppServer
		inputs["appServerVersion"] = state.AppServerVersion
		inputs["autoAssignElasticIps"] = state.AutoAssignElasticIps
		inputs["autoAssignPublicIps"] = state.AutoAssignPublicIps
		inputs["autoHealing"] = state.AutoHealing
		inputs["customConfigureRecipes"] = state.CustomConfigureRecipes
		inputs["customDeployRecipes"] = state.CustomDeployRecipes
		inputs["customInstanceProfileArn"] = state.CustomInstanceProfileArn
		inputs["customJson"] = state.CustomJson
		inputs["customSecurityGroupIds"] = state.CustomSecurityGroupIds
		inputs["customSetupRecipes"] = state.CustomSetupRecipes
		inputs["customShutdownRecipes"] = state.CustomShutdownRecipes
		inputs["customUndeployRecipes"] = state.CustomUndeployRecipes
		inputs["drainElbOnShutdown"] = state.DrainElbOnShutdown
		inputs["ebsVolumes"] = state.EbsVolumes
		inputs["elasticLoadBalancer"] = state.ElasticLoadBalancer
		inputs["installUpdatesOnBoot"] = state.InstallUpdatesOnBoot
		inputs["instanceShutdownTimeout"] = state.InstanceShutdownTimeout
		inputs["jvmOptions"] = state.JvmOptions
		inputs["jvmType"] = state.JvmType
		inputs["jvmVersion"] = state.JvmVersion
		inputs["name"] = state.Name
		inputs["stackId"] = state.StackId
		inputs["systemPackages"] = state.SystemPackages
		inputs["useEbsOptimizedInstances"] = state.UseEbsOptimizedInstances
	}
	s, err := ctx.ReadResource("aws:opsworks/javaAppLayer:JavaAppLayer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JavaAppLayer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *JavaAppLayer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *JavaAppLayer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Keyword for the application container to use. Defaults to "tomcat".
func (r *JavaAppLayer) AppServer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["appServer"])
}

// Version of the selected application container to use. Defaults to "7".
func (r *JavaAppLayer) AppServerVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["appServerVersion"])
}

// Whether to automatically assign an elastic IP address to the layer's instances.
func (r *JavaAppLayer) AutoAssignElasticIps() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoAssignElasticIps"])
}

// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
func (r *JavaAppLayer) AutoAssignPublicIps() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoAssignPublicIps"])
}

// Whether to enable auto-healing for the layer.
func (r *JavaAppLayer) AutoHealing() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoHealing"])
}

func (r *JavaAppLayer) CustomConfigureRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customConfigureRecipes"])
}

func (r *JavaAppLayer) CustomDeployRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customDeployRecipes"])
}

// The ARN of an IAM profile that will be used for the layer's instances.
func (r *JavaAppLayer) CustomInstanceProfileArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customInstanceProfileArn"])
}

// Custom JSON attributes to apply to the layer.
func (r *JavaAppLayer) CustomJson() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customJson"])
}

// Ids for a set of security groups to apply to the layer's instances.
func (r *JavaAppLayer) CustomSecurityGroupIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customSecurityGroupIds"])
}

func (r *JavaAppLayer) CustomSetupRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customSetupRecipes"])
}

func (r *JavaAppLayer) CustomShutdownRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customShutdownRecipes"])
}

func (r *JavaAppLayer) CustomUndeployRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customUndeployRecipes"])
}

// Whether to enable Elastic Load Balancing connection draining.
func (r *JavaAppLayer) DrainElbOnShutdown() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["drainElbOnShutdown"])
}

// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
func (r *JavaAppLayer) EbsVolumes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ebsVolumes"])
}

// Name of an Elastic Load Balancer to attach to this layer
func (r *JavaAppLayer) ElasticLoadBalancer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["elasticLoadBalancer"])
}

// Whether to install OS and package updates on each instance when it boots.
func (r *JavaAppLayer) InstallUpdatesOnBoot() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["installUpdatesOnBoot"])
}

// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
func (r *JavaAppLayer) InstanceShutdownTimeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["instanceShutdownTimeout"])
}

// Options to set for the JVM.
func (r *JavaAppLayer) JvmOptions() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["jvmOptions"])
}

// Keyword for the type of JVM to use. Defaults to `openjdk`.
func (r *JavaAppLayer) JvmType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["jvmType"])
}

// Version of JVM to use. Defaults to "7".
func (r *JavaAppLayer) JvmVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["jvmVersion"])
}

// A human-readable name for the layer.
func (r *JavaAppLayer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The id of the stack the layer will belong to.
func (r *JavaAppLayer) StackId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["stackId"])
}

// Names of a set of system packages to install on the layer's instances.
func (r *JavaAppLayer) SystemPackages() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["systemPackages"])
}

// Whether to use EBS-optimized instances.
func (r *JavaAppLayer) UseEbsOptimizedInstances() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["useEbsOptimizedInstances"])
}

// Input properties used for looking up and filtering JavaAppLayer resources.
type JavaAppLayerState struct {
	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer interface{}
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion interface{}
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps interface{}
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps interface{}
	// Whether to enable auto-healing for the layer.
	AutoHealing interface{}
	CustomConfigureRecipes interface{}
	CustomDeployRecipes interface{}
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn interface{}
	// Custom JSON attributes to apply to the layer.
	CustomJson interface{}
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds interface{}
	CustomSetupRecipes interface{}
	CustomShutdownRecipes interface{}
	CustomUndeployRecipes interface{}
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown interface{}
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes interface{}
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer interface{}
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot interface{}
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout interface{}
	// Options to set for the JVM.
	JvmOptions interface{}
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType interface{}
	// Version of JVM to use. Defaults to "7".
	JvmVersion interface{}
	// A human-readable name for the layer.
	Name interface{}
	// The id of the stack the layer will belong to.
	StackId interface{}
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages interface{}
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances interface{}
}

// The set of arguments for constructing a JavaAppLayer resource.
type JavaAppLayerArgs struct {
	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer interface{}
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion interface{}
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps interface{}
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps interface{}
	// Whether to enable auto-healing for the layer.
	AutoHealing interface{}
	CustomConfigureRecipes interface{}
	CustomDeployRecipes interface{}
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn interface{}
	// Custom JSON attributes to apply to the layer.
	CustomJson interface{}
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds interface{}
	CustomSetupRecipes interface{}
	CustomShutdownRecipes interface{}
	CustomUndeployRecipes interface{}
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown interface{}
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes interface{}
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer interface{}
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot interface{}
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout interface{}
	// Options to set for the JVM.
	JvmOptions interface{}
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType interface{}
	// Version of JVM to use. Defaults to "7".
	JvmVersion interface{}
	// A human-readable name for the layer.
	Name interface{}
	// The id of the stack the layer will belong to.
	StackId interface{}
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages interface{}
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances interface{}
}
