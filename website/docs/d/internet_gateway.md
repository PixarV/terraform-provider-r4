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
* `tags` - (Optional) A map of tags, each pair of which must exactly match
  a pair on the desired internet gateway.
* `filter` - (Optional) One or more name/value pairs to use as filters.
    * _Valid values_: See supported names and values in [EC2 API documentation][describe-igws]

## Attributes Reference

All arguments except `filter` block are also exported as
result attributes. This data source will complete the data by populating
any fields that are not included in the configuration with the data for
the selected internet gateway.

* `arn` - The Amazon Resource Name (ARN) of the internet gateway.
* `id` - The ID of the internet gateway.
* `owner_id` - The ID of the project that owns the internet gateway.

`attachments` are also exported with the following attributes, when there are relevant:
Each attachment supports the following:

* `state` - The current state of the attachment between the gateway and the VPC. Present only if a VPC is attached
* `vpc_id` - The ID of the attached VPC.
