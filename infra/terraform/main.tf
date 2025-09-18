provider "aws" {
  region = var.region
}

# --- EKS Cluster ---
module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = var.cluster_name
  cluster_version = "1.29"
  subnets         = var.subnets
  vpc_id          = var.vpc_id
}

# --- RDS (PostgreSQL) ---
resource "aws_db_instance" "postgres" {
  identifier        = "authen-postgres"
  allocated_storage = 20
  engine            = "postgres"
  engine_version    = "16.3"
  instance_class    = "db.t3.micro"
  name              = "authendb"
  username          = var.db_username
  password          = var.db_password
  publicly_accessible = false
  skip_final_snapshot = true
}

# --- Redis (Elasticache) ---
resource "aws_elasticache_cluster" "redis" {
  cluster_id           = "authen-redis"
  engine               = "redis"
  node_type            = "cache.t3.micro"
  num_cache_nodes      = 1
  parameter_group_name = "default.redis7"
}

output "postgres_endpoint" {
  value = aws_db_instance.postgres.address
}

output "redis_endpoint" {
  value = aws_elasticache_cluster.redis.cache_nodes[0].address
}
