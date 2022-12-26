---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_route_table"
description: |-
  Provides a resource to create a VPC routing table.
---

# Resource: aws_route_table

Provides a resource to create a VPC routing table.

~> **NOTE on Route Tables and Routes:** Terraform currently
provides both a standalone [Route resource](route.html) and a Route Table resource with routes
defined in-line. At this time you cannot use a Route Table with in-line routes
in conjunction with any Route resources. Doing so will cause
a conflict of rule settings and will overwrite rules.

~> **NOTE on `propagating_vgws` and the `aws_vpn_gateway_route_propagation` resource:**
If the `propagating_vgws` argument is present, it's not supported to _also_
define route propagations using `aws_vpn_gateway_route_propagation`, since
this resource will delete any propagating gateways not explicitly listed in
`propagating_vgws`. Omit this argument when defining route propagation using
the separate resource.

## Example Usage

```terraform
resource "aws_route_table" "example" {
  vpc_id = aws_vpc.example.id

  route {
    cidr_block = "10.0.1.0/24"
    gateway_id = aws_internet_gateway.example.id
  }

  route {
    ipv6_cidr_block        = "::/0"
    egress_only_gateway_id = aws_egress_only_internet_gateway.example.id
  }

  tags = {
    Name = "example"
  }
}
```

To subsequently remove all managed routes:

```terraform
resource "aws_route_table" "example" {
  vpc_id = aws_vpc.example.id

  route = []

  tags = {
    Name = "example"
  }
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required) The VPC ID.
* `route` - (Optional) A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.
* `propagating_vgws` - (Optional) A list of virtual gateways for propagation.

### route Argument Reference

This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).

One of the following destination arguments must be supplied:

* `cidr_block` - (Required) The CIDR block of the route.

One of the following target arguments must be supplied:

* `gateway_id` - (Optional) ID of an Internet gateway or virtual private gateway.
* `instance_id` - (Optional) ID of an EC2 instance.
* `network_interface_id` - (Optional) ID of an EC2 network interface.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

~> **NOTE:** Only the target that is entered is exported as a readable
attribute once the route resource is created.

* `id` - ID of the routing table.
* `arn` - The ARN of the route table.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

Exported but unsupported attributes:

* `carrier_gateway_id` - (Optional) Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
* `destination_prefix_list_id` - ID of a managed prefix list destination of the route. Always `""`.
* `ipv6_cidr_block` - The Ipv6 CIDR block of the route. Always `""`.
* `owner_id` - ID of the Croc Cloud account that owns the Default Network ACL. Always `""`.
* `route`
  * `core_network_arn` - The Amazon Resource Name (ARN) of a core network. Always `""`.
  * `egress_only_gateway_id` - ID of a VPC Egress Only Internet Gateway. Always `""`.
  * `nat_gateway_id` - ID of a VPC NAT gateway. Always `""`.
  * `transit_gateway_id` - ID of an EC2 Transit Gateway. Always `""`.
  * `vpc_endpoint_id` - ID of a VPC Endpoint. Always `""`.
  * `vpc_peering_connection_id` - ID of a VPC peering connection. Always `""`.

## Timeouts

`aws_route_table` provides the following [Timeouts](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts) configuration options:

- `create` - (Default `5 minutes`) Used for route creation
- `update` - (Default `2 minutes`) Used for route creation
- `delete` - (Default `5 minutes`) Used for route deletion

## Import

Route Tables can be imported using the route table `id`. For example, to import
route table `rtb-4e616f6d69`, use this command:

```
$ terraform import aws_route_table.public_rt rtb-4e616f6d69
```

[default-tags]: index.html#default_tags-configuration-block
