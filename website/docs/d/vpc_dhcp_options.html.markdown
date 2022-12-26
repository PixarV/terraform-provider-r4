---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_dhcp_options"
description: |-
  Retrieve information about an EC2 DHCP Options configuration
---

# Data Source: aws_vpc_dhcp_options

Retrieve information about an EC2 DHCP Options configuration.

## Example Usage

### Lookup by DHCP Options ID

```terraform
variable dopts_id {}

data "aws_vpc_dhcp_options" "example" {
  dhcp_options_id = var.dopts_id
}
```

### Lookup by Filter

```terraform
data "aws_vpc_dhcp_options" "example" {
  filter {
    name   = "key"
    values = ["domain-name"]
  }

  filter {
    name   = "value"
    values = ["example.com"]
  }
}
```

## Argument Reference

* `dhcp_options_id` - (Optional) The EC2 DHCP Options ID.
* `filter` - (Optional) List of custom filters as described below.

### filter

For more information about filtering, see the [EC2 API documentation][describe-dhcp-options].

* `name` - (Required) The name of the field to filter.
* `values` - (Required) Set of values for filtering.

## Attributes Reference

* `arn` - The ARN of the DHCP Options Set.
* `dhcp_options_id` - EC2 DHCP Options ID.
* `domain_name` - The suffix domain name to used when resolving non Fully Qualified Domain Names e.g., the `search` value in the `/etc/resolv.conf` file.
* `domain_name_servers` - List of name servers.
* `id` - EC2 DHCP Options ID.
* `netbios_name_servers` - List of NETBIOS name servers.
* `netbios_node_type` - The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
* `ntp_servers` - List of NTP servers.
* `tags` - A map of tags assigned to the resource.

Exported but unsupported attributes:

* `owner_id` - The ID of the CROC Cloud account that owns the DHCP options set.

[describe-dhcp-options]: https://docs.cloud.croc.ru/en/api/ec2/dhcp_options/DescribeDhcpOptions.html
