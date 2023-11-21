# API Documentation for Article Management Service

This document provides detailed instructions on how to interact with the Article Management Service. The service offers a simple and efficient way to manage articles, including creating articles, adding images to articles, and listing articles with or without images.

## Overview

The service exposes three primary endpoints:

- **Create Article**: To create a new article.
- **Add Image to Article**: To attach an image to an existing article.
- **List Articles**: To list all articles, with the option to filter articles based on whether they have images attached or not.

Below are the detailed instructions for each endpoint:

---

### Create Article

- **Endpoint**: `/create-article`
- **Method**: `POST`
- **Description**: This endpoint is used to create a new article.

#### Request Body

The request body should contain a JSON object with the following properties:

- `ID`: A unique identifier for the article.
- `Title`: The title of the article.
- `Description`: A description of the article (up to 4000 characters).
- `ExpireAt`: The expiration date and time of the article in ISO 8601 format.
- `ImagePaths`: An initial array of image paths (can be empty).

#### Example Request

```json
{
    "ID": "1",
    "Title": "My First Article",
    "Description": "This is a description of my first article.",
    "ExpireAt": "2024-01-01T00:00:00Z",
    "ImagePaths": []
}

### Add Image to Article

- **Endpoint**: `/add-image`
- **Method**: `POST`
- **Description**: Adds an image to an existing article by specifying the article's ID and the path of the image.

#### Request Body

The request should include a JSON object containing the following:

- `ArticleID`: The unique identifier of the article to which the image will be added.
- `ImagePath`: The file system path where the image is located.

#### Example Request

```json
{
    "ArticleID": "1",
    "ImagePath": "./images/lion.jpg"
}

#### Responses
Success (200 OK): Returns the updated article information with the added image path.
Article Not Found (404 Not Found): If the specified article does not exist, it returns an "Article not found" message.
Maximum Images Exceeded (400 Bad Request): If the article already has the maximum number of images allowed.

### List Articles

- **Endpoint**: `/list-articles`
- **Method**: `GET`
- **Description**: Lists all articles with the option to filter them based on whether they have images attached.

#### Query Parameters

- `withImages`: A boolean (`true` or `false`) to filter the articles. 
   - `true`: List only articles that have images attached.
   - `false`: List only articles that do not have any images attached.
   - Omitting this parameter will list all articles.

#### Example Requests

- To list articles with images: `http://localhost:8080/list-articles?withImages=true`
- To list articles without images: `http://localhost:8080/list-articles?withImages=false`

#### Response

- Returns a list of articles in JSON format, filtered according to the `withImages` query parameter.