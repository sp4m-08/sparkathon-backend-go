# UPC Backend Go

A Go-based backend service that enables barcode scanning, product lookup, and AI-powered conversations about products. Users can scan UPC codes to retrieve product information and engage in intelligent chat sessions about the products using an integrated LLM service.

## Server

- `https://upc-backend-go.onrender.com/`

## 📡 API Endpoints

### Product & Chat System

- `POST /products/scan` - Scan product UPC and initialize chat session
- `POST /products/chat` - Chat about scanned product
- `GET /products/chat/history/:session_id` - Get chat history for a session
- `GET /products/chat/health` - Check LLM service health

### Endpoint Details

#### 1. Scan Product and Initialize Chat

```
POST /products/scan
```

**Request Body:**

```json
{
  "upc": "8901023010415"
}
```

**Response:**

```json
{
  "session_id": "1a2b3c4d...",
  "product": {
    "id": 2,
    "ean": "8901023010415",
    "title": "Pack of 5 Cinthol Cool Soap..."
    //...
  }
}
```

#### 2. Chat About Product

```
POST /products/chat
```

**Request Body:**

```json
{
  "session_id": "1a2b3c4d...",
  "message": "What is this product used for?"
}
```

**Response:**

```json
{
  "response": "This is a pack of Cinthol Cool Soaps ideal for summer...",
  "session_id": "1a2b3c4d...",
  "history": [
    {
      "role": "user",
      "content": "What is this product used for?"
    },
    {
      "role": "assistant",
      "content": "This is a pack of Cinthol Cool Soaps..."
    }
  ]
}
```

#### 3. Get Chat History

```
GET /products/chat/history/:session_id
```

**Response:**

```json
{
  "session_id": "1a2b3c4d...",
  "history": [
    {
      "role": "system",
      "content": "Additional context provided: ..."
    },
    {
      "role": "user",
      "content": "What is this product used for?"
    },
    {
      "role": "assistant",
      "content": "This is a pack of Cinthol Cool Soaps..."
    }
  ]
}
```

#### 4. LLM Health Check

```
GET /products/chat/health
```

**Response:**

```json
{
  "status": "healthy",
  "database": "connected"
}
```

## 🚀 Features

- **Product Scanning System** - UPC barcode scanning and product lookup
- **AI-Powered Chat** - Intelligent product conversations using LLM
- **Session Management** - Persistent chat sessions with product context
- **Database Integration** - PostgreSQL with GORM for product storage
- **External LLM API** - Integration with Spark AI service
- **RESTful API** - Clean and well-structured REST endpoints
- **Error Handling** - Comprehensive error handling and logging
- **Environment Configuration** - Flexible configuration management
- **Health Monitoring** - LLM service health checks

## 📋 Prerequisites

Before running this project, make sure you have the following installed:

- Go 1.21 or higher
- Database (PostgreSQL/MongoDB/MySQL)
- Docker (optional, for containerized deployment)
- Git

## 🔧 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/sp4m-08/sparkathon-backend-go.git
cd sparkathon-backend-go
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
# Server Configuration
PORT=8080

# Database Configuration (Supabase PostgreSQL)
DATABASE_URL=postgresql://username:password@db.supabase.co:5432/postgres

# LLM Service Configuration
LLM_API_BASE=https://spark-ai-1rd9.onrender.com
```

### 4. Database Setup

```bash
# Set up your Supabase PostgreSQL connection
# Update DATABASE_URL in .env with your Supabase credentials
# The application will auto-migrate the Product model
```

### 5. Run the Application

```bash
# Development mode
go run main.go

# Or build and run
go build -o bin/server main.go
./bin/server
```

The server will start on `http://localhost:8080`

## Project Structure

```
upc-backend-sparkathon/
├── controllers/
│   └── productController.go  # Product and chat endpoints
├── models/
│   └── product.go           # Product model
├── routes/
│   └── productRoutes.go     # Route definitions
├── services/
│   ├── db.go               # Database initialization
│   ├── llm_client.go       # LLM API client
│   └── upc_api.go          # UPC API service
├── .env                    # Environment variables
├── .gitignore
├── go.mod
├── go.sum
├── main.go                 # Application entry point
└── README.md
```
