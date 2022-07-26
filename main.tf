variable "stage" {
    description = "The stage of the pipeline the application is deployed to."
    default = "dev" # dev, test, prod
    type = string
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

# Generate a random id for our deployment group
resource "random_string" "random" {
  length  = 16
  special = false
  upper   = false
}

provider "aws" {
  region  = "us-east-2"
  # The profile to use for the AWS CLI. Make sure you have this profile in your ~/.aws/credentials file.
  # This profile should have access to the S3 bucket you want to use for the deployment.
  # See: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html
  profile = "s3-access"
}

# A bucket to store our File Meta-Data
# TODO: This should be stored and retrieved from a Smart Contract.
resource "aws_s3_bucket" "meta-data" {
  bucket = "meta-data-bucket-${stage}-${random_string.random.result}"
  force_destroy = true
  tags = {
    Name        = "meta-data-bucket"
    Environment = "${stage}"
    Deployment_ID = "${random_string.random.result}"
  }
}

# A bucket to store our obao files
# TODO: This should be stored and retrieved from a Storage Provider.
resource "aws_s3_bucket" "obao-file" {
  bucket = "obao-file-bucket-${stage}-${random_string.random.result}"
  force_destroy = true
  tags = {
    Name        = "obao-file-bucket"
    Environment = "${stage}"
    Deployment_ID = "${random_string.random.result}"
  }
}

# A bucket to store our obao files
# TODO: This should be stored and retrieved from a Storage Provider.
resource "aws_s3_bucket" "endpoint" {
  bucket = "endpoint-bucket-${stage}-${random_string.random.result}"
  force_destroy = true
  tags = {
    Name        = "endpoint-bucket"
    Environment = "${stage}"
    Deployment_ID = "${random_string.random.result}"
  }
}

