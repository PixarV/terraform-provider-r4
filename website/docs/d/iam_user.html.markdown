---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_user"
description: |-
  Provides information about a IAM user.
---

[RFC3339 format]: https://datatracker.ietf.org/doc/html/rfc3339#section-5.8

# Data Source: aws_iam_user

Provides information about a IAM user.

## Example Usage

```terraform
data "aws_iam_user" "example" {
  user_name = "user"
}
```

## Argument Reference

* `user_name` - (Required) The name of the user.

## Attribute Reference

* `arn` - The Amazon Resource Name (ARN) of the user.
* `display_name` - The displayed name of the user.
* `email` - The email of the user.
todo * `enabled` - Indicates whether the user is locked.
* `id` - The ID of the user.
* `identity_provider` - The ID of the identity provider of the user. It is specified only for IdP users.
* `last_login_date` - The time when the user last logged in to the web interface in [RFC3339 format].
* `login` - The login of the user.
* `otp_required` -  Indicates whether the user is required to use two-factor authentication to log in to the web interface.
* `phone` - The phone number of the user.
* `update_date` - The time the user was last updated in [RFC3339 format].
* `user_id` - The ID of the user.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `path` - Path in which this user was created. Always `""`.
* `permissions_boundary` - The ARN of the policy that is used to set the permissions boundary for the user. Always `""`.
* `tags` - Map of tags assigned to the user. Always `{}`.
