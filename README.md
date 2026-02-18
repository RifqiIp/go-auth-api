# Go Auth API (Basic)

Simple authentication API built with **Golang**.  
This project is part of my learning journey transitioning from **Node.js backend** to **Golang backend**.

> ⚠️ Note: This project is still a **fake / mock implementation** (no database yet).  
> The main goal is to understand **HTTP flow, handler–service architecture, and Go fundamentals**.

---

## ✨ Features

- HTTP server using `net/http`
- Register endpoint (fake)
- Login endpoint (fake)
- Clean architecture:
  - Handler (HTTP layer)
  - Service (business logic)
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
│   │   └── response.go
│   │
│   ├── service/
│   │   ├── register.go
│   │   └── login.go
│   │
│   └── middleware/          # planned
│
├── go.mod
└── README.md
```

---

## 🚀 How to Run

Make sure Go is installed:

```bash
go version
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
  "email": "rifqi@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "message": "register success (fake)"
}
```

Status:
```
201 Created
```

---

### Login (Fake)
```http
POST /login
```

Body:
```json
{
  "email": "rifqi@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "message": "login success"
}
```

Status:
```
200 OK
```

---

## 🧠 What I Learned

- Go HTTP server basics
- Handler vs Service separation
- JSON encoding/decoding
- HTTP status codes in Go
- Backend project structuring

---

## 🛠️ Next Improvements

- JWT authentication
- Password hashing
- PostgreSQL integration
- Middleware (auth, logging)

---

## 👤 Author

Rifqi  
Backend Developer (Node.js → Golang)

---

## 📌 Notes

This project is intentionally simple and focused on fundamentals.
