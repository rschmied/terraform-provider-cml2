terraform {
  required_providers {
    cml2 = {
      version = "0.1.0-alpha-20-ge944c3e"
      source  = "registry.terraform.io/ciscodevnet/cml2"
    }
  }
  # experiments = [module_variable_optional_attrs]
}

variable "address" {
  description = "CML controller address"
  type        = string
  # default     = "https://192.168.122.245"
  default = "https://cml-controller.cml.lab"
}

variable "token" {
  description = "CML API token"
  type        = string
  default     = "qwe"
}

variable "toponame" {
  description = "topology name"
  type        = string
  default     = "absolute bananas"
}
variable "username" {
  description = "cml2 username"
  type        = string
  default     = "admin"
}
variable "password" {
  description = "cml2 password"
  type        = string
  sensitive   = true
  # default     = "absolute bananas"
}

variable "configs" {
  description = "configuration for nodes"
  type        = map(string)
  default = {
    "alpine-0" : "hostname duedelfirstofhername",
    "alpine-1" : "hostname ralph",
    "alpine-2" : "hostname ohmygod"
    "alpine-3" : "hostname karlheinz",
  }
}

# variable "special" {
#   description = "configuration for nodes"
#   type = map(object({
#     configuration : optional(string),
#     state : optional(string),
#     image_id : optional(string)
#   }))
#   default = {
#     "group[0-9" : {
#       "configuration" : "hostname grumbl",
#       # "state": null
#       # "state": "STARTED"
#       # "image_id": "alpine-3-10-base"
#       # "image_id": null
#     }
#   }
# }

provider "cml2" {
  address = var.address
  # token   = var.token
  username = var.username
  password = var.password
  # username = "qwe"
  # password = "qwe"
  cacert = file("ca.pem")
  # skip_verify = false
}
