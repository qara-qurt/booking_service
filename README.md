# Booking Service

## Description

This service allows booking rooms. It uses PostgreSQL as the database and is implemented in Go using the `chi` library for HTTP request routing and `pgx` for database interaction.

## Requirements

To run the application, you'll need the following dependencies:

- Go 1.19 or higher
- PostgreSQL
- Docker and Docker Compose
- Go libraries:
  - `chi`
  - `pgx`

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/qara-qurt/booking_service.git
   cd ./booking_service
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

## Running the Application

### Start the Application

To start the application using Docker Compose, run the following command:

```bash
make run
```

### Stop the Application

To stop the application, run the following command:

```bash
make stop
```

## Testing

### Running Tests

To run tests, use the following command:

```bash
make test
```

### Stopping the Test Environment

If you need to manually stop the test environment, use the following command:

```bash
make test_stop
```

## Available Routes

### 1. **Create Reservation**

- **Method:** `POST`
- **Path:** `/reservation`
- **Description:** Creates a new reservation.
- **Request Body:**
  ```json
  {
    "room_id": "string",
    "start_time": "2024-08-30T12:00:00Z",
    "end_time": "2024-08-30T14:00:00Z"
  }
  ```
- **Response:**
  - **201 Created:** Reservation was successfully created.
  - **400 Bad Request:** Invalid request payload, or start time is after end time.
  - **409 Conflict:** Room is already reserved for the given time range.
  - **500 Internal Server Error:** An error occurred while creating the reservation.

### 2. **Get Reservations by Room**

- **Method:** `GET`
- **Path:** `/reservation/{roomID}`
- **Description:** Retrieves all reservations for a specific room.
- **Path Parameter:**
  - `roomID` (string): The ID of the room.
- **Response:**
  - **200 OK:** A list of reservations for the specified room.
  - **400 Bad Request:** The `roomID` is missing from the URL.
  - **500 Internal Server Error:** An error occurred while retrieving the reservations.

## Example Requests

### 1. **Create Reservation**

**Request:**

```bash
curl -X POST http://localhost:3000/reservation -H "Content-Type: application/json" -d '{
  "room_id": "room123",
  "start_time": "2024-08-30T12:00:00Z",
  "end_time": "2024-08-30T14:00:00Z"
}'
```

**Response:**

```http
HTTP/1.1 201 Created
```

### 2. **Get Reservations by Room**

**Request:**

```bash
curl -X GET http://localhost:3000/reservation/room123
```

**Response:**

```json
[
  {
    "id": 1,
    "room_id": "room123",
    "start_time": "2024-08-30T12:00:00Z",
    "end_time": "2024-08-30T14:00:00Z"
  },
  {
    "id": 2,
    "room_id": "room123",
    "start_time": "2024-08-31T09:00:00Z",
    "end_time": "2024-08-31T11:00:00Z"
  }
]
```
