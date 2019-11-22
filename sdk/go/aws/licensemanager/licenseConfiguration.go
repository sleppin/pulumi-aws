// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a License Manager license configuration resource.
// 
// > **Note:** Removing the `licenseCount` attribute is not supported by the License Manager API - recreate the resource instead.
// 
// ## Rules
// 
// License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:
// 
// * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
// * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
// * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
// * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
// * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
// * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
// * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/licensemanager_license_configuration.html.markdown.
type LicenseConfiguration struct {
	s *pulumi.ResourceState
}

// NewLicenseConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLicenseConfiguration(ctx *pulumi.Context,
	name string, args *LicenseConfigurationArgs, opts ...pulumi.ResourceOpt) (*LicenseConfiguration, error) {
	if args == nil || args.LicenseCountingType == nil {
		return nil, errors.New("missing required argument 'LicenseCountingType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["licenseCount"] = nil
		inputs["licenseCountHardLimit"] = nil
		inputs["licenseCountingType"] = nil
		inputs["licenseRules"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["licenseCount"] = args.LicenseCount
		inputs["licenseCountHardLimit"] = args.LicenseCountHardLimit
		inputs["licenseCountingType"] = args.LicenseCountingType
		inputs["licenseRules"] = args.LicenseRules
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	s, err := ctx.RegisterResource("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LicenseConfiguration{s: s}, nil
}

// GetLicenseConfiguration gets an existing LicenseConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicenseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LicenseConfigurationState, opts ...pulumi.ResourceOpt) (*LicenseConfiguration, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["licenseCount"] = state.LicenseCount
		inputs["licenseCountHardLimit"] = state.LicenseCountHardLimit
		inputs["licenseCountingType"] = state.LicenseCountingType
		inputs["licenseRules"] = state.LicenseRules
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LicenseConfiguration{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LicenseConfiguration) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LicenseConfiguration) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Description of the license configuration.
func (r *LicenseConfiguration) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Number of licenses managed by the license configuration.
func (r *LicenseConfiguration) LicenseCount() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["licenseCount"])
}

// Sets the number of available licenses as a hard limit.
func (r *LicenseConfiguration) LicenseCountHardLimit() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["licenseCountHardLimit"])
}

// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
func (r *LicenseConfiguration) LicenseCountingType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["licenseCountingType"])
}

// Array of configured License Manager rules.
func (r *LicenseConfiguration) LicenseRules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["licenseRules"])
}

// Name of the license configuration.
func (r *LicenseConfiguration) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A mapping of tags to assign to the resource.
func (r *LicenseConfiguration) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering LicenseConfiguration resources.
type LicenseConfigurationState struct {
	// Description of the license configuration.
	Description interface{}
	// Number of licenses managed by the license configuration.
	LicenseCount interface{}
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit interface{}
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType interface{}
	// Array of configured License Manager rules.
	LicenseRules interface{}
	// Name of the license configuration.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a LicenseConfiguration resource.
type LicenseConfigurationArgs struct {
	// Description of the license configuration.
	Description interface{}
	// Number of licenses managed by the license configuration.
	LicenseCount interface{}
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit interface{}
	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType interface{}
	// Array of configured License Manager rules.
	LicenseRules interface{}
	// Name of the license configuration.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
