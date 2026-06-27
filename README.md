# SnippetBox

SnippetBox is a full-stack web application built with Go that allows users to create, store, and view text snippets. The project was developed by following **Let's Go** by Alex Edwards and showcases how to build secure, production-ready web applications using Go's standard library.

## ✨ Features

* User registration and authentication
* Login & Logout
* Session management
* Password hashing (bcrypt)
* CSRF protection
* Secure HTTPS configuration
* Create and view snippets
* Form validation
* SQL integration
* HTML template inheritance
* Custom middleware
* Static asset serving
* Structured logging
* Unit testing for templates

---

## 🛠 Tech Stack

* Go
* SQL
* HTML Templates
* CSS
* JavaScript
* net/http
* SCS Session Manager
* nosurf
* bcrypt

---

## 📂 Project Structure

```text
snippetbox/
│
├── cmd/
│   └── web/
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go
│       ├── middleware.go
│       ├── routes.go
│       ├── templates.go
│       └── templates_test.go
│
├── internal/
│   ├── assert/
│   ├── validator/
│   │   └── validator.go
│   └── models/
│       ├── errors.go
│       ├── snippets.go
│       └── users.go
│
├── tls/
│   ├── cert.pem
│   └── key.pem
│
├── ui/
│   ├── html/
│   │   ├── base.tmpl
│   │   ├── partials/
│   │   │   └── nav.tmpl
│   │   └── pages/
│   │       ├── home.tmpl
│   │       ├── view.tmpl
│   │       ├── create.tmpl
│   │       ├── signup.tmpl
│   │       └── login.tmpl
│   │
│   └── static/
│       ├── css/
│       ├── img/
│       └── js/
│
├── .vscode/
├── go.mod
└── go.sum
```

---

## 🌐 Routes

| Method | Endpoint             | Description                       |
| ------ | -------------------- | --------------------------------- |
| GET    | `/`                  | Display the home page             |
| GET    | `/snippet/view/{id}` | View a specific snippet           |
| GET    | `/snippet/create`    | Display the snippet creation form |
| POST   | `/snippet/create`    | Create a new snippet              |
| GET    | `/user/signup`       | Display the signup form           |
| POST   | `/user/signup`       | Register a new user               |
| GET    | `/user/login`        | Display the login form            |
| POST   | `/user/login`        | Authenticate and log in a user    |
| POST   | `/user/logout`       | Log out the authenticated user    |
| GET    | `/static/*`          | Serve CSS, JavaScript, and images |

---

## 🔒 Authentication & Security

* User registration
* User login/logout
* Secure password hashing using bcrypt
* Cookie-based session management
* CSRF protection
* HTTPS support
* Secure HTTP headers
* Request validation

---

## 🧩 Middleware

* Panic recovery
* Request logging
* Secure headers
* CSRF protection
* Authentication
* Authorization for protected routes

---

## 🗄 Database

### snippets

| Column  | Type      |
| ------- | --------- |
| id      | INTEGER   |
| title   | TEXT      |
| content | TEXT      |
| created | TIMESTAMP |
| expires | TIMESTAMP |

### users

| Column          | Type      |
| --------------- | --------- |
| id              | INTEGER   |
| name            | TEXT      |
| email           | TEXT      |
| hashed_password | BYTEA     |
| created         | TIMESTAMP |

---

## 🚀 Getting Started

Clone the repository:

```bash
git clone https://github.com/anikmahidul9/snippetbox.git
```

Install dependencies:

```bash
go mod tidy
```

Configure your PostgreSQL database connection.

Run the application:

```bash
go run ./cmd/web
```

Open your browser:

```
https://localhost:4000
```

---

## 📚 What I Learned

* Building web applications using Go
* RESTful routing
* HTML templating
* Middleware patterns
* Dependency injection using an application struct
* PostgreSQL integration
* User authentication
* Session management
* CSRF protection
* HTTPS/TLS
* Form validation
* Error handling
* Testing templates
* Writing clean and maintainable Go code

---

## 🔮 Future Improvements

* Edit snippets
* Delete snippets
* User profiles
* Search snippets
* Pagination
* REST API
* Docker support
* CI/CD with GitHub Actions
* Redis-backed session storage
* Rate limiting

---

## 📖 Acknowledgements

This project was built while studying **Let's Go** by **Alex Edwards**. It follows the architecture and best practices presented throughout the book while serving as a practical learning project for modern Go web development.
