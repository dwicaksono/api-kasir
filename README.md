# kasir-api

A RESTful API for a cashier (Point of Sale) system built with Go, following Clean Architecture principles.

## 📁 Project Structure

```
kasir-api/
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management (Viper)
│   ├── domain/
│   │   ├── product.go       # Product entity & interfaces
│   │   └── category.go      # Category entity & interfaces
│   ├── handler/
│   │   ├── product_handler.go   # HTTP handlers for products
│   │   └── category_handler.go  # HTTP handlers for categories
│   ├── repository/
│   │   ├── product_repository.go   # Database operations for products
│   │   └── category_repository.go  # Database operations for categories
│   └── usecase/
│       ├── product_usecase.go   # Business logic for products
│       └── category_usecase.go  # Business logic for categories
├── pkg/
│   └── database/
│       └── postgres.go      # Database connection utilities
├── .env                     # Environment variables (DO NOT COMMIT)
├── go.mod
└── go.sum
```

## 🏗️ Architecture Overview

This project follows **Clean Architecture** (also known as Hexagonal/Onion Architecture):

```
┌─────────────────────────────────────────────────────────┐
│                     HTTP Handlers                       │
│                   (internal/handler)                    │
├─────────────────────────────────────────────────────────┤
│                      Use Cases                          │
│                   (internal/usecase)                    │
├─────────────────────────────────────────────────────────┤
│                      Domain Layer                       │
│        Entities & Interfaces (internal/domain)          │
├─────────────────────────────────────────────────────────┤
│                     Repositories                        │
│                  (internal/repository)                  │
├─────────────────────────────────────────────────────────┤
│                       Database                          │
│                     (pkg/database)                      │
└─────────────────────────────────────────────────────────┘
```

### Layer Responsibilities

| Layer | Location | Responsibility |
|-------|----------|----------------|
| **Domain** | `internal/domain/` | Entities (structs) and interface definitions. No dependencies on other layers. |
| **Use Case** | `internal/usecase/` | Business logic. Orchestrates data flow between handlers and repositories. |
| **Repository** | `internal/repository/` | Data persistence. Implements domain interfaces using PostgreSQL. |
| **Handler** | `internal/handler/` | HTTP request/response handling. Parses input, calls use cases, returns JSON. |
| **Config** | `internal/config/` | Configuration loading via Viper (environment variables). |
| **Database** | `pkg/database/` | Shared database connection pool. |

### Dependency Flow

```
Handler → Usecase → Repository
           ↓
        Domain (interfaces)
```

> **Key Principle**: Dependencies point inward. The domain layer has no external dependencies.

## 🚀 Getting Started

### Prerequisites

- Go 1.22+ (uses `PathValue` for routing)
- PostgreSQL database

### Setup

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd kasir-api
   ```

2. **Configure environment variables**
   
   Create a `.env` file in the project root:
   ```env
   PORT=8080
   DB_CONN=postgresql://user:password@host:port/database
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```

   The server will start on the port specified in `.env` (default: `8080`).

## 📡 API Endpoints

### Health Check
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Check if API is running |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/products` | Get all products |
| GET | `/api/products/{id}` | Get product by ID |
| POST | `/api/products` | Create a new product |
| PUT | `/api/products/{id}` | Update a product |
| DELETE | `/api/products/{id}` | Delete a product |

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | Get all categories |
| GET | `/api/categories/{id}` | Get category by ID |
| POST | `/api/categories` | Create a new category |
| PUT | `/api/categories/{id}` | Update a category |
| DELETE | `/api/categories/{id}` | Delete a category |

## 🧪 Example Requests

### Create a Product
```bash
curl -X POST http://localhost:8080/api/produk \
     -H "Content-Type: application/json" \
     -d '{"nama": "Indomie Goreng", "harga": 4000, "stok": 100}'
```

### Get All Categories
```bash
curl http://localhost:8080/api/categories
```

### Update a Category
```bash
curl -X PUT http://localhost:8080/api/categories/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Makanan Ringan", "description": "Snack dan cemilan"}'
```

## 🔧 Technologies Used

- **Go 1.25** - Programming language
- **net/http** - Standard library HTTP server with Go 1.22+ routing
- **pgx/v5** - PostgreSQL driver
- **Viper** - Configuration management

## 📝 License

MIT License
