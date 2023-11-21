# Article Management Service

## Overview
The Article Management Service is a Go-based API designed for efficient management and retrieval of articles. Each article features a title, description, expiration date, and optional image attachments.

## Key Features
- **Article Storage**: Articles are stored in memory with unique identifiers, titles, descriptions (up to 4KB), expiration dates, and up to 3 image paths.
- **Image Path Attachment**: Ability to attach up to three image paths per article, each not exceeding 5MB.
- **Automatic Removal**: Automated removal of articles upon reaching their expiration date.
- **API Endpoints**: RESTful endpoints to create articles, add image paths, and list articles with filtering options based on image attachment.

## Project Structure
- `api/`: Contains the API handlers for various endpoints.
- `cmd/`: Includes the main application entry point.
- `config/`: Configuration-related files.
- `docs/`: Documentation files, including detailed API endpoint information.
- `pkg/`: Core packages including article definitions, constants, and storage handling.

## Usage
1. **Starting the Service**: Run the application from the `cmd` directory. This starts an HTTP server listening for requests related to article management.
2. **Interacting with the API**: Use tools like Postman or curl to interact with the available endpoints. The service supports operations like creating a new article, adding image paths to an article, and listing articles with various filters.
3. **API Documentation**: For detailed information about the API endpoints and how to use them, refer to `docs/API.md`.

## Note
This project is set up as a demonstration of Go programming capabilities and app design. It is structured for ease of understanding and further development.
