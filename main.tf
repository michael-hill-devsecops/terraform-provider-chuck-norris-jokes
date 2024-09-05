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
