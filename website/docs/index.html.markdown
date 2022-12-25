---
layout: "aws"
page_title: "Provider: AWS"
description: |-
  Use the Amazon Web Services (AWS) provider to interact with the many resources supported by AWS. You must configure the provider with the proper credentials before you can use it.
---

# AWS Provider

Fourth version

Use the Amazon Web Services (AWS) provider to interact with the
many resources supported by AWS. You must configure the provider
with the proper credentials before you can use it.

Use the navigation to the left to read about the available resources.

To learn the basics of Terraform using this provider, follow the
hands-on [get started tutorials](https://learn.hashicorp.com/tutorials/terraform/infrastructure-as-code?in=terraform/aws-get-started&utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) on HashiCorp's Learn platform. Interact with AWS services,
including Lambda, RDS, and IAM by following the [AWS services
tutorials](https://learn.hashicorp.com/collections/terraform/aws?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS).

## Example Usage

Terraform 0.13 and later:

```terraform
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "us-east-1"
}

# Create a VPC
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}
```

Terraform 0.12 and earlier:

```terraform
# Configure the AWS Provider
provider "aws" {
  version = "~> 3.0"
  region  = "us-east-1"
}

# Create a VPC
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}
```

## Authentication and Configuration

Configuration for the AWS Provider can be derived from several sources,
which are applied in the following order:

1. Parameters in the provider configuration
1. Environment variables
1. Shared credentials files
1. Shared configuration files
1. Container credentials
1. Instance profile credentials and region

This order matches the precedence used by the
[AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-precedence)
and the [AWS SDKs](https://aws.amazon.com/tools/).

The AWS Provider supports assuming an IAM role, either in
the provider configuration block parameter `assume_role`
or in [a named profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html).

The AWS Provider supports assuming an IAM role using [web identity federation and OpenID Connect (OIDC)](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html#cli-configure-role-oidc).
This can be configured either using environment variables or in a named profile.

When using a named profile, the AWS Provider also supports [sourcing credentials from an external process](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html).

### Provider Configuration

!> **Warning:** Hard-coded credentials are not recommended in any Terraform
configuration and risks secret leakage should this file ever be committed to a
public version control system.

Credentials can be provided by adding an `access_key`, `secret_key`, and optionally `token`, to the `aws` provider block.

Usage:

```terraform
provider "aws" {
  region     = "us-west-2"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}
```

Other settings related to authorization can be configured, such as:

* `profile`
* `shared_config_files`
* `shared_credentials_files`

### Environment Variables

Credentials can be provided by using the `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and optionally `AWS_SESSION_TOKEN` environment variables.
The region can be set using the `AWS_REGION` or `AWS_DEFAULT_REGION` environment variables.

For example:

```terraform
provider "aws" {}
```

```sh
$ export AWS_ACCESS_KEY_ID="anaccesskey"
$ export AWS_SECRET_ACCESS_KEY="asecretkey"
$ export AWS_REGION="us-west-2"
$ terraform plan
```

Other environment variables related to authorization are:

* `AWS_PROFILE`
* `AWS_CONFIG_FILE`
* `AWS_SHARED_CREDENTIALS_FILE`

