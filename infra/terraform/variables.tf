variable "region" {
  default = "ap-southeast-1"
}

variable "cluster_name" {
  default = "authen-eks"
}

variable "vpc_id" {}
variable "subnets" {
  type = list(string)
}

variable "db_username" {
  default = "admin"
}

variable "db_password" {
  default = "supersecret"
}
