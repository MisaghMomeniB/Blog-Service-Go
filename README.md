# ✍️ Blog Service (Go)

A modern **RESTful blog API** implemented in Go, complete with JWT authentication, CRUD operations, pagination, and basic tagging. Designed for backend developers to learn or extend a production-ready service.

---

## 📋 Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Tech Stack & Requirements](#tech-stack--requirements)  
4. [API Endpoints](#api-endpoints)  
5. [Installation & Run](#installation--run)  
6. [Usage Examples](#usage-examples)  
7. [Code Structure](#code-structure)  
8. [Security & Best Practices](#security--best-practices)  
9. [Enhancements](#enhancements)  
10. [Contributing](#contributing)  
11. [License](#license)

---

## 💡 Overview

This Go-based API implements a backend for a **blogging platform**, providing authentication, article management, tagging, and search. Built with performance and simplicity in mind—perfect for learning REST patterns, database integration, and secure design.

---

## ✅ Features

- 🔐 User registration & login with **JWT authentication**  
- 📝 CRUD operations for blog posts (Create / Read / Update / Delete)  
- 🗂️ Tag-based filtering & post searches  
- 📄 Pagination for article lists  
- 🧹 Input validation and error handling  
- ✅ Role-based access (author vs reader)

---

## 🛠️ Tech Stack & Requirements

- Go **1.18+** (modules enabled)  
- **Gin** or **Chi** router (fast HTTP handling)  
- **GORM** (PostgreSQL/MySQL/SQLite support)  
- **GitHub.com/golang-jwt/jwt** for token handling  
- **Validator.v10** for input validation

---

## 🔧 API Endpoints

| Endpoint               | Method | Description                                    | Auth Required |
|-----------------------|--------|------------------------------------------------|----------------|
| `/api/register`       | POST   | Create a new user                              | ❌              |
| `/api/login`          | POST   | Authenticate and obtain JWT                    | ❌              |
| `/api/posts`          | GET    | List posts (with pagination/query filters)     | ❌              |
| `/api/posts`          | POST   | Create a new post                              | ✅              |
| `/api/posts/{id}`     | GET    | Retrieve a specific post                       | ❌              |
| `/api/posts/{id}`     | PUT    | Update an existing post                        | ✅              |
| `/api/posts/{id}`     | DELETE | Delete a post                                  | ✅              |
| `/api/tags`           | GET    | Retrieve all tags                              | ❌              |

---

## ⚙️ Installation & Run

```bash
git clone https://github.com/MisaghMomeniB/Blog-Service-Go.git
cd Blog-Service-Go

go mod tidy
export DB_DSN="postgres://user:pass@localhost:5432/blogdb?sslmode=disable"
export JWT_SECRET="your_jwt_secret"
go run main.go
````

This initializes the API at `http://localhost:8080`.

---

## 🚀 Usage Examples

### Register & Login (CLI)

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret"}'

curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret"}'
```

### Create a Post (with JWT)

```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Hello Go","content":"My first post","tags":["go","api"]}'
```

---

## 📁 Code Structure

```
Blog-Service-Go/
├── cmd/               # main app entrypoint (main.go)
├── config/            # configuration, env loading
├── controllers/       # HTTP handlers
├── models/            # DB models (User, Post, Tag)
├── repositories/      # DB access logic using GORM
├── middleware/        # JWT auth, logging, validation
├── routes/            # Router setup
├── utils/             # Helpers (JWT, hashing)
├── go.mod
└── README.md
```

---

## 🔐 Security & Best Practices

* Uses **bcrypt** for secure password hashing
* Implements **JWT tokens** with expiration via `jwt-go`
* Sanitizes inputs and enforces schema validation
* Role-based authorization ensures users only modify their own content

---

## 🚧 Enhancements

* 🧠 Add **comments**, **likes**, and **user profiles**
* 📦 Implement role-based permissions (admin/editor pipelines)
* ♻️ Add **caching** (Redis) for list endpoints
* 📤 Add **media uploads**, image handling, and thumbnails
* 🧪 Write **unit tests** and **integration tests**

---

## 🤝 Contributing

Contributions welcome! Steps:

1. Fork the repo
2. Create a feature branch (`feature/…`)
3. Add changes/tests
4. Submit a detailed Pull Request

---

## 📄 License

Released under the **MIT License** — see `LICENSE` for details.
