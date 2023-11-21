Article Management Service
Overview
The Article Management Service is a robust and efficient service designed to store, manage, and automatically expire articles. It is built using Go, ensuring high performance and concurrency support. This service is ideal for applications that need to manage articles with associated images, offering functionalities like creation, updating, and automatic removal of articles upon expiration.

Features
Article Storage: Store articles with unique identifiers, titles, expiration dates, and descriptions (up to 4KB or 4000 characters).
Image Attachment: Support attaching up to 3 images per article (max size 5MB per image).
Efficient Retrieval: Quickly retrieve articles and their details.
Automatic Expiration: Automatically removes articles and their associated images on reaching the expiration date.
HTTP/GRPC Endpoints: Offers endpoints to create articles, add images to articles, and list article titles with or without attached images.

Endpoints
Create Article: Add a new article with the required details.
Add Image to Article: Attach images to an existing article by providing the image path.
List Article Titles: List titles of all articles, with options to filter articles with or without images.