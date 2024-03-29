# Terraform Provider for Cisco CML2

[![CodeQL](https://github.com/rschmied/terraform-provider-cml2/actions/workflows/codeql-analysis.yml/badge.svg?branch=dev)](https://github.com/rschmied/terraform-provider-cml2/actions/workflows/codeql-analysis.yml)

_This template repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). The template repository built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) can be found at [terraform-provider-scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding). See [Which SDK Should I Use?](https://www.terraform.io/docs/plugin/which-sdk.html) in the Terraform documentation for additional information._

This repository implements a [Terraform](https://www.terraform.io) provider for Cisco Modeling Labs version 2.3 and later. It's current state is "work-in-progress".  The current implementation provides:

- A resource and a data source (`internal/provider/`),
  - resource `cml2_lab` to create, update and destroy a lab based on a YAML topology file
  - update allows to modify state, e.g. from STOPPED to STARTED, ...
  - data source `cml2_lab_details` to retrieve operational state from a running lab.
- Examples (`examples/`) and generated documentation (`docs/`),
- Miscellaneous meta files.

Note:  The examples and docs as well as the tests are pretty much identical to
  the files found in the original templates...  See the [TODO.md](TODO.md) file!

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.17
- [CML2](https://cisco.com/go/cml) >= 2.3.0

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

### Environment

The provider is using environment variables for configuration.  Heres an example:

```shell
# for testing purposes, suggest to use direnv

# this is used by the cmlclient test command, ctest, not needed for the provider
# CML_LABID="ba35782a-06ee-4234-9de4-473ffe3c09e8"

CML_HOST="https://cml-controller.cml.lab"
CML_TOKEN=" valid-cmd-jwt "
CML_USERNAME=""
CML_PASSWORD=""
TF_VAR_token=$CML_TOKEN

export CML_HOST CML_USERNAME CML_PASSWORD TF_VAR_token
```

### HCL

> _NOTE:_ this is mostly outdated... Need to update!!

Here's a simple example that shows how to use the provider. It will import a
lab as defined in `topology.yml`, then start it and wait for it to converge
(e.g. all nodes report that they are BOOTED).  Then it will read the lab data
via `cml2_lab_details` using the ID from the import and print out the lab
details.  Only nodes with an IP address are considered for output (`only_with_ip = true`).

[![asciicast](https://asciinema.org/a/PfYfD1Br3QtytmR76kbGL1pva.svg)](https://asciinema.org/a/PfYfD1Br3QtytmR76kbGL1pva)

```hcl
terraform {
  required_providers {
    cml2 = {
      version = "0.1.0-alpha"
      source  = "registry.terraform.io/ciscodevnet/cml2"
    }
}

variable "address" {
  description = "CML controller address"
  type        = string
  default     = "https://192.168.122.245"
}

variable "token" {
  description = "CML API token"
  type        = string
}

variable "toponame" {
  description = "topology name"
  type        = string
  default     = "absolute bananas"
}

provider "cml2" {
  address = var.address
  token   = var.token
  # token       = null
  skip_verify = true
}

resource "cml2_lab" "bananas" {
  topology = templatefile("topology.yaml", { toponame = var.toponame })
  # start    = false
  # wait     = false
  # state    = "STARTED"
}

data "cml2_lab_details" "example" {
  id = cml2_lab.bananas.id
  only_with_ip = true
}

output "bla" {
  value = data.cml2_lab_details.example
}
```

Example `.terraformrc`:

```hcl
provider_installation {

  # dev_overrides {
  #   "hashicorp/cml2" = "/home/rschmied/Projects/terraform-provider-cml2"
  # }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  # direct {}
  filesystem_mirror {
    path = "/tmp/terraform"
    include = [ "registry.terraform.io/ciscodevnet/*" ]
  }
}
```

Sample filesystem mirror layout:

```plain
/tmp/terraform/
└── registry.terraform.io
    └── ciscodevnet
        └── cml2
            ├── terraform-provider-cml2_0.1.0-alpha_darwin_amd64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_darwin_arm64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_freebsd_386.zip
            ├── terraform-provider-cml2_0.1.0-alpha_freebsd_amd64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_freebsd_arm64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_freebsd_arm.zip
            ├── terraform-provider-cml2_0.1.0-alpha_linux_386.zip
            ├── terraform-provider-cml2_0.1.0-alpha_linux_amd64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_linux_arm64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_linux_arm.zip
            ├── terraform-provider-cml2_0.1.0-alpha_windows_386.zip
            ├── terraform-provider-cml2_0.1.0-alpha_windows_amd64.zip
            ├── terraform-provider-cml2_0.1.0-alpha_windows_arm64.zip
            └── terraform-provider-cml2_0.1.0-alpha_windows_arm.zip

3 directories, 14 files
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
