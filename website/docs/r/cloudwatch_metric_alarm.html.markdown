---
subcategory: "CloudWatch"
layout: "aws"
page_title: "AWS: aws_cloudwatch_metric_alarm"
description: |-
  Provides a cloudwatch metric alarm resource.
---

[metrics]: https://docs.k2.cloud/en/services/monitoring/metrics.html
[dimensions]: https://docs.k2.cloud/en/services/monitoring/metrics.html#dimensions

# Resource: aws_cloudwatch_metric_alarm

Provides a cloudwatch metric alarm resource.

## Example Usage

```terraform
variable instance_id {}

resource "aws_cloudwatch_metric_alarm" "example" {
  alarm_name          = "terraform-test-metric-alarm"
  alarm_description   = "This metric monitors EC2 CPU utilization"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  statistic           = "Average"
  evaluation_periods  = 2
  period              = 120
  threshold           = 80
  alarm_actions       = ["example@mail.com:EMAIL"]
  dimensions = {
    InstanceId = var.instance_id
  }
}
```

## Argument Reference

The following arguments are supported:

* `actions_enabled` - (Optional) Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
* `alarm_actions` - (Optional) The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action must be between 1-1024 characters in length. You can specify a maximum of 5 actions.
* `alarm_description` - (Optional) The description for the alarm. Must be between 1-255 characters in length.
* `alarm_name` - (Required) The descriptive name for the alarm. This name must be unique within the user's account. Must be between 1-255 characters in length.
* `comparison_operator` - (Required) The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`.
* `datapoints_to_alarm` - (Optional) The number of datapoints that must be breaching to trigger the alarm. Minimum value is 1.
* `dimensions` - (Required) The dimensions for the alarm's associated metric. You can specify a maximum of 10 dimensions. See docs for [dimensions][dimensions].
* `evaluation_periods` - (Required) The number of periods over which data is compared to the specified threshold. Minimum value is 1.
* `insufficient_data_actions` - (Optional) The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action must be between 1-1024 characters in length. You can specify a maximum of 5 actions.
* `metric_name` - (Required) The name for the alarm's associated metric. See docs for [supported metrics][metrics].
* `namespace` - (Required) The namespace for the alarm's associated metric. See docs for the [list of namespaces and supported metrics][metrics].
* `ok_actions` - (Optional) The list of actions to execute when this alarm transitions into an OK state from any other state. Each action must be between 1-1024 characters in length. You can specify a maximum of 5 actions.
* `period` - (Required) The period in seconds over which the specified `statistic` is applied. Value must be divisible by 60, minimum value is 60.
* `statistic` - (Required) The statistic to apply to the alarm's associated metric. Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
* `threshold` - (Required) The value against which the specified statistic is compared.
* `treat_missing_data` - (Optional) Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
* `unit` - (Optional) The unit for the alarm's associated metric. Valid values are `Percent`, `Bytes` and `Count`.

## Import

CloudWatch Metric Alarm can be imported using the `alarm_name`, e.g.,

```
$ terraform import aws_cloudwatch_metric_alarm.test terraform-test-metric-alarm
```
