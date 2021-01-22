terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "random_string" "random" {
  length = 10
}

module "test" {
  source = "./.."

  name = random_string.random.result
}