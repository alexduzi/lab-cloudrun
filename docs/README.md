# API Documentation

This directory contains the OpenAPI/Swagger documentation for the Weather API.

## Files

- `swagger.json` - OpenAPI specification in JSON format
- `swagger.yaml` - OpenAPI specification in YAML format
- `docs.go` - Generated Go code for Swagger integration

## Accessing the Documentation

When the application is running, you can access the interactive Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Regenerating Documentation

To regenerate the Swagger documentation after making changes to API annotations:

```bash
make swagger
```

Or manually:

```bash
swag init -g cmd/api/main.go -o docs
```

## API Endpoints

### Weather
- `GET /{cep}` - Get temperature by Brazilian postal code (CEP)
  - Success: 200 - Returns temperature in Celsius, Fahrenheit and Kelvin
  - Error: 404 - CEP not found
  - Error: 422 - Invalid CEP format

### Health
- `GET /health` - Health check endpoint
- `GET /readiness` - Readiness check endpoint
