// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES domain MAIL FROM resource.
    /// 
    /// &gt; **NOTE:** For the MAIL FROM domain to be fully usable, this resource should be paired with the `aws.ses.DomainIdentity` resource. To validate the MAIL FROM domain, a DNS MX record is required. To pass SPF checks, a DNS TXT record may also be required. See the [Amazon SES MAIL FROM documentation](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-set.html) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Example SES Domain Identity
    ///         var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new Aws.Ses.DomainIdentityArgs
    ///         {
    ///             Domain = "example.com",
    ///         });
    ///         var exampleMailFrom = new Aws.Ses.MailFrom("exampleMailFrom", new Aws.Ses.MailFromArgs
    ///         {
    ///             Domain = exampleDomainIdentity.Domain,
    ///             MailFromDomain = exampleDomainIdentity.Domain.Apply(domain =&gt; $"bounce.{domain}"),
    ///         });
    ///         // Example Route53 MX record
    ///         var exampleSesDomainMailFromMx = new Aws.Route53.Record("exampleSesDomainMailFromMx", new Aws.Route53.RecordArgs
    ///         {
    ///             Name = exampleMailFrom.MailFromDomain,
    ///             Records = 
    ///             {
    ///                 "10 feedback-smtp.us-east-1.amazonses.com",
    ///             },
    ///             Ttl = 600,
    ///             Type = "MX",
    ///             ZoneId = aws_route53_zone.Example.Id,
    ///         });
    ///         // Example Route53 TXT record for SPF
    ///         var exampleSesDomainMailFromTxt = new Aws.Route53.Record("exampleSesDomainMailFromTxt", new Aws.Route53.RecordArgs
    ///         {
    ///             Name = exampleMailFrom.MailFromDomain,
    ///             Records = 
    ///             {
    ///                 "v=spf1 include:amazonses.com -all",
    ///             },
    ///             Ttl = 600,
    ///             Type = "TXT",
    ///             ZoneId = aws_route53_zone.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class MailFrom : Pulumi.CustomResource
    {
        /// <summary>
        /// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
        /// </summary>
        [Output("behaviorOnMxFailure")]
        public Output<string?> BehaviorOnMxFailure { get; private set; } = null!;

        /// <summary>
        /// Verified domain name to generate DKIM tokens for.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
        /// </summary>
        [Output("mailFromDomain")]
        public Output<string> MailFromDomain { get; private set; } = null!;


        /// <summary>
        /// Create a MailFrom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MailFrom(string name, MailFromArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/mailFrom:MailFrom", name, args ?? new MailFromArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MailFrom(string name, Input<string> id, MailFromState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/mailFrom:MailFrom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MailFrom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MailFrom Get(string name, Input<string> id, MailFromState? state = null, CustomResourceOptions? options = null)
        {
            return new MailFrom(name, id, state, options);
        }
    }

    public sealed class MailFromArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
        /// </summary>
        [Input("behaviorOnMxFailure")]
        public Input<string>? BehaviorOnMxFailure { get; set; }

        /// <summary>
        /// Verified domain name to generate DKIM tokens for.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
        /// </summary>
        [Input("mailFromDomain", required: true)]
        public Input<string> MailFromDomain { get; set; } = null!;

        public MailFromArgs()
        {
        }
    }

    public sealed class MailFromState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
        /// </summary>
        [Input("behaviorOnMxFailure")]
        public Input<string>? BehaviorOnMxFailure { get; set; }

        /// <summary>
        /// Verified domain name to generate DKIM tokens for.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
        /// </summary>
        [Input("mailFromDomain")]
        public Input<string>? MailFromDomain { get; set; }

        public MailFromState()
        {
        }
    }
}
