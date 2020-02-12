# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UserPool(pulumi.CustomResource):
    admin_create_user_config: pulumi.Output[dict]
    """
    The configuration for AdminCreateUser requests.

      * `allowAdminCreateUserOnly` (`bool`) - Set to True if only the administrator is allowed to create user profiles. Set to False if users can sign themselves up via an app.
      * `inviteMessageTemplate` (`dict`) - The invite message template structure.
        * `emailMessage` (`str`) - The email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
        * `emailSubject` (`str`) - The subject line for the email message template. Conflicts with `email_verification_subject` argument.
        * `smsMessage` (`str`) - The SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.

      * `unusedAccountValidityDays` (`float`) - **DEPRECATED** Use password_policy.temporary_password_validity_days instead - The user account expiration limit, in days, after which the account is no longer usable.
    """
    alias_attributes: pulumi.Output[list]
    """
    Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `username_attributes`.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the user pool.
    """
    auto_verified_attributes: pulumi.Output[list]
    """
    The attributes to be auto-verified. Possible values: email, phone_number.
    """
    creation_date: pulumi.Output[str]
    """
    The date the user pool was created.
    """
    device_configuration: pulumi.Output[dict]
    """
    The configuration for the user pool's device tracking.

      * `challengeRequiredOnNewDevice` (`bool`) - Indicates whether a challenge is required on a new device. Only applicable to a new device.
      * `deviceOnlyRememberedOnUserPrompt` (`bool`) - If true, a device is only remembered on user prompt.
    """
    email_configuration: pulumi.Output[dict]
    """
    The Email Configuration.

      * `emailSendingAccount` (`str`) - Instruct Cognito to either use its built-in functional or Amazon SES to send out emails.
      * `replyToEmailAddress` (`str`) - The REPLY-TO email address.
      * `sourceArn` (`str`) - The ARN of the email source.
    """
    email_verification_message: pulumi.Output[str]
    """
    A string representing the email verification message. Conflicts with `verification_message_template` configuration block `email_message` argument.
    """
    email_verification_subject: pulumi.Output[str]
    """
    A string representing the email verification subject. Conflicts with `verification_message_template` configuration block `email_subject` argument.
    """
    endpoint: pulumi.Output[str]
    """
    The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
    """
    lambda_config: pulumi.Output[dict]
    """
    A container for the AWS Lambda triggers associated with the user pool.

      * `createAuthChallenge` (`str`) - The ARN of the lambda creating an authentication challenge.
      * `customMessage` (`str`) - A custom Message AWS Lambda trigger.
      * `defineAuthChallenge` (`str`) - Defines the authentication challenge.
      * `postAuthentication` (`str`) - A post-authentication AWS Lambda trigger.
      * `postConfirmation` (`str`) - A post-confirmation AWS Lambda trigger.
      * `preAuthentication` (`str`) - A pre-authentication AWS Lambda trigger.
      * `preSignUp` (`str`) - A pre-registration AWS Lambda trigger.
      * `preTokenGeneration` (`str`) - Allow to customize identity token claims before token generation.
      * `userMigration` (`str`) - The user migration Lambda config type.
      * `verifyAuthChallengeResponse` (`str`) - Verifies the authentication challenge response.
    """
    last_modified_date: pulumi.Output[str]
    """
    The date the user pool was last modified.
    """
    mfa_configuration: pulumi.Output[str]
    """
    Set to enable multi-factor authentication. Must be one of the following values (ON, OFF, OPTIONAL)
    """
    name: pulumi.Output[str]
    """
    The name of the attribute.
    """
    password_policy: pulumi.Output[dict]
    """
    A container for information about the user pool password policy.

      * `minimumLength` (`float`) - The minimum length of the password policy that you have set.
      * `requireLowercase` (`bool`) - Whether you have required users to use at least one lowercase letter in their password.
      * `require_numbers` (`bool`) - Whether you have required users to use at least one number in their password.
      * `require_symbols` (`bool`) - Whether you have required users to use at least one symbol in their password.
      * `requireUppercase` (`bool`) - Whether you have required users to use at least one uppercase letter in their password.
      * `temporaryPasswordValidityDays` (`float`) - In the password policy you have set, refers to the number of days a temporary password is valid. If the user does not sign-in during this time, their password will need to be reset by an administrator.
    """
    schemas: pulumi.Output[list]
    """
    A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.

      * `attributeDataType` (`str`) - The attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
      * `developerOnlyAttribute` (`bool`) - Specifies whether the attribute type is developer only.
      * `mutable` (`bool`) - Specifies whether the attribute can be changed once it has been created.
      * `name` (`str`) - The name of the attribute.
      * `numberAttributeConstraints` (`dict`) - Specifies the constraints for an attribute of the number type.
        * `maxValue` (`str`) - The maximum value of an attribute that is of the number data type.
        * `minValue` (`str`) - The minimum value of an attribute that is of the number data type.

      * `required` (`bool`) - Specifies whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
      * `stringAttributeConstraints` (`dict`) - -Specifies the constraints for an attribute of the string type.
        * `maxLength` (`str`) - The maximum length of an attribute value of the string type.
        * `minLength` (`str`) - The minimum length of an attribute value of the string type.
    """
    sms_authentication_message: pulumi.Output[str]
    """
    A string representing the SMS authentication message.
    """
    sms_configuration: pulumi.Output[dict]
    """
    The SMS Configuration.

      * `externalId` (`str`) - The external ID used in IAM role trust relationships. For more information about using external IDs, see [How to Use an External ID When Granting Access to Your AWS Resources to a Third Party](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html).
      * `snsCallerArn` (`str`) - The ARN of the Amazon SNS caller. This is usually the IAM role that you've given Cognito permission to assume.
    """
    sms_verification_message: pulumi.Output[str]
    """
    A string representing the SMS verification message. Conflicts with `verification_message_template` configuration block `sms_message` argument.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the User Pool.
    """
    user_pool_add_ons: pulumi.Output[dict]
    """
    Configuration block for user pool add-ons to enable user pool advanced security mode features.

      * `advancedSecurityMode` (`str`) - The mode for advanced security, must be one of `OFF`, `AUDIT` or `ENFORCED`.
    """
    username_attributes: pulumi.Output[list]
    """
    Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `alias_attributes`.
    """
    verification_message_template: pulumi.Output[dict]
    """
    The verification message templates configuration.

      * `defaultEmailOption` (`str`) - The default email option. Must be either `CONFIRM_WITH_CODE` or `CONFIRM_WITH_LINK`. Defaults to `CONFIRM_WITH_CODE`.
      * `emailMessage` (`str`) - The email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
      * `emailMessageByLink` (`str`) - The email message template for sending a confirmation link to the user, it must contain the `{##Click Here##}` placeholder.
      * `emailSubject` (`str`) - The subject line for the email message template. Conflicts with `email_verification_subject` argument.
      * `emailSubjectByLink` (`str`) - The subject line for the email message template for sending a confirmation link to the user.
      * `smsMessage` (`str`) - The SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.
    """
    def __init__(__self__, resource_name, opts=None, admin_create_user_config=None, alias_attributes=None, auto_verified_attributes=None, device_configuration=None, email_configuration=None, email_verification_message=None, email_verification_subject=None, lambda_config=None, mfa_configuration=None, name=None, password_policy=None, schemas=None, sms_authentication_message=None, sms_configuration=None, sms_verification_message=None, tags=None, user_pool_add_ons=None, username_attributes=None, verification_message_template=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Cognito User Pool resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] admin_create_user_config: The configuration for AdminCreateUser requests.
        :param pulumi.Input[list] alias_attributes: Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `username_attributes`.
        :param pulumi.Input[list] auto_verified_attributes: The attributes to be auto-verified. Possible values: email, phone_number.
        :param pulumi.Input[dict] device_configuration: The configuration for the user pool's device tracking.
        :param pulumi.Input[dict] email_configuration: The Email Configuration.
        :param pulumi.Input[str] email_verification_message: A string representing the email verification message. Conflicts with `verification_message_template` configuration block `email_message` argument.
        :param pulumi.Input[str] email_verification_subject: A string representing the email verification subject. Conflicts with `verification_message_template` configuration block `email_subject` argument.
        :param pulumi.Input[dict] lambda_config: A container for the AWS Lambda triggers associated with the user pool.
        :param pulumi.Input[str] mfa_configuration: Set to enable multi-factor authentication. Must be one of the following values (ON, OFF, OPTIONAL)
        :param pulumi.Input[str] name: The name of the attribute.
        :param pulumi.Input[dict] password_policy: A container for information about the user pool password policy.
        :param pulumi.Input[list] schemas: A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
        :param pulumi.Input[str] sms_authentication_message: A string representing the SMS authentication message.
        :param pulumi.Input[dict] sms_configuration: The SMS Configuration.
        :param pulumi.Input[str] sms_verification_message: A string representing the SMS verification message. Conflicts with `verification_message_template` configuration block `sms_message` argument.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the User Pool.
        :param pulumi.Input[dict] user_pool_add_ons: Configuration block for user pool add-ons to enable user pool advanced security mode features.
        :param pulumi.Input[list] username_attributes: Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `alias_attributes`.
        :param pulumi.Input[dict] verification_message_template: The verification message templates configuration.

        The **admin_create_user_config** object supports the following:

          * `allowAdminCreateUserOnly` (`pulumi.Input[bool]`) - Set to True if only the administrator is allowed to create user profiles. Set to False if users can sign themselves up via an app.
          * `inviteMessageTemplate` (`pulumi.Input[dict]`) - The invite message template structure.
            * `emailMessage` (`pulumi.Input[str]`) - The email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
            * `emailSubject` (`pulumi.Input[str]`) - The subject line for the email message template. Conflicts with `email_verification_subject` argument.
            * `smsMessage` (`pulumi.Input[str]`) - The SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.

          * `unusedAccountValidityDays` (`pulumi.Input[float]`) - **DEPRECATED** Use password_policy.temporary_password_validity_days instead - The user account expiration limit, in days, after which the account is no longer usable.

        The **device_configuration** object supports the following:

          * `challengeRequiredOnNewDevice` (`pulumi.Input[bool]`) - Indicates whether a challenge is required on a new device. Only applicable to a new device.
          * `deviceOnlyRememberedOnUserPrompt` (`pulumi.Input[bool]`) - If true, a device is only remembered on user prompt.

        The **email_configuration** object supports the following:

          * `emailSendingAccount` (`pulumi.Input[str]`) - Instruct Cognito to either use its built-in functional or Amazon SES to send out emails.
          * `replyToEmailAddress` (`pulumi.Input[str]`) - The REPLY-TO email address.
          * `sourceArn` (`pulumi.Input[str]`) - The ARN of the email source.

        The **lambda_config** object supports the following:

          * `createAuthChallenge` (`pulumi.Input[str]`) - The ARN of the lambda creating an authentication challenge.
          * `customMessage` (`pulumi.Input[str]`) - A custom Message AWS Lambda trigger.
          * `defineAuthChallenge` (`pulumi.Input[str]`) - Defines the authentication challenge.
          * `postAuthentication` (`pulumi.Input[str]`) - A post-authentication AWS Lambda trigger.
          * `postConfirmation` (`pulumi.Input[str]`) - A post-confirmation AWS Lambda trigger.
          * `preAuthentication` (`pulumi.Input[str]`) - A pre-authentication AWS Lambda trigger.
          * `preSignUp` (`pulumi.Input[str]`) - A pre-registration AWS Lambda trigger.
          * `preTokenGeneration` (`pulumi.Input[str]`) - Allow to customize identity token claims before token generation.
          * `userMigration` (`pulumi.Input[str]`) - The user migration Lambda config type.
          * `verifyAuthChallengeResponse` (`pulumi.Input[str]`) - Verifies the authentication challenge response.

        The **password_policy** object supports the following:

          * `minimumLength` (`pulumi.Input[float]`) - The minimum length of the password policy that you have set.
          * `requireLowercase` (`pulumi.Input[bool]`) - Whether you have required users to use at least one lowercase letter in their password.
          * `require_numbers` (`pulumi.Input[bool]`) - Whether you have required users to use at least one number in their password.
          * `require_symbols` (`pulumi.Input[bool]`) - Whether you have required users to use at least one symbol in their password.
          * `requireUppercase` (`pulumi.Input[bool]`) - Whether you have required users to use at least one uppercase letter in their password.
          * `temporaryPasswordValidityDays` (`pulumi.Input[float]`) - In the password policy you have set, refers to the number of days a temporary password is valid. If the user does not sign-in during this time, their password will need to be reset by an administrator.

        The **schemas** object supports the following:

          * `attributeDataType` (`pulumi.Input[str]`) - The attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
          * `developerOnlyAttribute` (`pulumi.Input[bool]`) - Specifies whether the attribute type is developer only.
          * `mutable` (`pulumi.Input[bool]`) - Specifies whether the attribute can be changed once it has been created.
          * `name` (`pulumi.Input[str]`) - The name of the attribute.
          * `numberAttributeConstraints` (`pulumi.Input[dict]`) - Specifies the constraints for an attribute of the number type.
            * `maxValue` (`pulumi.Input[str]`) - The maximum value of an attribute that is of the number data type.
            * `minValue` (`pulumi.Input[str]`) - The minimum value of an attribute that is of the number data type.

          * `required` (`pulumi.Input[bool]`) - Specifies whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
          * `stringAttributeConstraints` (`pulumi.Input[dict]`) - -Specifies the constraints for an attribute of the string type.
            * `maxLength` (`pulumi.Input[str]`) - The maximum length of an attribute value of the string type.
            * `minLength` (`pulumi.Input[str]`) - The minimum length of an attribute value of the string type.

        The **sms_configuration** object supports the following:

          * `externalId` (`pulumi.Input[str]`) - The external ID used in IAM role trust relationships. For more information about using external IDs, see [How to Use an External ID When Granting Access to Your AWS Resources to a Third Party](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html).
          * `snsCallerArn` (`pulumi.Input[str]`) - The ARN of the Amazon SNS caller. This is usually the IAM role that you've given Cognito permission to assume.

        The **user_pool_add_ons** object supports the following:

          * `advancedSecurityMode` (`pulumi.Input[str]`) - The mode for advanced security, must be one of `OFF`, `AUDIT` or `ENFORCED`.

        The **verification_message_template** object supports the following:

          * `defaultEmailOption` (`pulumi.Input[str]`) - The default email option. Must be either `CONFIRM_WITH_CODE` or `CONFIRM_WITH_LINK`. Defaults to `CONFIRM_WITH_CODE`.
          * `emailMessage` (`pulumi.Input[str]`) - The email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
          * `emailMessageByLink` (`pulumi.Input[str]`) - The email message template for sending a confirmation link to the user, it must contain the `{##Click Here##}` placeholder.
          * `emailSubject` (`pulumi.Input[str]`) - The subject line for the email message template. Conflicts with `email_verification_subject` argument.
          * `emailSubjectByLink` (`pulumi.Input[str]`) - The subject line for the email message template for sending a confirmation link to the user.
          * `smsMessage` (`pulumi.Input[str]`) - The SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.
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

            __props__['admin_create_user_config'] = admin_create_user_config
            __props__['alias_attributes'] = alias_attributes
            __props__['auto_verified_attributes'] = auto_verified_attributes
            __props__['device_configuration'] = device_configuration
            __props__['email_configuration'] = email_configuration
            __props__['email_verification_message'] = email_verification_message
            __props__['email_verification_subject'] = email_verification_subject
            __props__['lambda_config'] = lambda_config
            __props__['mfa_configuration'] = mfa_configuration
            __props__['name'] = name
            __props__['password_policy'] = password_policy
            __props__['schemas'] = schemas
            __props__['sms_authentication_message'] = sms_authentication_message
            __props__['sms_configuration'] = sms_configuration
            __props__['sms_verification_message'] = sms_verification_message
            __props__['tags'] = tags
            __props__['user_pool_add_ons'] = user_pool_add_ons
            __props__['username_attributes'] = username_attributes
            __props__['verification_message_template'] = verification_message_template
            __props__['arn'] = None
            __props__['creation_date'] = None
            __props__['endpoint'] = None
            __props__['last_modified_date'] = None
        super(UserPool, __self__).__init__(
            'aws:cognito/userPool:UserPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admin_create_user_config=None, alias_attributes=None, arn=None, auto_verified_attributes=None, creation_date=None, device_configuration=None, email_configuration=None, email_verification_message=None, email_verification_subject=None, endpoint=None, lambda_config=None, last_modified_date=None, mfa_configuration=None, name=None, password_policy=None, schemas=None, sms_authentication_message=None, sms_configuration=None, sms_verification_message=None, tags=None, user_pool_add_ons=None, username_attributes=None, verification_message_template=None):
        """
        Get an existing UserPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] admin_create_user_config: The configuration for AdminCreateUser requests.
        :param pulumi.Input[list] alias_attributes: Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `username_attributes`.
        :param pulumi.Input[str] arn: The ARN of the user pool.
        :param pulumi.Input[list] auto_verified_attributes: The attributes to be auto-verified. Possible values: email, phone_number.
        :param pulumi.Input[str] creation_date: The date the user pool was created.
        :param pulumi.Input[dict] device_configuration: The configuration for the user pool's device tracking.
        :param pulumi.Input[dict] email_configuration: The Email Configuration.
        :param pulumi.Input[str] email_verification_message: A string representing the email verification message. Conflicts with `verification_message_template` configuration block `email_message` argument.
        :param pulumi.Input[str] email_verification_subject: A string representing the email verification subject. Conflicts with `verification_message_template` configuration block `email_subject` argument.
        :param pulumi.Input[str] endpoint: The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
        :param pulumi.Input[dict] lambda_config: A container for the AWS Lambda triggers associated with the user pool.
        :param pulumi.Input[str] last_modified_date: The date the user pool was last modified.
        :param pulumi.Input[str] mfa_configuration: Set to enable multi-factor authentication. Must be one of the following values (ON, OFF, OPTIONAL)
        :param pulumi.Input[str] name: The name of the attribute.
        :param pulumi.Input[dict] password_policy: A container for information about the user pool password policy.
        :param pulumi.Input[list] schemas: A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
        :param pulumi.Input[str] sms_authentication_message: A string representing the SMS authentication message.
        :param pulumi.Input[dict] sms_configuration: The SMS Configuration.
        :param pulumi.Input[str] sms_verification_message: A string representing the SMS verification message. Conflicts with `verification_message_template` configuration block `sms_message` argument.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the User Pool.
        :param pulumi.Input[dict] user_pool_add_ons: Configuration block for user pool add-ons to enable user pool advanced security mode features.
        :param pulumi.Input[list] username_attributes: Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `alias_attributes`.
        :param pulumi.Input[dict] verification_message_template: The verification message templates configuration.

        The **admin_create_user_config** object supports the following:

          * `allowAdminCreateUserOnly` (`pulumi.Input[bool]`) - Set to True if only the administrator is allowed to create user profiles. Set to False if users can sign themselves up via an app.
          * `inviteMessageTemplate` (`pulumi.Input[dict]`) - The invite message template structure.
            * `emailMessage` (`pulumi.Input[str]`) - The email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
            * `emailSubject` (`pulumi.Input[str]`) - The subject line for the email message template. Conflicts with `email_verification_subject` argument.
            * `smsMessage` (`pulumi.Input[str]`) - The SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.

          * `unusedAccountValidityDays` (`pulumi.Input[float]`) - **DEPRECATED** Use password_policy.temporary_password_validity_days instead - The user account expiration limit, in days, after which the account is no longer usable.

        The **device_configuration** object supports the following:

          * `challengeRequiredOnNewDevice` (`pulumi.Input[bool]`) - Indicates whether a challenge is required on a new device. Only applicable to a new device.
          * `deviceOnlyRememberedOnUserPrompt` (`pulumi.Input[bool]`) - If true, a device is only remembered on user prompt.

        The **email_configuration** object supports the following:

          * `emailSendingAccount` (`pulumi.Input[str]`) - Instruct Cognito to either use its built-in functional or Amazon SES to send out emails.
          * `replyToEmailAddress` (`pulumi.Input[str]`) - The REPLY-TO email address.
          * `sourceArn` (`pulumi.Input[str]`) - The ARN of the email source.

        The **lambda_config** object supports the following:

          * `createAuthChallenge` (`pulumi.Input[str]`) - The ARN of the lambda creating an authentication challenge.
          * `customMessage` (`pulumi.Input[str]`) - A custom Message AWS Lambda trigger.
          * `defineAuthChallenge` (`pulumi.Input[str]`) - Defines the authentication challenge.
          * `postAuthentication` (`pulumi.Input[str]`) - A post-authentication AWS Lambda trigger.
          * `postConfirmation` (`pulumi.Input[str]`) - A post-confirmation AWS Lambda trigger.
          * `preAuthentication` (`pulumi.Input[str]`) - A pre-authentication AWS Lambda trigger.
          * `preSignUp` (`pulumi.Input[str]`) - A pre-registration AWS Lambda trigger.
          * `preTokenGeneration` (`pulumi.Input[str]`) - Allow to customize identity token claims before token generation.
          * `userMigration` (`pulumi.Input[str]`) - The user migration Lambda config type.
          * `verifyAuthChallengeResponse` (`pulumi.Input[str]`) - Verifies the authentication challenge response.

        The **password_policy** object supports the following:

          * `minimumLength` (`pulumi.Input[float]`) - The minimum length of the password policy that you have set.
          * `requireLowercase` (`pulumi.Input[bool]`) - Whether you have required users to use at least one lowercase letter in their password.
          * `require_numbers` (`pulumi.Input[bool]`) - Whether you have required users to use at least one number in their password.
          * `require_symbols` (`pulumi.Input[bool]`) - Whether you have required users to use at least one symbol in their password.
          * `requireUppercase` (`pulumi.Input[bool]`) - Whether you have required users to use at least one uppercase letter in their password.
          * `temporaryPasswordValidityDays` (`pulumi.Input[float]`) - In the password policy you have set, refers to the number of days a temporary password is valid. If the user does not sign-in during this time, their password will need to be reset by an administrator.

        The **schemas** object supports the following:

          * `attributeDataType` (`pulumi.Input[str]`) - The attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
          * `developerOnlyAttribute` (`pulumi.Input[bool]`) - Specifies whether the attribute type is developer only.
          * `mutable` (`pulumi.Input[bool]`) - Specifies whether the attribute can be changed once it has been created.
          * `name` (`pulumi.Input[str]`) - The name of the attribute.
          * `numberAttributeConstraints` (`pulumi.Input[dict]`) - Specifies the constraints for an attribute of the number type.
            * `maxValue` (`pulumi.Input[str]`) - The maximum value of an attribute that is of the number data type.
            * `minValue` (`pulumi.Input[str]`) - The minimum value of an attribute that is of the number data type.

          * `required` (`pulumi.Input[bool]`) - Specifies whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
          * `stringAttributeConstraints` (`pulumi.Input[dict]`) - -Specifies the constraints for an attribute of the string type.
            * `maxLength` (`pulumi.Input[str]`) - The maximum length of an attribute value of the string type.
            * `minLength` (`pulumi.Input[str]`) - The minimum length of an attribute value of the string type.

        The **sms_configuration** object supports the following:

          * `externalId` (`pulumi.Input[str]`) - The external ID used in IAM role trust relationships. For more information about using external IDs, see [How to Use an External ID When Granting Access to Your AWS Resources to a Third Party](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html).
          * `snsCallerArn` (`pulumi.Input[str]`) - The ARN of the Amazon SNS caller. This is usually the IAM role that you've given Cognito permission to assume.

        The **user_pool_add_ons** object supports the following:

          * `advancedSecurityMode` (`pulumi.Input[str]`) - The mode for advanced security, must be one of `OFF`, `AUDIT` or `ENFORCED`.

        The **verification_message_template** object supports the following:

          * `defaultEmailOption` (`pulumi.Input[str]`) - The default email option. Must be either `CONFIRM_WITH_CODE` or `CONFIRM_WITH_LINK`. Defaults to `CONFIRM_WITH_CODE`.
          * `emailMessage` (`pulumi.Input[str]`) - The email message template. Must contain the `{####}` placeholder. Conflicts with `email_verification_message` argument.
          * `emailMessageByLink` (`pulumi.Input[str]`) - The email message template for sending a confirmation link to the user, it must contain the `{##Click Here##}` placeholder.
          * `emailSubject` (`pulumi.Input[str]`) - The subject line for the email message template. Conflicts with `email_verification_subject` argument.
          * `emailSubjectByLink` (`pulumi.Input[str]`) - The subject line for the email message template for sending a confirmation link to the user.
          * `smsMessage` (`pulumi.Input[str]`) - The SMS message template. Must contain the `{####}` placeholder. Conflicts with `sms_verification_message` argument.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_create_user_config"] = admin_create_user_config
        __props__["alias_attributes"] = alias_attributes
        __props__["arn"] = arn
        __props__["auto_verified_attributes"] = auto_verified_attributes
        __props__["creation_date"] = creation_date
        __props__["device_configuration"] = device_configuration
        __props__["email_configuration"] = email_configuration
        __props__["email_verification_message"] = email_verification_message
        __props__["email_verification_subject"] = email_verification_subject
        __props__["endpoint"] = endpoint
        __props__["lambda_config"] = lambda_config
        __props__["last_modified_date"] = last_modified_date
        __props__["mfa_configuration"] = mfa_configuration
        __props__["name"] = name
        __props__["password_policy"] = password_policy
        __props__["schemas"] = schemas
        __props__["sms_authentication_message"] = sms_authentication_message
        __props__["sms_configuration"] = sms_configuration
        __props__["sms_verification_message"] = sms_verification_message
        __props__["tags"] = tags
        __props__["user_pool_add_ons"] = user_pool_add_ons
        __props__["username_attributes"] = username_attributes
        __props__["verification_message_template"] = verification_message_template
        return UserPool(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

