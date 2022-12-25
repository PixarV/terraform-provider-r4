---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "AWS: aws_ebs_volume"
description: |-
  Get information on an EBS volume.
---

# Data Source: aws_ebs_volume

Use this data source to get information about an EBS volume for use in other
resources.

## Example Usage

```terraform
data "aws_ebs_volume" "ebs_volume" {
  most_recent = true

  filter {
    name   = "volume-type"
    values = ["st2"]
  }

  filter {
    name   = "tag:Name"
    values = ["Example"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `most_recent` - (Optional) If more than one result is returned, use the most recent Volume.
* `filter` - (Optional) One or more name/value pairs to filter. Supported filters:
  * `attachment.instance-id` - The ID of the instance the volume is attached to.
  * `attachment.device` - The name of the device, attached to the instance (for example, disk1).
  * `attachment.status` - The attachment state.
  * `availability-zone` - The name of the availability zone, in which the volume was created.
  * `create-time` - The timestamp when the volume was created.
  * `size` - The size of the volume, in GiB.
  * `snapshot-id` - The ID of the snapshot from which the volume was created.
  * `status` - The state of the Volume.
  * `volume-id` - The ID of the volume the snapshot is for.
  * `volume-type` - The type of the volume. Can be `st2`, `gp2` or `io2`.
  * `tag:<tag-key>` – The key/value combination of a tag. Use the tag key in the filter name and the tag value as the filter value. 
  * `tag-key` - The key of a tag assigned to the resource. This filter is used to find all tagged resources with the specific key tag and any tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The volume ID (e.g., vol-59FCB34E).
* `volume_id` - The volume ID (e.g., vol-59FCB34E).
* `arn` - Amazon Resource Name (ARN) of the volume.
* `availability_zone` - The AZ where the EBS volume exists.
* `iops` - The amount of IOPS for the disk.
* `size` - The size of the drive in GiBs.
* `snapshot_id` - The snapshot_id the EBS volume is based off.
* `volume_type` - The type of EBS volume.
* `tags` - A map of tags for the resource.

Exported but unsupported attributes:

* `encrypted` - Whether the snapshot is encrypted. Always `false`.
* `kms_key_id` - The ARN for the KMS encryption key. Always `""`.
* `multi_attach_enabled` - Whether EBS Multi-Attach is enabled. Always `false`.
* `outpost_arn` - The ARN of the Outpost on which the snapshot is stored. Always `""`.
* `throughput` - The throughput that the volume supports, in MiB/s. Always 0.
