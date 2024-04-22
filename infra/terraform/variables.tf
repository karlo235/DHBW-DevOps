variable "resource_group_name" {
  description = "Name for the Azure resource group."
  type        = string
  default     = "example-dhbw-terraform"
}

variable "resource_group_location" {
  description = "Location for the Azure resource group."
  type        = string
  default     = "West Europe"
}

variable "environment_tag" {
  description = "Tag for the environment."
  type        = string
  default     = "Production"
}