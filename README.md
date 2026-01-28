# Kasir API (Go Backend)

A robust Point of Sale (POS) backend API built with Go, PostgreSQL (Supabase), and Clean Architecture.

## 🚀 Key Features
*   **Clean Architecture**: Separation of concerns (Domain, Repository, Service, Handler).
*   **Standard Go Layout**: scalable folder structure (`cmd/`, `internal/`).
*   **Context Propagation**: Proper timeout and cancellation handling.
*   **PostgreSQL**: Reliable data storage with Supabase.
*   **Configuration**: Environment-based config using Viper.

## 📂 Project Structure
```
kasir-api/
├── cmd/
│   ├── api/            # API Server Entry point
│   └── migrate/        # Database Migration Tool
├── internal/
│   ├── config/         # Configuration Loader (Viper)
│   ├── domain/         # Business Entities & Interfaces (Pure Go)
│   ├── handler/        # HTTP Handlers (Transport Layer)
│   ├── repository/     # Database Implementations (Data Layer)
│   └── service/        # Business Logic (Use Case Layer)
├── .env                # Environment Variables
├── go.mod              # Dependency Management
└── README.md           # Documentation
```

## 🛠 Prerequisites
*   [Go 1.22+](https://go.dev/)
*   [PostgreSQL](https://www.postgresql.org/) (or Supabase)

## ⚡️ Quick Start

### 1. Clone & Dependencies
```bash
git clone <repository-url>
cd kasir-api
go mod tidy
```

### 2. Configure Environment
Create a `.env` file in the root directory:
```env
PORT=:8080
DB_CONN=postgres://user:password@host:5432/dbname?sslmode=require
```

### 3. Run Migrations
Initialize the database schema:
```bash
go run cmd/migrate/main.go
```

### 4. Start Server
Run the API server:
```bash
go run cmd/api/main.go
```
The server will start at `http://localhost:8080`.

## 📡 API Endpoints

### Products
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/products` | Get all products (with Category Name) |
| `GET` | `/api/products/{id}` | Get product by ID |
| `POST` | `/api/products` | Create new product |
| `PUT` | `/api/products/{id}` | Update product |
| `DELETE` | `/api/products/{id}` | Delete product |

**Product Payload:**
```json
{
    "name": "Laptop",
    "description": "Gaming Laptop",
    "price": 15000000,
    "stock": 10,
    "category_id": 1
}
```

### Categories
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/categories` | Get all categories |
| `POST` | `/api/categories` | Create new category |
| ... | ... | (Standard CRUD) |

## 🏗 Architecture Decisions
1.  **Dependency Injection**: Dependencies are injected via constructors (e.g., `NewProductService(repo)`).
2.  **Interfaces**: Services depend on repository **interfaces**, not concrete structs, enabling easier unit testing.
3.  **Context**: `context.Context` is passed through all layers to allow request cancellation and timeout propagation to the database.
