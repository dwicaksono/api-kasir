# Kasir API (Go Backend)

A robust Point of Sale (POS) backend API built with Go, PostgreSQL (Supabase), and Clean Architecture.

## 🚀 Key Features
*   **Clean Architecture**: Separation of concerns (Domain, Usecase, Repository, Handler).
*   **Standard Go Layout**: Scalable folder structure (`cmd/`, `internal/`).
*   **Deployment Ready**: Compatible with standard Go build pipelines via root `main.go`.
*   **Context Propagation**: Proper timeout and cancellation handling.
*   **PostgreSQL**: Reliable data storage with Supabase (`pgx` driver).
*   **Configuration**: Environment-based config using Viper.

## 📂 Project Structure
```
kasir-api/
├── cmd/
│   ├── api/            # API Entry Point (Standard Layout)
│   └── migrate/        # Database Migration Tool
├── internal/
│   ├── app/            # Application Bootstrap Logic
│   ├── config/         # Configuration Loader (Viper)
│   ├── domain/         # Business Entities & Interfaces (Pure Go)
│   ├── handler/        # HTTP Handlers (Transport Layer)
│   ├── repository/     # Database Implementations (Data Layer)
│   └── usecase/        # Business Logic & Orchestration
├── pkg/                # Shared Packages (e.g., Database Connection)
├── .env                # Environment Variables
├── go.mod              # Dependency Management
├── main.go             # Root Entry Point (Platform Deployment)
└── README.md           # Documentation
```

## 🛠 Prerequisites
*   [Go 1.23+](https://go.dev/)
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
Run the API server (supports both methods):

**Method A: Standard Run (Root)**
```bash
go run main.go
```

**Method B: Command Dir**
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
1.  **Dependency Injection**: Dependencies are injected via constructors (e.g., `NewCategoryUsecase(repo)`).
2.  **Usecase Layer**: Contains business logic, orchestrating data flow between Handlers and Repositories.
3.  **Interfaces**: Usecases depend on Repository **interfaces**, allowing for easy mocking and testing.
4.  **Context**: `context.Context` is passed through all layers to ensure request lifecycle management.
