# Go Auth API (JWT & Role Based)

Authentication API built with **Golang** using the standard library (`net/http`).
This project is part of my learning journey transitioning from **Node.js backend** to **Golang backend**.

> ⚠️ Status: **Under Active Development**  
> PostgreSQL connection is already implemented and working.  
> Authentication flow (register/login/JWT) is being developed incrementally.

---

## 🎯 Project Goals

- Learn Go HTTP server fundamentals
- Apply clean architecture principles in Go
- Implement JWT authentication & RBAC step by step
- Integrate PostgreSQL without ORM (low-level control)
- Improve development workflow using hot reload (`air`)

---

## 🧱 Tech Stack

- Go 1.21+
- PostgreSQL
- net/http
- pgx (PostgreSQL driver)
- godotenv
- Air (live reload)

---

## 📁 Project Structure

```
go-auth-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── database.go
│   ├── handler/
│   │   └── (in development)
│   ├── middleware/
│   │   └── (in development)
│   └── service/
│       └── (in development)
├── .env.example
├── go.mod
└── README.md
```

---

## ⚙️ Environment Variables

Create a `.env` file in the project root.

Example (`.env.example`):

```
DATABASE_URL=postgres://<db_user>:<db_password>@localhost:5432/<db_name>
PORT=8080
```

⚠️ Do **NOT** commit your real `.env` file.

---

## 🗄️ Database

- PostgreSQL
- Connection handled via `pgxpool`
- Database connection is initialized once on app startup
- Connection tested successfully from Go

---

## ▶️ Running the App

Using `air` (recommended for development):

```
air
```

Or without hot reload:

```
go run ./cmd/api
```

Expected output:

```
Database connected
Server running on :8080
```

---

## 🚧 Current Progress

- [x] Project structure
- [x] Environment configuration
- [x] PostgreSQL connection (pgxpool)
- [x] Live reload with Air
- [ ] User repository
- [ ] Register API (DB-backed)
- [ ] Login API
- [ ] JWT authentication
- [ ] Role-based access control

---

## 🧠 Learning Focus

- Initialization order in Go applications
- Dependency boundaries (config vs handler vs service)
- Database connection lifecycle
- Clean, readable project structure
- Backend fundamentals over speed

---

## 👤 Author

Rifqi  
Backend Developer (Node.js → Golang)

---

## 📌 Notes

This project prioritizes **correctness and learning** over production readiness.
