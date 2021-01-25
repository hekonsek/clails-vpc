output "vpc_id" {
  value = module.test.vpc.vpc_id
}

output "vpc_private_subnets" {
  value = module.test.vpc.private_subnets
}

output "vpc" {
  value = module.test.vpc
}
