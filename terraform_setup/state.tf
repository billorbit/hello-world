terraform {
  backend "s3" {
    bucket = "billorbit-terraform-state-bucket"
    key = "terraform.tfstate"
    region = "ap-south-1"
  }
}