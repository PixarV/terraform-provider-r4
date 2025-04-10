---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "aws_internet_gateway"
description: |-
  Provides information about an internet gateway.
---

[describe-igws]: https://docs.k2.cloud/en/api/ec2/vpcs/DescribeInternetGateways.html

# Data Source: aws_internet_gateway

Provides information about an internet gateway.

## Example Usage

```terraform
data "aws_internet_gateway" "selected" {
  filter {
    name   = "attachment.vpc-id"
    values = ["vpc-12345678"]
  }
}
```

## Argument Reference

The arguments of this data source act as filters for querying the available
internet gateway in the current region. The given filters must match exactly one
internet gateway whose data will be exported as attributes.

* `internet_gateway_id` - (Optional) The ID of the internet gateway.
* `tags` - (Optional) Map of tags, each pair of which must exactly match
  a pair on the desired internet gateway.
* `filter` - (Optional) One or more name/value pairs to use as filters.
    * _Valid values_: See supported names and values in [EC2 API documentation][describe-igws]

## Attribute Reference

All arguments except `filter` block are also exported as
result attributes. This data source will complete the data by populating
any fields that are not included in the configuration with the data for
the selected internet gateway.

* `arn` - The Amazon Resource Name (ARN) of the internet gateway.
* `attachments` - List of VPC attachments to the internet gateway. The structure of this block is [described below](#attachments).
* `id` - The ID of the internet gateway.
* `owner_id` - The ID of the project that owns the internet gateway.

### attachments

* `state` - The current state of the attachment between the internet gateway and the VPC.
* `vpc_id` - The ID of the attached VPC.
