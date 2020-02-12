# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CapacityProvider(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) that identifies the capacity provider.
    """
    auto_scaling_group_provider: pulumi.Output[dict]
    """
    Nested argument defining the provider for the ECS auto scaling group. Defined below.

      * `autoScalingGroupArn` (`str`)
      * `managedScaling` (`dict`)
        * `maximumScalingStepSize` (`float`)
        * `minimumScalingStepSize` (`float`)
        * `status` (`str`)
        * `targetCapacity` (`float`)

      * `managedTerminationProtection` (`str`)
    """
    name: pulumi.Output[str]
    """
    The name of the capacity provider.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags.
    """
    def __init__(__self__, resource_name, opts=None, auto_scaling_group_provider=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_capacity_provider.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_scaling_group_provider: Nested argument defining the provider for the ECS auto scaling group. Defined below.
        :param pulumi.Input[str] name: The name of the capacity provider.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags.

        The **auto_scaling_group_provider** object supports the following:

          * `autoScalingGroupArn` (`pulumi.Input[str]`)
          * `managedScaling` (`pulumi.Input[dict]`)
            * `maximumScalingStepSize` (`pulumi.Input[float]`)
            * `minimumScalingStepSize` (`pulumi.Input[float]`)
            * `status` (`pulumi.Input[str]`)
            * `targetCapacity` (`pulumi.Input[float]`)

          * `managedTerminationProtection` (`pulumi.Input[str]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if auto_scaling_group_provider is None:
                raise TypeError("Missing required property 'auto_scaling_group_provider'")
            __props__['auto_scaling_group_provider'] = auto_scaling_group_provider
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(CapacityProvider, __self__).__init__(
            'aws:ecs/capacityProvider:CapacityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, auto_scaling_group_provider=None, name=None, tags=None):
        """
        Get an existing CapacityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the capacity provider.
        :param pulumi.Input[dict] auto_scaling_group_provider: Nested argument defining the provider for the ECS auto scaling group. Defined below.
        :param pulumi.Input[str] name: The name of the capacity provider.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags.

        The **auto_scaling_group_provider** object supports the following:

          * `autoScalingGroupArn` (`pulumi.Input[str]`)
          * `managedScaling` (`pulumi.Input[dict]`)
            * `maximumScalingStepSize` (`pulumi.Input[float]`)
            * `minimumScalingStepSize` (`pulumi.Input[float]`)
            * `status` (`pulumi.Input[str]`)
            * `targetCapacity` (`pulumi.Input[float]`)

          * `managedTerminationProtection` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["auto_scaling_group_provider"] = auto_scaling_group_provider
        __props__["name"] = name
        __props__["tags"] = tags
        return CapacityProvider(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

