// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Aws
{
    /// <summary>
    /// A Region represents any valid Amazon region that may be targeted with deployments.
    /// </summary>
    [EnumType]
    public readonly struct Region : IEquatable<Region>
    {
        private readonly string _value;

        private Region(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Region AFSouth1 { get; } = new Region("af-south-1");
        public static Region APEast1 { get; } = new Region("ap-east-1");
        public static Region APNortheast1 { get; } = new Region("ap-northeast-1");
        public static Region APNortheast2 { get; } = new Region("ap-northeast-2");
        public static Region APSouth1 { get; } = new Region("ap-south-1");
        public static Region APSoutheast2 { get; } = new Region("ap-southeast-2");
        public static Region APSoutheast1 { get; } = new Region("ap-southeast-1");
        public static Region CACentral { get; } = new Region("ca-central-1");
        public static Region CNNorth1 { get; } = new Region("cn-north-1");
        public static Region CNNorthwest1 { get; } = new Region("cn-northwest-1");
        public static Region EUCentral1 { get; } = new Region("eu-central-1");
        public static Region EUNorth1 { get; } = new Region("eu-north-1");
        public static Region EUWest1 { get; } = new Region("eu-west-1");
        public static Region EUWest2 { get; } = new Region("eu-west-2");
        public static Region EUWest3 { get; } = new Region("eu-west-3");
        public static Region EUSouth1 { get; } = new Region("eu-south-1");
        public static Region MESouth1 { get; } = new Region("me-south-1");
        public static Region SAEast1 { get; } = new Region("sa-east-1");
        public static Region USGovEast1 { get; } = new Region("us-gov-east-1");
        public static Region USGovWest1 { get; } = new Region("us-gov-west-1");
        public static Region USEast1 { get; } = new Region("us-east-1");
        public static Region USEast2 { get; } = new Region("us-east-2");
        public static Region USWest1 { get; } = new Region("us-west-1");
        public static Region USWest2 { get; } = new Region("us-west-2");

        public static bool operator ==(Region left, Region right) => left.Equals(right);
        public static bool operator !=(Region left, Region right) => !left.Equals(right);

        public static explicit operator string(Region value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Region other && Equals(other);
        public bool Equals(Region other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
