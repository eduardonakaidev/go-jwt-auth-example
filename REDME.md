# JWT Authentication Example Application

Welcome to the **JWT Authentication Example** application repository! This project demonstrates how to implement authentication using JSON Web Tokens (JWT) in a simple Go (Golang) application.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Setup Instructions](#setup-instructions)
- [Usage](#usage)
  - [Endpoints](#endpoints)
- [How It Works](#how-it-works)

---

## Overview
This application provides a basic implementation of JWT-based authentication in a Go web server. It includes routes that demonstrate both public and protected access, with the latter requiring a valid JWT token to access.

## Features
- User authentication using JWT tokens
- Middleware for token validation
- Public and secure endpoints
- Easy-to-understand code structure for learning purposes

## Setup Instructions

### Prerequisites
- [Go](https://golang.org/dl/) installed (version 1.17 or higher recommended)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/eduardonakaidev/jwt-auth-example.git
   cd jwt-auth-example
   ```
2. Install dependencies (if applicable):
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

The server will start at `http://localhost:3000`.

---

## Usage

### Endpoints

#### 1. **Authenticate User**
- **Endpoint**: `/api/auth`
- **Method**: `POST`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body** (JSON):
  ```json
  {
    "username": "eduardo",
    "password": "123456"
  }
  ```
- **Response** (Success):
  ```json
  {
    "token": "<JWT_TOKEN>",
    "claims": {
      "sub": "eduardo",
      "iat": 1713204230,
      "exp": 1713204245
    }
  }
  ```

#### 2. **Public Endpoint**
- **Endpoint**: `/api/public`
- **Method**: `GET`
- **Description**: A public endpoint accessible to everyone.
- **Response**:
  ```text
  Everyone can view this endpoint
  ```

#### 3. **Secure Endpoint**
- **Endpoint**: `/api/secure`
- **Method**: `GET`
- **Description**: A secure endpoint that requires a valid JWT token.
- **Headers**:
  ```
  X-Api-Token: <JWT_TOKEN>
  ```
- **Response** (Success):
  ```text
  You are authenticated
  ```
- **Response** (Failure):
  ```text
  not authorized
  ```

---

## How It Works

1. **Authentication**:
   - The `/api/auth` endpoint checks for valid credentials and generates a JWT token containing the user claims (`sub`, `iat`, and `exp`).

2. **Public Access**:
   - The `/api/public` endpoint can be accessed by anyone without any authentication.

3. **Protected Access**:
   - The `/api/secure` endpoint is protected by the `AuthMiddleware`. It validates the JWT token provided in the `X-Api-Token` header. Only requests with a valid token can access this route.

4. **Token Validation**:
   - The `AuthMiddleware` uses the `ValidateToken` utility to check the token's validity and expiration.
