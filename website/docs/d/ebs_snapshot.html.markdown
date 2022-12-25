---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "AWS: aws_ebs_snapshot"
description: |-
  Get information on an EBS Snapshot.
---

# Data Source: aws_ebs_snapshot

Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes

## Example Usage

```terraform
data "aws_ebs_snapshot" "ebs_snapshot" {
  most_recent = true
  owners      = ["self"]

  filter {
    name   = "volume-size"
    values = ["40"]
  }

  filter {
    name   = "tag:Name"
    values = ["Example"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `most_recent` - (Optional) If more than one result is returned, use the most recent snapshot.
* `owners` - (Optional) Returns the snapshots owned by the specified owner ID. Multiple owners can be specified.
* `snapshot_ids` - (Optional) Returns information on a specific snapshot ID.
* `restorable_by_user_ids` - (Optional) One or more Croc Cloud project IDs that can create volumes from the snapshot.
* `filter` - (Optional) One or more name/value pairs to filter off of. Supported filters:
  * `description` - The description of the snapshot. 
  * `owner-alias` - The alias of the image owner (for example, croc).
  * `owner-id` - The ID of the project that owns the snapshot.
  * `snapshot-id` - The ID of the snapshot.
  * `start-time` - The timestamp when the snapshot was initiated.
  * `status` - The state of the snapshot.
  * `volume-id` - The ID of the volume the snapshot is for.
  * `volume-size` - The size of the volume, in GiB.
  * `tag:<tag-key>` – The key/value combination of a tag. Use the tag key in the filter name and the tag value as the filter value. 
  * `tag-key` - The key of a tag assigned to the resource. This filter is used to find all tagged resources with the specific key tag and any tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the EBS Snapshot.
* `id` - The snapshot ID (e.g., snap-10F5F8DF).
* `snapshot_id` - The snapshot ID (e.g., snap-10F5F8DF).
* `description` - A description for the snapshot
* `owner_id` - The Croc Cloud project ID.
* `owner_alias` - The alias of the EBS snapshot owner.
* `volume_id` - The volume ID (e.g., vol-BE84BFA0).
* `volume_size` - The size of the drive in GiBs.
* `state` - The snapshot state.
* `tags` - A map of tags for the resource.

Exported but unsupported attributes:

* `data_encryption_key_id` - The data encryption key identifier for the snapshot. Always `""`.
* `encrypted` - Whether the snapshot is encrypted. Always `false`.
* `kms_key_id` - The ARN for the KMS encryption key. Always `""`.
* `outpost_arn` - The ARN of the Outpost on which the snapshot is stored. Always `""`.
* `storage_tier` - The storage tier in which the snapshot is stored. Always `""`.
