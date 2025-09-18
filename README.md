# 🚀 GoGinTemplateDDD

A production-ready template for building scalable microservices with **Go + Gin**, following **DDD (Domain Driven Design)** principles.  
Includes PostgreSQL, Redis, Kubernetes manifests, and Terraform for infrastructure management.

---

## 📂 Project Structure

```bash
.
├── cmd/                  # Entry point for the application
│   └── app/              # Main application entry
│       └── main.go       # Main file to start the service
│
├── Dockerfile            # Docker build file for containerizing the service
├── go.mod                # Go module definition
├── go.sum                # Go module dependency checksums
│
├── infra/                # Infrastructure layer (DB, Redis, K8s, Terraform)
│   ├── db/               # Database & cache connections + migration scripts
│   │   ├── postgresql/   # PostgreSQL setup & migrations
│   │   │   ├── migrations/   # Migration scripts (up/down SQL)
│   │   │   └── postgres.go   # PostgreSQL connection setup
│   │   └── redis/            # Redis client setup
│   │       └── redis.go
│   ├── k8s/              # Kubernetes manifests (YAML)
│   │   ├── deployment.yaml   # Deployment definition
│   │   ├── ingress.yaml      # Ingress for routing
│   │   └── service.yaml      # Service (ClusterIP/LoadBalancer)
│   └── terraform/        # Terraform IaC configuration
│       ├── main.tf           # Main infra definition
│       ├── outputs.tf        # Output variables
│       └── variables.tf      # Input variables
│
├── internal/             # Application business logic (DDD structure)
│   ├── app/              # App wiring (modules, routes, bootstrap)
│   │   ├── app.go
│   │   ├── module.go
│   │   └── routes.go
│   ├── config/           # Load & manage app configuration
│   │   ├── config.go
│   │   └── helper.go
│   ├── domain/           # Domain layer (Entities, Repositories, Services)
│   │   ├── entity/           # Core business entities
│   │   │   └── user_entity.go
│   │   ├── repository/       # Repository interfaces
│   │   │   └── user_repo_interface.go
│   │   └── service/          # Service interfaces
│   │       └── user_service_interface.go
│   ├── dto/              # Data Transfer Objects (request/response payloads)
│   │   └── auth_dto.go
│   ├── handler/          # HTTP handlers (controller layer)
│   │   ├── v1/               # Version 1 of API
│   │   │   └── user_handler.go
│   │   └── v2/               # Version 2 of API
│   │       └── user_handler.go
│   ├── models/           # Database models / ORM mapping
│   │   └── user_models.go
│   ├── repository/       # Repository implementations (Postgres, Redis, etc.)
│   │   └── user_repository_imple.go
│   └── service/          # Service implementations (business logic)
│       └── user_service_impl.go
│
├── README.md             # Project documentation
└── util/                 # Utility functions (helpers)
    ├── gen_token.go          # JWT token generation
    └── hash.go               # Hashing utilities (password, etc.)
```
⚙️ Features

✅ Go + Gin for fast HTTP APIs

✅ DDD architecture for maintainability

✅ PostgreSQL + Redis integration

✅ Migration scripts with up/down support

✅ Kubernetes manifests (Deployment, Service, Ingress)

✅ Terraform for cloud infrastructure

✅ JWT authentication support

✅ Modular design for scaling services
