# main.tf
#
# This Terraform configuration file demonstrates how to use the Chuck Norris
# provider. It specifies the provider, defines a resource to fetch a Chuck Norris
# joke, and outputs the joke as a Terraform output.

terraform {
  required_providers {
    chucknorris = {
      source  = "terraform.local/local/chucknorris"
      version = "1.0.0"
    }
  }
}

provider "chucknorris" {}

resource "chucknorris_joke" "example" {}

output "joke" {
  value = chucknorris_joke.example.joke
}
