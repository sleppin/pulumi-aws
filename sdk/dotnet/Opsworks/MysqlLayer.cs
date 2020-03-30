// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpsWorks
{
    /// <summary>
    /// Provides an OpsWorks MySQL layer resource.
    /// 
    /// &gt; **Note:** All arguments including the root password will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_mysql_layer.html.markdown.
    /// </summary>
    public partial class MysqlLayer : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name(ARN) of the layer.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically assign an elastic IP address to the layer's instances.
        /// </summary>
        [Output("autoAssignElasticIps")]
        public Output<bool?> AutoAssignElasticIps { get; private set; } = null!;

        /// <summary>
        /// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        /// </summary>
        [Output("autoAssignPublicIps")]
        public Output<bool?> AutoAssignPublicIps { get; private set; } = null!;

        /// <summary>
        /// Whether to enable auto-healing for the layer.
        /// </summary>
        [Output("autoHealing")]
        public Output<bool?> AutoHealing { get; private set; } = null!;

        [Output("customConfigureRecipes")]
        public Output<ImmutableArray<string>> CustomConfigureRecipes { get; private set; } = null!;

        [Output("customDeployRecipes")]
        public Output<ImmutableArray<string>> CustomDeployRecipes { get; private set; } = null!;

        /// <summary>
        /// The ARN of an IAM profile that will be used for the layer's instances.
        /// </summary>
        [Output("customInstanceProfileArn")]
        public Output<string?> CustomInstanceProfileArn { get; private set; } = null!;

        /// <summary>
        /// Custom JSON attributes to apply to the layer.
        /// </summary>
        [Output("customJson")]
        public Output<string?> CustomJson { get; private set; } = null!;

        /// <summary>
        /// Ids for a set of security groups to apply to the layer's instances.
        /// </summary>
        [Output("customSecurityGroupIds")]
        public Output<ImmutableArray<string>> CustomSecurityGroupIds { get; private set; } = null!;

        [Output("customSetupRecipes")]
        public Output<ImmutableArray<string>> CustomSetupRecipes { get; private set; } = null!;

        [Output("customShutdownRecipes")]
        public Output<ImmutableArray<string>> CustomShutdownRecipes { get; private set; } = null!;

        [Output("customUndeployRecipes")]
        public Output<ImmutableArray<string>> CustomUndeployRecipes { get; private set; } = null!;

        /// <summary>
        /// Whether to enable Elastic Load Balancing connection draining.
        /// </summary>
        [Output("drainElbOnShutdown")]
        public Output<bool?> DrainElbOnShutdown { get; private set; } = null!;

        /// <summary>
        /// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        /// </summary>
        [Output("ebsVolumes")]
        public Output<ImmutableArray<Outputs.MysqlLayerEbsVolumes>> EbsVolumes { get; private set; } = null!;

        /// <summary>
        /// Name of an Elastic Load Balancer to attach to this layer
        /// </summary>
        [Output("elasticLoadBalancer")]
        public Output<string?> ElasticLoadBalancer { get; private set; } = null!;

        /// <summary>
        /// Whether to install OS and package updates on each instance when it boots.
        /// </summary>
        [Output("installUpdatesOnBoot")]
        public Output<bool?> InstallUpdatesOnBoot { get; private set; } = null!;

        /// <summary>
        /// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        /// </summary>
        [Output("instanceShutdownTimeout")]
        public Output<int?> InstanceShutdownTimeout { get; private set; } = null!;

        /// <summary>
        /// A human-readable name for the layer.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Root password to use for MySQL.
        /// </summary>
        [Output("rootPassword")]
        public Output<string?> RootPassword { get; private set; } = null!;

        /// <summary>
        /// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
        /// </summary>
        [Output("rootPasswordOnAllInstances")]
        public Output<bool?> RootPasswordOnAllInstances { get; private set; } = null!;

        /// <summary>
        /// The id of the stack the layer will belong to.
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// Names of a set of system packages to install on the layer's instances.
        /// </summary>
        [Output("systemPackages")]
        public Output<ImmutableArray<string>> SystemPackages { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether to use EBS-optimized instances.
        /// </summary>
        [Output("useEbsOptimizedInstances")]
        public Output<bool?> UseEbsOptimizedInstances { get; private set; } = null!;


        /// <summary>
        /// Create a MysqlLayer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MysqlLayer(string name, MysqlLayerArgs args, CustomResourceOptions? options = null)
            : base("aws:opsworks/mysqlLayer:MysqlLayer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private MysqlLayer(string name, Input<string> id, MysqlLayerState? state = null, CustomResourceOptions? options = null)
            : base("aws:opsworks/mysqlLayer:MysqlLayer", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MysqlLayer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MysqlLayer Get(string name, Input<string> id, MysqlLayerState? state = null, CustomResourceOptions? options = null)
        {
            return new MysqlLayer(name, id, state, options);
        }
    }

    public sealed class MysqlLayerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically assign an elastic IP address to the layer's instances.
        /// </summary>
        [Input("autoAssignElasticIps")]
        public Input<bool>? AutoAssignElasticIps { get; set; }

        /// <summary>
        /// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        /// </summary>
        [Input("autoAssignPublicIps")]
        public Input<bool>? AutoAssignPublicIps { get; set; }

        /// <summary>
        /// Whether to enable auto-healing for the layer.
        /// </summary>
        [Input("autoHealing")]
        public Input<bool>? AutoHealing { get; set; }

        [Input("customConfigureRecipes")]
        private InputList<string>? _customConfigureRecipes;
        public InputList<string> CustomConfigureRecipes
        {
            get => _customConfigureRecipes ?? (_customConfigureRecipes = new InputList<string>());
            set => _customConfigureRecipes = value;
        }

        [Input("customDeployRecipes")]
        private InputList<string>? _customDeployRecipes;
        public InputList<string> CustomDeployRecipes
        {
            get => _customDeployRecipes ?? (_customDeployRecipes = new InputList<string>());
            set => _customDeployRecipes = value;
        }

        /// <summary>
        /// The ARN of an IAM profile that will be used for the layer's instances.
        /// </summary>
        [Input("customInstanceProfileArn")]
        public Input<string>? CustomInstanceProfileArn { get; set; }

        /// <summary>
        /// Custom JSON attributes to apply to the layer.
        /// </summary>
        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("customSecurityGroupIds")]
        private InputList<string>? _customSecurityGroupIds;

        /// <summary>
        /// Ids for a set of security groups to apply to the layer's instances.
        /// </summary>
        public InputList<string> CustomSecurityGroupIds
        {
            get => _customSecurityGroupIds ?? (_customSecurityGroupIds = new InputList<string>());
            set => _customSecurityGroupIds = value;
        }

        [Input("customSetupRecipes")]
        private InputList<string>? _customSetupRecipes;
        public InputList<string> CustomSetupRecipes
        {
            get => _customSetupRecipes ?? (_customSetupRecipes = new InputList<string>());
            set => _customSetupRecipes = value;
        }

        [Input("customShutdownRecipes")]
        private InputList<string>? _customShutdownRecipes;
        public InputList<string> CustomShutdownRecipes
        {
            get => _customShutdownRecipes ?? (_customShutdownRecipes = new InputList<string>());
            set => _customShutdownRecipes = value;
        }

        [Input("customUndeployRecipes")]
        private InputList<string>? _customUndeployRecipes;
        public InputList<string> CustomUndeployRecipes
        {
            get => _customUndeployRecipes ?? (_customUndeployRecipes = new InputList<string>());
            set => _customUndeployRecipes = value;
        }

        /// <summary>
        /// Whether to enable Elastic Load Balancing connection draining.
        /// </summary>
        [Input("drainElbOnShutdown")]
        public Input<bool>? DrainElbOnShutdown { get; set; }

        [Input("ebsVolumes")]
        private InputList<Inputs.MysqlLayerEbsVolumesArgs>? _ebsVolumes;

        /// <summary>
        /// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        /// </summary>
        public InputList<Inputs.MysqlLayerEbsVolumesArgs> EbsVolumes
        {
            get => _ebsVolumes ?? (_ebsVolumes = new InputList<Inputs.MysqlLayerEbsVolumesArgs>());
            set => _ebsVolumes = value;
        }

        /// <summary>
        /// Name of an Elastic Load Balancer to attach to this layer
        /// </summary>
        [Input("elasticLoadBalancer")]
        public Input<string>? ElasticLoadBalancer { get; set; }

        /// <summary>
        /// Whether to install OS and package updates on each instance when it boots.
        /// </summary>
        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        /// </summary>
        [Input("instanceShutdownTimeout")]
        public Input<int>? InstanceShutdownTimeout { get; set; }

        /// <summary>
        /// A human-readable name for the layer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Root password to use for MySQL.
        /// </summary>
        [Input("rootPassword")]
        public Input<string>? RootPassword { get; set; }

        /// <summary>
        /// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
        /// </summary>
        [Input("rootPasswordOnAllInstances")]
        public Input<bool>? RootPasswordOnAllInstances { get; set; }

        /// <summary>
        /// The id of the stack the layer will belong to.
        /// </summary>
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        [Input("systemPackages")]
        private InputList<string>? _systemPackages;

        /// <summary>
        /// Names of a set of system packages to install on the layer's instances.
        /// </summary>
        public InputList<string> SystemPackages
        {
            get => _systemPackages ?? (_systemPackages = new InputList<string>());
            set => _systemPackages = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether to use EBS-optimized instances.
        /// </summary>
        [Input("useEbsOptimizedInstances")]
        public Input<bool>? UseEbsOptimizedInstances { get; set; }

        public MysqlLayerArgs()
        {
        }
    }

    public sealed class MysqlLayerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name(ARN) of the layer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Whether to automatically assign an elastic IP address to the layer's instances.
        /// </summary>
        [Input("autoAssignElasticIps")]
        public Input<bool>? AutoAssignElasticIps { get; set; }

        /// <summary>
        /// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        /// </summary>
        [Input("autoAssignPublicIps")]
        public Input<bool>? AutoAssignPublicIps { get; set; }

        /// <summary>
        /// Whether to enable auto-healing for the layer.
        /// </summary>
        [Input("autoHealing")]
        public Input<bool>? AutoHealing { get; set; }

        [Input("customConfigureRecipes")]
        private InputList<string>? _customConfigureRecipes;
        public InputList<string> CustomConfigureRecipes
        {
            get => _customConfigureRecipes ?? (_customConfigureRecipes = new InputList<string>());
            set => _customConfigureRecipes = value;
        }

        [Input("customDeployRecipes")]
        private InputList<string>? _customDeployRecipes;
        public InputList<string> CustomDeployRecipes
        {
            get => _customDeployRecipes ?? (_customDeployRecipes = new InputList<string>());
            set => _customDeployRecipes = value;
        }

        /// <summary>
        /// The ARN of an IAM profile that will be used for the layer's instances.
        /// </summary>
        [Input("customInstanceProfileArn")]
        public Input<string>? CustomInstanceProfileArn { get; set; }

        /// <summary>
        /// Custom JSON attributes to apply to the layer.
        /// </summary>
        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("customSecurityGroupIds")]
        private InputList<string>? _customSecurityGroupIds;

        /// <summary>
        /// Ids for a set of security groups to apply to the layer's instances.
        /// </summary>
        public InputList<string> CustomSecurityGroupIds
        {
            get => _customSecurityGroupIds ?? (_customSecurityGroupIds = new InputList<string>());
            set => _customSecurityGroupIds = value;
        }

        [Input("customSetupRecipes")]
        private InputList<string>? _customSetupRecipes;
        public InputList<string> CustomSetupRecipes
        {
            get => _customSetupRecipes ?? (_customSetupRecipes = new InputList<string>());
            set => _customSetupRecipes = value;
        }

        [Input("customShutdownRecipes")]
        private InputList<string>? _customShutdownRecipes;
        public InputList<string> CustomShutdownRecipes
        {
            get => _customShutdownRecipes ?? (_customShutdownRecipes = new InputList<string>());
            set => _customShutdownRecipes = value;
        }

        [Input("customUndeployRecipes")]
        private InputList<string>? _customUndeployRecipes;
        public InputList<string> CustomUndeployRecipes
        {
            get => _customUndeployRecipes ?? (_customUndeployRecipes = new InputList<string>());
            set => _customUndeployRecipes = value;
        }

        /// <summary>
        /// Whether to enable Elastic Load Balancing connection draining.
        /// </summary>
        [Input("drainElbOnShutdown")]
        public Input<bool>? DrainElbOnShutdown { get; set; }

        [Input("ebsVolumes")]
        private InputList<Inputs.MysqlLayerEbsVolumesGetArgs>? _ebsVolumes;

        /// <summary>
        /// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        /// </summary>
        public InputList<Inputs.MysqlLayerEbsVolumesGetArgs> EbsVolumes
        {
            get => _ebsVolumes ?? (_ebsVolumes = new InputList<Inputs.MysqlLayerEbsVolumesGetArgs>());
            set => _ebsVolumes = value;
        }

        /// <summary>
        /// Name of an Elastic Load Balancer to attach to this layer
        /// </summary>
        [Input("elasticLoadBalancer")]
        public Input<string>? ElasticLoadBalancer { get; set; }

        /// <summary>
        /// Whether to install OS and package updates on each instance when it boots.
        /// </summary>
        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        /// </summary>
        [Input("instanceShutdownTimeout")]
        public Input<int>? InstanceShutdownTimeout { get; set; }

        /// <summary>
        /// A human-readable name for the layer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Root password to use for MySQL.
        /// </summary>
        [Input("rootPassword")]
        public Input<string>? RootPassword { get; set; }

        /// <summary>
        /// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
        /// </summary>
        [Input("rootPasswordOnAllInstances")]
        public Input<bool>? RootPasswordOnAllInstances { get; set; }

        /// <summary>
        /// The id of the stack the layer will belong to.
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("systemPackages")]
        private InputList<string>? _systemPackages;

        /// <summary>
        /// Names of a set of system packages to install on the layer's instances.
        /// </summary>
        public InputList<string> SystemPackages
        {
            get => _systemPackages ?? (_systemPackages = new InputList<string>());
            set => _systemPackages = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether to use EBS-optimized instances.
        /// </summary>
        [Input("useEbsOptimizedInstances")]
        public Input<bool>? UseEbsOptimizedInstances { get; set; }

        public MysqlLayerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class MysqlLayerEbsVolumesArgs : Pulumi.ResourceArgs
    {
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// For PIOPS volumes, the IOPS per disk.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The path to mount the EBS volume on the layer's instances.
        /// </summary>
        [Input("mountPoint", required: true)]
        public Input<string> MountPoint { get; set; } = null!;

        /// <summary>
        /// The number of disks to use for the EBS volume.
        /// </summary>
        [Input("numberOfDisks", required: true)]
        public Input<int> NumberOfDisks { get; set; } = null!;

        /// <summary>
        /// The RAID level to use for the volume.
        /// </summary>
        [Input("raidLevel")]
        public Input<string>? RaidLevel { get; set; }

        /// <summary>
        /// The size of the volume in gigabytes.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The type of volume to create. This may be `standard` (the default), `io1` or `gp2`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MysqlLayerEbsVolumesArgs()
        {
        }
    }

    public sealed class MysqlLayerEbsVolumesGetArgs : Pulumi.ResourceArgs
    {
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// For PIOPS volumes, the IOPS per disk.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The path to mount the EBS volume on the layer's instances.
        /// </summary>
        [Input("mountPoint", required: true)]
        public Input<string> MountPoint { get; set; } = null!;

        /// <summary>
        /// The number of disks to use for the EBS volume.
        /// </summary>
        [Input("numberOfDisks", required: true)]
        public Input<int> NumberOfDisks { get; set; } = null!;

        /// <summary>
        /// The RAID level to use for the volume.
        /// </summary>
        [Input("raidLevel")]
        public Input<string>? RaidLevel { get; set; }

        /// <summary>
        /// The size of the volume in gigabytes.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The type of volume to create. This may be `standard` (the default), `io1` or `gp2`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MysqlLayerEbsVolumesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class MysqlLayerEbsVolumes
    {
        public readonly bool? Encrypted;
        /// <summary>
        /// For PIOPS volumes, the IOPS per disk.
        /// </summary>
        public readonly int? Iops;
        /// <summary>
        /// The path to mount the EBS volume on the layer's instances.
        /// </summary>
        public readonly string MountPoint;
        /// <summary>
        /// The number of disks to use for the EBS volume.
        /// </summary>
        public readonly int NumberOfDisks;
        /// <summary>
        /// The RAID level to use for the volume.
        /// </summary>
        public readonly string? RaidLevel;
        /// <summary>
        /// The size of the volume in gigabytes.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The type of volume to create. This may be `standard` (the default), `io1` or `gp2`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private MysqlLayerEbsVolumes(
            bool? encrypted,
            int? iops,
            string mountPoint,
            int numberOfDisks,
            string? raidLevel,
            int size,
            string? type)
        {
            Encrypted = encrypted;
            Iops = iops;
            MountPoint = mountPoint;
            NumberOfDisks = numberOfDisks;
            RaidLevel = raidLevel;
            Size = size;
            Type = type;
        }
    }
    }
}
