# Choice Tech Project

## Prerequisites & Setup

1. **Create MySQL Database**
   - Ensure MySQL is installed and running.
   - Create a database named `choicetech`:
     ```sql
     CREATE DATABASE choicetech;
     ```

2. **Install Redis (Optional but Recommended)**
   - Redis is used for caching. If not installed, you can skip, but caching features will not work.
   - [Redis Quick Start Guide](https://redis.io/docs/getting-started/)

3. **Install Go Modules and Dependencies**
   - Open a terminal in your project directory and run:
     ```sh
     go mod tidy
     ```
   - This will download all necessary Go packages, including:
     - [gin-gonic/gin](https://github.com/gin-gonic/gin) (HTTP framework)
     - [gorm.io/gorm](https://gorm.io/) (ORM for MySQL)
     - [gorm.io/driver/mysql](https://gorm.io/docs/connecting_to_the_database.html#MySQL)
     - [github.com/xuri/excelize/v2](https://github.com/xuri/excelize) (Excel parsing)
     - [github.com/go-redis/redis/v8](https://github.com/go-redis/redis) (Redis client)

4. **Configure Environment (Optional)**
   - You can override MySQL/Redis connection settings using environment variables:
     - `MYSQL_DSN` (default: `root:root@tcp(127.0.0.1:3306)/choicetech`)
     - `REDIS_ADDR` (default: `localhost:6379`)
     - `REDIS_PASSWORD` (default: `""`)
     - `REDIS_DB` (default: `0`)

5. **Run the Project**
   - Start the server with:
     ```sh
     go run .
     ```

6. **API Usage**
   - Use Postman or similar tools to interact with the API endpoints.
   - See the included Postman collection for ready-to-use requests.

---

**Language:** Go  
**Database:** MySQL  
**Cache:** Redis (optional, but recommended for full functionality)

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
