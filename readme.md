# Choto-Link

Choto-Link is a URL shortening service built with Go and Redis. It provides a RESTful API for creating, managing, and accessing shortened URLs with features like rate limiting and custom short URLs.

## Features

- URL shortening with auto-generated or custom short URLs
- Configurable URL expiration (default 24 hours)
- Rate limiting to prevent abuse
- RESTful API for CRUD operations
- Redis-based storage for high performance
- Docker support for easy deployment

## Prerequisites

- Go
- Docker and Docker Compose (for containerized deployment)

## Installation

1. Clone the repository

2. Rename the `.env.example` file to `.env` and update with your credentials

3. Build and run using Docker Compose:

```bash
docker-compose up --build
```

The service will be available at `http://localhost:8080`

## API Endpoints

### Create Short URL

```http
POST /api/shorten
```

Request body:

```json
{
  "url": "https://example.com/very-long-url",
  "custom_short": "custom123", // optional
  "expiry": 3600 // optional, in seconds
}
```

### Get Original URL

```http
GET /api/shorten/:shortUrl
```

### Update Short URL

```http
PUT /api/shorten/:shortUrl
```

Request body:

```json
{
  "url": "https://example.com/new-url"
}
```

### Delete Short URL

```http
DELETE /api/shorten/:shortUrl
```

### Redirect to Original URL

```http
GET /:shortUrl
```

## FAQ

#### What is Choto?

Choto is a Bengali word that means "small" or "short". It's a common term used in Bangladesh to refer to something that is small in size or duration.
