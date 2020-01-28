# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UserPoolClient(pulumi.CustomResource):
    allowed_oauth_flows: pulumi.Output[list]
    """
    List of allowed OAuth flows (code, implicit, client_credentials).
    """
    allowed_oauth_flows_user_pool_client: pulumi.Output[bool]
    """
    Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
    """
    allowed_oauth_scopes: pulumi.Output[list]
    """
    List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
    """
    callback_urls: pulumi.Output[list]
    """
    List of allowed callback URLs for the identity providers.
    """
    client_secret: pulumi.Output[str]
    """
    The client secret of the user pool client.
    """
    default_redirect_uri: pulumi.Output[str]
    """
    The default redirect URI. Must be in the list of callback URLs.
    """
    explicit_auth_flows: pulumi.Output[list]
    """
    List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
    """
    generate_secret: pulumi.Output[bool]
    """
    Should an application secret be generated.
    """
    logout_urls: pulumi.Output[list]
    """
    List of allowed logout URLs for the identity providers.
    """
    name: pulumi.Output[str]
    """
    The name of the application client.
    """
    read_attributes: pulumi.Output[list]
    """
    List of user pool attributes the application client can read from.
    """
    refresh_token_validity: pulumi.Output[float]
    """
    The time limit in days refresh tokens are valid for.
    """
    supported_identity_providers: pulumi.Output[list]
    """
    List of provider names for the identity providers that are supported on this client.
    """
    user_pool_id: pulumi.Output[str]
    """
    The user pool the client belongs to.
    """
    write_attributes: pulumi.Output[list]
    """
    List of user pool attributes the application client can write to.
    """
    def __init__(__self__, resource_name, opts=None, allowed_oauth_flows=None, allowed_oauth_flows_user_pool_client=None, allowed_oauth_scopes=None, callback_urls=None, default_redirect_uri=None, explicit_auth_flows=None, generate_secret=None, logout_urls=None, name=None, read_attributes=None, refresh_token_validity=None, supported_identity_providers=None, user_pool_id=None, write_attributes=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Cognito User Pool Client resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_oauth_flows: List of allowed OAuth flows (code, implicit, client_credentials).
        :param pulumi.Input[bool] allowed_oauth_flows_user_pool_client: Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
        :param pulumi.Input[list] allowed_oauth_scopes: List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
        :param pulumi.Input[list] callback_urls: List of allowed callback URLs for the identity providers.
        :param pulumi.Input[str] default_redirect_uri: The default redirect URI. Must be in the list of callback URLs.
        :param pulumi.Input[list] explicit_auth_flows: List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
        :param pulumi.Input[bool] generate_secret: Should an application secret be generated.
        :param pulumi.Input[list] logout_urls: List of allowed logout URLs for the identity providers.
        :param pulumi.Input[str] name: The name of the application client.
        :param pulumi.Input[list] read_attributes: List of user pool attributes the application client can read from.
        :param pulumi.Input[float] refresh_token_validity: The time limit in days refresh tokens are valid for.
        :param pulumi.Input[list] supported_identity_providers: List of provider names for the identity providers that are supported on this client.
        :param pulumi.Input[str] user_pool_id: The user pool the client belongs to.
        :param pulumi.Input[list] write_attributes: List of user pool attributes the application client can write to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool_client.html.markdown.
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

            __props__['allowed_oauth_flows'] = allowed_oauth_flows
            __props__['allowed_oauth_flows_user_pool_client'] = allowed_oauth_flows_user_pool_client
            __props__['allowed_oauth_scopes'] = allowed_oauth_scopes
            __props__['callback_urls'] = callback_urls
            __props__['default_redirect_uri'] = default_redirect_uri
            __props__['explicit_auth_flows'] = explicit_auth_flows
            __props__['generate_secret'] = generate_secret
            __props__['logout_urls'] = logout_urls
            __props__['name'] = name
            __props__['read_attributes'] = read_attributes
            __props__['refresh_token_validity'] = refresh_token_validity
            __props__['supported_identity_providers'] = supported_identity_providers
            if user_pool_id is None:
                raise TypeError("Missing required property 'user_pool_id'")
            __props__['user_pool_id'] = user_pool_id
            __props__['write_attributes'] = write_attributes
            __props__['client_secret'] = None
        super(UserPoolClient, __self__).__init__(
            'aws:cognito/userPoolClient:UserPoolClient',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allowed_oauth_flows=None, allowed_oauth_flows_user_pool_client=None, allowed_oauth_scopes=None, callback_urls=None, client_secret=None, default_redirect_uri=None, explicit_auth_flows=None, generate_secret=None, logout_urls=None, name=None, read_attributes=None, refresh_token_validity=None, supported_identity_providers=None, user_pool_id=None, write_attributes=None):
        """
        Get an existing UserPoolClient resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_oauth_flows: List of allowed OAuth flows (code, implicit, client_credentials).
        :param pulumi.Input[bool] allowed_oauth_flows_user_pool_client: Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
        :param pulumi.Input[list] allowed_oauth_scopes: List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
        :param pulumi.Input[list] callback_urls: List of allowed callback URLs for the identity providers.
        :param pulumi.Input[str] client_secret: The client secret of the user pool client.
        :param pulumi.Input[str] default_redirect_uri: The default redirect URI. Must be in the list of callback URLs.
        :param pulumi.Input[list] explicit_auth_flows: List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY,  USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
        :param pulumi.Input[bool] generate_secret: Should an application secret be generated.
        :param pulumi.Input[list] logout_urls: List of allowed logout URLs for the identity providers.
        :param pulumi.Input[str] name: The name of the application client.
        :param pulumi.Input[list] read_attributes: List of user pool attributes the application client can read from.
        :param pulumi.Input[float] refresh_token_validity: The time limit in days refresh tokens are valid for.
        :param pulumi.Input[list] supported_identity_providers: List of provider names for the identity providers that are supported on this client.
        :param pulumi.Input[str] user_pool_id: The user pool the client belongs to.
        :param pulumi.Input[list] write_attributes: List of user pool attributes the application client can write to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool_client.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["allowed_oauth_flows"] = allowed_oauth_flows
        __props__["allowed_oauth_flows_user_pool_client"] = allowed_oauth_flows_user_pool_client
        __props__["allowed_oauth_scopes"] = allowed_oauth_scopes
        __props__["callback_urls"] = callback_urls
        __props__["client_secret"] = client_secret
        __props__["default_redirect_uri"] = default_redirect_uri
        __props__["explicit_auth_flows"] = explicit_auth_flows
        __props__["generate_secret"] = generate_secret
        __props__["logout_urls"] = logout_urls
        __props__["name"] = name
        __props__["read_attributes"] = read_attributes
        __props__["refresh_token_validity"] = refresh_token_validity
        __props__["supported_identity_providers"] = supported_identity_providers
        __props__["user_pool_id"] = user_pool_id
        __props__["write_attributes"] = write_attributes
        return UserPoolClient(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

