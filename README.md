# Go Auth API (JWT & Role Based)

Simple authentication API built with **Golang**.  
This project is part of my learning journey transitioning from **Node.js backend** to **Golang backend**.

> ⚠️ Note: This project is still a **fake / mock implementation** (no database yet).  
> The main goal is to understand **HTTP flow, clean architecture, JWT authentication, and middleware concepts in Go**.

---

## ✨ Features

- HTTP server using `net/http`
- Register endpoint (fake, in-memory)
- Login endpoint (JWT)
- Password hashing (bcrypt)
- JWT authentication
- Role-based access control (RBAC)
- Clean architecture:
  - Handler (HTTP layer)
  - Service (business logic)
  - Middleware (auth & role)
- JSON request & response
- Proper HTTP status codes

---

## 📁 Project Structure

```
go-auth-api/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── handler/
│   │   ├── register.go
│   │   ├── login.go
│   │   ├── profile.go
│   │   └── admin.go
│   │
│   ├── service/
│   │   ├── register.go
│   │   └── login.go
│   │
│   ├── middleware/
│   │   ├── jwt.go
│   │   └── admin.go
│   │
│   └── response/
│       └── response.go
│
├── pkg/
│   └── utils/
│       ├── jwt.go
│       └── password.go
│
├── .env
├── go.mod
└── README.md
```

---

## 🚀 How to Run

Make sure Go is installed:

```bash
go version
```

Create `.env` file:

```env
JWT_SECRET=your-super-secret-key
```

Run server:

```bash
go run cmd/api/main.go
```

Server runs at:

```
http://localhost:8080
```

---

## 🔗 API Endpoints

### Health Check
```http
GET /health
```

Response:
```
OK
```

---

### Register (Fake)
```http
POST /register
```

Body:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "message": "register success"
}
```

---

### Login (JWT)
```http
POST /login
```

Response:
```json
{
  "token": "JWT_TOKEN_HERE"
}
```

---

### Profile (Protected)
```http
GET /profile
Authorization: Bearer <JWT_TOKEN>
```

Response:
```json
{
  "message": "profile access granted",
  "data": {
    "email": "user@example.com"
  }
}
```

---

### Admin Dashboard (Admin Only)
```http
GET /admin
Authorization: Bearer <ADMIN_JWT_TOKEN>
```

Response:
```json
{
  "message": "admin dashboard",
  "data": {
    "status": "ok"
  }
}
```

---

## 🧠 What I Learned

- Go HTTP server fundamentals
- Clean architecture in Go
- Handler vs Service vs Middleware separation
- JWT authentication flow
- Context usage in Go
- Role-based access control (RBAC)
- Password hashing with bcrypt

---

## 🛠️ Next Improvements

- Database integration (PostgreSQL)
- Refresh token
- Logging middleware
- Request validation
- Unit testing

---

## 👤 Author

Rifqi  
Backend Developer (Node.js → Golang)

---

## 📌 Notes

This project focuses on **backend fundamentals and clean architecture**, not production readiness yet.
