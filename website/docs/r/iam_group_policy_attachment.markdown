---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_group_policy_attachment"
description: |-
  Attaches an IAM Policy to an IAM group.
---

# Resource: aws_iam_group_policy_attachment

Attaches an IAM Policy to an IAM group.

## Example Usage

```terraform
resource "aws_iam_group" "example" {
  name = "tf-group"
  type = "global"
}

resource "aws_iam_policy" "example" {
  name = "tf-policy"
  type = "global"

  policy = jsonencode(
    {
      Statement = [
        {
          Action = ["iam:ListUsers"],
        },
      ]
    }
  )
}

resource "aws_iam_group_policy_attachment" "example" {
  group_arn  = aws_iam_group.example.arn
  policy_arn = aws_iam_policy.example.arn
}
```

## Argument Reference

The following arguments are supported:

* `group_arn` - (Required) The Amazon Resource Name (ARN) of the group.
* `policy_arn` - (Required) The ARN of the attached policy.

## Attribute Reference

* `id` - The `group_arn` and `policy_arn` separated by a hash sign (`#`).

## Import

IAM group policy attachments can be imported using the `id`, e.g.,

```
$ terraform import aws_iam_group_policy_attachment.example arn:c2:iam::test.customer:group/group-example#arn:c2:iam:::policy/BackupFullAccess
```
