# Sweet Shop Management System

A robust backend API for managing a sweet shop, built with Go and Gin.

## Features

- **User Authentication**: Register and Login with JWT.
- **Sweets Management**: CRUD operations for sweets.
- **Inventory Management**: Purchase and Restock sweets.
- **Search**: Search sweets by name or category.
- **Database**: PostgreSQL with GORM.

## Setup

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System.git
    cd Sweet-Shop-Management-System
    ```

2.  **Setup Database**:
    Ensure you have Docker installed.
    ```bash
    docker-compose up -d
    ```

3.  **Environment Variables**:
    Copy `.env.example` to `.env` and update if necessary.
    ```bash
    cp .env.example .env
    ```

4.  **Run the Application**:
    ```bash
    go run cmd/server/main.go
    ```

## API Endpoints

### Auth
- `POST /api/auth/register`: Register a new user.
- `POST /api/auth/login`: Login and get a token.

### Sweets (Protected)
- `POST /api/sweets`: Add a sweet.
- `GET /api/sweets`: List all sweets.
- `GET /api/sweets/search?q=query`: Search sweets.
- `PUT /api/sweets/:id`: Update a sweet.
- `DELETE /api/sweets/:id`: Delete a sweet (Admin only).

### Inventory (Protected)
- `POST /api/sweets/:id/purchase`: Purchase a sweet.
- `POST /api/sweets/:id/restock`: Restock a sweet (Admin only).
