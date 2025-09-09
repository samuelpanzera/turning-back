# Product Overview

## Turning Back - Hexagonal Architecture

A Go-based REST API for managing "orçamentos" (budget quotes/estimates) built using Hexagonal Architecture (Ports & Adapters pattern). The application serves as a reference implementation of clean architecture principles in Go.

### Core Features
- Create, read, update, and delete budget quotes
- Multi-environment support (SQLite for development, PostgreSQL for production)
- Legacy endpoint compatibility (`/orcament`)
- Modern REST API (`/api/v1/orcamentos`)
- Health check and monitoring endpoints

### Business Domain
The application manages budget quotes with the following key attributes:
- Customer information (name, email, phone)
- Quote details (quantity of pieces, description)
- Optional file attachments
- Timestamps for creation and updates

### Architecture Goals
- Demonstrate hexagonal architecture implementation
- Maintain separation of concerns
- Enable easy testing and mocking
- Support multiple database backends
- Provide clean, maintainable code structure