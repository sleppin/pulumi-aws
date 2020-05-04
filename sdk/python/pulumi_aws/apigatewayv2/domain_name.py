# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DomainName(pulumi.CustomResource):
    api_mapping_selection_expression: pulumi.Output[str]
    """
    The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the domain name.
    """
    domain_name: pulumi.Output[str]
    """
    The domain name.
    """
    domain_name_configuration: pulumi.Output[dict]
    """
    The domain name configuration.

      * `certificate_arn` (`str`) - The ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source.
        Use the [`acm.Certificate`](https://www.terraform.io/docs/providers/aws/r/acm_certificate.html) resource to configure an ACM certificate.
      * `endpoint_type` (`str`) - The endpoint type. Valid values: `REGIONAL`.
      * `hosted_zone_id` (`str`) - The Amazon Route 53 Hosted Zone ID of the endpoint.
      * `security_policy` (`str`) - The Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
      * `targetDomainName` (`str`) - The target domain name.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the domain name.
    """
    def __init__(__self__, resource_name, opts=None, domain_name=None, domain_name_configuration=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Amazon API Gateway Version 2 domain name.
        More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).

        > **Note:** This resource establishes ownership of and the TLS settings for
        a particular domain name. An API stage can be associated with the domain name using the `apigatewayv2.ApiMapping` resource.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: The domain name.
        :param pulumi.Input[dict] domain_name_configuration: The domain name configuration.
        :param pulumi.Input[dict] tags: A map of tags to assign to the domain name.

        The **domain_name_configuration** object supports the following:

          * `certificate_arn` (`pulumi.Input[str]`) - The ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source.
            Use the [`acm.Certificate`](https://www.terraform.io/docs/providers/aws/r/acm_certificate.html) resource to configure an ACM certificate.
          * `endpoint_type` (`pulumi.Input[str]`) - The endpoint type. Valid values: `REGIONAL`.
          * `hosted_zone_id` (`pulumi.Input[str]`) - The Amazon Route 53 Hosted Zone ID of the endpoint.
          * `security_policy` (`pulumi.Input[str]`) - The Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
          * `targetDomainName` (`pulumi.Input[str]`) - The target domain name.
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

            if domain_name is None:
                raise TypeError("Missing required property 'domain_name'")
            __props__['domain_name'] = domain_name
            if domain_name_configuration is None:
                raise TypeError("Missing required property 'domain_name_configuration'")
            __props__['domain_name_configuration'] = domain_name_configuration
            __props__['tags'] = tags
            __props__['api_mapping_selection_expression'] = None
            __props__['arn'] = None
        super(DomainName, __self__).__init__(
            'aws:apigatewayv2/domainName:DomainName',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_mapping_selection_expression=None, arn=None, domain_name=None, domain_name_configuration=None, tags=None):
        """
        Get an existing DomainName resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_mapping_selection_expression: The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
        :param pulumi.Input[str] arn: The ARN of the domain name.
        :param pulumi.Input[str] domain_name: The domain name.
        :param pulumi.Input[dict] domain_name_configuration: The domain name configuration.
        :param pulumi.Input[dict] tags: A map of tags to assign to the domain name.

        The **domain_name_configuration** object supports the following:

          * `certificate_arn` (`pulumi.Input[str]`) - The ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source.
            Use the [`acm.Certificate`](https://www.terraform.io/docs/providers/aws/r/acm_certificate.html) resource to configure an ACM certificate.
          * `endpoint_type` (`pulumi.Input[str]`) - The endpoint type. Valid values: `REGIONAL`.
          * `hosted_zone_id` (`pulumi.Input[str]`) - The Amazon Route 53 Hosted Zone ID of the endpoint.
          * `security_policy` (`pulumi.Input[str]`) - The Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
          * `targetDomainName` (`pulumi.Input[str]`) - The target domain name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_mapping_selection_expression"] = api_mapping_selection_expression
        __props__["arn"] = arn
        __props__["domain_name"] = domain_name
        __props__["domain_name_configuration"] = domain_name_configuration
        __props__["tags"] = tags
        return DomainName(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
