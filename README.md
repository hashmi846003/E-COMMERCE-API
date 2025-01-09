
WEBSITE: https://roadmap.sh/projects/scalable-ecommerce-platform

E-Commerce API
This project is an API for an e-commerce platform built with Go. It provides functionality for user authentication, product management, and payment processing.
Project Structure:
ecommerce-api/
├── cmd/
│   └── main.go
├── controllers/
│   ├── userController.go
│   └── productController.go
├── models/
│   ├── userModel.go
│   └── productModel.go
├── routes/
│   └── routes.go
├── utils/
│   └── auth.go
├── services/
│   └── payment.go
├── .env
├── go.mod
├── go.sum
└── README.md
Requirements
Go 1.17 or later

MongoDB

A tool like Postman to test the API endpoints.

Installation
Clone the repository:

bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
Install the dependencies:

bash
go mod tidy
Create a .env file in the root directory with the following content:

plaintext
PORT=8080
DB_URI=mongodb://localhost:27017/ecommerce
JWT_SECRET=your_jwt_secret
Run the application:

bash
go run cmd/main.go
API Endpoints
User Endpoints
Register: POST /api/users/register

json
{
    "username": "example",
    "email": "example@example.com",
    "password": "yourpassword"
}
Login: POST /api/users/login

json
{
    "email": "example@example.com",
    "password": "yourpassword"
}
Product Endpoints
Add Product: POST /api/products (Admin only)

json
{
    "name": "Product Name",
    "price": 99.99,
    "description": "Product Description",
    "inStock": true
}
Get Products: GET /api/products

Payment Endpoint
Checkout: POST /api/products/checkout

json
{
    "amount": 150.00
}




