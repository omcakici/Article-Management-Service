Article Management Service
Overview
The Article Management Service is a Go-based API for storing, managing, and retrieving articles. Articles in this context are defined as items having a title, description, expiration date, and optionally, attached images.

Key Features
Article Storage: Articles are stored in memory with unique identifiers, titles, descriptions (up to 4KB), expiration dates, and up to 3 image paths.
Image Path Attachment: Ability to attach image paths to articles, with each image not exceeding 5MB.
Automatic Removal: Articles are automatically removed from storage upon reaching their expiration date.
API Endpoints: The service provides RESTful endpoints to create articles, add image paths, and list articles with options to filter based on image attachment.
Project Structure
api/: Contains the API handlers for the various endpoints.
cmd/: Includes the main application entry point.
config/: Configuration-related files.
docs/: Documentation files, including detailed API endpoint information.
pkg/: Core packages including article definitions, constants, and storage handling.
Usage
Starting the Service: Run the application from the cmd directory. This starts an HTTP server listening for requests related to article management.
Interacting with the API: Use tools like Postman or curl to interact with the available endpoints. The service supports operations like creating a new article, adding image paths to an article, and listing articles with various filters.
API Documentation: For detailed information about the API endpoints and how to use them, refer to docs/API.md.
Note
This project is set up as a demonstration of Go programming capabilities and app design. It is structured for ease of understanding and further development.

