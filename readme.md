# Dev Setup
## 1. create container postgres

```
docker run --name my-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
psql -U postgres
create database eventdb;
```

# 2. dev environment for the program

```
wgo run main.go
```


# User Endpoints Documentation

## 1. Register
### 1.1 Register User
#### Request
```json
{
    "username": "test",
    "password": "test"
}
```
#### Response
```json
{
    "id": 1,
    "username": "test",
    "password": "test"
}
```

### 1.2 Register User with existing username
#### Request
```json
{
    "username": "test",
    "password": "test"
}
```
#### Response
```json
{
    "error": "username already exists"
}
```

## 2. Login
### 2.1 Login User
#### Request
```json
{
    "username": "test",
    "password": "test"
}
```
#### Response
```json
{
    "id": 1,
    "username": "test",
    "password": "test"
}
```

### 2.2 Login User with invalid username
#### Request
```json
{
    "username": "test",
    "password": "test"
}
```
#### Response
```json
{
    "error": "invalid username or password"
}
```

## 3. Get Events
### 3.1 Get All Events
#### Request
```json
{}
```
#### Response
```json
[
    {
        "id": 1,
        "name": "test",
        "description": "test",
        "location": "test",
        "date_time": "2023-10-11T00:00:00Z",
        "user_id": 1
    }
]
```

### 3.2 Get Event by ID
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

## 4. Create Event
### 4.1 Create Event
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 4.2 Create Event with invalid date_time
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

## 5. Update Event
### 5.1 Update Event
#### Request
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 5.2 Update Event with invalid date_time
#### Request
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

## 6. Delete Event
### 6.1 Delete Event
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 6.2 Delete Event with invalid id
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "error": "invalid id"
}
```

# Event Endpoints Documentation

## 1. Register
### 1.1 Register Event
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 1.2 Register Event with invalid date_time
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

## 2. Login
### 2.1 Login Event
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 2.2 Login Event with invalid date_time
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

### 2.3 Login Event with invalid date_time
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

### 2.4 Login Event with invalid date_time
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

## 3. Get Registrations
### 3.1 Get All Registrations
#### Request
```json
{}
```
#### Response
```json
[
    {
        "id": 1,
        "event_id": 1,
        "user_id": 1
    }
]
```

### 3.2 Get Registration by ID
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```

## 4. Create Registration
### 4.1 Create Registration
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```

### 4.2 Create Registration with invalid event_id
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid event_id"
}
```

### 4.3 Create Registration with invalid user_id
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid user_id"
}
```

## 5. Update Registration
### 5.1 Update Registration
#### Request
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```

### 5.2 Update Registration with invalid event_id
#### Request
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid event_id"
}
```

### 5.3 Update Registration with invalid user_id
#### Request
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid user_id"
}
```

## 6. Delete Registration
### 6.1 Delete Registration
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```

### 6.2 Delete Registration with invalid id
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "error": "invalid id"
}
```

# Registration Endpoints Documentation

## 1. Register
### 1.1 Register Registration
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```

### 1.2 Register Registration with invalid event_id
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid event_id"
}
```

### 1.3 Register Registration with invalid user_id
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid user_id"
}
```

## 2. Login
### 2.1 Login Registration
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "id": 1,
    "event_id": 1,
    "user_id": 1
}
```

### 2.2 Login Registration with invalid event_id
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid event_id"
}
```

### 2.3 Login Registration with invalid user_id
#### Request
```json
{
    "event_id": 1,
    "user_id": 1
}
```
#### Response
```json
{
    "error": "invalid user_id"
}
```

## 3.   Get Events
### 3.1 Get All Events
#### Request
```json
{}
```
#### Response
```json
[
    {
        "id": 1,
        "name": "test",
        "description": "test",
        "location": "test",
        "date_time": "2023-10-11T00:00:00Z",
        "user_id": 1
    }
]
```

### 3.2 Get Event by ID
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

## 4. Create Event
### 4.1 Create Event
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 4.2 Create Event with invalid date_time
#### Request
```json
{
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

## 5. Update Event
### 5.1 Update Event
#### Request
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 5.2 Update Event with invalid date_time
#### Request
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z"
}
```
#### Response
```json
{
    "error": "invalid date_time format"
}
```

## 6. Delete Event
### 6.1 Delete Event
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "id": 1,
    "name": "test",
    "description": "test",
    "location": "test",
    "date_time": "2023-10-11T00:00:00Z",
    "user_id": 1
}
```

### 6.2 Delete Event with invalid id
#### Request
```json
{
    "id": 1
}
```
#### Response
```json
{
    "error": "invalid id"
}
```