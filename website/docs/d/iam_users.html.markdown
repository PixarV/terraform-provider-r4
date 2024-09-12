---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_users"
description: |-
  Provides ARNs and names of selected IAM users.
---

# Data Source: aws_iam_users

Provides ARNs (Amazon Resource Names) and names of selected IAM users.

## Example Usage

### All users

```terraform
data "aws_iam_users" "selected" {}
```

### Users filtered by name regex

```terraform
data "aws_iam_users" "selected" {
  name_regex = "user.*"
}
```

## Argument Reference

The following arguments are supported:

* `name_regex` - (Optional) A regex string to apply to the list of users returned by IAM API.

~> This filtering is done locally and could have a performance impact if the list of users is large.

## Attribute Reference

* `arns` - List of user ARNs.
* `id` - The region (e.g., `region-1`).
* `names` - List of usernames.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `path_prefix` - The path prefix for filtering the results.
