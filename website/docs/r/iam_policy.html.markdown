---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_policy"
description: |-
  Manages an IAM policy.
---

[iam-policies]: https://docs.cloud.croc.ru/en/services/iam/policies.html
[RFC3339 format]: https://datatracker.ietf.org/doc/html/rfc3339#section-5.8

# Resource: aws_iam_policy

Manages an IAM policy. For details about IAM policies, see the [user documentation][iam-policies].

## Example Usage

### Global Policy

```terraform
resource "aws_iam_policy" "example" {
  name        = "tf-policy-global"
  description = "tf-policy-global description"
  type        = "global"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode(
    {
      Statement = [
        {
          Action = ["iam:*"],
        },
      ]
    }
  )
}
```

### Project Policy

```terraform
resource "aws_iam_policy" "example" {
  name        = "tf-policy-project"
  description = "tf-policy-project description"
  type        = "project"

  policy = jsonencode(
    {
      Statement = [
        {
          Action = [
            "rc:GetConsole",
            "rc:ShareConsole",
          ]
        },
        {
          Action = [
            "ec2:DescribeInstances",
          ]
        },
      ]
    }
  )
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional, Editable) The description of the policy.
  The value must be no more than 1000 characters long.
* `name` - (Optional, Conflicts with `name_prefix`) The name of the policy. The value can only contain Latin letters, numbers, underscores (_),
  plus (+) and equals (=) signs, commas (,), periods (.), at symbols (@) and hyphens (-) (`^[\w+=,.@-]*$`).
  The value must be 1 to 128 characters long.
* `name_prefix` - (Optional, Conflicts with `name`) Creates a unique name beginning with the specified prefix.
  The value has the same character restrictions as `name`. The value must be 1 to 102 characters long.

~> If `name` and `name_prefix` are omitted, Terraform will assign a random, unique name prefixed with `terraform-`.

* `policy` - (Required, Editable) A string with policy-defined access rules in JSON format.
* `type` - (Required) The type of the policy. Valid values are `global`, `project`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - The Amazon Resource Name (ARN) of the policy
* `create_date` - The time when the policy was created in [RFC3339 format].
* `id` - The ARN of the policy.
* `owner` - The owner of the policy.
* `policy_id` - The ID of the policy.
* `update_date` - The time when the policy was last updated in [RFC3339 format].

->  **Unsupported attributes**
These attributes are currently unsupported:

* `path` - The path to the policy. Always `""`.

## Import

IAM policy can be imported using `arn`, e.g.,

* import predefined policy provided by cloud named `IAMFullAccess`:

```
$ terraform import aws_iam_policy.example arn:c2:iam:::policy/IAMFullAccess
```

* import policy provided by customer `test.customer` named `policy-example`:

```
$ terraform import aws_iam_policy.example arn:c2:iam::test.customer:policy/policy-example
```
