# 🛠️ Go Backend with Gin & GORM

This is a API built with [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/), and PostgreSQL. The project is designed for scalable deployment using AWS services like EC2, VPC, and S3.

## 🚀 Features

- 🔐 Secure VPC deployment (public/private subnet setup)
- 🗂️ File storage on AWS S3
- 📦 PostgreSQL integration with GORM ORM

## 📁 Project Structure

```
├── .gitignore
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
    └── main.go
├── config
    ├── config.go
    ├── database.go
    └── storage.go
├── docker-compose.yml
└── internal
    ├── handlers
        └── upload_handler.go
    ├── models
        └── doc_models.go
    ├── repo
        └── doc_db.go
    ├── routes
        └── routes.go
    └── storage
        └── storage.go
```

## ⚙️ Environment Variables

Make a `.env` file in the root directory with the following:

```env
PORT=8080

DB_HOST=your-private-ec2-ip
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydb

AWS_BUCKET=my-public-bucket-for-docs
AWS_REGION=us-west-2
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
```
---

Let me know if you want me to include Swagger/OpenAPI setup, Docker support, or examples of routes!
