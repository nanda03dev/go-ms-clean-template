# Golang Microservice Template

This repository provides a template for building a microservice in Go, following the Clean Architecture principles. It is designed to facilitate maintainability, testability, and scalability.

## Project Structure

```
project/
├── domain/                      
│   ├── user.go                  
│   └── user_repository.go       
├── application/                 
│   └── user_service.go          
├── infrastructure/              
│   ├── db/                     
│   │   ├── mongo.go             # MongoDB connection code
│   │   ├── postgres.go          # PostgreSQL connection code
│   │   └── db.go                # Common interface and connection logic
│   └── user_repository_impl.go  
├── interfaces/                  
│   ├── user_handler.go          
│   └── dto/
│       └── create_user_dto.go   
├── main.go                      
└── go.mod                       
```

## Layers Explained

- **Domain Layer**: Contains the core business models and domain logic. It defines entities and their relationships, encapsulating the business rules.

- **Application Layer**: Holds the business logic and use cases. It coordinates the flow of data between the domain and infrastructure layers, processing requests and responses.

- **Infrastructure Layer**: Manages data access and external integrations (e.g., database interactions, third-party APIs). It provides concrete implementations for the repository interfaces defined in the domain layer.

- **Interface Layer**: Contains HTTP handlers for routing incoming requests and managing responses. It may also include Data Transfer Objects (DTOs) for structured data interchange between client and server.

## Getting Started

### Prerequisites

- Go 1.XX or higher
- Dependencies specified in `go.mod`

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd project
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application

To start the application, run:
```bash
go run main.go
```

### Testing

To run tests, execute:
```bash
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Clean Architecture by Robert C. Martin](https://www.oreilly.com/library/view/clean-architecture-a/9780134494166/)
- [Golang Documentation](https://golang.org/doc/)

```
