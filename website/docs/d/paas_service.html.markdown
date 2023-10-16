---
subcategory: "PaaS"
layout: "aws"
page_title: "CROC Cloud: aws_paas_service"
description: |-
  Retrieves information about a PaaS service.
---

# Data Source: aws_paas_service

Retrieves information about a PaaS service.

## Example Usage

```terraform
data "aws_paas_service" "selected" {
  id = "fm-cluster-12345678"
}

output "paas_service_name" {
  value = data.aws_paas_service.selected.name
}
```

## Argument Reference

The following arguments are supported:

* `id` - The ID of the PaaS service (e.g. fm-cluster-12345678).
  * Required: Yes

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arbitrator_required`, `delete_interfaces_on_destroy` -> todo: remove from schema
* `auto_created_security_group_ids` - Set of the security group IDs that were created by CROC Cloud PaaS for the service.
* `backup_settings` - Backup settings for the service. The structure of this block is [described below](#backup-settings).
* `data_volume` - Service data volume parameters. The structure of this block is [described below](#datavolume).
* `endpoints`

### backup settings

The `backup_settings` block has the following structure:

* `bucket_name` - The physical name of the device.
* `enabled` - Map containing EBS information, if the device is EBS-based.

### data_volume

The `data_volume` block has the following structure:

* `bucket_name` - The physical name of the device.
* `enabled` - Map containing EBS information, if the device is EBS-based.