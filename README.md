# GoFr CRUD API with PostgreSQL

This is a simple CRUD (Create, Read, Update, Delete) API built using the GoFr framework in Go, with PostgreSQL as the database. The API manages a "customers" table with columns: `id`, `name`, `email`, and `phone`.


    
## Table of Contents
- [Requirements](#requirements)
- [Setup](#setup)
- [Configuration](#configuration)
- [Endpoints](#endpoints)
  - [1. Get all customers](#1-get-all-customers)
  - [2. Get a customer by ID](#2-get-a-customer-by-id)
  - [3. Create a new customer](#3-create-a-new-customer)
  - [4. Update a customer](#4-update-a-customer)
  - [5. Delete a customer](#5-delete-a-customer)
- [Run](#run)
## Requirements

Ensure you have the following installed:
- Go 
- GoFr framework
- PostgreSQL database
## Setup


1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gofr-crud-api.git
   ```
2. Change into the project directory:
   ```bash
   cd gofr-crud-api
   ```
3. Install dependencies:
    ```bash
    go mod tidy
    ```
## Configuration

Update the config/.env file with your PostgreSQL database connection details.

    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=admin
    DB_NAME=postgres
    DB_PORT=5432
    DB_DIALECT=postgres
## API Reference

#### Get all customers

```http
  GET /customer
```

![GETALL](https://github.com/ayushvaish2511/CRUD_GOFR/assets/72246792/934bf562-3a18-45af-be42-5688bc4c882d)


#### Get customer details by ID

```http
  GET /customer/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

![GETbyID](https://github.com/ayushvaish2511/CRUD_GOFR/assets/72246792/2919d6a9-afb2-43bd-be94-a0ed874bd876)

#### Create  New Customer

```http
  POST /customer
```
![CREATE](https://github.com/ayushvaish2511/CRUD_GOFR/assets/72246792/dd6ea632-055e-4b57-ba6a-3000bc26810c)

#### Update customer details by ID

```http
  PUT /customer/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

![UPDATE](https://github.com/ayushvaish2511/CRUD_GOFR/assets/72246792/6cf77442-b59a-4d8d-95e7-e2f5cb28d089)

#### Delete customer details by ID

```http
  DELETE /customer/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

![DELETE](https://github.com/ayushvaish2511/CRUD_GOFR/assets/72246792/ea8e39b5-bb99-4258-b512-258184a3edfc)


## Run 

Run the following command to start the API Server

```bash
  go run main.go
```

The server will start on http://localhost:9090
