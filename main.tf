terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

resource "random_string" "random" {
  length  = 16
  special = false
  upper   = false
}

provider "aws" {
  region  = "us-east-2"
  profile = "s3-access"
}

resource "aws_s3_bucket" "meta-data" {
  bucket = "meta-data-test-bucket-${random_string.random.result}"
  tags = {
    Name        = "meta-data-test-bucket-${random_string.random.result}"
    Environment = "test"
  }
}

resource "aws_s3_bucket" "obao" {
  bucket = "obao-test-bucket-${random_string.random.result}"
  tags = {
    Name        = "obao-test-bucket-${random_string.random.result}"
    Environment = "test"
  }
}

