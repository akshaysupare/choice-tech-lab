# Choice Tech Project

## Prerequisites
- Go 1.20+
- MySQL
- Redis

## Setup
1. Clone the repo.
2. Configure your MySQL and Redis connection in `config/config.go` or via environment variables.
3. Run:
   ```sh
   go mod tidy
   go run ./cmd/main.go
   ```

## API Endpoints

- `POST /import` — Upload Excel file (form-data, key: `file`)
- `GET /records` — Get all records
- `PUT /records/:id` — Update a record
- `DELETE /records/:id` — Delete a record

## Excel Format
Headers must be:
```
first_name, last_name, company_name, address, city, county, postal, phone, email, web
```

## Data Validation
- All columns must be present and non-empty.
- `email` must be a valid email address.
- `phone` must be a valid phone number (digits, spaces, dashes, plus allowed).

## Caching
- Data is cached in Redis for 5 minutes.
- If Redis is unavailable, the server will not start.

## Batch Import
- Data is inserted into MySQL in batches of 100 rows for efficiency.

## Error Handling
- All endpoints return clear error messages for invalid input or server errors.

## Postman Collection
- See `ChoiceTechProject.postman_collection.json` for ready-to-use API tests.
