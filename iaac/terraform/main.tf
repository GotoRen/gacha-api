provider "aws" {
  region  = var.region
  profile = var.profile
}

# terraform {
#   backend "s3" {
#     key    = "cta.tfstate"
#     region = "ap-northeast-1"
#   }
# }

locals {
  container_port = 8080
  created_by     = "Terraform"
  envs = [
    {
      name          = "staging",
      domain_prefix = "staging."
    },
    {
      name          = "production",
      domain_prefix = ""
    }
  ]
  domain          = "cta.tetsuya28.com"
  internal_domain = "cta.internal"
}

# 環境共通
data "aws_iam_policy_document" "ecr" {
  statement {
    sid       = "2"
    actions   = ["ecr:*"]
    resources = ["*"]
  }
  statement {
    sid       = "3"
    actions   = ["s3:*"]
    resources = ["*"]
  }
}

module "github_action_iam" {
  # source       = "github.com/tetsuya28/samples/terraform/modules/vpc"
  source     = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/iam"
  app        = var.app
  created_by = local.created_by
  name       = "github_action"
  policy     = data.aws_iam_policy_document.ecr.json
}

module "route53" {
  source     = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/route53"
  app        = var.app
  env        = "common"
  created_by = local.created_by
  domain     = local.domain
}

resource "aws_route53_zone" "internal" {
  name = local.internal_domain
  vpc {
    vpc_id = module.vpc.vpc_id
  }
  tags = {
    Name      = "${var.app}-internal"
    Env       = "common"
    CreatedBy = local.created_by
  }
}

module "vpc" {
  # source       = "github.com/tetsuya28/samples/terraform/modules/vpc"
  source       = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/vpc"
  app          = var.app
  created_by   = local.created_by
  vpc_cidr     = "172.16.0.0/16"
  subnet_count = 3
  nat          = false
}

module "aurora" {
  source           = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/aurora"
  app              = var.app
  env              = "common"
  created_by       = local.created_by
  internal_domain  = local.internal_domain
  internal_zone_id = aws_route53_zone.internal.zone_id
  engine           = "aurora-mysql"
  engine_version   = "5.7.mysql_aurora.2.09.1"
  family           = "aurora-mysql5.7"
  database_name    = var.app
  master_username  = "tetsuya"
  vpc_id           = module.vpc.vpc_id
  subnet_ids       = module.vpc.private_subnets
  source_sg_ids    = module.ecs.*.ec2_sg_id
  rds_params = {
    "character_set_client"     = "utf8mb4"
    "character_set_database"   = "utf8mb4"
    "character_set_connection" = "utf8mb4"
    "character_set_results"    = "utf8mb4"
    "character_set_server"     = "utf8mb4"
  }
}

# 環境依存
module "ecr" {
  count = length(local.envs)
  # source       = "github.com/tetsuya28/samples/terraform/modules/ecr"
  source     = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/ecr"
  app        = var.app
  env        = local.envs[count.index].name
  created_by = local.created_by
  name       = "cta-${local.envs[count.index].name}"
}

module "alb" {
  count = length(local.envs)
  # source       = "github.com/tetsuya28/samples/terraform/modules/alb"
  source     = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/alb"
  app        = var.app
  created_by = local.created_by
  env        = local.envs[count.index].name
  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.public_subnets
  domains    = ["${local.envs[count.index].domain_prefix}${local.domain}"]
  zone_id    = module.route53.zone_id
}

data "template_file" "container_def" {
  count    = length(module.ecr)
  template = file("container_def.json")
  vars = {
    app                 = var.app
    ecr_url             = module.ecr[count.index].repository.repository_url
    port                = local.container_port
    rds_endpoint        = module.aurora.rds_endpoint
    log_group_name      = module.log_group[count.index].name
    rds_master_password = module.aurora.rds_master_password
  }
}

module "log_group" {
  count      = length(local.envs)
  source     = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/cloudwatch_log"
  app        = var.app
  env        = local.envs[count.index].name
  created_by = local.created_by
  name       = "${var.app}/${local.envs[count.index].name}/ecs"
}

module "ecs" {
  count = length(local.envs)
  # source       = "github.com/tetsuya28/samples/terraform/modules/ecs"
  source           = "../../../../ghq/github.com/tetsuya28/samples/terraform/modules/ecs"
  app              = var.app
  env              = local.envs[count.index].name
  created_by       = local.created_by
  domain           = "${local.envs[count.index].domain_prefix}${local.domain}"
  vpc_id           = module.vpc.vpc_id
  subnet_ids       = module.vpc.public_subnets
  alb_sg_id        = module.alb[count.index].alb_sg_id
  task_role_policy = ""
  listener_arn     = module.alb[count.index].listener_arn
  container_def    = data.template_file.container_def[count.index].rendered
  container_port   = local.container_port
  container_name   = var.app
}
