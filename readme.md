# Clails - AWS VPC Terraform module

This is Terraform module for AWS VPC. It follows the following common design of VPC for AWS:
- 3 public subnets in 3 different AZs
- 3 private subnets in 3 different AZs
- NAT gateway for egres communication for private subnets 

<p align="center">
  <img src="./arch.png" width="300" />
</p>